package grade

import (
	"go-grading-api/internal/database"
)

func InsertGrade(g GradeResponse, homework, midterm, final float64) error {
	query := `
	INSERT INTO grades (student_id, homework, midterm, final, total, grade)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := database.DB.Exec(
		query,
		g.StudentID,
		homework,
		midterm,
		final,
		g.Total,
		g.Grade,
	)

	return err
}

func GetGradeByStudentID(studentID string) (*GradeResponse, error) {
	query := `
	SELECT student_id, total, grade
	FROM grades
	WHERE student_id = ?
	ORDER BY id DESC
	LIMIT 1
	`

	row := database.DB.QueryRow(query, studentID)

	var res GradeResponse
	err := row.Scan(&res.StudentID, &res.Total, &res.Grade)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
