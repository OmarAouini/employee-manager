package core

type Tabler interface {
	TableName() string
}

type Company struct {
}

func (Company) TableName() string {
	return "companies"
}
