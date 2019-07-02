package controllers

import (
	"net/http"

	"github.com/nandarimansyah/noteable_go/interfaces"

	"github.com/gorilla/mux"
	"github.com/nandarimansyah/noteable_go/helpers"
)

type MemberController struct {
	Service interfaces.IMemberService
}

func (c *MemberController) GetAllMember(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orgName := vars["name"]

	result := c.Service.GetAllMember(orgName)

	helpers.APIResponse(w, http.StatusOK, result)
	return
}
