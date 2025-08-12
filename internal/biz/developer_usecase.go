package biz

import (
	"context"
	"fmt"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type DeveloperUsecase struct {
	repo domain.DeveloperRepo // 依赖倒置：依赖领域层定义的接口
	log  *log.Helper          // 日志工具
}

func NewDeveloperUsecase(repo domain.DeveloperRepo, logger log.Logger) *DeveloperUsecase {
	return &DeveloperUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/industry")),
	}
}

// CreateDeveloper 创建
func (uc *DeveloperUsecase) CreateDeveloper(ctx context.Context, dev *domain.Developer) (*domain.Developer, error) {
	// 1. 调用领域模型的业务校验（封装核心规则）
	if err := dev.Validate(); err != nil {
		return nil, err
	}

	// 2. 业务规则：校验名称是否重复
	exists, err := uc.repo.ExistByName(ctx, dev.NameZh)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("开发商名称已存在")
	}

	// 3. 调用仓库接口完成数据持久化（不关心具体存储细节）
	return uc.repo.Create(ctx, dev)
}

// GetDeveloperByID 获取
func (uc *DeveloperUsecase) GetDeveloperByID(ctx context.Context, id uint32) (*domain.Developer, error) {
	return uc.repo.GetByID(ctx, int32(id))
}

// ListDevelopers 列表
func (uc *DeveloperUsecase) ListDevelopers(ctx context.Context, page, pageSize, countryID uint32, keyword string) ([]*domain.Developer, int64, error) {
	return uc.repo.List(ctx, page, pageSize, countryID, keyword)
}

// Update 更新开发商信息
func (uc *DeveloperUsecase) Update(ctx context.Context, dev *domain.Developer) (*domain.Developer, error) {
    // 1. 领域模型验证
    if err := dev.Validate(); err != nil {
        return nil, err
    }

    // 2. 检查名称唯一性（排除当前ID）
    exists, err := uc.repo.ExistByNameExcludeID(ctx, dev.NameZh, dev.ID)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, fmt.Errorf("开发商名称已存在")
    }

    // 3. 调用仓库层更新
    return uc.repo.Update(ctx, dev)
}
