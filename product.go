package main

import "github.com/oklog/ulid/v2"

type Product struct {
	Id    ulid.ULID
	Name  string
	Price int
}
