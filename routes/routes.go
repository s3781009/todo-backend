package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"todo-server/controllers"
	"todo-server/middleware"
)

func Setup(r *gin.Engine, connection *sql.DB) {

	r.GET("api/items", middleware.IsAuthorized(controllers.GetItems(connection)))

	r.POST("/api/items", middleware.IsAuthorized(controllers.PostItems(connection)))

	r.PATCH("/api/items", middleware.IsAuthorized(controllers.UpdateItem(connection)))

	r.DELETE("api/items", middleware.IsAuthorized(controllers.DeleteItem(connection)))

}
