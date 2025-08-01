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

type softwareRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

func NewSoftwareRepo(db *gorm.DB, logger log.Logger) domain.SoftwareRepo {
	return &softwareRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/software_repo")),
	}
}

// 实现Create方法
func (s *softwareRepo) Create(ctx context.Context, software *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {
	// 领域模型转数据模型
	dataModel := &model.IsmsSoftware{
		NameZh:      software.NameZh,
		NameEn:      software.NameEn,
		DeveloperID: software.DeveloperID,
		Version:     software.Version,
		Description: software.Description,
		CountryID:   software.CountryID,
		Status:      software.Status,
	}

	// 插入数据
	err := s.query.IsmsSoftware.WithContext(ctx).Create(dataModel)
	if err != nil {
		return nil, fmt.Errorf("创建软件失败: %w", err)
	}

	// 更新领域模型的ID和时间
	software.ID = dataModel.ID
	software.CreatedAt = dataModel.CreatedAt
	software.UpdatedAt = dataModel.UpdatedAt

	// 处理关联关系（行业和操作系统）
	if len(software.IndustryIDs) > 0 {
		for _, id := range software.IndustryIDs {
			err := s.query.IsmsSoftwareIndustry.WithContext(ctx).Create(&model.IsmsSoftwareIndustry{
				SoftwareID: int32(software.ID),
				IndustryID: int32(id),
			})
			if err != nil {
				return software, fmt.Errorf("关联行业失败: %w", err)
			}
		}
	}

	if len(software.OsIDs) > 0 {
		for _, id := range software.OsIDs {
			err := s.query.IsmsSoftwareOS.WithContext(ctx).Create(&model.IsmsSoftwareOS{
				SoftwareID: software.ID,
				OsID:       int32(id),
			})
			if err != nil {
				return software, fmt.Errorf("关联操作系统失败: %w", err)
			}
		}
	}

	return software, nil
}

// 实现Update方法
func (s *softwareRepo) Update(ctx context.Context, software *domain.IsmsSoftware) error {
	// 领域模型转数据模型
	dataModel := &model.IsmsSoftware{
		ID:          software.ID,
		NameZh:      software.NameZh,
		NameEn:      software.NameEn,
		DeveloperID: software.DeveloperID,
		Version:     software.Version,
		Description: software.Description,
		CountryID:   software.CountryID,
		Status:      software.Status,
	}

	// 更新数据
	_, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.ID.Eq(dataModel.ID)).
		Updates(dataModel)
	if err != nil {
		return fmt.Errorf("更新软件失败: %w", err)
	}

	// 更新关联关系（先删除旧关系，再添加新关系）
	if len(software.IndustryIDs) > 0 {
		// 删除旧关系
		_, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
			Where(s.query.IsmsSoftwareIndustry.SoftwareID.Eq(int32(software.ID))).
			Delete()
		if err != nil {
			return fmt.Errorf("删除旧行业关联失败: %w", err)
		}

		// 添加新关系
		for _, id := range software.IndustryIDs {
			err := s.query.IsmsSoftwareIndustry.WithContext(ctx).Create(&model.IsmsSoftwareIndustry{
				SoftwareID: int32(software.ID),
				IndustryID: int32(id),
			})
			if err != nil {
				return fmt.Errorf("关联新行业失败: %w", err)
			}
		}
	}

	if len(software.OsIDs) > 0 {
		// 删除旧关系
		_, err := s.query.IsmsSoftwareOS.WithContext(ctx).
			Where(s.query.IsmsSoftwareOS.SoftwareID.Eq(int32(software.ID))).
			Delete()
		if err != nil {
			return fmt.Errorf("删除旧操作系统关联失败: %w", err)
		}

		// 添加新关系
		for _, id := range software.OsIDs {
			err := s.query.IsmsSoftwareOS.WithContext(ctx).Create(&model.IsmsSoftwareOS{
				SoftwareID: software.ID,
				OsID:       int32(id),
			})
			if err != nil {
				return fmt.Errorf("关联新操作系统失败: %w", err)
			}
		}
	}

	return nil
}

