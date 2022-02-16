package errors

import "fmt"

type APIError struct {
	Service string
	Detail  string
	Code    string
}

func (A APIError) Error() string {
	s := fmt.Sprintf("code=%s details=%s service=%s", A.Code, A.Detail, A.Service)
	return s
}
