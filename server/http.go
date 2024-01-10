package server

import (
	v1Api "app/api/v1"
	"app/configs"
	"app/internal/adapter/driver"
	"embed"
	"io/fs"
	nethttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

//go:embed dist
var content embed.FS

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *configs.ApplicationConfig, auth *driver.LoginHandler, message *driver.MessageHandler, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Server.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.Server.HTTP.Addr))
	}
	if c.Server.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Server.HTTP.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	v1Api.RegisterAuthenticationHTTPServer(srv, auth)
	v1Api.RegisterMessageHTTPServer(srv, message)
	openAPIhandler := handleSwaggerUI(configs.OpenAPI)
	srv.HandlePrefix("/q/", openAPIhandler)
	return srv
}

func handleSwaggerUI(file []byte) nethttp.Handler {
	router := mux.NewRouter()
	fsys, _ := fs.Sub(content, "dist")
	sh := nethttp.StripPrefix("/q/swagger-ui", nethttp.FileServer(nethttp.FS(fsys)))
	router.HandleFunc("/q/openapi.yaml", byteHandler(file))
	router.PathPrefix("/q/swagger-ui").Handler(sh)
	return router
}

func byteHandler(b []byte) nethttp.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write(b)
	}
}
