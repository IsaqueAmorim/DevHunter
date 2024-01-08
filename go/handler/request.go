package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"romote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParameterRequired("Role", "string")
	}
	if r.Company == "" {
		return errParameterRequired("Company", "string")
	}
	if r.Location == "" {
		return errParameterRequired("Location", "string")
	}
	if r.Link == "" {
		return errParameterRequired("Link", "string")
	}
	if r.Remote == nil {
		return errParameterRequired("Remote", "bool")
	}
	if r.Salary == 0.0 {
		return errParameterRequired("Salary", "float64")
	}

	return nil
}

func errParameterRequired(name, typ string) error {
	return fmt.Errorf("Params: %s (type: %s) is required", name, typ)
}
