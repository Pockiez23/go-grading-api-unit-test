package student

import "go-grading-api/internal/database"

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
