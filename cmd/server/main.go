package main

import (
	"errors"
	"go-grading-api/internal/auth"
	"go-grading-api/internal/database"
	"go-grading-api/internal/grade"
	"go-grading-api/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", auth.LoginHandler)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// TODO: uncomment this when we want to see user profile
			protected.GET("/profile", func(c *gin.Context) {
				username, _ := c.Get("username")
				role, _ := c.Get("role")

				c.JSON(200, gin.H{
					"username": username,
					"role":     role,
				})
			})

			protected.POST("/grade/calculate",
				middleware.RequireRole("instructor"),
				grade.CalculateGradeHandler)

			protected.POST("/grade/submit",
				middleware.RequireRole("instructor"),
				grade.SubmitGradeHandler)

			protected.GET("/grade/:studentId",
				grade.GetGradeHandler)
		}
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(errors.New("failed to run server"))
		return
	}
}
