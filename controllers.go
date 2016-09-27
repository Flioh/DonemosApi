package main

import (
	"gopkg.in/mgo.v2"
)

type Controller struct {
	session *mgo.Session
}

func NewController(s *mgo.Session) *Controller {
	return &Controller{s}
}
