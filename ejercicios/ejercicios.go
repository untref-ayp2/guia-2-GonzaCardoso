package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

func InvertirCadena(cadena string) string {
	var pila stack.Stack
	var nueva string
	for _, v := range cadena {
		pila.Push(string(v))
	}
	for !pila.IsEmpty() {
		caracter, _ := pila.Pop()
		nueva += caracter.(string)
	}

	return nueva
}

/*
	func Palindromo(string) bool {
		//TODO
	}
*/
func Balanceada(corchetes string) bool {
	var pila queue.Queue
	//var p string
	var c string
	//var l string
	for _, v := range corchetes {
		pila.Enqueue(string(v))
	}
	for !pila.IsEmpty() {
		caracter, _ := pila.Dequeue()
		if (caracter.(string) == "[" && len(c)%2 == 0) || (caracter.(string) == "]" && len(c)%2 != 0) {
			c += caracter.(string)
		} 
	}
	return false
}

/*
func UnirColas(q1, q2 queue.Queue) queue.Queue {
	// TODO
}
*/
