package main
import (
"fmt"
"net/http"
)

var urls = []string{
  "http://www.google.com/",
  "http://golang.org/",
  "http://blog.golang.org/",
}

func main() {
  // Executes an HTTP HEAD request for all URL's
  // and returns the HTTP status string or an error string.
  for _, url := range urls {
    resp, err := http.Head(url)
    if err != nil {
    fmt.Println("Error:", url, err)
  }
  fmt.Println(url, ": ", resp.Status)
  }
}
