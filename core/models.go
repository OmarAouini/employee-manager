package core

import "time"

type Tabler interface {
	TableName() string
}

type Company struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	VatCode     string    `json:"vat_code"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Company) TableName() string {
	return "companies"
}

type Project struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Expenses  float32   `json:"expenses"`
	Incomes   float32   `json:"incomes"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Project) TableName() string {
	return "projects"
}

type Employee struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Age           int       `json:"age"`
	BirthDate     time.Time `json:"birth_date"`
	TaxCode       string    `json:"tax_code"`
	Qualification string    `json:"qualification"`
	Salary        float32   `json:"salary"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employees"
}
