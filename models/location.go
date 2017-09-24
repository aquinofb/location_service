package models

type Location struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Types []string `json:"types"`
  Lat string `json:"lat"`
  Lng string `json:"lng"`
}
