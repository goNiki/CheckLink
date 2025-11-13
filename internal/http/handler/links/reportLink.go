package links

import (
	"goNiki/CheckLink/internal/dto"
	"goNiki/CheckLink/internal/infrastructure/logger/sl"
	"goNiki/CheckLink/pkg/errorsAPP"
	"io"
	"net/http"

	"github.com/go-chi/render"
)

func (h *handler) GetReportLinks(w http.ResponseWriter, r *http.Request) {

	var req dto.ReqGetReportLinks

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		h.log.Error("Decode Error", sl.Error(err))
		render.Status(r, http.StatusBadRequest)
		return
	}

	if len(req.LinksList) == 0 {
		h.log.Error("lists_link = 0", sl.Error(errorsAPP.ErrInvalidValidation))
		render.Status(r, http.StatusBadRequest)
		return
	}

	reader, err := h.report.CreateReport(r.Context(), req.LinksList)
	if err != nil {
		h.log.Error("Error create Report", sl.Error(err))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="report.pdf"`)
	render.Status(r, http.StatusOK)

	_, err = io.Copy(w, reader)
	if err != nil {
		h.log.Error("Error writing response", sl.Error(err))
		return
	}
}
