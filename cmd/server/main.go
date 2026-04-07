package main

import (
	"errors"
	"go-grading-api/config"
	"go-grading-api/internal/auth"
	"go-grading-api/internal/grade"
	"go-grading-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.InitDB()

	// Initialize repository
	gradeRepo := &grade.GradeRepository{}

	// Initialize service with repository
	gradeService := grade.NewGradeService(gradeRepo)

	// Initialize handlers with dependency injection
	gradeHandler := grade.NewHandler(gradeService)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", auth.LoginHandler)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile", func(c *gin.Context) {
				username := "Jaonine"
				role := "student admin"

				c.JSON(200, gin.H{
					"username": username,
					"role":     role,
				})
			})

			protected.POST("/grade/submit",
				middleware.RequireRole("instructor"),
				gradeHandler.SubmitGradeHandler)

			protected.GET("/grade/:studentId",
				gradeHandler.GetGradeHandler)
		}
	}

	err := r.Run(":3000")
	if err != nil {
		log.Fatal(errors.New("failed to run server"))
		return
	}
}
