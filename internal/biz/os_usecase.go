package biz

import (
	"context"
	"fmt"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type OSUsecase struct {
	repo domain.OSRepo
	log  *log.Helper
}

func NewOSUsecase(repo domain.OSRepo, logger log.Logger) *OSUsecase {
	return &OSUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/os")),
	}
}

// CreateOS 创建操作系统
func (uc *OSUsecase) CreateOS(ctx context.Context, os *domain.OS) (*domain.OS, error) {
	// 领域模型校验
	if err := os.Validate(); err != nil {
		return nil, err
	}

	// 检查唯一性
	exists, err := uc.repo.ExistByNameVersionArch(ctx, os.Name, os.Version, os.Architecture, 0)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("同名同版本同架构的操作系统已存在")
	}

	return uc.repo.Create(ctx, os)
}

// UpdateOS 更新操作系统
func (uc *OSUsecase) UpdateOS(ctx context.Context, os *domain.OS) (*domain.OS, error) {
	// 领域模型校验
	if err := os.Validate(); err != nil {
		return nil, err
	}

	// 检查唯一性（排除当前ID）
	exists, err := uc.repo.ExistByNameVersionArch(ctx, os.Name, os.Version, os.Architecture, os.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("同名同版本同架构的操作系统已存在")
	}

	return uc.repo.Update(ctx, os)
}

// DeleteOS 删除操作系统
func (uc *OSUsecase) DeleteOS(ctx context.Context, id int32) error {
	// 实际项目中应检查是否有关联数据
	return uc.repo.Delete(ctx, id)
}

// GetOSByID 查询单个操作系统
func (uc *OSUsecase) GetOSByID(ctx context.Context, id int32) (*domain.OS, error) {
	return uc.repo.GetByID(ctx, id)
}

// ListOS 分页查询操作系统列表
func (uc *OSUsecase) ListOS(ctx context.Context, page, pageSize uint32, keyword, manufacturer string) ([]*domain.OS, int64, error) {
	return uc.repo.List(ctx, page, pageSize, keyword, manufacturer)
}