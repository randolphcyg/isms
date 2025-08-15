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

type osRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

func NewOSRepo(db *gorm.DB, logger log.Logger) domain.OSRepo {
	return &osRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/os_repo")),
	}
}

// Create 创建操作系统
func (r *osRepo) Create(ctx context.Context, os *domain.OS) (*domain.OS, error) {
	dataModel := &model.IsmsOS{
		Name:         os.Name,
		Version:      os.Version,
		Architecture: os.Architecture,
		Manufacturer: os.Manufacturer,
		ReleaseYear:  os.ReleaseYear,
		Description:  os.Description,
	}

	err := r.query.IsmsOS.WithContext(ctx).Create(dataModel)
	if err != nil {
		return nil, fmt.Errorf("创建操作系统失败: %w", err)
	}

	os.ID = dataModel.ID
	os.CreatedAt = dataModel.CreatedAt
	os.UpdatedAt = dataModel.UpdatedAt
	return os, nil
}

// Update 更新操作系统
func (r *osRepo) Update(ctx context.Context, os *domain.OS) (*domain.OS, error) {
	dataModel := &model.IsmsOS{
		ID:           os.ID,
		Name:         os.Name,
		Version:      os.Version,
		Architecture: os.Architecture,
		Manufacturer: os.Manufacturer,
		ReleaseYear:  os.ReleaseYear,
		Description:  os.Description,
	}

	err := r.query.IsmsOS.WithContext(ctx).Save(dataModel)
	if err != nil {
		return nil, fmt.Errorf("更新操作系统失败: %w", err)
	}

	return &domain.OS{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		Version:      dataModel.Version,
		Architecture: dataModel.Architecture,
		Manufacturer: dataModel.Manufacturer,
		ReleaseYear:  dataModel.ReleaseYear,
		Description:  dataModel.Description,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}, nil
}

// Delete 删除操作系统
func (r *osRepo) Delete(ctx context.Context, id int32) error {
	_, err := r.query.IsmsOS.WithContext(ctx).
		Where(query.IsmsOS.ID.Eq(id)).
		Delete()
	if err != nil {
		return fmt.Errorf("删除操作系统失败: %w", err)
	}
	return nil
}

// GetByID 查询单个操作系统
func (r *osRepo) GetByID(ctx context.Context, id int32) (*domain.OS, error) {
	q := r.query.IsmsOS.WithContext(ctx)
	q = q.Where(r.query.IsmsOS.ID.Eq(id))

	dataModel, err := q.First()
	if err != nil {
		return nil, fmt.Errorf("查询操作系统失败: %w", err)
	}

	return &domain.OS{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		Version:      dataModel.Version,
		Architecture: dataModel.Architecture,
		Manufacturer: dataModel.Manufacturer,
		ReleaseYear:  dataModel.ReleaseYear,
		Description:  dataModel.Description,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}, nil
}

// List 分页查询操作系统列表
func (r *osRepo) List(ctx context.Context, page, pageSize uint32, keyword, manufacturer string) ([]*domain.OS, int64, error) {
	q := r.query.IsmsOS.WithContext(ctx)

	// 条件筛选
	if keyword != "" {
		q = q.Where(r.query.IsmsOS.Name.Like("%" + keyword + "%")).
			Or(r.query.IsmsOS.Version.Like("%" + keyword + "%"))
	}
	if manufacturer != "" {
		// 需要正确构造manufacturer查询条件
		if keyword != "" {
			// 如果已经有keyword条件，需要使用嵌套条件
			q = q.Where(r.query.IsmsOS.Manufacturer.Eq(manufacturer))
		} else {
			// 如果没有keyword条件，直接添加manufacturer条件
			q = q.Where(r.query.IsmsOS.Manufacturer.Eq(manufacturer))
		}
	}

	// 总数查询
	total, err := q.Count()
	if err != nil {
		return nil, 0, fmt.Errorf("查询操作系统总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	dataModels, err := q.
		Offset(int(offset)).
		Limit(int(pageSize)).
		Order(r.query.IsmsOS.Name.Asc()).
		Order(r.query.IsmsOS.Version.Asc()).
		Find()
	if err != nil {
		return nil, 0, fmt.Errorf("查询操作系统列表失败: %w", err)
	}

	// 转换领域模型
	osList := make([]*domain.OS, 0, len(dataModels))
	for _, m := range dataModels {
		// 添加nil检查避免空指针异常
		if m == nil {
			continue
		}
		osList = append(osList, &domain.OS{
			ID:           m.ID,
			Name:         m.Name,
			Version:      m.Version,
			Architecture: m.Architecture,
			Manufacturer: m.Manufacturer,
			ReleaseYear:  m.ReleaseYear,
			Description:  m.Description,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
		})
	}

	return osList, total, nil
}

// ExistByNameVersionArch 检查同名同版本同架构的操作系统是否存在
func (r *osRepo) ExistByNameVersionArch(ctx context.Context, name, version, architecture string, excludeID int32) (bool, error) {
	q := r.query.IsmsOS.WithContext(ctx).
		Where(query.IsmsOS.Name.Eq(name)).
		Where(query.IsmsOS.Version.Eq(version)).
		Where(query.IsmsOS.Architecture.Eq(architecture))

	if excludeID > 0 {
		q = q.Where(query.IsmsOS.ID.Neq(excludeID))
	}

	count, err := q.Count()
	if err != nil {
		return false, fmt.Errorf("查询操作系统失败: %w", err)
	}
	return count > 0, nil
}
