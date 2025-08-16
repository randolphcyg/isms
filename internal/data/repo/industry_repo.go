package repo

import (
	"context"
	"fmt"

	"isms/internal/data/model"
	"isms/internal/data/query"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// industryRepo 实现 domain.IndustryRepo 接口
type industryRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

// NewIndustryRepo 创建仓库实例
func NewIndustryRepo(db *gorm.DB, logger log.Logger) domain.IndustryRepo {
	return &industryRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/industry_repo")),
	}
}

// GetAllCategories 查询所有大类（去重）
func (r *industryRepo) GetAllCategories(ctx context.Context) ([]*domain.IndustryCategory, error) {
	var modelCats []*model.IsmsIndustry
	err := r.db.WithContext(ctx).
		Table("isms_industry").
		Select("DISTINCT category_code, category_name").
		Order("category_code ASC").
		Scan(&modelCats).Error
	if err != nil {
		r.log.Errorf("查询大类失败: %v", err)
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	// 转换：model → domain（确保字段匹配）
	domainCats := make([]*domain.IndustryCategory, 0, len(modelCats))
	for _, m := range modelCats {
		domainCats = append(domainCats, &domain.IndustryCategory{
			CategoryCode: m.CategoryCode,
			CategoryName: m.CategoryName,
		})
	}
	return domainCats, nil
}

// GetSubcategoriesByCode 根据大类编码查询小类
func (r *industryRepo) GetSubcategoriesByCode(ctx context.Context, categoryCode string) ([]*domain.IsmsIndustry, error) {
	// 调用生成的查询器查询数据模型
	modelClassifications, err := r.query.IsmsIndustry.WithContext(ctx).
		Where(r.query.IsmsIndustry.CategoryCode.Eq(categoryCode)).
		Order(r.query.IsmsIndustry.SubcategoryCode.Asc()).
		Find()
	if err != nil {
		r.log.Errorf("查询小类失败: %v", err)
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	// 转换：model → domain（重点处理 ID 字段）
	domainClassifications := make([]*domain.IsmsIndustry, 0, len(modelClassifications))
	for _, m := range modelClassifications {
		domainClassifications = append(domainClassifications, &domain.IsmsIndustry{
			ID:              m.ID, // 此处 ID 类型必须一致（均为 uint64）
			CategoryCode:    m.CategoryCode,
			CategoryName:    m.CategoryName,
			SubcategoryCode: m.SubcategoryCode,
			SubcategoryName: m.SubcategoryName,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
		})
	}
	return domainClassifications, nil
}

// GetCategoryByCode 校验大类是否存在
func (r *industryRepo) GetCategoryByCode(ctx context.Context, categoryCode string) (*domain.IndustryCategory, error) {
	modelCat, err := r.query.IsmsIndustry.WithContext(ctx).
		Where(r.query.IsmsIndustry.CategoryCode.Eq(categoryCode)).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("大类不存在: %w", err)
		}
		r.log.Errorf("查询大类失败: %v", err)
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	return &domain.IndustryCategory{
		CategoryCode: modelCat.CategoryCode,
		CategoryName: modelCat.CategoryName,
	}, nil
}

// GetIndustriesByIDs 根据ID列表批量查询行业信息
func (r *industryRepo) GetIndustriesByIDs(ctx context.Context, ids []int32) ([]*domain.IsmsIndustry, error) {
	// 调用生成的查询器查询数据模型
	modelIndustries, err := r.query.IsmsIndustry.WithContext(ctx).
		Where(r.query.IsmsIndustry.ID.In(ids...)).
		Find()
	if err != nil {
		r.log.Errorf("批量查询行业信息失败: %v", err)
		return nil, fmt.Errorf("查询失败: %w", err)
	}

	// 转换：model → domain
	industryList := make([]*domain.IsmsIndustry, 0, len(modelIndustries))
	for _, m := range modelIndustries {
		industryList = append(industryList, &domain.IsmsIndustry{
			ID:              m.ID,
			CategoryCode:    m.CategoryCode,
			CategoryName:    m.CategoryName,
			SubcategoryCode: m.SubcategoryCode,
			SubcategoryName: m.SubcategoryName,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
		})
	}
	return industryList, nil
}

// Count 统计行业总数
func (r *industryRepo) Count(ctx context.Context) (int64, error) {
	count, err := r.query.IsmsIndustry.WithContext(ctx).Count()
	if err != nil {
		return 0, fmt.Errorf("统计行业总数失败: %w", err)
	}
	return count, nil
}
