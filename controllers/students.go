package controllers

import (
	"example/crud-api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateStudentInput struct {
	LastName  string  `json:"firstName"`
	FirstName string  `json:"lastName"`
	Average   float64 `json:"average"`
}

type UpdateStudentInput struct {
	LastName  string  `json:"firstName"`
	FirstName string  `json:"lastName"`
	Average   float64 `json:"average"`
}

// Get all students
func GetAllStudents(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var students []models.Student
	db.Find(&students)
	c.JSON(http.StatusOK, gin.H{"data": students})
}

// Create new student
func CreateStudent(c *gin.Context) {
	// Validating student input
	var input CreateStudentInput

	// catch error
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create student
	student := models.Student{FirstName: input.FirstName, LastName: input.LastName, Average: input.Average}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

// Find student by id
func FindStudentById(c *gin.Context) {
	var student models.Student

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// update Student by id
func UpdateStudentById(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var student models.Student
	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Student
	updatedInput.LastName = input.FirstName
	updatedInput.FirstName = input.LastName
	updatedInput.Average = input.Average

	db.Model(&student).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// delete student by id
func DeleteStudentById(c *gin.Context) {
	// Get student by id if exist
	db := c.MustGet("db").(*gorm.DB)
	var student models.Student

	if err := db.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"data": "Student was deleted successfully"})
}
