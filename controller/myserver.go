package main

import (
	// "crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	// "strconv"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/sxs123/transFile/model"
	"os"
	"text/template"
	"time"
)

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%q", "See you again ,Golang!")
}
func file_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("chick chick guguda")
		// crutime := time.Now().Unix()
		// h := md5.New()
		// io.WriteString(h, strconv.FormatInt(crutime, 10))
		// token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("../views/index.html")
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "I'm sorry ,but something error happened.")
			return
		}
		f, err := os.Create("../files/" + handler.Filename)
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "I'm sorry ,but something error happened.")
			return
		}
		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			fmt.Fprint(w, "I'm sorry ,but something error happened.")
		} else {
			f := model.File{1, handler.Filename, "http://localhost:3000/files/" + handler.Filename, time.Now().String()}
			f.Create()
			fmt.Println("redirect to host:3000/files")
			http.Redirect(w, r, "/files_list", http.StatusFound)
		}
	}
}

func list_handler(w http.ResponseWriter, r *http.Request) {
	objs, err := model.Get()
	if err != nil {
		log.Fatalf("error happened in list_handler:%v", err)
		http.Redirect(w, r, "/404", http.StatusFound)
	}
	t, _ := template.ParseFiles("../views/files.tpl")
	if err := t.Execute(w, objs); err != nil {
		panic(err)
	}

}

func error_handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("../views/404.tpl")
	t.Execute(w, nil)
}

func main() {
	fmt.Println("Start Serving ...")
	// model.Init_db()
	mux := http.NewServeMux()
	mux.HandleFunc("/", root_handler)
	mux.HandleFunc("/create_file", file_handler)
	mux.HandleFunc("/404", error_handler)
	// http.HandleFunc("/files_list", list_handler)
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("../files"))))
	mux.HandleFunc("/files_list", list_handler)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
