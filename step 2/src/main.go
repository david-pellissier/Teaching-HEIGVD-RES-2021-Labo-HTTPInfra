package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Fortune struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var msg []Fortune

func homePage(w http.ResponseWriter, r *http.Request){
    //m := Fortune
    json.NewEncoder(w).Encode(msg)

}

func handleRequests() {
    http.HandleFunc("/", homePage)//.Methods("GET")
    log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
    msg =  []Fortune{
        Fortune{
            Title:   "In the future",
            Desc:    "You will die some day",
            Content: "Yes you will",
        },
        Fortune{
            Title:   "Today",
            Desc:    "You have eaten",
            Content: "Am I right ?",
        },
    }
    handleRequests()
}