package auth

type Create struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type PhoneRequest struct {
	PhoneNumber string `json:"phone_number"`
}
