package main

import (
	"testing"
)

func TestInterp(t *testing.T) {

	t.Run("numC", func(t *testing.T) {
		tst := numC{
			n: 4,
		}
		exp := serialize(interp(tst, baseEnv))
		if exp != "4" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	//correctly receives error (which causes whole program to halt)
	/*
		t.Run("idC", func(t *testing.T) {
			tst := idC{
				id: "hello",
			}
			exp := serialize(interp(tst, baseEnv))
			if exp != "hello" {
				t.Fatalf("Got: %v\n", exp)
			}
		})
	*/
	t.Run("strC", func(t *testing.T) {
		tst := strC{
			s: "hello there",
		}
		exp := serialize(interp(tst, baseEnv))
		if exp != "hello there" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	t.Run("lookup1", func(t *testing.T) {
		tst := idC{
			id: "+",
		}
		exp := lookup(tst.id, baseEnv)
		if (exp != primV{"+"}) {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	t.Run("interp1", func(t *testing.T) {
		tst := idC{
			id: "+",
		}
		exp := serialize(interp(tst, baseEnv))
		if exp != "#<primop>" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	/*
		t.Run("ifC", func(t *testing.T) {
			tst := ifC{
				cond: idC{"true"},
				then: numC{1},
				els:  numC{2},
			}
			exp := serialize(interp(tst, baseEnv))
			if exp != "1" {
				t.Fatalf("Got: %v\n", exp)
			}
		})
	*/
	/*t.Run("lamC", func(t *testing.T) {
	    tst := lamC{

	    }
	    exp := interp(tst)
	    if exp != "1"{
	        t.Fatalf("Got: %v\n", exp)
	    }
	})*/

	/*
		t.Run("appC", func(t *testing.T) {
			argList := make([]ExprC, 2)

			tst := appC{
				fun:  idC{"+"},
				args: append(argList, numC{1}, numC{2}),
			}
			exp := interp(tst)
			if exp != "1" {
				t.Fatalf("Got: %v\n", exp)
			}
		})
	*/
}
