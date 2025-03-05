package main

import (
	"fmt"
	"social_todo_list/middleware"
	gin_handler "social_todo_list/modules/item/transport/gin"

	"github.com/gin-gonic/gin"
)

// Clean Architecture

func main() {
	r := gin.Default()
	store := middleware.NewLimiterStore()
	r.Use(middleware.RateLimiterMiddleware(store))
	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", gin_handler.CreateTodoItem(db))
			items.GET("", gin_handler.GetTodoItems(db))
			items.GET("/:id", gin_handler.GetTodoItemId(db))
			items.PATCH("/:id", gin_handler.UpdateTodoItem(db))
			items.DELETE("/:id", gin_handler.DeleteTodoItem(db))
		}
	}

	port := fmt.Sprintf(":%d", config.Server.Port)
	r.Run(port)
}
