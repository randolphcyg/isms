package biz

import (
	"context"

	v1 "isms/api/isms/v1"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义业务错误（替换原有的用户相关错误）
var (
	ErrCategoryNotFound = errors.NotFound(v1.ErrorReason_CATEGORY_NOT_FOUND.String(), "行业大类不存在")
	ErrCategoryInvalid  = errors.NotFound(v1.ErrorReason_CATEGORY_INVALID.String(), "无效的大类编码")
)

// IndustryUsecase 行业分类业务用例
type IndustryUsecase struct {
	repo domain.IndustryRepo // 依赖仓库接口
	log  *log.Helper         // 日志工具
}

// NewIndustryUsecase 创建业务用例实例
func NewIndustryUsecase(repo domain.IndustryRepo, logger log.Logger) *IndustryUsecase {
	return &IndustryUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "biz/industry")),
	}
}

// ListAllCategories 业务方法：查询所有行业大类
func (uc *IndustryUsecase) ListAllCategories(ctx context.Context) ([]*domain.IndustryCategory, error) {
	uc.log.WithContext(ctx).Info("开始查询所有行业大类")

	// 调用仓库接口获取数据
	categories, err := uc.repo.GetAllCategories(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("查询所有大类失败: %v", err)
		return nil, err
	}

	return categories, nil
}

// GetSubcategories 业务方法：根据大类编码查询小类（包含业务校验）
func (uc *IndustryUsecase) GetSubcategories(ctx context.Context, categoryCode string) ([]*domain.IsmsIndustry, error) {
	uc.log.WithContext(ctx).Infof("开始查询大类[%s]下的小类", categoryCode)

	// 1. 业务参数校验（修正原有的错误逻辑，增加参数合法性检查）
	if categoryCode == "" {
		uc.log.WithContext(ctx).Error("大类编码为空")
		return nil, ErrCategoryInvalid
	}

	// 2. 校验大类是否存在（补充业务逻辑，确保大类有效）
	_, err := uc.repo.GetCategoryByCode(ctx, categoryCode)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("大类[%s]不存在: %v", categoryCode, err)
		return nil, ErrCategoryNotFound
	}

	// 3. 调用仓库接口查询小类
	subcategories, err := uc.repo.GetSubcategoriesByCode(ctx, categoryCode)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("查询大类[%s]的小类失败: %v", categoryCode, err)
		return nil, err
	}

	return subcategories, nil
}
