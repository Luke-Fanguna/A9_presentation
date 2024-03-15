package main

import (
	"log"
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

type lamC struct {
	args []idC
	body ExprC
}

type appC struct {
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

func (exp ifC) init() string {
	return "1"
}

func (exp lamC) init() string {
	return "1"
}

func (exp appC) init() string {
	return "1"
}

// values
type Value interface {
	val() string
}

type numV struct {
	n int
}

type strV struct {
	s string
}

type boolV struct {
	b bool
}

type closV struct {
	args []idC
	body ExprC
	env  Env
}

type primV struct {
	name string
}

func (exp numV) val() string {
	return "1"
}
func (exp boolV) val() string {
	return "1"
}
func (exp strV) val() string {
	return "1"
}
func (exp closV) val() string {
	return "1"
}
func (exp primV) val() string {
	return "1"
}

// environment structs
type binding struct {
	name string
	val  Value
}

type Env struct {
	bindings []binding
}

// global var
var bindings []binding = make([]binding, 2)
var baseEnv Env = Env{append(bindings,
	binding{"true", boolV{true}},
	binding{"false", boolV{false}},
	binding{"+", primV{"+"}},
	binding{"-", primV{"-"}},
	binding{"*", primV{"*"}},
	binding{"/", primV{"/"}},
	binding{"<=", primV{"<="}},
	binding{"equal?", primV{"equal?"}},
	binding{"error", primV{"error"}})}

//functions

func serialize(val Value) string {
	switch v := val.(type) {
	case numV:
		return strconv.Itoa(v.n)
	case strV:
		return v.s
	case boolV:
		if v.b {
			return "true"
		} else {
			return "false"
		}
	case closV:
		return "#<procedure>"
	case primV:
		return "#<primop>"
	default:
		return "unimplemented"
	}
}

func lookup(sym string, env Env) Value {
	for i := len(env.bindings) - 1; i >= 0; i-- {
		if (env.bindings[i]).name == sym {
			return env.bindings[i].val
		}
	}
	log.Fatalf("GOAZO9 symbol not found :(")
	return strV{"GOAZO9 symbol not found"}
}

// interp
func interp(e ExprC, env Env) Value {
	switch v := e.(type) {
	case numC:
		return numV{v.n}
	case idC:
		//need to implement lookup
		return lookup(v.id, env)

	case strC:
		return strV{v.s}
	case ifC:
		if (interp(v.cond, env) == boolV{true}) {
			return interp(v.then, env)
		} else {
			return interp(v.els, env)
		}
	default:
		return strV{"unimplemented"}
	}
}

// possibly add i/o with os import?
func main() {

	//fmt.Println(interp(numC{5}))
}
