package grade

func CalculateGrade(homework, midterm, final float64) (float64, string) {
	total := homework*0.3 + midterm*0.3 + final*0.4

	var grade string
	if total >= 80 {
		grade = "A"
	} else if total >= 70 {
		grade = "B"
	} else if total >= 60 {
		grade = "C"
	} else if total >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	return total, grade
}
