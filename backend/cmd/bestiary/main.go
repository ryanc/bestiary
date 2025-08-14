package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path"
	"sort"
	"syscall"
)

const blobPath string = "/blob/"
const port int = 3000

var httpBlobPath string = path.Join("/api", blobPath)

func main() {
	mux := http.NewServeMux()

	dir := http.Dir("." + blobPath)
	fs := http.FileServer(dir)

	mux.Handle(httpBlobPath+"/", http.StripPrefix(httpBlobPath+"/", fs))
	mux.HandleFunc("GET /api/beasts", beastsHandler)

	go func() {
		slog.Info("starting server", "address", port)
		http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func beastsHandler(w http.ResponseWriter, r *http.Request) {
	var blobs []string

	files, err := os.ReadDir("." + blobPath)
	if err != nil {
		fmt.Printf("error: reading directory failed: %s", err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		ii, _ := files[i].Info()
		ji, _ := files[j].Info()
		return ii.ModTime().After(ji.ModTime())
	})

	for _, file := range files {
		path := path.Join(blobPath, file.Name())
		uri, _ := url.JoinPath("http://localhost:8080/api", path)
		blobs = append(blobs, uri)
	}

	body, err := json.Marshal(blobs)
	if err != nil {
		fmt.Printf("error: JSON marshal failed: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
