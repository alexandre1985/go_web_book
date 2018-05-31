package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

// Store notes
var noteStore = make(map[string]Note)

var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	// decode from json to note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	id++
	note.CreatedOn = time.Now()

	id_s := strconv.Itoa(id)

	noteStore[id_s] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var arrayNotes []Note

	for _, note := range noteStore {
		arrayNotes = append(arrayNotes, note)
	}

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(arrayNotes)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	target_id := vars["id"]
	if target, ok := noteStore[target_id]; ok {
		var newNote Note
		// decode from json to note
		err := json.NewDecoder(r.Body).Decode(&newNote)
		if err != nil {
			panic(err)
		}

		newNote.CreatedOn = target.CreatedOn
		delete(noteStore, target_id)
		noteStore[target_id] = newNote
	} else {
		log.Printf("Could not find key of Note %s to delete", target_id)
	}
	w.WriteHeader(http.StatusNoContent)
}

// HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	target_id := vars["id"]
	if _, ok := noteStore[target_id]; ok {
		delete(noteStore, target_id)
	} else {
		log.Printf("Could not find key of Note %s to delete", target_id)	
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
