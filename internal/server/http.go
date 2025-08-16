package server

import (
	sHttp "net/http"
	"os"

	v1 "isms/api/isms/v1"
	"isms/internal/conf"
	"isms/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, data *conf.Data,
	industryService *service.IndustryService,
	developerService *service.DeveloperService,
	softwareService *service.SoftwareService,
	countryService *service.CountryService,
	osService *service.OSService,
	dashboardService *service.DashboardService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}),
			handlers.AllowCredentials(),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// 1. 暴露根目录的 openapi.yaml
	yamlPath := data.Openapi.Path
	if yamlPath == "" {
		log.Errorw("配置文件中未设置 openapi.path")
		// 可选：设置默认路径作为兜底
		yamlPath = "openapi.yaml"
	}
	srv.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile(yamlPath)
		if err != nil {
			sHttp.Error(w, "openapi.yaml not found", sHttp.StatusNotFound)
			return
		}

		// 设置响应头和状态码（使用标准库 http.StatusOK）
		w.Header().Set("Content-Type", "application/x-yaml")
		w.WriteHeader(sHttp.StatusOK)
		_, _ = w.Write(content)
	})

	// 2. 注册 Swagger UI，指向 openapi.yaml
	srv.HandlePrefix("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/openapi.yaml"), // 指向暴露的 YAML 文档
	))

	// 3. 注册业务接口
	v1.RegisterIndustryHTTPServer(srv, industryService)
	v1.RegisterDeveloperHTTPServer(srv, developerService)
	v1.RegisterSoftwareHTTPServer(srv, softwareService)
	v1.RegisterCountryHTTPServer(srv, countryService)
	v1.RegisterOSHTTPServer(srv, osService)
	v1.RegisterDashboardHTTPServer(srv, dashboardService)
	return srv
}
