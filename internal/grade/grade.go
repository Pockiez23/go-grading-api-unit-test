// grade.go สำหรับการทำ Unit Test และ Mockdata เพื่อใช้ในการทดสอบ
package grade

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculateGrade_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		homework float64
		midterm  float64
		final    float64
		expected string
	}{
		{"Grade A", 80, 70, 90, "A"},
		{"Grade B", 70, 70, 75, "B"},
		{"Grade C", 60, 60, 65, "C"},
		{"Grade D", 50, 50, 55, "D"},
		{"Grade F", 40, 40, 45, "F"},
		{"Invalid Negative Score", -10, 70, 90, "Invalid"},
		{"Invalid Out of Range Score", 110, 70, 90, "Invalid"},
		{"Boundary Grade A (80)", 80, 80, 80, "A"},
		{"Boundary Grade B (70)", 70, 70, 70, "B"},
		{"Boundary Grade C (60)", 60, 60, 60, "C"},
		{"Boundary Grade D (50)", 50, 50, 50, "D"},
		{"Boundary Grade F (49)", 49, 49, 49, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, grade := CalculateGrade(tt.homework, tt.midterm, tt.final)
			assert.Equal(t, tt.expected, grade)
		})
	}
}

type MockRepository struct{}

func (m *MockRepository) GetGradeByStudentID(studentID string) (*Response, error) {
	return &Response{
		StudentID: studentID,
		Total:     85,
		Grade:     "A",
	}, nil
}

func (m *MockRepository) InsertGrade(g Response, homework, midterm, final float64) error {
	return nil
}

func TestCheckGrade(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewGradeService(mockRepo)

	res, err := service.CheckGrade("65001")

	assert.NoError(t, err)
	assert.Equal(t, "65001", res.StudentID)
	assert.Equal(t, "A", res.Grade)
}

// MockService simulates business logic for handler tests
type MockService struct{}

func (m *MockService) CheckGrade(studentID string) (*Response, error) {
	return &Response{
		StudentID: studentID,
		Total:     90,
		Grade:     "A",
	}, nil
}

func (m *MockService) SubmitGrade(req Request) (*Response, error) {
	return &Response{
		StudentID: req.StudentID,
		Total:     85,
		Grade:     "A",
	}, nil
}

func TestGetGradeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{}
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/grade/:studentId", handler.GetGradeHandler)

	req, _ := http.NewRequest("GET", "/grade/65001", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
