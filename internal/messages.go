package messages

const (
	InvalidRequest  = "Invalid request payload."
	SomethingErrors = "Something errors."

	// register user
	RegisterSuccessfully = "User registered successfully."

	// login
	LoginFailedInternal       = "An unexpected error occurred. Please try again later."
	LoginIncorrectCredentials = "Login failed. Please check your credentials."
	LoginSuccesfully          = "Login successfully."
)
