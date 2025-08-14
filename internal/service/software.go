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
// 修改结构体定义，添加industryUsecase依赖
type SoftwareService struct {
	pb.UnimplementedSoftwareServer

	uc         *biz.SoftwareUsecase // 依赖 biz 层的业务用例
	industryUc *biz.IndustryUsecase // 行业用例
	log        *log.Helper          // 日志工具
}

func NewSoftwareService(uc *biz.SoftwareUsecase, industryUc *biz.IndustryUsecase, logger log.Logger) *SoftwareService {
	return &SoftwareService{
		uc:         uc,
		industryUc: industryUc,
		log:        log.NewHelper(log.With(logger, "module", "service/software")),
	}
}

func (s *SoftwareService) CreateSoftware(ctx context.Context, req *pb.CreateSoftwareReq) (*pb.CreateSoftwareResp, error) {
	// 参数校验
	if req.NameEn == "" {
		return nil, status.Errorf(codes.InvalidArgument, "软件英文名称不能为空")
	}

	// 转换API请求到领域模型
	var releaseYear, releaseMonth, releaseDay *int32
	var sizeBytes *int64
	if req.ReleaseYear != 0 {
		releaseYear = &req.ReleaseYear
	}
	if req.ReleaseMonth != 0 {
		releaseMonth = &req.ReleaseMonth
	}
	if req.ReleaseDay != 0 {
		releaseDay = &req.ReleaseDay
	}
	if req.SizeBytes != 0 {
		sizeBytes = &req.SizeBytes
	}

	// 系统要求字段
	var cpuReq *string
	var memoryMinGb *float64
	var diskMinGb *float64
	var sysReqOther *string

	if req.CpuReq != "" {
		cpuReq = &req.CpuReq
	}
	if req.MemoryMinGb > 0 {
		memoryMinGb = &req.MemoryMinGb
	}
	if req.DiskMinGb > 0 {
		diskMinGb = &req.DiskMinGb
	}
	if req.SysReqOther != "" {
		sysReqOther = &req.SysReqOther
	}

	// 转换API请求到领域模型
	domainSoftware := &domain.IsmsSoftware{
		NameZh:       req.NameZh,
		NameEn:       req.NameEn,
		Version:      req.Version,
		ReleaseYear:  releaseYear,
		ReleaseMonth: releaseMonth,
		ReleaseDay:   releaseDay,
		Description:  &req.Description,

		// 系统要求字段
		CPUReq:      cpuReq,
		MemoryMinGb: memoryMinGb,
		DiskMinGb:   diskMinGb,
		SysReqOther: sysReqOther,

		CountryID:   req.CountryId,
		DeveloperID: req.DeveloperId,
		IndustryIDs: req.IndustryIds,
		SizeBytes:   sizeBytes,
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
		Id:      int64(createdSoftware.ID),
		Message: "软件创建成功",
	}, nil
}

func (s *SoftwareService) UpdateSoftware(ctx context.Context, req *pb.UpdateSoftwareReq) (*pb.UpdateSoftwareResp, error) {
	s.log.WithContext(ctx).Infof("收到更新软件请求: ID=%d", req.Id)

	// 参数校验
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "软件ID不能为空")
	}

	// 转换API请求到领域模型
	var releaseYear, releaseMonth, releaseDay *int32
	if req.ReleaseYear != 0 {
		releaseYear = &req.ReleaseYear
	}
	if req.ReleaseMonth != 0 {
		releaseMonth = &req.ReleaseMonth
	}
	if req.ReleaseDay != 0 {
		releaseDay = &req.ReleaseDay
	}

	// 转换API请求到领域模型
	domainSoftware := &domain.IsmsSoftware{
		ID:           int32(req.Id),
		NameZh:       req.NameZh,
		NameEn:       req.NameEn,
		Version:      req.Version,
		ReleaseYear:  releaseYear,
		ReleaseMonth: releaseMonth,
		ReleaseDay:   releaseDay,
		Description:  &req.Description,
		CountryID:    int32(req.CountryId),
		Status:       req.Status,
		OsIDs:        req.OsIds,
		SizeBytes:    &req.SizeBytes,
	}

	// 调用biz层核心业务逻辑
	err := s.uc.UpdateSoftware(ctx, domainSoftware)
	if err != nil {
		s.log.WithContext(ctx).Errorf("更新软件失败: %v", err)
		return nil, status.Errorf(codes.Internal, "更新软件失败: %v", err)
	}

	return &pb.UpdateSoftwareResp{Success: true, Message: "软件更新成功"}, nil
}

