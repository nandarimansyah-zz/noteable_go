package interfaces

import "net/http"

type IMemberController interface {
	GetAllMember(w http.ResponseWriter, r *http.Request)
}
