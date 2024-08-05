package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
    "io"
)

func init() {
    if len(os.Args) != 2 { die("Usage: "+os.Args[0]+" <url>") }
}

func main() {

    // Parse url
    urlString := os.Args[1]
    url, err := url.Parse(urlString)
    if err != nil { die("Could not parse url: '"+urlString+"'") }

    if url.Scheme == "" { die("Missing protocol scheme ('http://', 'https://')") }

    // Craft client
    client := &http.Client { 
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

    // Craft header
    header := make(http.Header)
    header.Set("User-Agent", "gurl/1.0")

    // Craft request
    request := &http.Request {
        Method: "GET",
        URL: url,
        Header: header,
    }

    // Do request, handle errors
    response, err := client.Do(request)
    if err != nil { handleRequestErrors(request, err) }
    defer response.Body.Close()

    // Print response to stdout
    body, err := io.ReadAll(response.Body)
    if err != nil { die(err.Error()) }
    fmt.Print(string(body))

}

