package grade

type GradeRequest struct {
	StudentID string  `json:"studentId"`
	Homework  float64 `json:"homework"`
	Midterm   float64 `json:"midterm"`
	Final     float64 `json:"final"`
}

type GradeResponse struct {
	StudentID string  `json:"studentId"`
	Total     float64 `json:"total"`
	Grade     string  `json:"grade"`
}
