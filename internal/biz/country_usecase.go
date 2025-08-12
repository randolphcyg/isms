package biz

import (
	"context"
	"fmt"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type CountryUsecase struct {
	repo domain.CountryRepo
	log  *log.Helper
}

func NewCountryUsecase(repo domain.CountryRepo, logger log.Logger) *CountryUsecase {
	return &CountryUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/country")),
	}
}

// CreateCountry 创建国家
func (uc *CountryUsecase) CreateCountry(ctx context.Context, country *domain.Country) (*domain.Country, error) {
	// 领域模型校验
	if err := country.Validate(); err != nil {
		return nil, err
	}

	// 检查代码唯一性
	exists, err := uc.repo.ExistByCode(ctx, country.IsoCode, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("国家代码已存在")
	}

	return uc.repo.Create(ctx, country)
}

// UpdateCountry 更新国家
func (uc *CountryUsecase) UpdateCountry(ctx context.Context, country *domain.Country) (*domain.Country, error) {
	// 领域模型校验
	if err := country.Validate(); err != nil {
		return nil, err
	}

	// 检查代码唯一性（排除当前ID）
	exists, err := uc.repo.ExistByCode(ctx, country.IsoCode, country.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("国家代码已存在")
	}

	return uc.repo.Update(ctx, country)
}

// DeleteCountry 删除国家
func (uc *CountryUsecase) DeleteCountry(ctx context.Context, id int32) error {
	// 实际项目中应检查是否有关联数据
	return uc.repo.Delete(ctx, id)
}

// GetCountryByID 查询单个国家
func (uc *CountryUsecase) GetCountryByID(ctx context.Context, id uint32) (*domain.Country, error) {
	return uc.repo.GetByID(ctx, int32(id))
}

// ListCountries 分页查询国家列表
func (uc *CountryUsecase) ListCountries(ctx context.Context, page, pageSize uint32, keyword, continent string) ([]*domain.Country, int64, error) {
	return uc.repo.List(ctx, page, pageSize, keyword, continent)
}
