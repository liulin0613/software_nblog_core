package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
	"time"
)

func Cors() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"x-request-with", "content-type", "Origin", "Accept", "token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
