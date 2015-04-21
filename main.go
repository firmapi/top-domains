package main

import (
  "./vendor/martini"
  "encoding/csv"
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "strconv"
)

var alexa_list []*Domain

type WorkingResponse struct {
  Status string `json:"status"`
}

type ErrorResponse struct {
  Status  string `json:"status"`
  Message string `json:"message"`
}

type DomainRankResponse struct {
  Status   string `json:"status"`
  Found    bool   `json:"found"`
  Domain   string `json:"domain"`
  Position int    `json:"position"`
}

type Domain struct {
  Position int
  Value    string
}

// Extracts the list from the csv and loads it into memory.
//
func buildAlexList() {
  csvfile, err := os.Open("list.csv")

  if err != nil {
    fmt.Println(err)
    return
  }

  defer csvfile.Close()

  reader := csv.NewReader(csvfile)

  reader.FieldsPerRecord = -1

  rawCSVdata, err := reader.ReadAll()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  for _, each := range rawCSVdata {
    position, _ := strconv.Atoi(each[0])
    alexa_list = append(alexa_list, &Domain{Position: position, Value: each[1]})
  }
}

// Returns the position of the supplied domain or 0 if it wasn't found.
//
func domainPosition(domain string) (bool, int) {
  for _, element := range alexa_list {
    if element.Value == domain {
      return true, element.Position
    }
  }

  return false, 0
}

func main() {
  buildAlexList()

  fmt.Println(len(alexa_list))

  m := martini.Classic()
  m.Get("/", func() (int, string) {
    response := &WorkingResponse{
      Status: "success",
    }
    string_response, _ := json.Marshal(response)

    return 200, string(string_response)
  })

  m.Get("/rank", func(r *http.Request) (int, string) {
    domain := r.URL.Query().Get("domain")

    if domain == "" {
      response := &ErrorResponse{
        Status:  "error",
        Message: "Query param 'domain' required.",
      }

      string_response, _ := json.Marshal(response)

      return 500, string(string_response)
    }

    found, position := domainPosition(domain)

    response := &DomainRankResponse{
      Status:   "success",
      Domain:   domain,
      Found:    found,
      Position: position,
    }

    string_response, _ := json.Marshal(response)

    return 200, string(string_response)
  })

  m.Run()
}
