package prompthandler

import (
	"chatgpt-challenge/delivery/http_server/http_io"
	"net/http"
)

// GetAll godoc
// @Summary 	 Get a list of all prompts
// @Tags         Prompts
// @Produce      json
// @Success      200  {object} http_io.Envelope{data=[]promptparam.GetAllResponse,error=nil}
// @Router       /prompts [get].
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res := h.promptSvc.GetAll()

	http_io.WriteJSON(w, http.StatusOK, http_io.Envelope{Data: res}, nil)
}
