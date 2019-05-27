package main

//blocking call

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
    links := []string{
      "http://google.com",
      "http://facebook.com",
      "http://amazon.com",
      "http://golang.org",
      "http://stackoverflow.com",
      "http://nosuchweb.com",
      "http://bolor-toli.com",
    }

    c := make(chan string)

    for _, link := range links {
      go checkLink(link,c )
    }

    for l := range c {
      go func (link string) {
        time.Sleep(time.Second * 3)
        checkLink(link, c)
      }(l)
    }
 }

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link, "might be down!")
    c <- link
    return
  }

  fmt.Println(link, "is up!")
  c <- link
}
