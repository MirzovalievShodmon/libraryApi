package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MirzovalievShodmon/libraryApi/services"
)

func GetBorrowRecordsHandler(w http.ResponseWriter, r *http.Request) {
	records, err := services.GetAllBorrowRecords()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}
