package main

import (
	"fmt"
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

func (num numC) init() string {
	return strconv.Itoa(num.n)
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
var bindings []binding = make([]binding, 0)
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

func primop(prim primV, a ExprC, b ExprC, env Env) Value {
	a2 := interp(a, env)
	b2 := interp(b, env)
	if prim.name == "equal?" {
		return boolV{a2 == b2}
	}

	//rest of functions involve numbers so check
	switch a2.(type) {
	case numV:
		//nothing
	default:
		log.Fatalf("GOAZO9 primitive operation %v requires integers", prim.name)
	}
	switch b2.(type) {
	case numV:
		//nothing
	default:
		log.Fatalf("GOAZO9 primitive operation %v requires integers", prim.name)
	}

	//convert to numV
	a3 := a2.(numV).n
	b3 := b2.(numV).n

	//going to add the other operations in a bit
	switch prim.name {
	case "+":
		return numV{a3 + b3}
	case "-":
		return numV{a3 - b3}
	case "*":
		return numV{a3 * b3}
	case "/":
		return numV{a3 / b3}
	case "<=":
		return boolV{a3 <= b3}
	default:
		log.Fatalf("GOAZO9 primitive operation not found %v", prim.name)
		return strV{"GOAZO9 primitive operation not found " + prim.name}
	}
}

// extend env given a list of symbols and their corresponding values
func multiExtend2(env Env, syms []string, vals []Value) Env {
	for i := 0; i < len(syms); i++ {
		env = extendEnv(env, syms[i], vals[i])
	}
	return env
}

// extendEnv
func extendEnv(env Env, sym string, val Value) Env {
	newBinding := binding{sym, val}
	env.bindings = append(env.bindings, newBinding)
	return env
}

// interpret list of arguments to an appC
func interpAppArgs(args []ExprC, env Env) []Value {
	var interprettedArgs []Value = make([]Value, len(args))
	for i := 0; i < len(args); i++ {
		interprettedArgs[i] = interp(args[i], env)
	}
	return interprettedArgs
}

func getIdcSyms(ids []idC) []string {
	var names []string = make([]string, len(ids))
	for i := 0; i < len(ids); i++ {
		names[i] = ids[i].id
	}
	return names
}

// interp
func interp(e ExprC, env Env) Value {
	switch v := e.(type) {
	case numC:
		return numV{v.n}
	case idC:
		return lookup(v.id, env)
	case strC:
		return strV{v.s}
	case lamC:
		return closV{v.args, v.body, env}
	case ifC:
		//need to implement primitive boolean operations
		if (interp(v.cond, env) == boolV{true}) {
			return interp(v.then, env)
		} else {
			return interp(v.els, env)
		}
	case appC:
		//get the function definition
		var fun Value = interp(v.fun, env)
		//check if primitive or closure
		switch fun.(type) {
		case primV:
			//cast
			funPrim := fun.(primV)
			return primop(funPrim, v.args[0], v.args[1], env)
		case closV:
			funClos := fun.(closV)
			a_vals := interpAppArgs(v.args, env)
			env2 := multiExtend2(funClos.env, getIdcSyms(funClos.args), a_vals)
			return interp(funClos.body, env2)
		default:
			log.Fatalf("GOAZO9 trying to apply non function")
			return strV{"GOAZO9 trying to apply non function"}
		}
	default:
		return strV{"unimplemented"}
	}
}

// possibly add i/o with os import?
func main() {

	var args []ExprC = make([]ExprC, 2)
	args[0] = numC{2}
	args[1] = numC{3}

	fmt.Println(args[0])
}
