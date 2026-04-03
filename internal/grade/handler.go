package grade

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateGradeHandler(c *gin.Context) {
	var req GradeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	total, grade := CalculateGrade(req.Homework, req.Midterm, req.Final)

	res := GradeResponse{
		StudentID: req.StudentID,
		Total:     total,
		Grade:     grade,
	}

	c.JSON(http.StatusOK, res)
}

func SubmitGradeHandler(c *gin.Context) {
	var req GradeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	total, gradeLetter := CalculateGrade(req.Homework, req.Midterm, req.Final)

	res := GradeResponse{
		StudentID: req.StudentID,
		Total:     total,
		Grade:     gradeLetter,
	}

	err := InsertGrade(res, req.Homework, req.Midterm, req.Final)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to save grade"})
		return
	}

	c.JSON(200, res)
}

func GetGradeHandler(c *gin.Context) {
	studentID := c.Param("studentId")

	grade, err := GetGradeByStudentID(studentID)
	if err != nil {
		c.JSON(404, gin.H{"error": "grade not found"})
		return
	}

	c.JSON(200, grade)
}
