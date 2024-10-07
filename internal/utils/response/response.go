package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	// return Response{
	// 	Status: StatusError,
	// 	Error:  errs.Error(),
	// }

	var errMasgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMasgs = append(errMasgs, err.Field()+" is required")
		case "gte":
			errMasgs = append(errMasgs, err.Field()+" must be greater than or equal to "+err.Param())
		case "lte":
			errMasgs = append(errMasgs, err.Field()+" must be less than or equal to "+err.Param())

		}
	}

	return Response{
		Status: StatusError,
		Error:  errMasgs[0],
	}

}
