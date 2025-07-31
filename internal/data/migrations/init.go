package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 数据库连接配置（替换为实际环境参数）
	dsn := "root:j*mPRCA2g$y^@tcp(127.0.0.1:3306)/industrial_software_db?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	// 初始化代码生成器配置
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/data/query",                       // 查询代码输出路径
		ModelPkgPath:      "./internal/data/model",                       // 模型结构体输出路径
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface, // 生成默认查询和接口
		FieldNullable:     true,                                          // 生成可为空字段（对应数据库NULLable）
		FieldWithIndexTag: true,                                          // 生成字段索引标签（便于ORM映射）
		FieldWithTypeTag:  true,                                          // 生成字段类型标签（建议开启）
	})

	// 关联数据库实例
	g.UseDB(db)

	// 生成所有表的模型（按实际表名添加）
	industryTable := g.GenerateModel("isms_industry")                  // 行业表
	countryTable := g.GenerateModel("isms_country")                    // 国家表
	osTable := g.GenerateModel("isms_os")                              // 操作系统表
	developerTable := g.GenerateModel("isms_developer")                // 开发商表
	softwareTable := g.GenerateModel("isms_software")                  // 主软件表
	softwareIndustryTable := g.GenerateModel("isms_software_industry") // 软件-行业关联表
	softwareOsTable := g.GenerateModel("isms_software_os")             // 软件-操作系统关联表

	// 应用基础CRUD生成（会生成Insert/Update/Delete/Query等方法）
	g.ApplyBasic(
		industryTable,
		countryTable,
		osTable,
		developerTable,
		softwareTable,
		softwareIndustryTable,
		softwareOsTable,
	)

	// 执行生成
	g.Execute()
}
