package repo

import (
	"context"
	"fmt"

	"isms/internal/data/model"
	"isms/internal/data/query"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type developerRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

func NewDeveloperRepo(db *gorm.DB, logger log.Logger) domain.DeveloperRepo {
	return &developerRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/developer_repo")),
	}
}

// Create 创建
func (r *developerRepo) Create(ctx context.Context, dev *domain.Developer) (*domain.Developer, error) {
	// 领域模型→数据模型（ORM实体）
	dataModel := &model.IsmsDeveloper{
		NameZh:      dev.NameZh,
		NameEn:      dev.NameEn,
		CountryID:   dev.CountryID,
		Website:     dev.Website,
		Description: dev.Description,
	}

	// 调用生成的query插入数据
	err := r.query.IsmsDeveloper.WithContext(ctx).Create(dataModel)
	if err != nil {
		return nil, err
	}

	// 数据模型→领域模型（返回给上层）
	dev.ID = dataModel.ID
	dev.CreatedAt = dataModel.CreatedAt
	dev.UpdatedAt = dataModel.UpdatedAt
	return dev, nil
}

// ExistByName 检查名称是否存在
func (r *developerRepo) ExistByName(ctx context.Context, nameZh string) (bool, error) {
	count, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.NameZh.Eq(nameZh)).
		Count()
	if err != nil {
		return false, fmt.Errorf("查询开发商名称失败: %w", err)
	}
	return count > 0, nil
}

// ExistByNameExcludeID 检查名称是否存在（排除指定ID）
func (r *developerRepo) ExistByNameExcludeID(ctx context.Context, nameZh string, excludeID int32) (bool, error) {
	count, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.NameZh.Eq(nameZh),
			r.query.IsmsDeveloper.ID.Neq(excludeID)).
		Count()
	if err != nil {
		return false, fmt.Errorf("查询开发商名称失败: %w", err)
	}
	return count > 0, nil
}

// GetByID 根基ID查询
func (r *developerRepo) GetByID(ctx context.Context, id int32) (*domain.Developer, error) {
	dataModel, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.ID.Eq(id)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("开发商ID不存在: %d", id)
		}
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	return &domain.Developer{
		ID:          dataModel.ID,
		NameZh:      dataModel.NameZh,
		NameEn:      dataModel.NameEn,
		CountryID:   dataModel.CountryID,
		Website:     dataModel.Website,
		Description: dataModel.Description,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}, nil
}

// List 分页查询
func (r *developerRepo) List(
	ctx context.Context,
	page, pageSize, countryID uint32,
	keyword string,
) ([]*domain.Developer, int64, error) {
	// 1. 基础查询：指定表名，绑定上下文
	db := r.db.WithContext(ctx).Table("isms_developer")

	// 2. 构建筛选条件
	if countryID > 0 {
		db = db.Where("country_id = ?", countryID)
	}
	if keyword != "" {
		db = db.Where("name_zh LIKE ?", "%"+keyword+"%").
			Or("name_en LIKE ?", "%"+keyword+"%")
	}

	// 3. 查询总条数（用于分页）
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询总数失败: %w", err)
	}

	// 4. 分页查询
	offset := (page - 1) * pageSize
	var modelDevs []*model.IsmsDeveloper
	if err := db.Offset(int(offset)).Limit(int(pageSize)).Find(&modelDevs).Error; err != nil {
		return nil, 0, fmt.Errorf("查询数据失败: %w", err)
	}

	// 5. 转换数据模型到领域模型
	domainDevs := make([]*domain.Developer, 0, len(modelDevs))

	// 批量查询国家信息（优化N+1查询问题）
	countryIDs := make([]int32, 0)
	countryIDSet := make(map[int32]struct{})
	for _, m := range modelDevs {
		if m.CountryID != 0 {
			if _, exists := countryIDSet[m.CountryID]; !exists {
				countryIDSet[m.CountryID] = struct{}{}
				countryIDs = append(countryIDs, m.CountryID)
			}
		}
	}

	// 构建国家ID到名称的映射
	countryMap := make(map[int32]string)
	if len(countryIDs) > 0 {
		countries, err := r.query.IsmsCountry.WithContext(ctx).
			Select(r.query.IsmsCountry.ID, r.query.IsmsCountry.NameZh).
			Where(r.query.IsmsCountry.ID.In(countryIDs...)). // 修正：使用r.query实例
			Find()
		if err != nil {
			r.log.Error("批量查询国家信息失败: %v", err)
			return nil, 0, fmt.Errorf("批量查询国家信息失败: %w", err)
		}

		for _, c := range countries {
			if c.NameZh != "" {
				countryMap[c.ID] = c.NameZh
			} else {
				countryMap[c.ID] = "未知国家"
				r.log.Warn("国家ID为%d的名称为空", c.ID)
			}
		}
	}

	for _, m := range modelDevs {
		var countryNameZh string
		if m.CountryID != 0 {
			if name, ok := countryMap[m.CountryID]; ok {
				countryNameZh = name
			} else {
				countryNameZh = "未知国家"
				r.log.Warn("未找到国家ID为%d的记录", m.CountryID)
			}
		} else {
			countryNameZh = ""
		}

		domainDevs = append(domainDevs, &domain.Developer{
			ID:            m.ID,
			NameZh:        m.NameZh,
			NameEn:        m.NameEn,
			CountryID:     m.CountryID,
			CountryNameZh: countryNameZh,
			Website:       m.Website,
			Description:   m.Description,
			CreatedAt:     m.CreatedAt,
			UpdatedAt:     m.UpdatedAt,
		})
	}

	return domainDevs, total, nil
}

// Update 更新开发商信息
func (r *developerRepo) Update(ctx context.Context, dev *domain.Developer) (*domain.Developer, error) {
	dataModel := &model.IsmsDeveloper{
		ID:          dev.ID,
		NameZh:      dev.NameZh,
		NameEn:      dev.NameEn,
		CountryID:   dev.CountryID,
		Website:     dev.Website,
		Description: dev.Description,
		// 注意：不要更新CreatedAt字段
	}

	// 使用Save方法更新记录（会更新所有字段）
	err := r.query.IsmsDeveloper.WithContext(ctx).Save(dataModel)
	if err != nil {
		r.log.Errorf("更新开发商失败: %v", err)
		return nil, err
	}

	// 转换回领域模型并返回
	return &domain.Developer{
		ID:          dataModel.ID,
		NameZh:      dataModel.NameZh,
		NameEn:      dataModel.NameEn,
		CountryID:   dataModel.CountryID,
		Website:     dataModel.Website,
		Description: dataModel.Description,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}, nil
}

// Count 统计开发商总数
func (r *developerRepo) Count(ctx context.Context) (int64, error) {
	count, err := r.query.IsmsDeveloper.WithContext(ctx).Count()
	if err != nil {
		return 0, fmt.Errorf("统计开发商总数失败: %w", err)
	}
	return count, nil
}
