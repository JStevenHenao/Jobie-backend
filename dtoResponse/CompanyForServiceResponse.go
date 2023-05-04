package dtoresponse

import "tidy/models"

type CompanyForServiceResponse struct {
	Companies []models.Company `json:"companies"`
}
