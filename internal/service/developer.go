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
		Website:     req.Website,
		Description: req.Description,
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
		CountryName: createdDev.CountryName, // 从领域模型获取冗余的国家名称
		Website:     createdDev.Website,
		Description: createdDev.Description,
		CreatedAt:   createdDev.CreatedAt.Format(time.RFC3339), // 标准化时间格式
		UpdatedAt:   createdDev.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *DeveloperService) GetDeveloper(ctx context.Context, req *pb.GetDeveloperReq) (*pb.DeveloperResp, error) {
	return &pb.DeveloperResp{}, nil
}
func (s *DeveloperService) ListDevelopers(ctx context.Context, req *pb.ListDevelopersReq) (*pb.ListDevelopersResp, error) {
	return &pb.ListDevelopersResp{}, nil
}
