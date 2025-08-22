package repo

import (
	"context"
	"fmt"
	"strings"
	"time"

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
		DeveloperID:  software.DeveloperID,
		SizeBytes:    software.SizeBytes,
		Status:       software.Status,
		BitWidths:    &bitWidths,
		SourceURL:    software.SourceUrl,
		DownloadLink: software.DownloadLink,
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

	// 构建更新字段映射
	updateFields := map[string]interface{}{
		"name_zh":       software.NameZh,
		"name_en":       software.NameEn,
		"version":       software.Version,
		"release_year":  software.ReleaseYear,
		"release_month": software.ReleaseMonth,
		"release_day":   software.ReleaseDay,
		"description":   software.Description,
		"cpu_req":       software.CPUReq,
		"memory_min_gb": software.MemoryMinGb,
		"disk_min_gb":   software.DiskMinGb,
		"sys_req_other": software.SysReqOther,
		"developer_id":  software.DeveloperID,
		"size_bytes":    software.SizeBytes,
		"status":        software.Status,
		"bit_widths":    &bitWidths,
		"source_url":    software.SourceUrl,
		"download_link": software.DownloadLink,
	}

	// 更新数据
	_, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.ID.Eq(software.ID)).
		Updates(updateFields)
	if err != nil {
		return fmt.Errorf("更新软件失败: %w", err)
	}

	// 更新关联关系（通过计算差异来精确处理）
	// 处理行业关联
	{
		// 查询当前关联的行业ID
		currentIndustryRelations, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
			Where(s.query.IsmsSoftwareIndustry.SoftwareID.Eq(int32(software.ID))).
			Find()
		if err != nil {
			return fmt.Errorf("查询当前行业关联失败: %w", err)
		}

		currentIndustryIDs := make([]int32, 0, len(currentIndustryRelations))
		for _, rel := range currentIndustryRelations {
			currentIndustryIDs = append(currentIndustryIDs, rel.IndustryID)
		}

		// 计算差异
		toAddIndustries, toDeleteIndustries := calculateDifferences(currentIndustryIDs, software.IndustryIDs)

		// 删除需要删除的关联
		if len(toDeleteIndustries) > 0 {
			_, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
				Where(s.query.IsmsSoftwareIndustry.SoftwareID.Eq(int32(software.ID)),
					s.query.IsmsSoftwareIndustry.IndustryID.In(toDeleteIndustries...)).
				Delete()
			if err != nil {
				return fmt.Errorf("删除行业关联失败: %w", err)
			}
		}

		// 添加需要新增的关联
		for _, id := range toAddIndustries {
			err := s.query.IsmsSoftwareIndustry.WithContext(ctx).Create(&model.IsmsSoftwareIndustry{
				SoftwareID: software.ID,
				IndustryID: id,
			})
			if err != nil {
				return fmt.Errorf("关联新行业失败: %w", err)
			}
		}
	}

	// 处理操作系统关联
	{
		// 查询当前关联的操作系统ID
		currentOSRelations, err := s.query.IsmsSoftwareOS.WithContext(ctx).
			Where(s.query.IsmsSoftwareOS.SoftwareID.Eq(software.ID)).
			Find()
		if err != nil {
			return fmt.Errorf("查询当前操作系统关联失败: %w", err)
		}

		currentOSIDs := make([]int32, 0, len(currentOSRelations))
		for _, rel := range currentOSRelations {
			currentOSIDs = append(currentOSIDs, rel.OsID)
		}

		// 计算差异
		toAddOS, toDeleteOS := calculateDifferences(currentOSIDs, software.OsIDs)

		// 删除需要删除的关联
		if len(toDeleteOS) > 0 {
			_, err := s.query.IsmsSoftwareOS.WithContext(ctx).
				Where(s.query.IsmsSoftwareOS.SoftwareID.Eq(software.ID),
					s.query.IsmsSoftwareOS.OsID.In(toDeleteOS...)).
				Delete()
			if err != nil {
				return fmt.Errorf("删除操作系统关联失败: %w", err)
			}
		}

		// 添加需要新增的关联
		for _, id := range toAddOS {
			err := s.query.IsmsSoftwareOS.WithContext(ctx).Create(&model.IsmsSoftwareOS{
				SoftwareID: software.ID,
				OsID:       id,
			})
			if err != nil {
				return fmt.Errorf("关联新操作系统失败: %w", err)
			}
		}
	}

	return nil
}

