package api

import (
	"fmt"
	"net/http"
)

type Response struct {
	Code 	int
	Message string
	Data 	any
}

func DeleteResponse(id string) Response {
	return Response{
		Code: http.StatusOK,
		Message: fmt.Sprintf("deleted %s", id),
	}
}



