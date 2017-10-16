package repository

type AdminID int

type Admin struct {
	ID          AdminID
	Mail        string
	Hashed      string
	DisplayName string
}

type AdminRepository interface {
	FindByMail(mail string) (*Admin, error)
	FindByID(ID AdminID) (*Admin, error)
	Create(mail string, password string) (*Admin, error)
}
