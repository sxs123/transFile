package main

import (
	// "crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/astaxie/beego/session"
	"github.com/sxs123/transFile/current/model"
	"os"
	"text/template"
	"time"
)

// var globalSessions *session.Manager

// func init() {
// 	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
// 	go globalSessions.GC()
// }

type Result struct {
	Code   int
	Data   []interface{}
	Status string
}

func validate_login(w http.ResponseWriter, r *http.Request) {
	// uname, _ := r.Cookie("username")
	tt, err := r.Cookie("transfile")
	if tt == nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if tt != nil {
		if tt.Value != "babysqualyou,sasexymonthfuck" || err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
}

func login_handler(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.tpl")
		t.Execute(w, nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		pwd := r.FormValue("password")
		if email == "xuesong_smile@163.com" && pwd == "weshallgo" {
			// fmt.Println("username:", r.Form["username"])
			// sess.Set("username", r.Form["username"])
			// fmt.Println("password:", r.Form["password"])
			COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second // 单位：秒。
			maxAge := int(COOKIE_MAX_MAX_AGE)
			coo := &http.Cookie{Name: "transfile", MaxAge: maxAge, Value: "babysqualyou,sasexymonthfuck", Path: "/", HttpOnly: false}
			http.SetCookie(w, coo)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	t, _ := template.ParseFiles("views/index.tpl")
	t.Execute(w, nil)
}
func file_handler(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/create_file.html")
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
		f, err := os.Create("./files/" + handler.Filename)
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "I'm sorry ,but something error happened.")
			return
		}
		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			fmt.Fprint(w, "I'm sorry ,but something error happened.")
		} else {
			f := model.File{1, handler.Filename, "/files/" + handler.Filename, time.Now().String()}
			f.Create()
			fmt.Println("redirect to host:3000/files")
			http.Redirect(w, r, "/files_list", http.StatusFound)
		}
	}
}

func list_handler(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	objs, err := model.Get()
	if err != nil {
		log.Fatalf("error happened in list_handler:%v", err)
		http.Redirect(w, r, "/404", http.StatusFound)
	}
	t, _ := template.ParseFiles("views/files.tpl")
	if err := t.Execute(w, objs); err != nil {
		panic(err)
	}

}

func error_handler(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	t, _ := template.ParseFiles("../views/404.tpl")
	t.Execute(w, nil)
}

func remove_file(w http.ResponseWriter, r *http.Request) {
	validate_login(w, r)
	r.ParseForm()
	file := r.FormValue("file")
	fileId := r.FormValue("fileId")
	id, _ := strconv.Atoi(fileId)
	os.Remove("./files/" + file)
	ret := Result{}
	if err := model.DestroyOne(id); err != nil {
		ret.Code = 500
		ret.Status = "Destroy error!"
		rr, _ := json.Marshal(ret)
		w.Header().Set("Content-Type", "application/json")
		w.Write(rr)
	}
	ret.Code = 200
	ret.Status = "Destroy success!"
	rr, _ := json.Marshal(ret)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rr)
}

func main() {
	fmt.Println("Start Serving ...")
	// model.Init_db()
	mux := http.NewServeMux()
	mux.HandleFunc("/", root_handler)
	mux.HandleFunc("/login", login_handler)
	mux.HandleFunc("/create_file", file_handler)
	mux.HandleFunc("/404", error_handler)
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files"))))
	mux.HandleFunc("/files_list", list_handler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/file/delete", remove_file)
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
