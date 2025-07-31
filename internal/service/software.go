package service

import (
	"context"

	pb "isms/api/isms/v1"
	"isms/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// SoftwareService 软件领域服务（关联行业逻辑、开发商逻辑）
type SoftwareService struct {
	pb.UnimplementedSoftwareServer

	uc  *biz.SoftwareUsecase // 依赖 biz 层的业务用例
	log *log.Helper          // 日志工具
}

func NewSoftwareService(uc *biz.SoftwareUsecase, logger log.Logger) *SoftwareService {
	return &SoftwareService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/software")),
	}
}

func (s *SoftwareService) CreateSoftware(ctx context.Context, req *pb.CreateSoftwareReq) (*pb.CreateSoftwareResp, error) {
	return &pb.CreateSoftwareResp{}, nil
}
func (s *SoftwareService) UpdateSoftware(ctx context.Context, req *pb.UpdateSoftwareReq) (*pb.UpdateSoftwareResp, error) {
	return &pb.UpdateSoftwareResp{}, nil
}
func (s *SoftwareService) ListSoftware(ctx context.Context, req *pb.ListSoftwareReq) (*pb.ListSoftwareResp, error) {
	return &pb.ListSoftwareResp{}, nil
}
func (s *SoftwareService) GetSoftwareById(ctx context.Context, req *pb.GetSoftwareByIdReq) (*pb.IsmsSoftware, error) {
	return &pb.IsmsSoftware{}, nil
}
