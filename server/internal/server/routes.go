package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/markbates/goth/gothic"
	"github.com/mathletedev/decosavvy/internal/config"
)

func (s *Server) RegisterRoutes(allowedOrigins []string) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
	}))

	r.Get("/api/hello", s.HandleHello)
	r.Get("/api/me", s.HandleMe)
	r.Get("/api/products", s.HandleProducts)
	r.Get("/api/cart", s.HandleCart)
	r.Post("/api/add-to-cart", s.HandleAddToCart)
	r.Post("/api/remove-from-cart", s.HandleRemoveFromCart)

	r.Get("/auth/{provider}", s.HandleAuth)
	r.Get("/auth/{provider}/callback", s.HandleAuthCallback)
	r.Get("/signout/{provider}", s.HandleSignout)

	return r
}

func (s *Server) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func (s *Server) HandleMe(w http.ResponseWriter, r *http.Request) {
	id, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := s.db.ReadUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) HandleProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.db.ReadProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (s *Server) HandleCart(w http.ResponseWriter, r *http.Request) {
	id, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = s.db.ReadUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	products, err := s.db.ReadCart(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (s *Server) HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	id, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = s.db.ReadUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.db.AddToCart(id, string(body))
}

func (s *Server) HandleRemoveFromCart(w http.ResponseWriter, r *http.Request) {
	id, err := gothic.GetFromSession("user", r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, err = s.db.ReadUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s.db.RemoveFromCart(id, string(body))
}

func (s *Server) HandleAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}

func (s *Server) HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, r)
		return
	}

	rows, err := s.db.Query(context.Background(), "SELECT id FROM users WHERE email=$1;", user.Email)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, r)
		return
	}

	defer rows.Close()

	var id string
	if rows.Next() {
		rows.Scan(&id)
	} else {
		// create user if one doesn't exist
		id, err = s.db.CreateUser(user.Email, user.AvatarURL)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, r)
			return
		}
	}

	// set user id in session
	gothic.StoreInSession("user", id, r, w)

	http.Redirect(w, r, config.WebUrl, http.StatusFound)
}

func (s *Server) HandleSignout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", config.WebUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