// calculateDifferences 计算新旧ID列表的差异
func calculateDifferences(oldIDs, newIDs []int32) (toAdd, toDelete []int32) {
	oldSet := make(map[int32]bool)
	newSet := make(map[int32]bool)

	// 构建旧ID集合
	for _, id := range oldIDs {
		oldSet[id] = true
	}

	// 构建新ID集合
	for _, id := range newIDs {
		newSet[id] = true
	}

	// 计算需要新增的ID
	for id := range newSet {
		if !oldSet[id] {
			toAdd = append(toAdd, id)
		}
	}

	// 计算需要删除的ID
	for id := range oldSet {
		if !newSet[id] {
			toDelete = append(toDelete, id)
		}
	}

	return toAdd, toDelete
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

	// 查询关联的开发商
	developer, err := s.query.IsmsDeveloper.WithContext(ctx).
		Where(s.query.IsmsDeveloper.ID.Eq(dataModel.DeveloperID)).
		First()
	if err != nil {
		return nil, fmt.Errorf("查询软件操作系统关联失败: %w", err)
	}

	countryID := int32(0)
	countryID = developer.CountryID

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
		CountryID:    countryID,
		Status:       dataModel.Status,
		SizeBytes:    dataModel.SizeBytes,
		IndustryIDs:  industryIDs,
		OsIDs:        osIDs,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
		BitWidths:    bitWidths,
		SourceUrl:    dataModel.SourceURL,
		DownloadLink: dataModel.DownloadLink,
	}, nil
}