func (s *softwareRepo) FindByID(ctx context.Context, id uint32) (*domain.IsmsSoftware, error) {
	dataModel, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.ID.Eq(int32(id))).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("软件ID不存在: %d", id)
		}
		return nil, fmt.Errorf("查询软件失败: %w", err)
	}

	// 查询关联的行业
	industryRelations, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
		Where(s.query.IsmsSoftwareIndustry.SoftwareID.Eq(int32(id))).
		Find()
	if err != nil {
		return nil, fmt.Errorf("查询软件行业关联失败: %w", err)
	}

	industryIDs := make([]int32, 0, len(industryRelations))
	for _, rel := range industryRelations {
		industryIDs = append(industryIDs, rel.IndustryID)
	}

	// 查询关联的操作系统
	osRelations, err := s.query.IsmsSoftwareOS.WithContext(ctx).
		Where(s.query.IsmsSoftwareOS.SoftwareID.Eq(int32(id))).
		Find()
	if err != nil {
		return nil, fmt.Errorf("查询软件操作系统关联失败: %w", err)
	}

	// 将osIDs类型从[]int64改为[]int32
	osIDs := make([]int32, 0, len(osRelations))
	for _, rel := range osRelations {
		osIDs = append(osIDs, rel.OsID)
	}

	// 转换数据模型到领域模型
	return &domain.IsmsSoftware{
		ID:          dataModel.ID,
		NameZh:      dataModel.NameZh,
		NameEn:      dataModel.NameEn,
		DeveloperID: dataModel.DeveloperID,
		Version:     dataModel.Version,
		Description: dataModel.Description,
		CountryID:   dataModel.CountryID,
		Status:      dataModel.Status,
		IndustryIDs: industryIDs,
		OsIDs:       osIDs,
		CreatedAt:   dataModel.CreatedAt,
		UpdatedAt:   dataModel.UpdatedAt,
	}, nil
}

func (s *softwareRepo) List(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	// 基础查询
	db := s.db.WithContext(ctx).Table("isms_software")

	// 构建筛选条件
	if opts.DeveloperID > 0 {
		db = db.Where("developer_id = ?", opts.DeveloperID)
	}
	if opts.CategoryCode != "" {
		// 这里需要关联查询行业分类
		db = db.Joins("JOIN isms_software_industry ON isms_software.id = isms_software_industry.software_id")
		db = db.Joins("JOIN isms_industry ON isms_software_industry.industry_id = isms_industry.id")
		db = db.Where("isms_industry.category_code = ?", opts.CategoryCode)
	}
	if opts.Status > 0 {
		db = db.Where("status = ?", opts.Status)
	}
	if opts.Keyword != "" {
		db = db.Where("name_zh LIKE ?", "%"+opts.Keyword+"%").
			Or("name_en LIKE ?", "%"+opts.Keyword+"%")
	}

	// 查询总条数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询软件总数失败: %w", err)
	}

	// 分页查询
	offset := (opts.Page - 1) * opts.PageSize
	var modelSoftwares []*model.IsmsSoftware
	if err := db.Offset(int(offset)).Limit(int(opts.PageSize)).Find(&modelSoftwares).Error; err != nil {
		return nil, 0, fmt.Errorf("查询软件列表失败: %w", err)
	}

	// 转换数据模型到领域模型
	domainSoftwares := make([]*domain.IsmsSoftware, 0, len(modelSoftwares))
	for _, m := range modelSoftwares {
		domainSoftwares = append(domainSoftwares, &domain.IsmsSoftware{
			ID:          m.ID,
			NameZh:      m.NameZh,
			NameEn:      m.NameEn,
			DeveloperID: m.DeveloperID,
			Version:     m.Version,
			Description: m.Description,
			CountryID:   m.CountryID,
			Status:      m.Status,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		})
	}

	return domainSoftwares, total, nil
}

func (s *softwareRepo) ExistByNameAndVersion(ctx context.Context, name, version string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *softwareRepo) ExistByID(ctx context.Context, id uint32) (bool, error) {
	//TODO implement me
	panic("implement me")
}
