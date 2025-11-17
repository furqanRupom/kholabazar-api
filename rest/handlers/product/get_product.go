package product

import (
	"kholabazar/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	pId, err := strconv.Atoi(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Please give me valid JSON")
		return
	}

	product,err := h.svc.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Something went wrong!")
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found!")
		return
	}
	utils.SendData(w,http.StatusOK,product)
}
