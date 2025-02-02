package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/more", MoreHandler)
	router.POST("/feed", FeedHandler)
	router.POST("/recommend", RecommendHandler)
	router.POST("/views", ViewsHandler)
	router.POST("/watchingnow", WatchingNowHandler)
	router.GET("/metrics", PrometheusHandler())
	router.GET("/health_check", HealthCheckHandler)
}
