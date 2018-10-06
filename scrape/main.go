package main

import (
    "fmt"
    "net/http"
    "html/template"
    "io/ioutil"
    "encoding/xml"
    "sync"
    )

var wg sync.WaitGroup

var locationMap = map[string]string {"https://auburn.craigslist.org/": "auburn "...}

var totalRecovers int = 0
var successfulReads int = 0

type Listings struct {
    Links []string `xml:"item>link"`
    Titles []string `xml:"item>title"`
    Descriptions []string `xml:"item>description"`
    Dates []string `xml:"item>date"`
}

type Listing struct {
    Title string
    Description string
    Date string
}

type ListAggPage struct {
        Title string
        Listings map[string]Listing
        SearchRequest string
}

func cleanUp(link string) {
    defer wg.Done()
    if r:= recover(); r!= nil {
        totalRecovers++
//      recoverMap <- link
    }
}

func cityRoutine(c chan Listings, link string) {
    defer cleanUp(link)

    var i Listings
    address := link + "search/sss?format=rss&query=motorhome"
    resp, rErr := http.Get(address)
    if(rErr != nil) {
        fmt.Println("Fatal error has occurs while getting response.")
        fmt.Println(rErr);
    }

    bytes, bErr := ioutil.ReadAll(resp.Body)
    if(bErr != nil) {
        fmt.Println("Fatal error has occurs while getting bytes.")
        fmt.Println(bErr);
    }
    xml.Unmarshal(bytes, &i)
    resp.Body.Close()
    c <- i
    successfulReads++
}

func listingAggHandler(w http.ResponseWriter, r *http.Request) {
    queue := make(chan Listings, 99999)
    listing_map := make(map[string]Listing)

    request_queue := make(chan string)
    for i := 0; i < 20; i++ {
        go func() {
            for {
                key := <- request_queue
                cityRoutine(queue, key)                
            }
        }()
    }
    
    for key, _ := range locationMap {
        wg.Add(1)
        request_queue <- key
    }

    wg.Wait()
    close(request_queue)
    close(queue)

    for elem := range queue { 
        for index, _ := range elem.Links {
        listing_map[elem.Links[index]] = Listing{elem.Titles[index * 2], elem.Descriptions[index], elem.Dates[index]}
        }
    }

    p := ListAggPage{Title: "Craigslist Aggregator", Listings: listing_map}
    t, _ := template.ParseFiles("basictemplating.html")
    fmt.Println(t.Execute(w, p))

    fmt.Println("Successfully loaded: ", successfulReads)       
    fmt.Println("Recovered from: ", totalRecovers)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", listingAggHandler)
    http.ListenAndServe(":8000", nil) 
}
