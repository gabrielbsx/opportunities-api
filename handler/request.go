package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func errParamIsRequired(param string) error {
	return fmt.Errorf("%s is required", param)
}

func (r *CreateOpeningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("malformed request")
	}

	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("all fields are required")
	}

	if r.Role == "" {
		return errParamIsRequired("role")
	}

	if r.Company == "" {
		return errParamIsRequired("company")
	}

	if r.Location == "" {
		return errParamIsRequired("location")
	}

	if r.Link == "" {
		return errParamIsRequired("link")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary")
	}

	return nil
}
