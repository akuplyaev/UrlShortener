package handlers
import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"UrlShortener/storage"
)
//http://localhost:8080/a?url=http%3A%2F%2Fgoogle.com%2F%3Fq%3Dgolang
//храним значения key=shortUrl=value=longUrl
func AddShortUrl (response http.ResponseWriter,request *http.Request){
	params:=request.URL.Query()
	longUrl:=params.Get("url")
	if longUrl == "" {  //придумать как проверить этот url на корекктность
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}
	shortUrl:=storage.FindShortUrl(longUrl)
	response.WriteHeader(http.StatusOK)
	fmt.Fprintf(response,"%s",shortUrl)
}
func RedirectUrl(responce http.ResponseWriter,request *http.Request) {
	vars := mux.Vars(request)
	shortUrl := vars["shortUrl"]
    longUrl:=storage.FindLongUrl(shortUrl)
	if longUrl==""{
          http.Error(responce, "Bad Request ",http.StatusBadRequest)
	}
	http.RedirectHandler(longUrl,302)
}