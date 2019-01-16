package main			//We declare this file as an independent, executable script rather than a shared library of function definitions.

import "fmt"		     	//We import our library for standard i/o, which we use for printing lines to the terminal.
import "math"		     	//We import our math library for the Sin trigonometry function we use in this example.

const s string = "constant"  	//We use "const" to declare a constant value. Here, we declare a constant string value.

func main()	{	     	//We declare our main to run when this file is built and executed, and define it inside these curly brackets.

     fmt.Println(s)	     	//We print the constant string we defined outside the scope of this function.

     const n = 500000000     	//We can use numeric consts, which are initially typeless until explicitly typecasted or used in a context where a specific type is required or expected by a function.

     const d = 3e20 / n	     	//We can assign using hexidecimal notation.
     fmt.Println(d)   	     	//We print the result of the above expression, "6e+11". 'd' does not have a specific type at this point.

     fmt.Println(int64(d))   	//We print out the new string representation of 'd', which changed after casting d to type int64

     fmt.Println(math.Sin(n))	//We do not need to explicitly cast, however. Our const 'n' is casted to a type float64 here just by being passed to math.Sin, which expects a parameter of that type.
     
}