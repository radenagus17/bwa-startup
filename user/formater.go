package user

type Formater struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user User, token string)Formater{
	formatter := Formater{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: token,
	}

	return formatter
}