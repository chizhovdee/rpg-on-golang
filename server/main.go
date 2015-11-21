package main
import (
	"net/http"
	"html/template"
	"path/filepath"
)

func main(){
	mux := http.NewServeMux()

	// статичные файлы
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles(filepath.Join("views", "index.html")))

	templ.Execute(w, nil)
}

