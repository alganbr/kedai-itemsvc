package controllers

import (
	"github.com/alganbr/kedai-itemsvc/internal/managers"
	"github.com/alganbr/kedai-itemsvc/internal/models"
	"github.com/alganbr/kedai-utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type IItemController interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Patch(c *gin.Context)
}

type ItemController struct {
	basketManager managers.IItemManager
}

func NewItemController(itemManager managers.IItemManager) IItemController {
	return &ItemController{
		basketManager: itemManager,
	}
}

// Get godoc
// @Description  Get item by ID
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "Item ID"
// @Success      200  {object}  models.Item
// @Router       /item/{id} [get]
func (ctrl *ItemController) Get(c *gin.Context) {
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, getErr := ctrl.basketManager.Get(id)
	if getErr != nil {
		c.AbortWithStatusJSON(getErr.Code, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Create godoc
// @Description  Create a new item
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        request  body      models.ItemRq  true  "Item Request"
// @Success      201      {object}  models.Item
// @Router       /item [post]
func (ctrl *ItemController) Create(c *gin.Context) {
	var rq models.ItemRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	result, saveErr := ctrl.basketManager.Create(&rq)
	if saveErr != nil {
		c.AbortWithStatusJSON(saveErr.Code, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Update godoc
// @Description  Update existing item
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        id       path      int            true  "Item ID"
// @Param        request  body      models.ItemRq  true  "Item Request"
// @Success      200      {object}  models.Item
// @Router       /item/{id} [put]
func (ctrl *ItemController) Update(c *gin.Context) {
	var rq models.ItemRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, updateErr := ctrl.basketManager.Update(id, &rq)
	if updateErr != nil {
		c.AbortWithStatusJSON(updateErr.Code, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Patch godoc
// @Description  Update existing item
// @Tags         Item
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Param        request  body      models.ItemRq  true  "Item Request"
// @Success      200      {object}  models.Item
// @Router       /item/{id} [patch]
func (ctrl *ItemController) Patch(c *gin.Context) {
	var rq models.ItemRq
	if bindErr := c.ShouldBindJSON(&rq); bindErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: bindErr.Error(),
		})
		return
	}
	id, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &errors.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID",
		})
		return
	}
	result, patchErr := ctrl.basketManager.Patch(id, &rq)
	if patchErr != nil {
		c.AbortWithStatusJSON(patchErr.Code, patchErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
