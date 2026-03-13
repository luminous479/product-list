package database


type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"pass"`
	IsShopWoner bool   `json:"woner"`
}

var users []User

func CreateUser(u User) User {
	if u.ID != 0 {
		return u
	}
	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func GetUserByEmail(email string) *User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}	
	}
	return nil
}