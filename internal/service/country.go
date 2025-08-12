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

type CountryService struct {
	pb.UnimplementedCountryServer

	uc  *biz.CountryUsecase
	log *log.Helper
}

func NewCountryService(uc *biz.CountryUsecase, logger log.Logger) *CountryService {
	return &CountryService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/country")),
	}
}

// CreateCountry 创建国家
func (s *CountryService) CreateCountry(ctx context.Context, req *pb.CreateCountryReq) (*pb.CountryResp, error) {
	s.log.WithContext(ctx).Infof("创建国家: %s(%s)", req.NameZh, req.Code)

	country := &domain.Country{
		NameZh:    req.NameZh,
		NameEn:    req.NameEn,
		IsoCode:   req.Code,
		Continent: req.Continent,
	}

	created, err := s.uc.CreateCountry(ctx, country)
	if err != nil {
		if strings.Contains(err.Error(), "已存在") {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "创建国家失败: %v", err)
	}

	return s.toCountryResp(created), nil
}

// UpdateCountry 更新国家
func (s *CountryService) UpdateCountry(ctx context.Context, req *pb.UpdateCountryReq) (*pb.CountryResp, error) {
	s.log.WithContext(ctx).Infof("更新国家: ID=%d, %s", req.Id, req.NameZh)

	country := &domain.Country{
		ID:        int32(req.Id),
		NameZh:    req.NameZh,
		NameEn:    req.NameEn,
		IsoCode:   req.Code,
		Continent: req.Continent,
	}

	updated, err := s.uc.UpdateCountry(ctx, country)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新国家失败: %v", err)
	}

	return s.toCountryResp(updated), nil
}

// DeleteCountry 删除国家
func (s *CountryService) DeleteCountry(ctx context.Context, req *pb.DeleteCountryReq) (*pb.DeleteCountryResp, error) {
	s.log.WithContext(ctx).Infof("删除国家: ID=%d", req.Id)

	err := s.uc.DeleteCountry(ctx, int32(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除国家失败: %v", err)
	}

	return &pb.DeleteCountryResp{Success: true, Message: "删除成功"}, nil
}

// GetCountry 查询单个国家
func (s *CountryService) GetCountry(ctx context.Context, req *pb.GetCountryReq) (*pb.CountryResp, error) {
	country, err := s.uc.GetCountryByID(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "国家不存在: %v", err)
	}

	return s.toCountryResp(country), nil
}

// ListCountries 分页查询国家列表
func (s *CountryService) ListCountries(ctx context.Context, req *pb.ListCountriesReq) (*pb.ListCountriesResp, error) {
	countries, total, err := s.uc.ListCountries(ctx, req.Page, req.PageSize, req.Keyword, req.Continent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "查询国家列表失败: %v", err)
	}

	items := make([]*pb.CountryResp, 0, len(countries))
	for _, c := range countries {
		items = append(items, s.toCountryResp(c))
	}

	return &pb.ListCountriesResp{
		Items:    items,
		Total:    uint32(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

// 转换领域模型到API响应
func (s *CountryService) toCountryResp(country *domain.Country) *pb.CountryResp {
	return &pb.CountryResp{
		Id:        uint32(country.ID),
		NameZh:    country.NameZh,
		NameEn:    country.NameEn,
		Code:      country.IsoCode,
		Continent: country.Continent,
		CreatedAt: country.CreatedAt.Format(time.RFC3339),
		UpdatedAt: country.UpdatedAt.Format(time.RFC3339),
	}
}
