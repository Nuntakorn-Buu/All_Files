package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const portNumber = ":5000" //ส่วนที่เพิ่มเข้ามา

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	handler := newHandler(db)

	r := gin.New()

	r.GET("/users", handler.listUsersHandler)
	r.POST("/users", handler.createUserHandler)
	r.DELETE("/users/:id", handler.deleteUserHandler)

	r.Run(portNumber)//ส่วนที่เพิ่มเข้ามา
	// r.Run()
}

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

type User struct {
	ID     string `json:"id"`
	Username  string `json:"username"`
	Email  string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) listUsersHandler(c *gin.Context) {
	var users []User

	if result := h.db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func (h *Handler) createUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func (h *Handler) deleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}