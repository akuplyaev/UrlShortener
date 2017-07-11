package main
import (
	"github.com/gorilla/mux"
	"net/http"
	"UrlShortener/handlers"
)
func main()  {
	router :=mux.NewRouter()
	router.HandleFunc("/a", handlers.AddShortUrl)
	router.HandleFunc("/s/{shortUrl}", handlers.RedirectUrl)
	http.ListenAndServe(":8080",router)
}

