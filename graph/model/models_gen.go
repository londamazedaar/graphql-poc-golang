// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Address struct {
	Road     string `json:"road"`
	Locality string `json:"locality"`
}

type BitcoinPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Base     string `json:"base"`
}

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type DataBitcoinPrice struct {
	Data *BitcoinPrice `json:"data"`
}

type Train struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}