package service

import (
	"context"
	"strings"
	"time"

	pb "isms/api/isms/v1"
	"isms/internal/biz"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeveloperService struct {
	pb.UnimplementedDeveloperServer

	uc  *biz.DeveloperUsecase // 依赖biz层的用例接口
	log *log.Helper           // 日志工具
}

func NewDeveloperService(uc *biz.DeveloperUsecase, logger log.Logger) *DeveloperService {
	return &DeveloperService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/developer")),
	}
}

func (s *DeveloperService) CreateDeveloper(ctx context.Context, req *pb.CreateDeveloperReq) (*pb.DeveloperResp, error) {
	// 1. 日志记录详细请求参数
	s.log.WithContext(ctx).Infof(
		"收到创建开发商请求: 中文名称=%s, 英文名称=%s, 国家ID=%d",
		req.NameZh, req.NameEn, req.CountryId,
	)

	// 2. 参数基础校验（补充proto validate未覆盖的业务校验）
	if req.CountryId == 0 {
		s.log.WithContext(ctx).Warn("国家ID不能为空")
		return nil, status.Errorf(codes.InvalidArgument, "国家ID不能为空")
	}

	// 3. 转换API请求到领域模型（注意类型匹配，避免隐式转换错误）
	domainDev := &domain.Developer{
		NameZh:      req.NameZh,
		NameEn:      req.NameEn,
		CountryID:   int32(req.CountryId), // 确保类型匹配domain定义
		Website:     &req.Website,
		Description: &req.Description,
		// 时间字段由数据库自动生成，无需手动设置
	}

	// 4. 调用biz层核心业务逻辑
	createdDev, err := s.uc.CreateDeveloper(ctx, domainDev)
	if err != nil {
		// 分类错误日志，便于排查
		if strings.Contains(err.Error(), "已存在") {
			s.log.WithContext(ctx).Warnf("创建开发商失败: %v", err)
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		s.log.WithContext(ctx).Errorf("创建开发商异常: %v", err)
		return nil, status.Errorf(codes.Internal, "创建开发商失败，请稍后重试")
	}

	// 5. 转换领域模型到API响应（完整映射字段）
	return &pb.DeveloperResp{
		Id:          uint32(createdDev.ID), // 转换为proto定义的uint32
		NameZh:      createdDev.NameZh,
		NameEn:      createdDev.NameEn,
		CountryId:   uint32(createdDev.CountryID),
		Website:     *createdDev.Website,
		Description: *createdDev.Description,
		CreatedAt:   createdDev.CreatedAt.Format(time.RFC3339), // 标准化时间格式
		UpdatedAt:   createdDev.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DeveloperService) GetDeveloper(ctx context.Context, req *pb.GetDeveloperReq) (*pb.DeveloperResp, error) {
	s.log.WithContext(ctx).Infof("收到查询开发商请求: ID=%d", req.Id)

	// 调用biz层获取领域模型
	developer, err := s.uc.GetDeveloperByID(ctx, req.Id)
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询开发商失败: %v", err)
		return nil, status.Errorf(codes.NotFound, "开发商不存在")
	}

	// 转换领域模型到API响应
	return &pb.DeveloperResp{
		Id:          uint32(developer.ID),
		NameZh:      developer.NameZh,
		NameEn:      developer.NameEn,
		CountryId:   uint32(developer.CountryID),
		Website:     *developer.Website,
		Description: *developer.Description,
		CreatedAt:   developer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   developer.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DeveloperService) ListDevelopers(ctx context.Context, req *pb.ListDevelopersReq) (*pb.ListDevelopersResp, error) {
	s.log.WithContext(ctx).Infof("收到开发商列表请求: 页码=%d, 每页数量=%d, 国家ID=%d, 关键词=%s",
		req.Page, req.PageSize, req.CountryId, req.Keyword)

	// 调用biz层获取领域模型列表
	developers, total, err := s.uc.ListDevelopers(ctx, req.Page, req.PageSize, req.CountryId, req.Keyword)
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询开发商列表失败: %v", err)
		return nil, status.Errorf(codes.Internal, "查询开发商列表失败")
	}

	// 转换领域模型列表到API响应
	respDevelopers := make([]*pb.DeveloperResp, 0, len(developers))
	for _, dev := range developers {
		respDevelopers = append(respDevelopers, &pb.DeveloperResp{
			Id:          uint32(dev.ID),
			NameZh:      dev.NameZh,
			NameEn:      dev.NameEn,
			CountryId:   uint32(dev.CountryID),
			Website:     *dev.Website,
			Description: *dev.Description,
			CreatedAt:   dev.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   dev.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.ListDevelopersResp{
		Total:    uint32(total),
		Items:    respDevelopers,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *DeveloperService) UpdateDevelopers(ctx context.Context, req *pb.UpdateDeveloperReq) (*pb.DeveloperResp, error) {
    s.log.WithContext(ctx).Infof("收到更新开发商请求: ID=%d, 中文名称=%s", req.Id, req.NameZh)

    // 转换API请求到领域模型
    domainDev := &domain.Developer{
        ID:          int32(req.Id),
        NameZh:      req.NameZh,
        NameEn:      req.NameEn,
        CountryID:   int32(req.CountryId),
        Website:     &req.Website,
        Description: &req.Description,
    }

    // 调用biz层更新方法
    updatedDev, err := s.uc.Update(ctx, domainDev)
    if err != nil {
        s.log.WithContext(ctx).Errorf("更新开发商失败: %v", err)
        return nil, status.Errorf(codes.Internal, "更新开发商失败")
    }

    // 转换领域模型到API响应
    return &pb.DeveloperResp{
        Id:          uint32(updatedDev.ID),
        NameZh:      updatedDev.NameZh,
        NameEn:      updatedDev.NameEn,
        CountryId:   uint32(updatedDev.CountryID),
        Website:     *updatedDev.Website,
        Description: *updatedDev.Description,
        CreatedAt:   updatedDev.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   updatedDev.UpdatedAt.Format(time.RFC3339),
    }, nil
}
