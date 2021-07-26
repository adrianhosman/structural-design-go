package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adrianhosman/structural-design-go/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseError(statusCode, errors.Cause(err)))
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.BuildAPIResponseSuccess(statusCode, data))
}

func (h *Handler) CalculateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	businessID := vars["business_id"]
	if businessID == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("business_id must provided"))
		return
	}
	result, err := h.usecase.CalculateInvoiceData(businessID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	writeSuccessResponse(w, result)
}
