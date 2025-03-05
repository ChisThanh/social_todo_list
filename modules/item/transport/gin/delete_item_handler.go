package gin_handler

import (
	"net/http"
	"social_todo_list/common"
	"social_todo_list/modules/item/biz"
	"social_todo_list/modules/item/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteTodoItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStorage(db)
		business := biz.NewDeleteItemBiz(store)

		if err := business.DeleteItem(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
