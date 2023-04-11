package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/near129/practice_todo_app/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetALLTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
	return r
}
