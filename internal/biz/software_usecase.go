package biz

import (
	"context"
	"fmt"

	"isms/internal/domain"

	"github.com/go-kratos/kratos/v2/log"
)

type SoftwareUsecase struct {
	repo          domain.SoftwareRepo  // 依赖倒置：依赖领域层定义的接口
	countryRepo   domain.CountryRepo   // 添加国家仓储依赖
	developerRepo domain.DeveloperRepo // 添加开发商仓储依赖
	log           *log.Helper          // 日志工具
}

func NewSoftwareUsecase(repo domain.SoftwareRepo, countryRepo domain.CountryRepo, developerRepo domain.DeveloperRepo, logger log.Logger) *SoftwareUsecase {
	return &SoftwareUsecase{
		repo:          repo,
		countryRepo:   countryRepo,
		developerRepo: developerRepo,
		log:           log.NewHelper(log.With(logger, "module", "biz/software")),
	}
}

// CreateSoftware 创建工业软件
func (uc *SoftwareUsecase) CreateSoftware(ctx context.Context, software *domain.IsmsSoftware) (*domain.IsmsSoftware, error) {
	// 业务规则校验
	if software.NameEn == "" {
		return nil, fmt.Errorf("软件英文名称不能为空")
	}
	if software.DeveloperID == 0 {
		return nil, fmt.Errorf("开发商ID不能为空")
	}
	if software.Version == "" {
		return nil, fmt.Errorf("版本号不能为空")
	}

	fmt.Println("@@@@@@@@")
	fmt.Println(software.IndustryIDs)

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
	software, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 填充国家和开发商名称
	if err := uc.fillCountryAndDeveloperNames(ctx, software); err != nil {
		uc.log.Warnf("填充软件ID=%d的国家和开发商名称失败: %v", id, err)
		// 不返回错误，继续处理
	}

	return software, nil
}

func (uc *SoftwareUsecase) ListSoftware(ctx context.Context, opts domain.ListSoftwareOptions) ([]*domain.IsmsSoftware, int64, error) {
	// 分页参数校验
	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 || opts.PageSize > 100 {
		opts.PageSize = 10
	}

	softwares, total, err := uc.repo.List(ctx, opts)
	if err != nil {
		return nil, 0, err
	}

	// 批量填充国家和开发商名称
	if err := uc.fillCountryAndDeveloperNamesBatch(ctx, softwares); err != nil {
		uc.log.Warnf("批量填充软件列表的国家和开发商名称失败: %v", err)
		// 不返回错误，继续处理
	}

	return softwares, total, nil
}

// fillCountryAndDeveloperNames 填充单个软件的国家和开发商名称
func (uc *SoftwareUsecase) fillCountryAndDeveloperNames(ctx context.Context, software *domain.IsmsSoftware) error {
	// 获取国家名称
	if software.CountryID > 0 {
		country, err := uc.countryRepo.GetByID(ctx, software.CountryID)
		if err != nil {
			return fmt.Errorf("查询国家ID=%d失败: %w", software.CountryID, err)
		}
		software.CountryName = country.NameZh
	}

	// 获取开发商名称
	if software.DeveloperID > 0 {
		developer, err := uc.developerRepo.GetByID(ctx, software.DeveloperID)
		if err != nil {
			return fmt.Errorf("查询开发商ID=%d失败: %w", software.DeveloperID, err)
		}
		software.DeveloperName = developer.NameZh
	}

	return nil
}

// fillCountryAndDeveloperNamesBatch 批量填充软件列表的国家和开发商名称
func (uc *SoftwareUsecase) fillCountryAndDeveloperNamesBatch(ctx context.Context, softwares []*domain.IsmsSoftware) error {
	if len(softwares) == 0 {
		return nil
	}

	// 收集所有需要查询的国家ID和开发商ID
	countryIDSet := make(map[int32]struct{})
	developerIDSet := make(map[int32]struct{})

	for _, sw := range softwares {
		if sw.CountryID > 0 {
			countryIDSet[sw.CountryID] = struct{}{}
		}
		if sw.DeveloperID > 0 {
			developerIDSet[sw.DeveloperID] = struct{}{}
		}
	}

	// 批量查询国家信息
	countryMap := make(map[int32]string)
	if len(countryIDSet) > 0 {
		countryIDs := make([]int32, 0, len(countryIDSet))
		for id := range countryIDSet {
			countryIDs = append(countryIDs, id)
		}

		for _, id := range countryIDs {
			country, err := uc.countryRepo.GetByID(ctx, id)
			if err != nil {
				uc.log.Warnf("查询国家ID=%d失败: %v", id, err)
				continue
			}
			countryMap[id] = country.NameZh
		}
	}

	// 批量查询开发商信息
	developerMap := make(map[int32]string)
	if len(developerIDSet) > 0 {
		developerIDs := make([]int32, 0, len(developerIDSet))
		for id := range developerIDSet {
			developerIDs = append(developerIDs, id)
		}

		for _, id := range developerIDs {
			developer, err := uc.developerRepo.GetByID(ctx, id)
			if err != nil {
				uc.log.Warnf("查询开发商ID=%d失败: %v", id, err)
				continue
			}
			developerMap[id] = developer.NameZh
		}
	}

	// 填充软件列表中的国家和开发商名称
	for _, sw := range softwares {
		if sw.CountryID > 0 {
			if name, ok := countryMap[sw.CountryID]; ok {
				sw.CountryName = name
			}
		}
		if sw.DeveloperID > 0 {
			if name, ok := developerMap[sw.DeveloperID]; ok {
				sw.DeveloperName = name
			}
		}
	}

	return nil
}
