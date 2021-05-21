package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "strings"
    "time"
)

type Fortune struct {
    Title string `json:"Title"`
    Content string `json:"content"`
}

var msg []Fortune

// Code repris de https://github.com/bmc/fortune-go
// Pour les parties de lecture et parsing du fichier fortunes

// Given a path representing a fortune file, load the file, parse it,
// an return an array of fortune strings.
func readFortuneFile(fortuneFile string) ([]string, error) {
    content, err := ioutil.ReadFile(fortuneFile)
    var fortunes []string = nil
    if err == nil {
        fortunes = strings.Split(string(content), "\n%\n")
    }
    return fortunes, err
}

// Given a path representing a fortune file, load the file, parse it,
// choose a random fortune, and display it to standard output. Returns
// a Go error object on error or nil on success.
func findAndPrint(fortuneFile string) error {
    msg = make([]Fortune, 0)
    fortunes, err := readFortuneFile(fortuneFile)

    if err == nil {
        rand.Seed(time.Now().UTC().UnixNano())
        min := 1
        max := 10
        r := rand.Intn(max - min + 1) + min

        for i := 0; i < r; i++ {
            rand.Seed(time.Now().UTC().UnixNano())
            f := rand.Int() % len(fortunes)
            title := fmt.Sprintf("fortune %d", i + 1)
            msg = append(msg, Fortune{Title: title,Content: fortunes[f]})
        }
    }
    fmt.Sprintln("message length : ", len(msg))
    return err
}

func homePage(w http.ResponseWriter, r *http.Request){
    errf := findAndPrint("fortunes")
    if errf != nil { errf.Error() }

    enc := json.NewEncoder(w)
    enc.SetEscapeHTML(true)
    enc.SetIndent("", "    ")

    errj := enc.Encode(msg)
    if errj != nil { errj.Error() }
}

func handleRequests() {
    r := mux.NewRouter()
    r.HandleFunc("/", homePage).Methods("GET")
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
    fmt.Println("Server is running on port 80")
    handleRequests()
}
