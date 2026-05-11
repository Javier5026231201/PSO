package main

import (
	"log"
	"net/http"

	"mahasiswa_app/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Database
	config.ConnectDB()

	r := gin.Default()

	// JURUS PAMUNGKAS: Panggil file frontend satu per satu agar tidak bentrok dengan API
	r.StaticFile("/", "../frontend/index.html")
	r.StaticFile("/index.html", "../frontend/index.html")
	r.StaticFile("/dashboard.html", "../frontend/dashboard.html")
	r.StaticFile("/script.js", "../frontend/script.js")
	r.StaticFile("/style.css", "../frontend/style.css")

	

	// --- MIDDLEWARE CORS MANUAL (CARA PALING AMPUH) ---
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Jika browser melakukan "Preflight Request" (OPTIONS), langsung beri status OK (204)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// --- LOGIN ---
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Format salah"})
			return
		}

		var id int
		err := config.DB.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", req.Username, req.Password).Scan(&id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Username atau password salah"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login berhasil"})
	})

	// --- NOTES ---
	r.GET("/notes", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, title, content FROM notes")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal ambil data"})
			return
		}
		defer rows.Close()

		var notes []gin.H
		for rows.Next() {
			var id int
			var title, content string
			rows.Scan(&id, &title, &content)
			notes = append(notes, gin.H{"id": id, "title": title, "content": content})
		}
		if notes == nil {
			notes = []gin.H{}
		}
		c.JSON(http.StatusOK, notes)
	})

	r.POST("/notes", func(c *gin.Context) {
		var data struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		c.BindJSON(&data)
		_, err := config.DB.Exec("INSERT INTO notes(title, content) VALUES(?,?)", data.Title, data.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal simpan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Berhasil"})
	})

	r.DELETE("/notes/:id", func(c *gin.Context) {
		id := c.Param("id")
		config.DB.Exec("DELETE FROM notes WHERE id=?", id)
		c.JSON(http.StatusOK, gin.H{"message": "Terhapus"})
	})

	// --- SCHEDULES ---
	r.GET("/schedules", func(c *gin.Context) {
		rows, err := config.DB.Query("SELECT id, course, day, time, room FROM schedules")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal ambil data"})
			return
		}
		defer rows.Close()

		var schedules []gin.H
		for rows.Next() {
			var id int
			var course, day, time, room string
			rows.Scan(&id, &course, &day, &time, &room)
			schedules = append(schedules, gin.H{"id": id, "course": course, "day": day, "time": time, "room": room})
		}
		if schedules == nil {
			schedules = []gin.H{}
		}
		c.JSON(http.StatusOK, schedules)
	})

	r.POST("/schedules", func(c *gin.Context) {
		var data struct {
			Course string `json:"course"`
			Day    string `json:"day"`
			Time   string `json:"time"`
			Room   string `json:"room"`
		}
		c.BindJSON(&data)
		_, err := config.DB.Exec("INSERT INTO schedules(course, day, time, room) VALUES(?,?,?,?)", data.Course, data.Day, data.Time, data.Room)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal simpan"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Berhasil"})
	})

	r.DELETE("/schedules/:id", func(c *gin.Context) {
		id := c.Param("id")
		config.DB.Exec("DELETE FROM schedules WHERE id=?", id)
		c.JSON(http.StatusOK, gin.H{"message": "Terhapus"})
	})

	log.Println("Server running on port 8080")
	r.Run(":8080")
}