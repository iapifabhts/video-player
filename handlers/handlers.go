package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type item struct {
	Poster string `json:"poster"`
	Video  string `json:"video"`
}

var items = []item{}

var fileServer = http.FileServer(http.Dir(os.TempDir()))

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		
		if r.Method == http.MethodOptions {
			return
		}
		
		next(w, r)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
	fileServer.ServeHTTP(w, r)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(items)
	w.Write(response)
}

func UploadItem(w http.ResponseWriter, r *http.Request) {
	var item item
	json.NewDecoder(r.Body).Decode(&item)
	r.Body.Close()
	items = append(items, item)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	uploadedFile, headers, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	ext := strings.Split(headers.Filename, ".")[1]
	name := genString()
	
	if err = os.Mkdir(fmt.Sprintf("%s/%s", os.TempDir(), name), 0660); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	var file *os.File
	if file, err = os.Create(
		fmt.Sprintf("%s/%s/%s.%s", os.TempDir(), name, name, ext),
	); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if _, err = file.ReadFrom(uploadedFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if fileType[ext] == "video" {
		go covert(name, name+"."+ext)
		w.Write([]byte(fmt.Sprintf("/%s/.m3u8", name)))
		return
	}
	
	w.Write([]byte(fmt.Sprintf("/%s/%s.%s", name, name, ext)))
}

var fileType = map[string]string{
	"png": "image",
	"jpg": "image",
	"mp4": "video",
}

func covert(filepath, filename string) {
	if err := exec.Command(
		"ffmpeg", "-i",
		fmt.Sprintf("%s/%s/%s", os.TempDir(), filepath, filename),
		"-g", "60", "-hls_time", "60",
		fmt.Sprintf("%s/%s/.m3u8", os.TempDir(), filepath),
	).Run(); err != nil {
		log.Println(err)
	}
	log.Println("end")
}

func genString() (str string) {
	rand.Seed(time.Now().UnixNano())
	chars := "qwertyuiopasdfghjklzxcvbnm0123456789"
	for i := 0; i < 32; i++ {
		str += string(chars[rand.Intn(len(chars))])
	}
	return
}
