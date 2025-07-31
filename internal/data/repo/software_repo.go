package repo

import (
	"context"

	"isms/internal/data/query"
	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type softwareRepo struct {
	db    *gorm.DB
	query *query.Query
	log   *log.Helper
}

func NewSoftwareRepo(db *gorm.DB, logger log.Logger) domain.SoftwareRepo {
	return &softwareRepo{
		db:    db,
		query: query.Use(db),
		log:   log.NewHelper(log.With(logger, "module", "data/software_repo")),
	}
}

func (s softwareRepo) Create(ctx context.Context, software *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {
	//TODO implement me
	panic("implement me")
}

func (s softwareRepo) Update(ctx context.Context, software *domain.IsmsSoftware) error {
	//TODO implement me
	panic("implement me")
}

func (s softwareRepo) FindByID(ctx context.Context, id int64) (*domain.IsmsSoftware, error) {
	//TODO implement me
	panic("implement me")
}

func (s softwareRepo) List(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	//TODO implement me
	panic("implement me")
}
