package http

import "time"

type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Value   interface{} `json:"value"`
	}

	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	User struct {
		Id              string    `json:"id"`
		FirstName       string    `json:"first_name"`
		LastName        string    `json:"last_name"`
		Email           string    `json:"email"`
		BirthDay        time.Time `json:"birthday"`
		Role            string    `json:"role"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		IsEmailVerified bool      `json:"is_email_verified"`
	}

	Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8,lte=104"`
	}

	AddUser struct {
		FirstName string    `json:"first_name" binding:"required"`
		LastName  string    `json:"last_name" binding:"required"`
		Email     string    `json:"email" binding:"required,email"`
		Password  string    `json:"password" binding:"required,gte=8,lte=104"`
		BirthDay  time.Time `json:"birthday" binding:"required"`
	}

	UpdateUser struct {
		FirstName string    `json:"first_name" binding:"required"`
		LastName  string    `json:"last_name" binding:"required"`
		BirthDay  time.Time `json:"birthday" binding:"required"`
	}

	UpdatePassword struct {
		OldPassword string `json:"old_password" binding:"required,gte=8,lte=104"`
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=104"`
	}

	ForgotPassword struct {
		Email string `json:"email" binding:"required,email"`
	}

	ResetPassword struct {
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=104"`
	}
)
