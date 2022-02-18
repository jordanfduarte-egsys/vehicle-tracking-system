package domain

/**
* Methods used for goroutine of requests
* @package domain
* @author Jordan Duarte
**/

import (
    "fmt"
    "bytes"
    "net/http"
    "time"
)

type Site struct {
    URL string
    Buffer []byte
}

func sending(wId int, url string, buffer []byte) int {
    fmt.Printf("Working ID %d\n", wId)
    client := &http.Client{}

    fmt.Printf("Sending to %s\n", url)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(buffer))
    if err != nil {
        fmt.Printf("%s", err)
    }

    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
       req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err2 := client.Do(req)
    var statusCode int
    if err2 != nil {
        fmt.Printf("Status returned %s of hostname %s\n", err2.Error(), url)
        statusCode = 400
    } else {
        fmt.Printf("Status returned %s of hostname %s\n", resp.Status, url)
        statusCode = resp.StatusCode
        defer resp.Body.Close()
    }
    return statusCode
}

func Crawl(wId int, jobs <-chan Site) {
    for site := range jobs {
        fmt.Print("First try\n")
        var statusCode = sending(wId, site.URL, site.Buffer)

        if statusCode != 200 {
            time.Sleep(1 * time.Second)
            fmt.Print("Second try\n")
            statusCode = sending(wId, site.URL, site.Buffer)
        }

        if statusCode != 200 {
            time.Sleep(5 * time.Second)
            fmt.Print("Third attempt\n")
            statusCode = sending(wId, site.URL, site.Buffer)
        }

        if statusCode != 200 {
            time.Sleep(15 * time.Second)
            fmt.Print("Fourth attempt\n")
            statusCode = sending(wId, site.URL, site.Buffer)
        }
    }
}