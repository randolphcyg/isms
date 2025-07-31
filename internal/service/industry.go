package service

import (
	"context"

	pb "isms/api/isms/v1"
	"isms/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type IndustryService struct {
	pb.UnimplementedIndustryServer

	uc  *biz.IndustryUsecase // 依赖 biz 层的业务用例
	log *log.Helper          // 日志工具
}

func NewIndustryService(uc *biz.IndustryUsecase, logger log.Logger) *IndustryService {
	return &IndustryService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/industry")),
	}
}

func (s *IndustryService) GetSubcategories(ctx context.Context, req *pb.GetSubcategoriesReq) (*pb.GetSubcategoriesResp, error) {
	s.log.WithContext(ctx).Infof("接收查询大类[%s]小类的请求", req.CategoryCode)

	// 1. 调用 biz 层获取领域模型
	domainClassifications, err := s.uc.GetSubcategories(ctx, req.CategoryCode)
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询大类[%s]小类失败: %v", req.CategoryCode, err)
		return nil, err
	}

	// 2. 转换领域模型到 API 模型
	apiClassifications := make([]*pb.IsmsIndustry, 0, len(domainClassifications))
	for _, dc := range domainClassifications {
		apiClassifications = append(apiClassifications, &pb.IsmsIndustry{
			Id:              dc.ID,
			CategoryCode:    dc.CategoryCode,
			CategoryName:    dc.CategoryName,
			SubcategoryCode: dc.SubcategoryCode,
			SubcategoryName: dc.SubcategoryName,
		})
	}

	s.log.WithContext(ctx).Infof("查询大类[%s]小类成功，共 %d 条数据", req.CategoryCode, len(apiClassifications))

	return &pb.GetSubcategoriesResp{Subcategories: apiClassifications}, nil
}
func (s *IndustryService) ListCategories(ctx context.Context, req *pb.ListCategoriesReq) (*pb.ListCategoriesResp, error) {
	s.log.WithContext(ctx).Info("接收查询所有行业大类请求")

	// 1. 调用 biz 层获取领域模型（修正：移除硬编码，实际调用业务逻辑）
	domainCategories, err := s.uc.ListAllCategories(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询行业大类失败: %v", err)
		return nil, err // 错误会自动转换为 API 错误码（需配合 errorx 或中间件）
	}

	// 2. 转换领域模型到 API 模型（核心：domain → v1）
	categories := make([]*pb.IndustryCategory, 0, len(domainCategories))
	for _, dc := range domainCategories {
		categories = append(categories, &pb.IndustryCategory{
			CategoryCode: dc.CategoryCode, // 大类编码（如"B"）
			CategoryName: dc.CategoryName, // 大类名称（如"采矿业"）
		})
	}

	s.log.WithContext(ctx).Infof("查询行业大类成功，共 %d 条数据", len(categories))

	return &pb.ListCategoriesResp{Categories: categories}, nil
}
