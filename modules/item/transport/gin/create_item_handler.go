package gin_handler

import (
	"net/http"
	"social_todo_list/common"
	"social_todo_list/modules/item/biz"
	"social_todo_list/modules/item/model"
	"social_todo_list/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTodoItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewInvalidParamError(err))
			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewCreateItemBiz(store)

		if err := business.CreateItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.Id))
	}
}
