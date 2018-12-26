package main

import "fmt"

type empty interface {}

func whatIsThis(i interface {})  {
	switch v := i.(type) {
	case int:
		fmt.Printf("This is a int: %d\n", v)
	case string:
		fmt.Printf("This is a string: %s\n", v)
	case uint32:
		fmt.Printf("This is a uint32: %d\n", v)
	default:
		fmt.Printf("I don't know what this is : %v\n", v)
	}

	fmt.Printf("Assertion Stype\n")
	// type assert style
	switch i.(type) {
	case int:
		fmt.Printf("This is a int: %d\n", i.(int))
	case string:
		fmt.Printf("This is a string: %s\n", i.(string))
	case uint32:
		fmt.Printf("This is a uint32: %d\n", i.(uint32))
	default:
		fmt.Printf("I don't know what this is \n",)
	}
	fmt.Printf("\n\n")


	/*
	fmt.Printf("%T\n", i)     // %T  tells the type
	*/
}

func main()  {

	whatIsThis(42)
	whatIsThis(uint32(42))
	whatIsThis("42")
	whatIsThis([...]string{"A","B","C"})

}