func (s *softwareRepo) List(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	q := s.query.IsmsSoftware.WithContext(ctx)

	// 构建筛选条件
	if opts.DeveloperID > 0 {
		q = q.Where(s.query.IsmsSoftware.DeveloperID.Eq(opts.DeveloperID))
	}
	if opts.Status != "" {
		q = q.Where(s.query.IsmsSoftware.Status.Eq(opts.Status))
	}
	// 添加国家筛选条件
	if opts.CountryID > 0 {
		// 先查询该国家下所有开发商的ID
		developerIDs, err := s.query.IsmsDeveloper.WithContext(ctx).
			Select(s.query.IsmsDeveloper.ID).
			Where(s.query.IsmsDeveloper.CountryID.Eq(opts.CountryID)).
			Find()
		if err != nil {
			return nil, 0, fmt.Errorf("查询国家下开发商失败: %w", err)
		}

		// 提取开发商ID到切片中
		ids := make([]int32, 0, len(developerIDs))
		for _, dev := range developerIDs {
			ids = append(ids, dev.ID)
		}

		// 筛选这些开发商开发的软件
		q = q.Where(s.query.IsmsSoftware.DeveloperID.In(ids...))
	}

	if opts.IndustryID > 0 {
		softwareIDs, err := s.query.IsmsSoftwareIndustry.WithContext(ctx).
			Select(s.query.IsmsSoftwareIndustry.SoftwareID).
			Where(s.query.IsmsSoftwareIndustry.IndustryID.Eq(opts.IndustryID)).
			Find()
		if err != nil {
			return nil, 0, fmt.Errorf("查询软件行业关联失败: %w", err)
		}

		// 提取软件ID到切片中
		ids := make([]int32, 0, len(softwareIDs))
		for _, id := range softwareIDs {
			ids = append(ids, id.SoftwareID)
		}

		// 将子查询结果应用到主查询上
		q = q.Where(s.query.IsmsSoftware.ID.In(ids...))
	}

	if opts.Keyword != "" {
		q = q.Where(s.query.IsmsSoftware.NameEn.Like("%" + opts.Keyword + "%"))
	}

	// 添加按创建时间倒序排序
	q = q.Order(s.query.IsmsSoftware.CreatedAt.Desc())

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

	// 收集开发商ID用于批量查询国家信息
	developerIDs := make([]int32, 0, len(modelSoftwares))
	developerIDMap := make(map[int32]bool) // 用于去重
	for _, m := range modelSoftwares {
		if !developerIDMap[m.DeveloperID] {
			developerIDs = append(developerIDs, m.DeveloperID)
			developerIDMap[m.DeveloperID] = true
		}
	}

	// 批量查询开发商信息
	developerMap := make(map[int32]*model.IsmsDeveloper)
	if len(developerIDs) > 0 {
		developers, err := s.query.IsmsDeveloper.WithContext(ctx).
			Where(s.query.IsmsDeveloper.ID.In(developerIDs...)).
			Find()
		if err != nil {
			return nil, 0, fmt.Errorf("查询开发商信息失败: %w", err)
		}

		for _, dev := range developers {
			developerMap[dev.ID] = dev
		}
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

		countryID := int32(0)
		if dev, ok := developerMap[m.DeveloperID]; ok {
			countryID = dev.CountryID
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
			CountryID:    countryID,
			Status:       m.Status,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
			BitWidths:    bitWidths,
			SourceUrl:    m.SourceURL,
			DownloadLink: m.DownloadLink,
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
		Where(s.query.IsmsSoftware.NameEn.Eq(name)).
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

// Count 统计软件总数
func (s *softwareRepo) Count(ctx context.Context) (int64, error) {
	count, err := s.query.IsmsSoftware.WithContext(ctx).Count()
	if err != nil {
		return 0, fmt.Errorf("统计软件总数失败: %w", err)
	}
	return count, nil
}

// CountByIndustry 统计各行业软件数量
func (s *softwareRepo) CountByIndustry(ctx context.Context, topN int32) ([]*domain.IndustryStatResult, error) {
	var results []struct {
		IndustryID   int32  `gorm:"column:industry_id"`
		IndustryName string `gorm:"column:subcategory_name"`
		Count        int64  `gorm:"column:count"`
	}

	// 执行连接查询统计各行业软件数量
	err := s.db.WithContext(ctx).
		Table("isms_software AS s"). // 给表设置别名
		Select("isi.industry_id, i.subcategory_name, COUNT(s.id) as count").
		Joins("JOIN isms_software_industry isi ON s.id = isi.software_id").
		Joins("JOIN isms_industry i ON isi.industry_id = i.id").
		Group("isi.industry_id, i.subcategory_name").
		Order("count DESC").
		Limit(int(topN)).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("统计各行业软件数量失败: %w", err)
	}

	// 构建结果
	var industryStats []*domain.IndustryStatResult
	for _, result := range results {
		industryStats = append(industryStats, &domain.IndustryStatResult{
			IndustryID:    result.IndustryID,
			IndustryName:  result.IndustryName,
			SoftwareCount: result.Count,
		})
	}

	return industryStats, nil
}

// CountByStatus 统计各状态软件数量
func (s *softwareRepo) CountByStatus(ctx context.Context) ([]*domain.StatusStatResult, error) {
	var results []struct {
		Status string `gorm:"column:status"`
		Count  int64  `gorm:"column:count"`
	}

	// 按状态分组统计软件数量
	err := s.db.WithContext(ctx).
		Table("isms_software AS s"). // 给表设置别名
		Select("status, COUNT(s.id) as count").
		Group("status").
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("统计各状态软件数量失败: %w", err)
	}

	// 构建结果
	var statusStats []*domain.StatusStatResult
	for _, result := range results {
		// 根据状态值设置标签
		var label string
		switch result.Status {
		case "active":
			label = "有效"
		case "inactive":
			label = "下架"
		case "testing":
			label = "测试中"
		case "discontinued":
			label = "停止维护"
		default:
			label = result.Status
		}

		statusStats = append(statusStats, &domain.StatusStatResult{
			Status:        result.Status,
			StatusLabel:   label,
			SoftwareCount: result.Count,
		})
	}

	return statusStats, nil
}

// CountByTimeRange 统计指定时间范围内的软件数量
func (s *softwareRepo) CountByTimeRange(ctx context.Context, start, end time.Time) (int64, error) {
	var count int64
	count, err := s.query.IsmsSoftware.WithContext(ctx).
		Where(s.query.IsmsSoftware.CreatedAt.Between(start, end)).
		Count()
	if err != nil {
		return 0, fmt.Errorf("统计指定时间范围内的软件数量失败: %w", err)
	}
	return count, nil
}

// CountByCountry 统计各国软件数量
func (s *softwareRepo) CountByCountry(ctx context.Context, topN int32) ([]*domain.CountryStatResult, error) {
	var results []struct {
		CountryID     int32  `gorm:"column:country_id"`
		CountryNameZh string `gorm:"column:name_zh"`
		CountryNameEn string `gorm:"column:name_en"`
		Count         int64  `gorm:"column:count"`
	}

	// 执行连接查询统计各国软件数量
	err := s.db.WithContext(ctx).
		Table("isms_software AS s"). // 给表设置别名
		Select("c.id as country_id, c.name_zh, c.name_en, COUNT(s.id) as count").
		Joins("JOIN isms_developer d ON s.developer_id = d.id"). // 先关联开发商表
		Joins("JOIN isms_country c ON d.country_id = c.id").     // 再通过开发商关联国家表
		Group("c.id, c.name_zh, c.name_en").
		Order("count DESC").
		Limit(int(topN)).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("统计各国软件数量失败: %w", err)
	}

	// 构建结果
	var countryStats []*domain.CountryStatResult
	for _, result := range results {
		countryStats = append(countryStats, &domain.CountryStatResult{
			CountryID:     result.CountryID,
			CountryNameZh: result.CountryNameZh,
			CountryNameEn: result.CountryNameEn,
			SoftwareCount: result.Count,
		})
	}

	return countryStats, nil
}

