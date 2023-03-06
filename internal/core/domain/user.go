package domain

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type UserProfile struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
	School    string `json:"school"`
}

func (u *UserProfile) Pointers() []interface{} {
	return []interface{}{
		&u.ID,
		&u.UserName,
		&u.FirstName,
		&u.LastName,
		&u.Phone,
		&u.Address,
		&u.City,
		&u.School,
	}
}
