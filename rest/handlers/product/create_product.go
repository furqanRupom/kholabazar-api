package product

import (
	"encoding/json"
	"fmt"
	"kholabazar/domain"
	"kholabazar/utils"
	"net/http"
)

type ReqProduct struct {
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Please give me valid JSON")
		return
	}
	createProduct, err := h.svc.Create(domain.Product{
		Name:        req.Name,
		Image:       req.Image,
		Price:       req.Price,
		Description: req.Description,
		Category:    req.Category,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Production creation failed!")
		return
	}
	utils.SendData(w,http.StatusCreated,createProduct)

}
