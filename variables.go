package main		//We declare this file as an independent, executable script rather than a shared library of function definitions.

import "fmt"		//We import our library for standard i/o, which we use for printing lines to the terminal.

func main()	{	//We declare our main to run when this script is built and executed, and define it in the curly brackets.

     var a = "initial" //We declare a variable named 'a' and initialize it to a string value, "initial".
     fmt.Println(a)    //We use the "print line" function from our imported library to display the value we assigned to 'a'.

     var b,c int = 1,2 //We declare two variables, specify their type(int), and initialize them by assigning each their respective values.
     fmt.Println(b, c) //We print the values of b and c on their own line. Note the function accepts multiple args, and integers as args.

     var d = true      //We do not need to specify the type, Go will infer the type of a declared variable upon initialization.
     fmt.Println(d)    //We print out the boolean we just specified above.

     var e int	       //We can specify the type without initializing.
     fmt.Println(e)    //We find that regardless of type, uninitialized variables are automatically zero-valued.

     f := "short"      //We can use the operator := as shorthand for declaration & initialization.
     fmt.Println(f)    //We print the value we just assigned to a new variable, 'f'.
}