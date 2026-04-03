package auth

// GetUserByUsername returns a user by username, mocks a users account for demo purposes actually fetches from a database
func GetUserByUsername(username string) (*User, error) {
	// Role is instructor
	if username == "KrittiyaB" {
		return &User{ID: 1, Username: "KrittiyaB", Password: "1234", Role: "instructor"}, nil
	}

	// Role is student
	if username == "John" {
		return &User{ID: 2, Username: "John", Password: "1234", Role: "student"}, nil
	}

	return nil, nil
}
