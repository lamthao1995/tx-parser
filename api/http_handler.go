package api

import (
	"encoding/json"
	"log"
	"net/http"
	"tx-parser/domain"
)

type Handler struct {
	parser domain.Parser
}

func NewHandler(parser domain.Parser) *Handler {
	return &Handler{parser: parser}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/currentBlock", h.getCurrentBlock)
	mux.HandleFunc("/subscribe", h.subscribe)
	mux.HandleFunc("/transactions", h.getTransactions)
}

func (h *Handler) getCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block, err := h.parser.GetCurrentBlock()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    map[string]int{"currentBlock": block},
	})
}

func (h *Handler) subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "address is required",
		})
		return
	}

	err := h.parser.Subscribe(address)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "Address subscribed successfully",
	})
}

func (h *Handler) getTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		respondWithJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Message: "address is required",
		})
		return
	}

	transactions, err := h.parser.GetTransactions(address)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    transactions,
	})
}

// respondWithJSON is a helper function to send JSON responses
func respondWithJSON(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error encoding response:", err)
	}
}
