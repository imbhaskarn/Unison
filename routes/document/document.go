package document

import (
	"Unison/handlers/document"
	"Unison/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/", middlewares.AuthRequired(), documentHandlers.CreateDocument)
	router.GET("/:id", middlewares.AuthRequired(), documentHandlers.GetDocument)
	router.PUT("/:id", middlewares.AuthRequired(), documentHandlers.UpdateDocument)
	router.DELETE("/:id", middlewares.AuthRequired(), documentHandlers.DeleteDocument)
}
