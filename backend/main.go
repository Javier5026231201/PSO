package main

import (
	"mahasiswa_app/backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/login", func(c *gin.Context) {
		var u struct{ Username, Password string }
		c.BindJSON(&u)
		var id int
		err := config.DB.QueryRow("SELECT id FROM users WHERE username=$1 AND password=$2", u.Username, u.Password).Scan(&id)
		if err != nil {
			c.JSON(401, gin.H{"error": "Gagal"})
			return
		}
		c.JSON(200, gin.H{"id": id, "username": u.Username})
	})

	r.POST("/notes", func(c *gin.Context) {
		var n struct {
			UserID                         int
			Title, Content, Category, Date string
		}
		c.BindJSON(&n)
		_, err := config.DB.Exec("INSERT INTO notes (user_id, title, content, category, created_at) VALUES ($1, $2, $3, $4, $5)",
			n.UserID, n.Title, n.Content, n.Category, n.Date)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "saved"})
	})

	r.GET("/notes", func(c *gin.Context) {
		uid := c.Query("user_id")
		cat := c.Query("category")
		query := "SELECT title, category, created_at FROM notes WHERE user_id = $1"
		if cat != "" {
			query += " AND category = '" + cat + "'"
		}

		rows, _ := config.DB.Query(query, uid)
		var notes []interface{}
		for rows.Next() {
			var t, ct, dt string
			rows.Scan(&t, &ct, &dt)
			notes = append(notes, gin.H{"title": t, "category": ct, "date": dt})
		}
		c.JSON(200, notes)
	})

	r.Run(":8080")
}
