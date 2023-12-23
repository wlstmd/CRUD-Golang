package types

type APIResponse struct {
	Result      int64  `json:"result"`
	Description string `json:"description"`
}

type CreateUserResponse struct {
	*UserResponse
}

type GetUserResponse struct {
	*UserResponse
}

type UpdateUserResponse struct {
	*UserResponse
}

type DeleteUserResponse struct {
	*UserResponse
}

func NewAPIResponse(description string, result int64) *APIResponse {
	return &APIResponse{
		Result:      result,
		Description: description,
	}
}
