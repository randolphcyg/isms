package biz

import (
	"context"
	"fmt"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type SoftwareUsecase struct {
	repo domain.SoftwareRepo // 依赖倒置：依赖领域层定义的接口
	log  *log.Helper         // 日志工具
}

func NewSoftwareUsecase(repo domain.SoftwareRepo, logger log.Logger) *SoftwareUsecase {
	return &SoftwareUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/software")),
	}
}

// CreateSoftware 实现“创建工业软件”的完整业务流程
// 优化CreateSoftware方法，添加业务校验
func (uc *SoftwareUsecase) CreateSoftware(ctx context.Context, software *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {
	// 业务规则校验
	if software.NameZh == "" {
		return nil, fmt.Errorf("软件名称不能为空")
	}
	if software.DeveloperID == 0 {
		return nil, fmt.Errorf("开发商ID不能为空")
	}
	if software.Version == "" {
		return nil, fmt.Errorf("版本号不能为空")
	}

	// 调用领域模型的业务校验
	if err := software.Validate(); err != nil {
		return nil, err
	}

	// 检查软件是否已存在（根据名称和版本）
	exists, err := uc.repo.ExistByNameAndVersion(ctx, software.NameZh, software.Version)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("该软件版本已存在")
	}

	// 调用仓库接口完成数据持久化
	return uc.repo.Create(ctx, software)
}

func (uc *SoftwareUsecase) UpdateSoftware(ctx context.Context, software *domain.IsmsSoftware) error {
	// 业务规则校验
	if software.ID == 0 {
		return fmt.Errorf("软件ID不能为空")
	}

	// 检查软件是否存在
	exists, err := uc.repo.ExistByID(ctx, uint32(software.ID))
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("软件不存在")
	}

	// 调用领域模型的业务校验
	if err := software.Validate(); err != nil {
		return err
	}

	// 调用仓库接口更新数据
	return uc.repo.Update(ctx, software)
}

func (uc *SoftwareUsecase) GetSoftwareByID(ctx context.Context, id uint32) (*domain.IsmsSoftware, error) {
	return uc.repo.FindByID(ctx, id)
}

func (uc *SoftwareUsecase) ListSoftware(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	// 分页参数校验
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 || opts.PageSize > 100 {
		opts.PageSize = 10
	}

	return uc.repo.List(ctx, opts)
}
