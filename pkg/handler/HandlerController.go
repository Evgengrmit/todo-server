package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.SignIn)
		auth.POST("/sign_in", h.SignUp)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.CreateList)
			lists.GET("/", h.GetAllLists)
			lists.GET("/:id", h.GetListById)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)
		}
		items := api.Group(":id/items")
		{
			items.POST("/", h.CreateItem)
			items.GET("/", h.GetAllItems)
			items.GET("/:item_id", h.GetItemById)
			items.PUT("/:item_id", h.UpdateItem)
			items.DELETE("/:item_id", h.DeleteItem)
		}
	}
	return router
}
