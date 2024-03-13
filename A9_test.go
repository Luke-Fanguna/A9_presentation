package main

import (
    "testing"
)

func TestInterp(t *testing.T) {
    t.Run("numC", func(t *testing.T) {
        tst := numC{
            n:4,
        }
        exp := tst.interp()
        if exp != "4" {
            t.Fatalf("Got: %v\n", exp)
        }
    })

    t.Run("idC", func(t *testing.T) {
        ts := idC{
            id:"hello",
        }
        exp := ts.interp()
        if exp != "hello" {
            t.Fatalf("Got: %v\n", exp)
        }
    })
}
