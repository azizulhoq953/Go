package main
import "fmt"
func main(){

//primitive data type
//rune,byte=alias

//composite Data type
//array


//var students [10]string

//students[0]="Azizul"
//students[1]="Galib"
//students[2]="Safin"

//fmt.Println(students[..])
//fmt.Println(len(students))
//fmt.Println(students[0])

/*short hand array, string literals
In computer science, 
a literal is a notation for representing a fixed value 
in source code. ... An anonymous function is a literal
for the function type. In contrast to literals, variables
or constants are symbols that can take on one of a class 
of fixed values, the constant being constrained not to
change.*/
// implicit=অন্তর্নিহিত
students :=[...]string{"Azizul", "Safin","Tamanna", "Galib",}

fmt.Println(students,len(students))

}