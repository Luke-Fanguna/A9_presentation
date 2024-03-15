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

	t.Run("ifC", func(t *testing.T) {
		var args []ExprC = make([]ExprC, 0)
		args = append(args, numC{2})
		args = append(args, numC{3})

		tst := ifC{
			cond: appC{idC{"<="}, args},
			then: numC{1},
			els:  numC{2},
		}

		exp := serialize(interp(tst, baseEnv))
		if exp != "1" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	t.Run("ifC2", func(t *testing.T) {
		var args []ExprC = make([]ExprC, 2)
		args[0] = numC{3}
		args[1] = numC{2}

		var args2 []ExprC = make([]ExprC, 2)
		args2[0] = numC{10}
		args2[1] = numC{3}

		tst := ifC{
			cond: appC{idC{"<="}, args},
			then: numC{3},
			els:  appC{idC{"+"}, args2},
		}

		exp := serialize(interp(tst, baseEnv))
		if exp != "13" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

	t.Run("appC", func(t *testing.T) {
		var params []idC = make([]idC, 2)
		params[0] = idC{"x"}
		params[1] = idC{"y"}

		var args []ExprC = make([]ExprC, 2)
		args[0] = idC{"x"}
		args[1] = idC{"y"}
		var lam = lamC{params, appC{idC{"*"}, args}}

		var args2 []ExprC = make([]ExprC, 2)
		args[0] = numC{3}
		args[1] = numC{4}
		var prog = appC{lam, args2}
		exp := serialize(interp(prog, baseEnv))
		if exp != "12" {
			t.Fatalf("Got: %v\n", exp)
		}
	})

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
