package core

// main core interface that provide behaviours
type Core interface {
	GetCompanies() ([]Company, error)
}

type CoreV1 struct{}

func (c *CoreV1) GetCompanies() ([]Company, error) {
	var companies []Company

	query := DB.Find(&companies)
	if query.Error != nil {
		return nil, query.Error
	}
	return companies, nil
}
