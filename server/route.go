package server

import (
	"github.com/KvrocksLabs/kvrocks-controller/consts"
	"github.com/KvrocksLabs/kvrocks-controller/metadata/memory"
	"github.com/KvrocksLabs/kvrocks-controller/server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoute(engine *gin.Engine) {
	storage := memory.NewMemStorage()
	_ = storage.CreateNamespace("test-ns")
	_ = storage.CreateCluster("test-ns", "test-cluster")
	_ = storage.CreateShard("test-ns", "test-cluster", "test-shard")
	engine.Use(func(c *gin.Context) {
		c.Set(consts.ContextKeyStorage, storage)
		c.Next()
	})

	apiV1 := engine.Group("/api/v1/")
	{
		namespaces := apiV1.Group("namespaces")
		{
			namespaces.GET("", handlers.ListNamespace)
			namespaces.POST("/:namespace", handlers.CreateNamespace)
			namespaces.DELETE("/:namespace", handlers.RemoveNamespace)
		}

		clusters := namespaces.Group("/:namespace/clusters")
		{
			clusters.GET("", handlers.ListCluster)
			clusters.POST("/:cluster", handlers.CreateCluster)
			clusters.DELETE("/:cluster", handlers.RemoveCluster)
		}

		shards := clusters.Group("/:cluster/shards")
		{
			shards.GET("", handlers.ListShard)
			shards.GET("/:shard", handlers.GetShard)
			shards.POST("/:shard", handlers.CreateShard)
			shards.DELETE("/:shard", handlers.RemoveShard)
			shards.POST("/:shard/slots", handlers.AddShardSlots)
		}

		nodes := shards.Group("/:shard/nodes")
		{
			nodes.GET("", handlers.ListNode)
			nodes.POST("/:id", handlers.CreateNode)
			nodes.DELETE("/:id", handlers.RemoveNode)
		}
	}
}
