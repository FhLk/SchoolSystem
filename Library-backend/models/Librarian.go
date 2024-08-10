package models

type Librarian struct {
	Id         string `gorm:"column:Id;primaryKey" json:"id"`
	FirstName  string `gorm:"column:FirstN_name" json:"firstName"`
	MiddleName string `gorm:"column:Middle_name" json:"middleName"`
	LastName   string `gorm:"column:Last_name" json:"lastName"`
	Email      string `gorm:"column:Email" json:"email"`
	Role       string `gorm:"column:Role" json:"role"`
	Account    string `gorm:"column:Account_Id" json:"account"`
}

type Account struct {
	AccountId string `gorm:"column:Account_Id;primaryKey" json:"account_id"`
	Username  string `gorm:"column:Username" json:"username"`
	Password  string `gorm:"column:Password" json:"password"`
}
