package models

type Book struct {
	BookId        string  `gorm:"column:Id;primary_key" json:"id"`
	BookImg       string  `gorm:"column:Img" json:"img"`
	BookTitle     string  `gorm:"column:Title" json:"title"`
	BookCategory  string  `gorm:"column:Category" json:"category"`
	BookAuthor    string  `gorm:"column:Author" json:"author"`
	BookPublisher string  `gorm:"column:Publisher" json:"publisher"`
	BookYear      int     `gorm:"column:Year" json:"year"`
	BookStatus    float64 `gorm:"column:Status" json:"status"`
}
