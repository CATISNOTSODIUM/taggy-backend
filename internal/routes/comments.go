package routes

import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	comments "github.com/CATISNOTSODIUM/taggy-backend/internal/handlers/comments"
)


func CommentRoutes(r chi.Router) {
	r.Post("/comments/create", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleCreate(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	})
	r.Post("/comments", func(w http.ResponseWriter, req *http.Request) {
			response, _ := comments.HandleRetrieve(w, req)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	})
	
}