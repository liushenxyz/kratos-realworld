package auth

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

func CreateTokenString(secret, username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"time":     time.Now(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func JWTAuth(secret, typ string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				auths := strings.SplitN(tokenString, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], typ) {
					return nil, errors.Unauthorized("token", "jwt token missing")
				}
				jwtToken := auths[1]

				token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.Unauthorized("token", "unexpected signing method")
					}
					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					if u, ok := claims["username"]; ok {
						// put CurrentUser into ctx
						ctx = NewContext(ctx, &CurrentUser{Username: u.(string)})
					}
				} else {
					return nil, errors.Unauthorized("token", "token Invalid")
				}
			}
			return handler(ctx, req)
		}
	}
}

// CurrentUser is the type of value stored in the Contexts.
type CurrentUser struct {
	Username string
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// userKey is the key for user.User values in Contexts. It is
// unexported; clients use user.NewContext and user.FromContext
// instead of using this key directly.
var userKey key

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, u *CurrentUser) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) (*CurrentUser, bool) {
	u, ok := ctx.Value(userKey).(*CurrentUser)
	return u, ok
}