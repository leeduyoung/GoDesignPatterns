package main

type Element interface {
	Accept(Visitor)
}
