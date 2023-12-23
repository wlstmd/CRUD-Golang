package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type UserResponse struct {
	*APIResponse
	*User
}
