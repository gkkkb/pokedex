package request

import "github.com/google/uuid"

func CreateRequestID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return uid.String()
}