func (s *SoftwareService) ListSoftware(ctx context.Context, req *pb.ListSoftwareReq) (*pb.ListSoftwareResp, error) {
	s.log.WithContext(ctx).Infof("收到软件列表请求: 页码=%d, 每页数量=%d", req.Page, req.PageSize)

	// 构建查询选项
	opts := domain.ListSoftwareOptions{
		Page:         req.Page,
		PageSize:     req.PageSize,
		CategoryCode: req.CategoryCode,
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
		// 安全处理可选的日期字段
		var releaseYear, releaseMonth, releaseDay int32
		var sizeBytes int64
		if sw.ReleaseYear != nil {
			releaseYear = *sw.ReleaseYear
		}
		if sw.ReleaseMonth != nil {
			releaseMonth = *sw.ReleaseMonth
		}
		if sw.ReleaseDay != nil {
			releaseDay = *sw.ReleaseDay
		}
		if sw.SizeBytes != nil {
			sizeBytes = *sw.SizeBytes
		}

		var cpuReq, sysReqOther string
		var memoryMinGb, diskMinGb float64
		if sw.CPUReq != nil {
			cpuReq = *sw.CPUReq
		}
		if sw.MemoryMinGb != nil {
			memoryMinGb = *sw.MemoryMinGb
		}
		if sw.DiskMinGb != nil {
			diskMinGb = *sw.DiskMinGb
		}
		if sw.SysReqOther != nil {
			sysReqOther = *sw.SysReqOther
		}

		// 转换行业详情
		industryDetails := make([]*pb.IsmsIndustry, 0, len(sw.IndustryDetails))
		for _, detail := range sw.IndustryDetails {
			industryDetails = append(industryDetails, &pb.IsmsIndustry{
				Id:              detail.Id,
				CategoryCode:    detail.CategoryCode,
				CategoryName:    detail.CategoryName,
				SubcategoryCode: detail.SubcategoryCode,
				SubcategoryName: detail.SubcategoryName,
			})
		}

		respSoftwares = append(respSoftwares, &pb.IsmsSoftware{
			Id:              int64(sw.ID),
			NameZh:          sw.NameZh,
			NameEn:          sw.NameEn,
			Version:         sw.Version,
			ReleaseYear:     releaseYear,
			ReleaseMonth:    releaseMonth,
			ReleaseDay:      releaseDay,
			Description:     *sw.Description,
			CpuReq:          cpuReq,
			MemoryMinGb:     memoryMinGb,
			DiskMinGb:       diskMinGb,
			SysReqOther:     sysReqOther,
			CountryId:       sw.CountryID,
			CountryName:     sw.CountryName,
			IndustryIds:     sw.IndustryIDs,
			IndustryDetails: industryDetails,
			DeveloperId:     sw.DeveloperID,
			DeveloperName:   sw.DeveloperName,
			SizeBytes:       sizeBytes,
			Status:          sw.Status,
			OsIds:           sw.OsIDs,
			CreatedAt:       sw.CreatedAt.Format(time.RFC3339),
			UpdatedAt:       sw.UpdatedAt.Format(time.RFC3339),
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
	sw, err := s.uc.GetSoftwareByID(ctx, uint32(req.Id))
	if err != nil {
		s.log.WithContext(ctx).Errorf("查询软件失败: %v", err)
		return nil, status.Errorf(codes.NotFound, "软件不存在")
	}

	// 安全处理可选的日期字段
	var releaseYear, releaseMonth, releaseDay int32
	var sizeBytes int64
	if sw.ReleaseYear != nil {
		releaseYear = *sw.ReleaseYear
	}
	if sw.ReleaseMonth != nil {
		releaseMonth = *sw.ReleaseMonth
	}
	if sw.ReleaseDay != nil {
		releaseDay = *sw.ReleaseDay
	}
	if sw.SizeBytes != nil {
		sizeBytes = *sw.SizeBytes
	}

	var cpuReq, sysReqOther string
	var memoryMinGb, diskMinGb float64
	if sw.CPUReq != nil {
		cpuReq = *sw.CPUReq
	}
	if sw.MemoryMinGb != nil {
		memoryMinGb = *sw.MemoryMinGb
	}
	if sw.DiskMinGb != nil {
		diskMinGb = *sw.DiskMinGb
	}
	if sw.SysReqOther != nil {
		sysReqOther = *sw.SysReqOther
	}

	// 转换行业详情
	industryDetails := make([]*pb.IsmsIndustry, 0, len(sw.IndustryDetails))
	for _, detail := range sw.IndustryDetails {
		industryDetails = append(industryDetails, &pb.IsmsIndustry{
			Id:              detail.Id,
			CategoryCode:    detail.CategoryCode,
			CategoryName:    detail.CategoryName,
			SubcategoryCode: detail.SubcategoryCode,
			SubcategoryName: detail.SubcategoryName,
		})
	}

	// 转换领域模型到API响应
	return &pb.IsmsSoftware{
		Id:              int64(sw.ID),
		NameZh:          sw.NameZh,
		NameEn:          sw.NameEn,
		Version:         sw.Version,
		ReleaseYear:     releaseYear,
		ReleaseMonth:    releaseMonth,
		ReleaseDay:      releaseDay,
		Description:     *sw.Description,
		CpuReq:          cpuReq,
		MemoryMinGb:     memoryMinGb,
		DiskMinGb:       diskMinGb,
		SysReqOther:     sysReqOther,
		CountryId:       sw.CountryID,
		CountryName:     sw.CountryName,
		DeveloperId:     sw.DeveloperID,
		DeveloperName:   sw.DeveloperName,
		SizeBytes:       sizeBytes,
		Status:          sw.Status,
		OsIds:           sw.OsIDs,
		IndustryIds:     sw.IndustryIDs,
		IndustryDetails: industryDetails,
		CreatedAt:       sw.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       sw.UpdatedAt.Format(time.RFC3339),
	}, nil
}
