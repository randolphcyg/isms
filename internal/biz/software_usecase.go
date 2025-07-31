package biz

import (
	"context"

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
func (uc *SoftwareUsecase) CreateSoftware(ctx context.Context, dev *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {

	return uc.repo.Create(ctx, dev)
}
