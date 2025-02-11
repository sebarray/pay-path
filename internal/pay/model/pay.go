package model

type Pay struct {
	ID          string `json:"id"`
	Tag         string `json:"tag"`
	AddressBank string `json:"addressBank"`
	Rgb         string `json:"rgb"`
}



type PayList struct {
Id string `json:"id"`
PayList []Pay `json:"payList"`
}