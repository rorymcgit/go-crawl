package main

import (
  "fmt"
  "flag"
  "os"
  "net/http"
  "io/ioutil"
)

func main() {
  flag.Parse()
  args := flag.Args()
  fmt.Println(args)

  if len(args) < 1 {
    fmt.Println("You did not specify start page")
    os.Exit(1)
  }
  retrieve(args[0])
}

func retrieve(uri string) {
  resp, err := http.Get(uri)
  if err != nil {
    return
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  fmt.Println(string(body))
}
