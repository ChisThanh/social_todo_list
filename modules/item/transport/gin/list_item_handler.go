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

func GetTodoItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		paging.Process()

		var filter model.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewListItemBiz(store)

		items, err := business.ListItems(c.Request.Context(), &paging, &filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(items, paging, filter))
	}
}
