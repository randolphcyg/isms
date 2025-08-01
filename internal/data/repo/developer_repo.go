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

// Create 实现领域层的Create方法（转换领域模型→数据模型）
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

// ExistByName （检查名称是否存在）
func (r *developerRepo) ExistByName(ctx context.Context, nameZh string) (bool, error) {
	count, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.NameZh.Eq(nameZh)).
		Count()
	if err != nil {
		return false, fmt.Errorf("查询开发商名称失败: %w", err)
	}
	return count > 0, nil
}

// 修改 GetByID 方法的参数类型
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

// 修改 List 方法的参数类型
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
	for _, m := range modelDevs {
		domainDevs = append(domainDevs, &domain.Developer{
			ID:          m.ID,
			NameZh:      m.NameZh,
			NameEn:      m.NameEn,
			CountryID:   m.CountryID,
			Website:     m.Website,
			Description: m.Description,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		})
	}

	return domainDevs, total, nil
}
