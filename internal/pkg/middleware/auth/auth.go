package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
)

var secret []byte = []byte("123")

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2022, 6, 30, 0, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func JWTAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenString := tr.RequestHeader().Get("Authorization")
				// spew.Dump(aa)
				fmt.Printf("%v\r\n", tokenString)

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return secret, nil
				})

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					fmt.Println(claims["foo"], claims["nbf"])
				} else {
					fmt.Println(err)
				}

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
