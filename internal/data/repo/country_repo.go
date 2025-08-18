package repo

import (
	"context"
	"fmt"

	"isms/internal/data/model"
	"isms/internal/data/query"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type countryRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

func NewCountryRepo(db *gorm.DB, logger log.Logger) domain.CountryRepo {
	return &countryRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/country_repo")),
	}
}

// Create 创建国家
func (r *countryRepo) Create(ctx context.Context, country *domain.Country) (*domain.Country, error) {
	var continentPtr *string
	if country.Continent != "" {
		continentPtr = &country.Continent
	}

	dataModel := &model.IsmsCountry{
		NameZh:    country.NameZh,
		NameEn:    country.NameEn,
		IsoCode:   country.IsoCode,
		Continent: continentPtr, // 使用可能为nil的指针
	}

	err := r.query.IsmsCountry.WithContext(ctx).Create(dataModel)
	if err != nil {
		return nil, fmt.Errorf("创建国家失败: %w", err)
	}

	country.ID = dataModel.ID
	country.CreatedAt = dataModel.CreatedAt
	country.UpdatedAt = dataModel.UpdatedAt
	return country, nil
}

// Update 更新国家
func (r *countryRepo) Update(ctx context.Context, country *domain.Country) (*domain.Country, error) {
	var continentPtr *string
	if country.Continent != "" {
		continentPtr = &country.Continent
	}

	dataModel := &model.IsmsCountry{
		ID:        country.ID,
		NameZh:    country.NameZh,
		NameEn:    country.NameEn,
		IsoCode:   country.IsoCode,
		Continent: continentPtr, // 使用可能为nil的指针
	}

	err := r.query.IsmsCountry.WithContext(ctx).Save(dataModel)
	if err != nil {
		return nil, fmt.Errorf("更新国家失败: %w", err)
	}

	// 安全地处理continent字段
	continent := ""
	if dataModel.Continent != nil {
		continent = *dataModel.Continent
	}

	return &domain.Country{
		ID:        dataModel.ID,
		NameZh:    dataModel.NameZh,
		NameEn:    dataModel.NameEn,
		IsoCode:   dataModel.IsoCode,
		Continent: continent, // 使用安全的字符串值
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}, nil
}

// Delete 删除国家
func (r *countryRepo) Delete(ctx context.Context, id int32) error {
	_, err := r.query.IsmsCountry.WithContext(ctx).
		Where(query.IsmsCountry.ID.Eq(id)).
		Delete()
	if err != nil {
		return fmt.Errorf("删除国家失败: %w", err)
	}
	return nil
}

// GetByID 查询单个国家
func (r *countryRepo) GetByID(ctx context.Context, id int32) (*domain.Country, error) {
	q := r.query.IsmsCountry.WithContext(ctx)
	q = q.Where(r.query.IsmsCountry.ID.Eq(id))

	dataModel, err := q.First()
	if err != nil {
		return nil, fmt.Errorf("查询国家失败: %w", err)
	}

	continent := ""
	if dataModel.Continent != nil {
		continent = *dataModel.Continent
	}

	return &domain.Country{
		ID:        dataModel.ID,
		NameZh:    dataModel.NameZh,
		NameEn:    dataModel.NameEn,
		IsoCode:   dataModel.IsoCode,
		Continent: continent,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}, nil
}

// List 分页查询国家列表
func (r *countryRepo) List(ctx context.Context, page, pageSize uint32, keyword, continent string) ([]*domain.Country, int64, error) {
	q := r.query.IsmsCountry.WithContext(ctx)

	// 条件筛选
	if keyword != "" {
		q = q.Where(query.IsmsCountry.NameZh.Like("%" + keyword + "%")).
			Or(query.IsmsCountry.NameEn.Like("%" + keyword + "%")).
			Or(query.IsmsCountry.IsoCode.Like("%" + keyword + "%"))
	}
	if continent != "" {
		q = q.Where(query.IsmsCountry.Continent.Eq(continent))
	}

	// 总数查询
	total, err := q.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("查询国家总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	dataModels, err := q.
		Offset(int(offset)).
		Limit(int(pageSize)).
		Order(r.query.IsmsCountry.ID.Asc()).
		Find()
	if err != nil {
		return nil, 0, fmt.Errorf("查询国家列表失败: %w", err)
	}

	// 转换领域模型
	countries := make([]*domain.Country, 0, len(dataModels))
	for _, m := range dataModels {
		// 添加nil检查避免空指针异常
		if m == nil {
			continue
		}
		continent := ""
		if m.Continent != nil {
			continent = *m.Continent
		}
		countries = append(countries, &domain.Country{
			ID:        m.ID,
			NameZh:    m.NameZh,
			NameEn:    m.NameEn,
			IsoCode:   m.IsoCode,
			Continent: continent,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		})
	}

	return countries, total, nil
}

// ExistByCode 检查国家代码是否存在
func (r *countryRepo) ExistByCode(ctx context.Context, code string, excludeID int32) (bool, error) {
	q := r.query.IsmsCountry.WithContext(ctx).
		Where(r.query.IsmsCountry.IsoCode.Eq(code))

	if excludeID > 0 {
		q = q.Where(r.query.IsmsCountry.ID.Neq(excludeID))
	}

	count, err := q.Count()
	if err != nil {
		return false, fmt.Errorf("查询国家代码失败: %w", err)
	}
	return count > 0, nil
}

// Count 统计国家总数
func (r *countryRepo) Count(ctx context.Context) (int64, error) {
	count, err := r.query.IsmsCountry.WithContext(ctx).Count()
	if err != nil {
		return 0, fmt.Errorf("统计国家总数失败: %w", err)
	}
	return count, nil
}
