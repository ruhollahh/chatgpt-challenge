package laptophandler

import (
	"chatgpt-challenge/delivery/http_server/http_io"
	"net/http"
)

// GetAll godoc
// @Summary 	 Get a list of all laptops
// @Tags         Laptops
// @Produce      json
// @Success      200  {object} http_io.Envelope{data=[]laptopparam.GetAllResponse,error=nil}
// @Router       /laptops [get].
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.laptopSvc.GetAll()

	http_io.WriteJSON(w, http.StatusOK, http_io.Envelope{Data: res}, nil)
}
