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
		Website:     &dev.Website,
		Description: &dev.Description,
	}

	// 调用生成的query插入数据
	err := r.query.IsmsDeveloper.WithContext(ctx).Create(dataModel)
	if err != nil {
		return nil, err
	}

	// 数据模型→领域模型（返回给上层）
	dev.ID = uint32(dataModel.ID)
	dev.CreatedAt = dataModel.CreatedAt
	dev.UpdatedAt = dataModel.UpdatedAt
	return dev, nil
}

// GetByID 根据ID查询开发商详情
func (r *developerRepo) GetByID(ctx context.Context, id uint32) (*domain.Developer, error) {
	// 1. 查询数据模型（ORM实体）
	dataModel, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.ID.Eq(int32(id))). // 注意：若数据库ID是int64，需转换
		First()                                        // 查询单条记录

	// 2. 处理查询错误（如记录不存在）
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("开发商ID不存在: %d", id)
		}
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	// 3. 数据模型 → 领域模型（转换字段）
	return &domain.Developer{
		ID:        uint32(dataModel.ID),
		NameZh:    dataModel.NameZh,
		NameEn:    dataModel.NameEn,
		CountryID: dataModel.CountryID,
		// 若需要国家名称，需关联查询国家表（此处仅基础转换）
		Website:     *dataModel.Website,     // 解引用指针（确保非空，或加判断）
		Description: *dataModel.Description, // 解引用指针
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}, nil
}

// List 分页查询开发商（优化版，参考industry领域实现风格）
func (r *developerRepo) List(
	ctx context.Context,
	page, pageSize uint32,
	countryID uint32,
	keyword string,
) ([]*domain.Developer, int64, error) {
	// 1. 基础查询：指定表名，绑定上下文（与industryRepo风格一致）
	db := r.db.WithContext(ctx).Table("isms_developer") // 显式指定表名，避免模型与表名映射问题

	// 2. 构建筛选条件（参考gorm原生语法，清晰直观）
	if countryID > 0 {
		db = db.Where("country_id = ?", countryID) // 国家ID筛选
	}
	if keyword != "" {
		// 关键词模糊匹配（中文/英文名称，OR条件）
		db = db.Where("name_zh LIKE ?", "%"+keyword+"%").
			Or("name_en LIKE ?", "%"+keyword+"%")
	}

	// 3. 查询总数（单独计数，与列表查询分离，更高效）
	var total int64
	if err := db.Count(&total).Error; err != nil {
		r.log.Errorf("查询开发商总数失败: %v", err)
		return nil, 0, fmt.Errorf("查询总数失败: %w", err)
	}

	// 4. 分页查询列表（计算偏移量，限制返回条数）
	offset := int((page - 1) * pageSize)
	limit := int(pageSize)
	var dataModels []*model.IsmsDeveloper
	if err := db.
		Offset(offset).
		Limit(limit).
		Order("id DESC"). // 按ID倒序（最新的在前，可按需调整）
		Find(&dataModels).Error; err != nil {
		r.log.Errorf("查询开发商列表失败: %v", err)
		return nil, 0, fmt.Errorf("查询列表失败: %w", err)
	}

	// 5. 模型转换（model → domain，严格处理指针字段）
	domainList := make([]*domain.Developer, 0, len(dataModels))
	for _, dm := range dataModels {
		// 安全处理指针字段（避免nil解引用）
		website := ""
		if dm.Website != nil {
			website = *dm.Website
		}
		description := ""
		if dm.Description != nil {
			description = *dm.Description
		}

		domainList = append(domainList, &domain.Developer{
			ID:          uint32(dm.ID),
			NameZh:      dm.NameZh,
			NameEn:      dm.NameEn,
			CountryID:   dm.CountryID,
			Website:     website,
			Description: description,
			CreatedAt:   dm.CreatedAt,
			UpdatedAt:   dm.UpdatedAt,
		})
	}

	return domainList, total, nil
}

// ExistByName （检查名称是否存在）
func (r *developerRepo) ExistByName(ctx context.Context, nameZh string) (bool, error) {
	count, err := r.query.IsmsDeveloper.WithContext(ctx).
		Where(r.query.IsmsDeveloper.NameZh.Eq(nameZh)).
		Count() // 查询符合条件的记录数
	if err != nil {
		return false, fmt.Errorf("查询名称是否存在失败: %w", err)
	}
	return count > 0, nil
}
