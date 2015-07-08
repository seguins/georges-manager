package main


/* Vulcand types */

type BackendVulcand struct {
  Type string
}

type ServerVulcand struct {
  Url string `json:"URL"`
}

type FrontendVulcand struct {
  Type string
  BackendId string
  Route string
}
