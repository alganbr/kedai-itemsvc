package routes

import "github.com/alganbr/kedai-itemsvc/internal/controllers"

type ItemRoutes struct {
	router         Router
	itemController controllers.IItemController
}

func NewItemRoutes(router Router, itemController controllers.IItemController) ItemRoutes {
	return ItemRoutes{
		router:         router,
		itemController: itemController,
	}
}

func (r *ItemRoutes) Setup() {
	itemGroup := r.router.Path.Group("/item")
	itemGroup.GET("/:id", r.itemController.Get)
	itemGroup.POST("", r.itemController.Create)
	itemGroup.PUT("/:id", r.itemController.Update)
	itemGroup.PATCH("/:id", r.itemController.Patch)
}
