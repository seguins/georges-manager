package main

import (
  "fmt"
)

func main() {

  fmt.Println("Configuration begin")

  vulcand := VulcandManager{[]string{"http://172.17.0.8:2379"}, "my_test"}

  vulcand.SetServer("http://172.17.0.9/", 0)
  vulcand.SetFrontend("Path(`/`)")

  fmt.Println("Done")

}
