package repo

import (
	"context"
	"fmt"
	"strings"

	v1 "isms/api/isms/v1"
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

func (s *softwareRepo) Create(ctx context.Context, software *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {
	var bitWidths string
	if len(software.BitWidths) > 0 {
		bitWidths = strings.Join(software.BitWidths, ",")
	}
	// 领域模型转数据模型
	dataModel := &model.IsmsSoftware{
		NameZh:       software.NameZh,
		NameEn:       software.NameEn,
		Version:      software.Version,
		ReleaseYear:  software.ReleaseYear,
		ReleaseMonth: software.ReleaseMonth,
		ReleaseDay:   software.ReleaseDay,
		CPUReq:       software.CPUReq,
		MemoryMinGb:  software.MemoryMinGb,
		DiskMinGb:    software.DiskMinGb,
		SysReqOther:  software.SysReqOther,
		Description:  software.Description,
		CountryID:    software.CountryID,
		DeveloperID:  software.DeveloperID,
		SizeBytes:    software.SizeBytes,
		Status:       software.Status,
		BitWidths:    &bitWidths,
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
				SoftwareID: software.ID,
				IndustryID: id,
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
				OsID:       id,
			})
			if err != nil {
				return software, fmt.Errorf("关联操作系统失败: %w", err)
			}
		}
	}

	return software, nil
}

