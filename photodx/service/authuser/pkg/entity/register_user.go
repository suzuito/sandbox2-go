package entity

type RequestRegisterUserResultCode int

const (
	// UserCreationRequest was created
	RequestRegisterUserResultCodeCreated RequestRegisterUserResultCode = 1
	// Email already existed
	RequestRegisterUserResultCodeEmailAlreadyExisted RequestRegisterUserResultCode = 2
)
