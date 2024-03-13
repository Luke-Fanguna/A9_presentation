package main

import (
	"fmt"
    "strconv"
)

type ExprC interface {
    interp() string
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

func (num numC) interp() string {
    return strconv.Itoa(num.n)
}

func (exp idC) interp() string {
    return exp.id
}

func (exp strC) interp() string {
    return exp.s
}

/*func (e interp string {
    switch v := e.(type){
    case numC:
        return "numC"
    case idC:
        return "idC"
    case strC:
        return "strC"
    default:
        return "unimplemented"
    }
}
*/
func main() {
	fmt.Println("Hello World")
	fmt.Println("Semicolon")
}