func (s *softwareRepo) Update(ctx context.Context, software *domain.IsmsSoftware) error {
	var bitWidths string
	if len(software.BitWidths) > 0 {
		bitWidths = strings.Join(software.BitWidths, ",")
	}
	// 领域模型转数据模型
	dataModel := &model.IsmsSoftware{
		ID:           software.ID,
		NameZh:       software.NameZh,
		NameEn:       software.NameEn,
		Version:      software.Version,
		ReleaseYear:  software.ReleaseYear,
		ReleaseMonth: software.ReleaseMonth,
		ReleaseDay:   software.ReleaseDay,
		Description:  software.Description,
		CPUReq:       software.CPUReq,
		MemoryMinGb:  software.MemoryMinGb,
		DiskMinGb:    software.DiskMinGb,
		SysReqOther:  software.SysReqOther,
		CountryID:    software.CountryID,
		DeveloperID:  software.DeveloperID,
		SizeBytes:    software.SizeBytes,
		Status:       software.Status,
		BitWidths:    &bitWidths,
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

	var bitWidths []string
	if dataModel.BitWidths != nil {
		bitWidths = strings.Split(*dataModel.BitWidths, ",")
	}

	// 转换数据模型到领域模型
	return &domain.IsmsSoftware{
		ID:           dataModel.ID,
		NameZh:       dataModel.NameZh,
		NameEn:       dataModel.NameEn,
		DeveloperID:  dataModel.DeveloperID,
		Version:      dataModel.Version,
		ReleaseYear:  dataModel.ReleaseYear,
		ReleaseMonth: dataModel.ReleaseMonth,
		ReleaseDay:   dataModel.ReleaseDay,
		Description:  dataModel.Description,
		CPUReq:       dataModel.CPUReq,
		MemoryMinGb:  dataModel.MemoryMinGb,
		DiskMinGb:    dataModel.DiskMinGb,
		SysReqOther:  dataModel.SysReqOther,
		CountryID:    dataModel.CountryID,
		Status:       dataModel.Status,
		SizeBytes:    dataModel.SizeBytes,
		IndustryIDs:  industryIDs,
		OsIDs:        osIDs,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
		BitWidths:    bitWidths,
	}, nil
}

func (s *softwareRepo) List(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	// 基础查询（只查询软件表本身，不关联其他表）
	q := s.query.IsmsSoftware.WithContext(ctx)

	// 构建筛选条件（不关联行业表）
	if opts.DeveloperID > 0 {
		q = q.Where(s.query.IsmsSoftware.DeveloperID.Eq(opts.DeveloperID))
	}
	if opts.Status != "" {
		q = q.Where(s.query.IsmsSoftware.Status.Eq(opts.Status))
	}
	if opts.Keyword != "" {
		q = q.Where(s.query.IsmsSoftware.NameZh.Like("%"+opts.Keyword+"%"),
			s.query.IsmsSoftware.NameEn.Like("%"+opts.Keyword+"%"))
	}

	// 查询总条数（不关联其他表）
	total, err := q.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("查询软件总数失败: %w", err)
	}

	// 分页查询（只查询软件表本身）
	modelSoftwares, err := q.Offset(int((opts.Page - 1) * opts.PageSize)).
		Limit(int(opts.PageSize)).
		Find()
	if err != nil {
		return nil, 0, fmt.Errorf("查询软件列表失败: %w", err)
	}

	// 收集软件ID用于批量查询关联关系
	softwareIDs := make([]int32, 0, len(modelSoftwares))
	for _, m := range modelSoftwares {
		softwareIDs = append(softwareIDs, m.ID)
	}

	// 批量查询行业关联
	industryRelations := make(map[int32][]int32)          // softwareID -> industryIDs
	industryDetails := make(map[int32][]*v1.IsmsIndustry) // softwareID -> industryDetails
	if len(softwareIDs) > 0 {
		relations, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
			Where(s.query.IsmsSoftwareIndustry.SoftwareID.In(softwareIDs...)).
			Find()
		if err != nil {
			return nil, 0, fmt.Errorf("查询软件行业关联失败: %w", err)
		}

		// 用于去重industryID
		industryIDSet := make(map[int32]map[int32]bool)

		for _, rel := range relations {
			// 初始化map
			if industryIDSet[rel.SoftwareID] == nil {
				industryIDSet[rel.SoftwareID] = make(map[int32]bool)
			}

			// 如果该industryID还没有添加到该softwareID的列表中
			if !industryIDSet[rel.SoftwareID][rel.IndustryID] {
				industryRelations[rel.SoftwareID] = append(industryRelations[rel.SoftwareID], rel.IndustryID)
				industryIDSet[rel.SoftwareID][rel.IndustryID] = true

				industryDetails[rel.SoftwareID] = append(industryDetails[rel.SoftwareID], &v1.IsmsIndustry{
					Id: rel.IndustryID,
				})
			}
		}
	}

	var domainSoftwares []*domain.IsmsSoftware
	for _, m := range modelSoftwares {
		var bitWidths []string
		if m.BitWidths != nil {
			bitWidths = strings.Split(*m.BitWidths, ",")
		}

		sw := &domain.IsmsSoftware{
			ID:           m.ID,
			NameZh:       m.NameZh,
			NameEn:       m.NameEn,
			DeveloperID:  m.DeveloperID,
			Version:      m.Version,
			ReleaseYear:  m.ReleaseYear,
			ReleaseMonth: m.ReleaseMonth,
			ReleaseDay:   m.ReleaseDay,
			Description:  m.Description,
			CPUReq:       m.CPUReq,
			MemoryMinGb:  m.MemoryMinGb,
			DiskMinGb:    m.DiskMinGb,
			SysReqOther:  m.SysReqOther,
			SizeBytes:    m.SizeBytes,
			CountryID:    m.CountryID,
			Status:       m.Status,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
			BitWidths:    bitWidths,
		}

		// 设置行业ID
		if ids, ok := industryRelations[m.ID]; ok {
			sw.IndustryIDs = ids
		}

		// 设置行业详情
		if details, ok := industryDetails[m.ID]; ok {
			sw.IndustryDetails = details
		}

		domainSoftwares = append(domainSoftwares, sw)
	}

	return domainSoftwares, total, nil
}

// ExistByNameAndVersion 检查具有指定名称和版本的软件是否存在
func (s *softwareRepo) ExistByNameAndVersion(ctx context.Context, name, version string) (bool, error) {
	var count int64
	count, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.NameZh.Eq(name)).
		Where(s.query.IsmsSoftware.Version.Eq(version)).
		Count()
	if err != nil {
		return false, fmt.Errorf("检查软件名称和版本是否存在失败: %w", err)
	}
	return count > 0, nil
}

// ExistByID 检查具有指定ID的软件是否存在
func (s *softwareRepo) ExistByID(ctx context.Context, id uint32) (bool, error) {
	var count int64
	count, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.ID.Eq(int32(id))).
		Count()
	if err != nil {
		return false, fmt.Errorf("检查软件ID是否存在失败: %w", err)
	}
	return count > 0, nil
}
