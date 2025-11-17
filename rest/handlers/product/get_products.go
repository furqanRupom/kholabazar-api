package product

import (
	"kholabazar/utils"
	"net/http"
	"strconv"
	"sync"
)

var count int64

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageStr := reqQuery.Get("page")
	limitStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	productList, err := h.svc.List(page, limit)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to load product list")
		return
	}

	/*
		wg := WaitGroup {
		 noCopy : noCopy{},
		 state : atomic.uint64 {
		  _ : noCopy
		  _ : align64
		  v : uint64
		 },
		 sema : 0
		}

	*/
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		count1, _ := h.svc.Count()
		count = count1
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		count1, _ := h.svc.Count()
		count = count1
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		count1, _ := h.svc.Count()
		count = count1
	}()
	wg.Wait()

	utils.SendPage(w, productList, page, limit, count)
}
