package main

import (
	"fmt"
	"strconv"
	//"os"
)

type ExprC interface {
	init() string
}

type numC struct {
	n int
}

type idC struct {
	id string
}

type strC struct {
	s string
}

type ifC struct {
	cond ExprC
	then ExprC
	els  ExprC
}

func (num numC) init() string {
	return strconv.Itoa(num.n)
}

func (exp idC) init() string {
	return exp.id
}

func (exp strC) init() string {
	return exp.s
}

func (exp ifC) init() string {
	return interp(exp.cond)
}

func interp(e ExprC) string {
	switch v := e.(type) {
	case numC:
		return strconv.Itoa(v.n)
	case idC:
		return v.id
	case strC:
		return v.s
	case ifC:
		if interp(v.cond) == "true" {
			return interp(v.then)
		} else {
			return interp(v.els)
		}
	default:
		return "unimplemented"
	}
}

// possibly add i/o with os import?
func main() {
	fmt.Println(interp(numC{5}))
}
