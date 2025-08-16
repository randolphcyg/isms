package biz

import (
	"context"
	"time"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type DashboardUsecase struct {
	softwareRepo   domain.SoftwareRepo
	developerRepo  domain.DeveloperRepo
	industryRepo   domain.IndustryRepo
	countryRepo    domain.CountryRepo
	log            *log.Helper
}

func NewDashboardUsecase(
	softwareRepo domain.SoftwareRepo,
	developerRepo domain.DeveloperRepo,
	industryRepo domain.IndustryRepo,
	countryRepo domain.CountryRepo,
	logger log.Logger,
) *DashboardUsecase {
	return &DashboardUsecase{
		softwareRepo:  softwareRepo,
		developerRepo: developerRepo,
		industryRepo:  industryRepo,
		countryRepo:   countryRepo,
		log:           log.NewHelper(log.With(logger, "module", "biz/dashboard")),
	}
}

// GetOverviewStats 获取概览统计数据
func (uc *DashboardUsecase) GetOverviewStats(ctx context.Context) (*domain.OverviewStats, error) {
	// 获取软件总数
	softwareCount, err := uc.softwareRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取开发商总数
	developerCount, err := uc.developerRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取行业总数
	industryCount, err := uc.industryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取国家总数
	countryCount, err := uc.countryRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 获取最近30天新增软件数
	newSoftwareCount, err := uc.softwareRepo.CountByTimeRange(ctx, time.Now().AddDate(0, 0, -30), time.Now())
	if err != nil {
		return nil, err
	}

	return &domain.OverviewStats{
		TotalSoftware:    softwareCount,
		TotalDevelopers:  developerCount,
		TotalIndustries:  industryCount,
		TotalCountries:   countryCount,
		NewSoftwareCount: newSoftwareCount,
		LastUpdated:      time.Now(),
	}, nil
}

// GetSoftwareByIndustryStats 获取软件按行业分布统计
func (uc *DashboardUsecase) GetSoftwareByIndustryStats(ctx context.Context, topN int32) ([]*domain.IndustryStatItem, error) {
	items, err := uc.softwareRepo.CountByIndustry(ctx, topN)
	if err != nil {
		return nil, err
	}

	// 获取软件总数用于计算百分比
	total, err := uc.softwareRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	var result []*domain.IndustryStatItem
	for _, item := range items {
		percentage := 0.0
		if total > 0 {
			percentage = float64(item.SoftwareCount) / float64(total) * 100
		}

		result = append(result, &domain.IndustryStatItem{
			IndustryID:   item.IndustryID,
			IndustryName: item.IndustryName,
			SoftwareCount: item.SoftwareCount,
			Percentage:   percentage,
		})
	}

	return result, nil
}

// GetSoftwareByCountryStats 获取软件按国家分布统计
func (uc *DashboardUsecase) GetSoftwareByCountryStats(ctx context.Context, topN int32) ([]*domain.CountryStatItem, error) {
	items, err := uc.softwareRepo.CountByCountry(ctx, topN)
	if err != nil {
		return nil, err
	}

	// 获取软件总数用于计算百分比
	total, err := uc.softwareRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	var result []*domain.CountryStatItem
	for _, item := range items {
		percentage := 0.0
		if total > 0 {
			percentage = float64(item.SoftwareCount) / float64(total) * 100
		}

		result = append(result, &domain.CountryStatItem{
			CountryID:      item.CountryID,
			CountryNameZh:  item.CountryNameZh,
			CountryNameEn:  item.CountryNameEn,
			SoftwareCount:  item.SoftwareCount,
			Percentage:     percentage,
		})
	}

	return result, nil
}

// GetSoftwareByDeveloperStats 获取软件按开发商分布统计
func (uc *DashboardUsecase) GetSoftwareByDeveloperStats(ctx context.Context, topN int32) ([]*domain.DeveloperStatItem, error) {
	items, err := uc.softwareRepo.CountByDeveloper(ctx, topN)
	if err != nil {
		return nil, err
	}

	// 获取软件总数用于计算百分比
	total, err := uc.softwareRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	var result []*domain.DeveloperStatItem
	for _, item := range items {
		percentage := 0.0
		if total > 0 {
			percentage = float64(item.SoftwareCount) / float64(total) * 100
		}

		result = append(result, &domain.DeveloperStatItem{
			DeveloperID:     item.DeveloperID,
			DeveloperNameZh: item.DeveloperNameZh,
			DeveloperNameEn: item.DeveloperNameEn,
			SoftwareCount:   item.SoftwareCount,
			Percentage:      percentage,
		})
	}

	return result, nil
}

// GetSoftwareTrendStats 获取软件按年份发布趋势统计
func (uc *DashboardUsecase) GetSoftwareTrendStats(ctx context.Context, recentYears int32) ([]*domain.TrendStatItem, error) {
	items, err := uc.softwareRepo.CountByYear(ctx, recentYears)
	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	var result []*domain.TrendStatItem
	for _, item := range items {
		result = append(result, &domain.TrendStatItem{
			Year:          item.Year,
			SoftwareCount: item.SoftwareCount,
		})
	}

	return result, nil
}

// GetSoftwareByStatusStats 获取软件按状态分布统计
func (uc *DashboardUsecase) GetSoftwareByStatusStats(ctx context.Context) ([]*domain.StatusStatItem, error) {
	items, err := uc.softwareRepo.CountByStatus(ctx)
	if err != nil {
		return nil, err
	}

	// 获取软件总数用于计算百分比
	total, err := uc.softwareRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	var result []*domain.StatusStatItem
	for _, item := range items {
		percentage := 0.0
		if total > 0 {
			percentage = float64(item.SoftwareCount) / float64(total) * 100
		}

		result = append(result, &domain.StatusStatItem{
			Status:       item.Status,
			StatusLabel:  item.StatusLabel,
			SoftwareCount: item.SoftwareCount,
			Percentage:   percentage,
		})
	}

	return result, nil
}