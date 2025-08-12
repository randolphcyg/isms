package domain

import (
	"context"
	"fmt"
	"time"
)

// Country 国家领域模型
type Country struct {
	ID        int32     `json:"id"`
	NameZh    string    `json:"name_zh"`   // 中文名称
	NameEn    string    `json:"name_en"`   // 英文名称
	IsoCode   string    `json:"iso_code"`  // 国家代码(ISO标准)
	Continent string    `json:"continent"` // 所属大洲
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate 业务规则校验
func (c *Country) Validate() error {
	if len(c.NameZh) == 0 || len(c.NameZh) > 100 {
		return fmt.Errorf("中文名称长度必须为1-100字符")
	}
	if len(c.NameEn) == 0 || len(c.NameEn) > 100 {
		return fmt.Errorf("英文名称长度必须为1-100字符")
	}
	if len(c.IsoCode) < 2 || len(c.IsoCode) > 10 {
		return fmt.Errorf("国家代码长度必须为2-10字符")
	}
	if len(c.Continent) == 0 || len(c.Continent) > 50 {
		return fmt.Errorf("大洲名称长度必须为1-50字符")
	}
	return nil
}

// CountryRepo 国家仓库接口
type CountryRepo interface {
	Create(ctx context.Context, country *Country) (*Country, error)
	Update(ctx context.Context, country *Country) (*Country, error)
	Delete(ctx context.Context, id int32) error
	GetByID(ctx context.Context, id int32) (*Country, error)
	List(ctx context.Context, page, pageSize uint32, keyword, continent string) ([]*Country, int64, error)
	ExistByCode(ctx context.Context, code string, excludeID int32) (bool, error)
}
