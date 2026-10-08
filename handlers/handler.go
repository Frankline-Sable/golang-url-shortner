package handlers

import (
	"fmt"
	"golang-url-shortner/models"
	"golang-url-shortner/storage"
	"html/template"
	"net/http"
)

type Handler struct {
	store *storage.Store
}

func NewHandler(store *storage.Store) *Handler {
	return &Handler{store: store}
}

// ShortenUrl handles POST requests to shorten a URL
func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	// parse the longURL request body
	r.ParseForm()
	longUrl := r.FormValue("url")
	if longUrl == "" {
		http.Error(w, "miss url params", http.StatusBadRequest)
		return
	}
	// generate short url
	shortURL := models.GenerateShortURL()
	h.store.AddToStore(shortURL, longUrl)

	// responding
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://shorty.url/" + shortURL))
}

// RedirectURL handles GET request to redirect to the originalURL
func (h *Handler) RedirectURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// extract short url from the request
	shortURL := r.URL.Path[1:]
	fmt.Println("ShortURL: ", shortURL)

	// retrieve long url from the store
	longURL, exists := h.store.GetFromStore(shortURL)

	if !exists {

		http.Error(w, "t1: URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)

}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		println("whats happpening")
		http.NotFound(w, r)

		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

}
