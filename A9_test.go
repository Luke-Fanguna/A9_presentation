package main

import (
    "testing"
)

func TestInterp(t *testing.T) {
    t.Run("numC", func(t *testing.T) {
        tst := numC{
            n:4,
        }
        exp := interp(tst)
        if exp != "4" {
            t.Fatalf("Got: %v\n", exp)
        }
    })

    t.Run("idC", func(t *testing.T) {
        tst := idC{
            id:"hello",
        }
        exp := interp(tst)
        if exp != "hello" {
            t.Fatalf("Got: %v\n", exp)
        }
    })
}
