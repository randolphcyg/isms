package service

import (
	"context"

	pb "isms/api/isms/v1"
	"isms/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer
	uc  *biz.DashboardUsecase
	log *log.Helper
}

func NewDashboardService(uc *biz.DashboardUsecase, logger log.Logger) *DashboardService {
	return &DashboardService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/dashboard")),
	}
}

func (s *DashboardService) GetOverviewStats(ctx context.Context, req *pb.GetOverviewStatsReq) (*pb.GetOverviewStatsResp, error) {
	result, err := s.uc.GetOverviewStats(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetOverviewStatsResp{
		TotalSoftware:    result.TotalSoftware,
		TotalDevelopers:  result.TotalDevelopers,
		TotalIndustries:  result.TotalIndustries,
		TotalCountries:   result.TotalCountries,
		NewSoftwareCount: result.NewSoftwareCount,
		LastUpdated:      result.LastUpdated.Unix(),
	}, nil
}

func (s *DashboardService) GetSoftwareByIndustryStats(ctx context.Context, req *pb.GetSoftwareByIndustryStatsReq) (*pb.GetSoftwareByIndustryStatsResp, error) {
	topN := req.TopN
	if topN <= 0 {
		topN = 10 // 默认值
	}
	result, err := s.uc.GetSoftwareByIndustryStats(ctx, topN)
	if err != nil {
		return nil, err
	}

	var items []*pb.IndustryStatItem
	for _, stat := range result {
		items = append(items, &pb.IndustryStatItem{
			IndustryId:    stat.IndustryID,
			IndustryName:  stat.IndustryName,
			SoftwareCount: stat.SoftwareCount,
			Percentage:    stat.Percentage,
		})
	}

	return &pb.GetSoftwareByIndustryStatsResp{
		Items: items,
	}, nil
}

func (s *DashboardService) GetSoftwareByCountryStats(ctx context.Context, req *pb.GetSoftwareByCountryStatsReq) (*pb.GetSoftwareByCountryStatsResp, error) {
	topN := req.TopN
	if topN <= 0 {
		topN = 10 // 默认值
	}
	result, err := s.uc.GetSoftwareByCountryStats(ctx, topN)
	if err != nil {
		return nil, err
	}

	var items []*pb.CountryStatItem
	for _, stat := range result {
		items = append(items, &pb.CountryStatItem{
			CountryId:     stat.CountryID,
			CountryNameZh: stat.CountryNameZh,
			CountryNameEn: stat.CountryNameEn,
			SoftwareCount: stat.SoftwareCount,
			Percentage:    stat.Percentage,
		})
	}

	return &pb.GetSoftwareByCountryStatsResp{
		Items: items,
	}, nil
}

func (s *DashboardService) GetSoftwareByDeveloperStats(ctx context.Context, req *pb.GetSoftwareByDeveloperStatsReq) (*pb.GetSoftwareByDeveloperStatsResp, error) {
	topN := req.TopN
	if topN <= 0 {
		topN = 10 // 默认值
	}
	result, err := s.uc.GetSoftwareByDeveloperStats(ctx, topN)
	if err != nil {
		return nil, err
	}

	var items []*pb.DeveloperStatItem
	for _, stat := range result {
		items = append(items, &pb.DeveloperStatItem{
			DeveloperId:     stat.DeveloperID,
			DeveloperNameZh: stat.DeveloperNameZh,
			DeveloperNameEn: stat.DeveloperNameEn,
			SoftwareCount:   stat.SoftwareCount,
			Percentage:      stat.Percentage,
		})
	}

	return &pb.GetSoftwareByDeveloperStatsResp{
		Items: items,
	}, nil
}

func (s *DashboardService) GetSoftwareTrendStats(ctx context.Context, req *pb.GetSoftwareTrendStatsReq) (*pb.GetSoftwareTrendStatsResp, error) {
	recentYears := req.RecentYears
	if recentYears <= 0 {
		recentYears = 5 // 默认值
	}
	result, err := s.uc.GetSoftwareTrendStats(ctx, recentYears)
	if err != nil {
		return nil, err
	}

	var items []*pb.TrendStatItem
	for _, stat := range result {
		items = append(items, &pb.TrendStatItem{
			Year:          stat.Year,
			SoftwareCount: stat.SoftwareCount,
		})
	}

	return &pb.GetSoftwareTrendStatsResp{
		Items: items,
	}, nil
}

func (s *DashboardService) GetSoftwareByStatusStats(ctx context.Context, req *pb.GetSoftwareByStatusStatsReq) (*pb.GetSoftwareByStatusStatsResp, error) {
	result, err := s.uc.GetSoftwareByStatusStats(ctx)
	if err != nil {
		return nil, err
	}

	var items []*pb.StatusStatItem
	for _, stat := range result {
		items = append(items, &pb.StatusStatItem{
			Status:        stat.Status,
			StatusLabel:   stat.StatusLabel,
			SoftwareCount: stat.SoftwareCount,
			Percentage:    stat.Percentage,
		})
	}

	return &pb.GetSoftwareByStatusStatsResp{
		Items: items,
	}, nil
}
