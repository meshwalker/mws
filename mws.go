package main

import (
	"github.com/gin-gonic/gin"
	//log "github.com/Sirupsen/logrus"
	db "meshwalker.com/mws/pkg/database"
	"meshwalker.com/mws/customer"
	"meshwalker.com/mws/cluster"
	"meshwalker.com/mws/info"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	db.InitDb()	// Init RethinkDB

/*
	// Boutique - App-Store
	boutique := router.Group("/boutique")
	{
		boutique.GET("/", boutique.ListAllApps)
		boutique.GET("/:app_id", boutique.GetApp)
		boutique.GET("/categories", boutique.GetCategories)
		boutique.GET("/categories/:categorie")
	}*/

	// Customer
	customerRoutes := router.Group("/customer")
	{
		customerRoutes.GET("/:id", customer.GetById)
		customerRoutes.POST("/", customer.Create)
		customerRoutes.PUT("/:id", customer.Update)
		customerRoutes.DELETE("/:id", customer.Delete)
	}

	router.Any("/*action", info.Info)

	// Cluster health check (ddns)
	router.GET("getmyip", cluster.GetMyIP)
	router.GET("/health/:cid", cluster.HealthCheck)

	// Cluster
	clusterRoutes := router.Group("/cluster")
	{
		clusterRoutes.GET("/:cid", cluster.GetById)
		clusterRoutes.POST("/", cluster.Create)
		clusterRoutes.PUT("/:cid", cluster.Update)
		clusterRoutes.DELETE("/:cid", cluster.Delete)
		clusterRoutes.POST("/:cid/nodes/", cluster.AddNode)
		clusterRoutes.DELETE("/:cid/nodes/:nid", cluster.RemoveNode)
	}

	router.Run(":8877")
}