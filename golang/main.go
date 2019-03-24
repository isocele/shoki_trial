package main

import (
  "log"
  "net/http"
  "io/ioutil"
  "bytes"
  "fmt"
  "github.com/gocolly/colly"
  "encoding/json"
)

type Response struct {
 rating string
 count string
}

func main() {
 var res Response
  http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
   (w).Header().Set("Access-Control-Allow-Origin", "*")
   (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
   (w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
   body, err := ioutil.ReadAll(r.Body)
   if err != nil {
       log.Printf("Error reading body: %v", err)
       http.Error(w, "can't read body", http.StatusBadRequest)
       return
    }

   r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

   c := colly.NewCollector()

   c.OnHTML(".restaurants-detail-overview-cards-RatingsOverviewCard__overallRating--nohTl", func(e *colly.HTMLElement) {
    link := e.Attr("href")
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    res.rating = e.Text

   })
   c.OnHTML(".restaurants-detail-overview-cards-RatingsOverviewCard__ratingCount--DFxkG", func(e *colly.HTMLElement) {
    link := e.Attr("href")
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    res.count = e.Text
   })

   buf := new(bytes.Buffer)
   buf.ReadFrom(r.Body)
   newStr := buf.String()

   c.Visit(newStr)
   jsonres, err := json.Marshal(map[string]interface{}{
           "data": "Data received",
           "rating": res.rating,
           "count": res.count,
       })

   fmt.Println(string(jsonres))
   // fmt.Printf("ok -%s", string(jsonres))
   w.Header().Set("Content-Type", "application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(jsonres)
  }))

  log.Println("Now server is running on port 8083")
  http.ListenAndServe(":8083", nil)
}
