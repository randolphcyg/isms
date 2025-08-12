package data

import (
	"isms/internal/conf"
	"isms/internal/data/repo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData,
	wire.FieldsOf(new(*Data), "db"),
	repo.NewIndustryRepo,
	repo.NewDeveloperRepo,
	repo.NewSoftwareRepo,
	repo.NewCountryRepo,
)

// Data .
type Data struct {
	Db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db := InitDB(c.Database)
	return &Data{
		Db: db,
	}, cleanup, nil
}
