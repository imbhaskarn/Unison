package documentHandlers

import (
	"Unison/db"

	"github.com/gin-gonic/gin"
)

func CreateDocument(c *gin.Context) {

	var resBody map[string]any
	if err := c.ShouldBindJSON(&resBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	title, ok := resBody["title"].(string)
	if !ok || title == "" {
		c.JSON(400, gin.H{"error": "Title is required"})
		return
	}
	content, ok := resBody["content"].(string)
	if !ok || content == "" {
		c.JSON(400, gin.H{"error": "Content is required"})
		return
	}

	data, err := db.DB.Exec("INSERT INTO documents (title, content, created_at, updated_at) VALUES ($1, $2, NOW(), NOW())", title, content)

	if err != nil {
		c.JSON(500, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	println("Document created successfully:", data)

	id, err := data.LastInsertId()
	if err != nil {
		// If LastInsertId is not supported (e.g., PostgreSQL), you may need to query the ID separately
		c.JSON(200, gin.H{
			"message": "Document created successfully",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Document created successfully",
		"id":      id,
	})
}

func GetDocument(c *gin.Context) {

}

func UpdateDocument(c *gin.Context) {

}

func DeleteDocument(c *gin.Context) {

}
