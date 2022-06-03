// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type RealWorldHTTPServer interface {
	AddComments(context.Context, *AddCommentsRequest) (*AddCommentsReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)
	DeleteComments(context.Context, *DeleteCommentsRequest) (*DeleteCommentsReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*FavoriteArticleReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*FeedArticlesReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*FollowUserReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileReply, error)
	GetTags(context.Context, *GetTagsRequest) (*GetTagsReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Registration(context.Context, *RegistrationRequest) (*RegistrationReply, error)
	UnfavoriteArticle(context.Context, *UnfavoriteArticleRequest) (*UnfavoriteArticleReply, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
}

func RegisterRealWorldHTTPServer(s *http.Server, srv RealWorldHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/login", _RealWorld_Login0_HTTP_Handler(srv))
	r.POST("/api/users", _RealWorld_Registration0_HTTP_Handler(srv))
	r.GET("/api/user", _RealWorld_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/api/user", _RealWorld_UpdateUser0_HTTP_Handler(srv))
	r.GET("/api/profiles/{username}", _RealWorld_GetProfile0_HTTP_Handler(srv))
	r.POST("/api/profiles/{username}/follow", _RealWorld_FollowUser0_HTTP_Handler(srv))
	r.DELETE("/api/profiles/{username}/follow", _RealWorld_UnfollowUser0_HTTP_Handler(srv))
	r.GET("/api/articles", _RealWorld_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _RealWorld_FeedArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}", _RealWorld_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/articles", _RealWorld_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/{slug}", _RealWorld_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}", _RealWorld_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/comments", _RealWorld_AddComments0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}/comments", _RealWorld_GetComments0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/comments/{id}", _RealWorld_DeleteComments0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/favorite", _RealWorld_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/favorite", _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv))
	r.GET("/api/tags", _RealWorld_GetTags0_HTTP_Handler(srv))
}

func _RealWorld_Login0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_Registration0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegistrationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/Registration")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Registration(ctx, req.(*RegistrationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegistrationReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetCurrentUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/GetCurrentUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentUserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/UpdateUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetProfile0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/GetProfile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/FollowUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowUser(ctx, req.(*FollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FollowUserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnfollowUser0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfollowUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/UnfollowUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfollowUser(ctx, req.(*UnfollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnfollowUserReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_ListArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/ListArticles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FeedArticles0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/FeedArticles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FeedArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/GetArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_CreateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/CreateArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UpdateArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/UpdateArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/DeleteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_AddComments0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/AddComments")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComments(ctx, req.(*AddCommentsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddCommentsReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetComments0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCommentsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/GetComments")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetComments(ctx, req.(*GetCommentsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCommentsReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_DeleteComments0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/DeleteComments")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComments(ctx, req.(*DeleteCommentsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteCommentsReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_FavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/FavoriteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FavoriteArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_UnfavoriteArticle0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnfavoriteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/UnfavoriteArticle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnfavoriteArticle(ctx, req.(*UnfavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnfavoriteArticleReply)
		return ctx.Result(200, reply)
	}
}

func _RealWorld_GetTags0_HTTP_Handler(srv RealWorldHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/realworld.v1.RealWorld/GetTags")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTagsReply)
		return ctx.Result(200, reply)
	}
}

type RealWorldHTTPClient interface {
	AddComments(ctx context.Context, req *AddCommentsRequest, opts ...http.CallOption) (rsp *AddCommentsReply, err error)
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *CreateArticleReply, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *DeleteArticleReply, err error)
	DeleteComments(ctx context.Context, req *DeleteCommentsRequest, opts ...http.CallOption) (rsp *DeleteCommentsReply, err error)
	FavoriteArticle(ctx context.Context, req *FavoriteArticleRequest, opts ...http.CallOption) (rsp *FavoriteArticleReply, err error)
	FeedArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *FeedArticlesReply, err error)
	FollowUser(ctx context.Context, req *FollowUserRequest, opts ...http.CallOption) (rsp *FollowUserReply, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *GetArticleReply, err error)
	GetComments(ctx context.Context, req *GetCommentsRequest, opts ...http.CallOption) (rsp *GetCommentsReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *GetCurrentUserReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *GetProfileReply, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *GetTagsReply, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *ListArticlesReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Registration(ctx context.Context, req *RegistrationRequest, opts ...http.CallOption) (rsp *RegistrationReply, err error)
	UnfavoriteArticle(ctx context.Context, req *UnfavoriteArticleRequest, opts ...http.CallOption) (rsp *UnfavoriteArticleReply, err error)
	UnfollowUser(ctx context.Context, req *UnfollowUserRequest, opts ...http.CallOption) (rsp *UnfollowUserReply, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *UpdateArticleReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserReply, err error)
}

type RealWorldHTTPClientImpl struct {
	cc *http.Client
}

func NewRealWorldHTTPClient(client *http.Client) RealWorldHTTPClient {
	return &RealWorldHTTPClientImpl{client}
}

func (c *RealWorldHTTPClientImpl) AddComments(ctx context.Context, in *AddCommentsRequest, opts ...http.CallOption) (*AddCommentsReply, error) {
	var out AddCommentsReply
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/AddComments"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*CreateArticleReply, error) {
	var out CreateArticleReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/CreateArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*DeleteArticleReply, error) {
	var out DeleteArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/DeleteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) DeleteComments(ctx context.Context, in *DeleteCommentsRequest, opts ...http.CallOption) (*DeleteCommentsReply, error) {
	var out DeleteCommentsReply
	pattern := "/api/articles/{slug}/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/DeleteComments"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...http.CallOption) (*FavoriteArticleReply, error) {
	var out FavoriteArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/FavoriteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*FeedArticlesReply, error) {
	var out FeedArticlesReply
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/FeedArticles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...http.CallOption) (*FollowUserReply, error) {
	var out FollowUserReply
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/FollowUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*GetArticleReply, error) {
	var out GetArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/GetArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...http.CallOption) (*GetCommentsReply, error) {
	var out GetCommentsReply
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/GetComments"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*GetCurrentUserReply, error) {
	var out GetCurrentUserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/GetCurrentUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*GetProfileReply, error) {
	var out GetProfileReply
	pattern := "/api/profiles/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/GetProfile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*GetTagsReply, error) {
	var out GetTagsReply
	pattern := "/api/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/GetTags"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*ListArticlesReply, error) {
	var out ListArticlesReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/ListArticles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) Registration(ctx context.Context, in *RegistrationRequest, opts ...http.CallOption) (*RegistrationReply, error) {
	var out RegistrationReply
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/Registration"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfavoriteArticle(ctx context.Context, in *UnfavoriteArticleRequest, opts ...http.CallOption) (*UnfavoriteArticleReply, error) {
	var out UnfavoriteArticleReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/UnfavoriteArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...http.CallOption) (*UnfollowUserReply, error) {
	var out UnfollowUserReply
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/UnfollowUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*UpdateArticleReply, error) {
	var out UpdateArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/UpdateArticle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RealWorldHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserReply, error) {
	var out UpdateUserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/realworld.v1.RealWorld/UpdateUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}