package user_request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=150"`
	Age      int8   `json:"age" binding:"required,min=4,max=140"`
}

type UserUpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"omitempty,min=4,max=150"`
	Age  int8   `json:"age" binding:"omitempty,min=4,max=140"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
