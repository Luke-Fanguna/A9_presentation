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
	ExprC
	cond ExprC
	then ExprC
	els  ExprC
}

type lamC struct {
	ExprC
	args []idC
	body ExprC
}

type appC struct {
	ExprC
	fun  ExprC
	args []ExprC
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

/*
func (exp ifC) init() string {
	return interp(exp.cond)
}


func (exp lamC) init() string {
	return interp(exp.body)
}

func (exp appC) init() string {
	return interp(exp.fun)
}
*/

// values
type Value interface {
	val()
}

type numV struct {
	Value
	n int
}

type strV struct {
	Value
	s string
}

type boolV struct {
	Value
	b bool
}

type closV struct {
	Value
	args []idC
	body ExprC
	env  Env
}

type primV struct {
	Value
	name string
}

// environment structs
type binding struct {
	name string
	val  Value
}

type Env struct {
	bindings []binding
}

// interp
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
