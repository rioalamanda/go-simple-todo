package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rioalamanda/go-simple-todo/model"
	"github.com/rioalamanda/go-simple-todo/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			var err error
			name := r.URL.Query().Get("name")
			data := []views.PostRequest{}
			if name != "" {
				data, err = model.ReadByName(name)
			} else {
				data, err = model.ReadAll()

			}
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}
