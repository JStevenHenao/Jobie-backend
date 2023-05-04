package dtoRequest

type CompanyServiceRequest struct {
	IdCompany  int64   `json:"id"`
	IdServices []int64 `json:"idServices"`
}

func (c *CompanyServiceRequest) GetIdCompany() int64 {
	return c.IdCompany
}

func (c *CompanyServiceRequest) GetIdServices() []int64 {
	return c.IdServices
}
