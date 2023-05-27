package userModel

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{
		ID:    "1",
		Name:  "John Doe",
		Email: "john@example.com",
	},
	{
		ID:    "2",
		Name:  "Jane Smith",
		Email: "jane@example.com",
	},
}

func GetUsers() []User {
	return users
}

func GetUserByID(id string) *User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

