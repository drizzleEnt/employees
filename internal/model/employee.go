package model

type Employee struct {
	Id          int         `json:"-" db:"id"`
	Name        string      `json:"name" db:"name"`
	Surname     string      `json:"surname" db:"surname"`
	Phone       string      `json:"phone" db:"phone"`
	CompanyId   int         `json:"companyId" db:"company"`
	Passport    Passport    `db:""`
	Departament Departament `db:""`
}

type Passport struct {
	Type   string `json:"type" db:"type"`
	Number string `json:"number" db:"number"`
}

type Departament struct {
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
}
