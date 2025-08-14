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

type OSService struct {
	pb.UnimplementedOSServer

	uc  *biz.OSUsecase
	log *log.Helper
}

func NewOSService(uc *biz.OSUsecase, logger log.Logger) *OSService {
	return &OSService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/os")),
	}
}

// CreateOS 创建操作系统
func (s *OSService) CreateOS(ctx context.Context, req *pb.CreateOSReq) (*pb.OSResp, error) {
	s.log.WithContext(ctx).Infof("创建操作系统: %s %s(%s)", req.Name, req.Version, req.Architecture)

	os := &domain.OS{
		Name:         req.Name,
		Version:      req.Version,
		Architecture: req.Architecture,
		Manufacturer: &req.Manufacturer,
		ReleaseYear:  &req.ReleaseYear,
		Description:  &req.Description,
	}

	// 处理可选字段
	if req.Manufacturer == "" {
		os.Manufacturer = nil
	}
	if req.ReleaseYear == 0 {
		os.ReleaseYear = nil
	}
	if req.Description == "" {
		os.Description = nil
	}

	created, err := s.uc.CreateOS(ctx, os)
	if err != nil {
		if strings.Contains(err.Error(), "已存在") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "创建操作系统失败: %v", err)
	}

	return s.toOSResp(created), nil
}

// UpdateOS 更新操作系统
func (s *OSService) UpdateOS(ctx context.Context, req *pb.UpdateOSReq) (*pb.OSResp, error) {
	s.log.WithContext(ctx).Infof("更新操作系统: ID=%d, %s %s(%s)", req.Id, req.Name, req.Version, req.Architecture)

	os := &domain.OS{
		ID:           int32(req.Id),
		Name:         req.Name,
		Version:      req.Version,
		Architecture: req.Architecture,
		Manufacturer: &req.Manufacturer,
		ReleaseYear:  &req.ReleaseYear,
		Description:  &req.Description,
	}

	// 处理可选字段
	if req.Manufacturer == "" {
		os.Manufacturer = nil
	}
	if req.ReleaseYear == 0 {
		os.ReleaseYear = nil
	}
	if req.Description == "" {
		os.Description = nil
	}

	updated, err := s.uc.UpdateOS(ctx, os)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新操作系统失败: %v", err)
	}

	return s.toOSResp(updated), nil
}

// DeleteOS 删除操作系统
func (s *OSService) DeleteOS(ctx context.Context, req *pb.DeleteOSReq) (*pb.DeleteOSResp, error) {
	s.log.WithContext(ctx).Infof("删除操作系统: ID=%d", req.Id)

	err := s.uc.DeleteOS(ctx, int32(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除操作系统失败: %v", err)
	}

	return &pb.DeleteOSResp{Success: true, Message: "删除成功"}, nil
}

// GetOS 查询单个操作系统
func (s *OSService) GetOS(ctx context.Context, req *pb.GetOSReq) (*pb.OSResp, error) {
	os, err := s.uc.GetOSByID(ctx, int32(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "操作系统不存在: %v", err)
	}

	return s.toOSResp(os), nil
}

// ListOS 分页查询操作系统列表
func (s *OSService) ListOS(ctx context.Context, req *pb.ListOSReq) (*pb.ListOSResp, error) {
	osList, total, err := s.uc.ListOS(ctx, req.Page, req.PageSize, req.Keyword, req.Manufacturer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询操作系统列表失败: %v", err)
	}

	items := make([]*pb.OSResp, 0, len(osList))
	for _, os := range osList {
		items = append(items, s.toOSResp(os))
	}

	return &pb.ListOSResp{
		Items:    items,
		Total:    uint32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// 转换领域模型到API响应
func (s *OSService) toOSResp(os *domain.OS) *pb.OSResp {
	resp := &pb.OSResp{
		Id:           uint32(os.ID),
		Name:         os.Name,
		Version:      os.Version,
		Architecture: os.Architecture,
		CreatedAt:    os.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    os.UpdatedAt.Format(time.RFC3339),
	}

	// 处理可选字段
	if os.Manufacturer != nil {
		resp.Manufacturer = *os.Manufacturer
	}
	if os.ReleaseYear != nil {
		resp.ReleaseYear = *os.ReleaseYear
	}
	if os.Description != nil {
		resp.Description = *os.Description
	}

	return resp
}
