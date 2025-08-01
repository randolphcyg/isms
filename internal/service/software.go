package service

import (
	"context"
	"time"

	pb "isms/api/isms/v1"
	"isms/internal/biz"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//s.log.WithContext(ctx).Infof("收到创建软件请求: 名称=%s, 开发商ID=%d", req.Name, req.DeveloperId)

	// 参数校验
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "软件名称不能为空")
	}

	// 转换API请求到领域模型
	domainSoftware := &domain.IsmsSoftware{
		NameZh:      req.Name,
		NameEn:      req.NameEn,
		Version:     req.Version,
		Description: &req.Description,
		CountryID:   int32(req.CountryId),
		Status:      req.Status,
		OsIDs:       req.OsIds,
	}

	// 调用biz层核心业务逻辑
	createdSoftware, err := s.uc.CreateSoftware(ctx, domainSoftware)
	if err != nil {
		s.log.WithContext(ctx).Errorf("创建软件失败: %v", err)
		return nil, status.Errorf(codes.Internal, "创建软件失败: %v", err)
	}

	return &pb.CreateSoftwareResp{
		Id: int64(createdSoftware.ID),
	}, nil
}

func (s *SoftwareService) UpdateSoftware(ctx context.Context, req *pb.UpdateSoftwareReq) (*pb.UpdateSoftwareResp, error) {
	s.log.WithContext(ctx).Infof("收到更新软件请求: ID=%d", req.Id)

	// 参数校验
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "软件ID不能为空")
	}

	// 转换API请求到领域模型
	domainSoftware := &domain.IsmsSoftware{
		ID:          int32(req.Id),
		NameZh:      req.Name,
		NameEn:      req.NameEn,
		Version:     req.Version,
		Description: &req.Description,
		CountryID:   int32(req.CountryId),
		Status:      req.Status,
		OsIDs:       req.OsIds,
	}

	// 调用biz层核心业务逻辑
	err := s.uc.UpdateSoftware(ctx, domainSoftware)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新软件失败: %v", err)
		return nil, status.Errorf(codes.Internal, "更新软件失败: %v", err)
	}

	return &pb.UpdateSoftwareResp{Success: true}, nil
}

func (s *SoftwareService) ListSoftware(ctx context.Context, req *pb.ListSoftwareReq) (*pb.ListSoftwareResp, error) {
	s.log.WithContext(ctx).Infof("收到软件列表请求: 页码=%d, 每页数量=%d", req.Page, req.PageSize)

	// 构建查询选项
	opts := domain.ListSoftwareOptions{
		Page:         req.Page,
		PageSize:     req.PageSize,
		CategoryCode: req.Category,
		Status:       req.Status,
		Keyword:      req.Keyword,
	}

	// 调用biz层核心业务逻辑
	softwares, total, err := s.uc.ListSoftware(ctx, opts)
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询软件列表失败: %v", err)
		return nil, status.Errorf(codes.Internal, "查询软件列表失败: %v", err)
	}

	// 转换领域模型列表到API响应
	respSoftwares := make([]*pb.IsmsSoftware, 0, len(softwares))
	for _, sw := range softwares {
		respSoftwares = append(respSoftwares, &pb.IsmsSoftware{
			Id:          int64(sw.ID),
			Name:        sw.NameZh,
			NameEn:      sw.NameEn,
			Version:     sw.Version,
			Description: *sw.Description,
			CountryId:   int64(sw.CountryID),
			Status:      sw.Status,
			OsIds:       sw.OsIDs,
			CreatedAt:   sw.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   sw.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.ListSoftwareResp{
		Total:    total,
		Items:    respSoftwares,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *SoftwareService) GetSoftwareById(ctx context.Context, req *pb.GetSoftwareByIdReq) (*pb.IsmsSoftware, error) {
	s.log.WithContext(ctx).Infof("收到查询软件请求: ID=%d", req.Id)

	// 调用biz层核心业务逻辑
	software, err := s.uc.GetSoftwareByID(ctx, uint32(req.Id))
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询软件失败: %v", err)
		return nil, status.Errorf(codes.NotFound, "软件不存在")
	}

	// 转换领域模型到API响应
	return &pb.IsmsSoftware{
		Id:          int64(software.ID),
		Name:        software.NameZh,
		NameEn:      software.NameEn,
		Version:     software.Version,
		Description: *software.Description,
		CountryId:   int64(software.CountryID),
		Status:      software.Status,
		OsIds:       software.OsIDs,
		CreatedAt:   software.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   software.UpdatedAt.Format(time.RFC3339),
	}, nil
}
