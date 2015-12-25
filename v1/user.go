package herohunt

type User struct {
	ID        string
	Email     string
	ScreeName string
	Password  string
}

func Login(ID, Password string) (*LoginSession, error) {
	return nil, nil
}
