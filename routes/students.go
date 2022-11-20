package routes

import (
	"example/crud-api-go/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupStudentsRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.GET("/students", controllers.GetAllStudents)
	router.POST("/students", controllers.CreateStudent)
	router.GET("/students/:id", controllers.FindStudentById)
	router.PATCH("/students/:id", controllers.UpdateStudentById)
	router.DELETE("/students/:id", controllers.DeleteStudentById)
	return router
}