// CountByDeveloper 统计各开发商软件数量
func (s *softwareRepo) CountByDeveloper(ctx context.Context, topN int32) ([]*domain.DeveloperStatResult, error) {
	var results []struct {
		DeveloperID     int32  `gorm:"column:developer_id"`
		DeveloperNameZh string `gorm:"column:name_zh"`
		DeveloperNameEn string `gorm:"column:name_en"`
		Count           int64  `gorm:"column:count"`
	}

	// 执行连接查询统计各开发商软件数量
	err := s.db.WithContext(ctx).
		Table("isms_software AS s"). // 给表设置别名
		Select("d.id as developer_id, d.name_zh, d.name_en, COUNT(s.id) as count").
		Joins("JOIN isms_developer d ON s.developer_id = d.id").
		Group("d.id, d.name_zh, d.name_en").
		Order("count DESC").
		Limit(int(topN)).
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("统计各开发商软件数量失败: %w", err)
	}

	// 构建结果
	var developerStats []*domain.DeveloperStatResult
	for _, result := range results {
		developerStats = append(developerStats, &domain.DeveloperStatResult{
			DeveloperID:     result.DeveloperID,
			DeveloperNameZh: result.DeveloperNameZh,
			DeveloperNameEn: result.DeveloperNameEn,
			SoftwareCount:   result.Count,
		})
	}

	return developerStats, nil
}

// CountByYear 统计各年份软件数量
func (s *softwareRepo) CountByYear(ctx context.Context, recentYears int32) ([]*domain.TrendStatResult, error) {
	var results []struct {
		Year  int32 `gorm:"column:release_year"`
		Count int64 `gorm:"column:count"`
	}

	// 计算起始年份
	currentYear := int32(time.Now().Year())
	startYear := currentYear - recentYears + 1

	// 按发布年份分组统计软件数量
	err := s.db.WithContext(ctx).
		Table("isms_software AS s"). // 给表设置别名
		Select("release_year, COUNT(s.id) as count").
		Where("release_year IS NOT NULL AND release_year > 0 AND release_year >= ?", startYear).
		Group("release_year").
		Order("release_year").
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("统计各年份软件数量失败: %w", err)
	}

	// 构建结果
	var trendStats []*domain.TrendStatResult
	for _, result := range results {
		trendStats = append(trendStats, &domain.TrendStatResult{
			Year:          result.Year,
			SoftwareCount: result.Count,
		})
	}

	return trendStats, nil
}
