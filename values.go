package main				/*We specify this package to be an executable program.*/

import "fmt"	     	     		/*We import the fmt package for standard I/O functions*/

func main(){	     	     		/*We declare the entry point for this program as a function with no parameters*/

     fmt.Println("go" + "lang")		/*We can use the + operator to concatenate strings*/		

     fmt.Println("1+1 =", 1+1)		/*We can concatenate the results of an int operation to a string*/
     fmt.Println("7.0/3.0 =", 7.0/3.0)	/*We can also do the same with floats, boolean values, or logical comparisons*/

     fmt.Println(true && false)
     fmt.Println(true || false)
     fmt.Println(!true)
}