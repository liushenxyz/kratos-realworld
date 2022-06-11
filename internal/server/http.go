package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	v1 "realworld/api/realworld/v1"
	"realworld/internal/conf"
	"realworld/internal/pkg/middleware/auth"
	"realworld/internal/service"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/realworld.v1.RealWorld/Login":        {},
		"/realworld.v1.RealWorld/Registration": {},
		"/realworld.v1.RealWorld/GetArticle":   {},
		"/realworld.v1.RealWorld/GetTags":      {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(confServer *conf.Server, confAuth *conf.Auth, realworld *service.RealWorldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(confAuth.Secret, confAuth.Typ)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if confServer.Http.Network != "" {
		opts = append(opts, http.Network(confServer.Http.Network))
	}
	if confServer.Http.Addr != "" {
		opts = append(opts, http.Address(confServer.Http.Addr))
	}
	if confServer.Http.Timeout != nil {
		opts = append(opts, http.Timeout(confServer.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, realworld)

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	return srv
}
