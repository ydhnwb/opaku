package _user

import "github.com/ydhnwb/opaku-dummy-backend/entity"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

func NewUserResponse(user entity.User) UserResponse {
	print("kcal")
	print(user.ID)
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}
