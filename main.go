package main

import (
       "fmt"
       "net/http"
       "html/template"
)

var templates *template.Template
var homeTemplate *template.Template

func main(){
  // http.HandleFunc("/",Homefunc)
  http.HandleFunc("/login ",Loginfunc)
  http.ListenAndServe(":8086",nil)
}

//HandleFunc Handles / url and ask user name
func Homefunc(w http.ResponseWriter, r *http.Request){
  if r.Method == "GET"{
    homeTemplate.Execute(w,nil)
  }
}

//HANDLE THE /Loginfunc and shows the profile of the logged in user on a GET Request
//HANDLE LOGIC PROCESS ON POST REQUEST
func Loginfunc(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w,"This is the profile page ")
}

// func PopulateTemplate()  {
//   templates,err:=tsemplate.ParseGlob("*./templates/*.html")
//    if err != nil {
//      fmt.Println("main.go:PopulateTemplate:",err)
//      os.Exit(1)
//    }
// homeTemplate = templates.Lookup("tasks.html")
// }
