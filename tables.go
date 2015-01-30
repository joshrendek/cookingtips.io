package main

import (
	_ "gopkg.in/pg.v2"
)

type Page struct {
	Id           int64
	Title        string
	Instructions []string
	Youtubes     []string
	Articles     []string
	Tags         []string
}
