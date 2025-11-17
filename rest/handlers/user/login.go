package user

import (
	"encoding/json"
	"fmt"
	"kholabazar/utils"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid data!")
		return
	}
	usr, err := h.svc.Find(req.Email, req.Password)
	fmt.Println(err)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "User find failed!")
	}
	if usr == nil {
		utils.SendError(w, http.StatusNotFound, "User not found!")
		return
	}
	accessToken, err := utils.CreateJWT(utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	}, h.conf.JWTSecret)

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server")
		return
	}
	utils.SendData(w, http.StatusOK, accessToken)

}
