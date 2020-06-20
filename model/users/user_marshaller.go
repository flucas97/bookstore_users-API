package users

import "encoding/json"

type PublicUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    string `json:"status"`
}

type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    string `json:"status"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:        user.ID,
			FirstName: user.FirstName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Status:    user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		return nil
	}

	return privateUser
}
