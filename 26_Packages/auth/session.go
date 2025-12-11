package auth

// private because of small letter
func extractSession() string {
	return "loggedin"
}

func GetSession() string {
	return extractSession()
}
