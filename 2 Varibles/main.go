package main

import (
	"fmt"
)

const PublicVarible string = "jhwasdgfjkasf"// Public

func main(){
	var username string = "Jon"
	fmt.Println(username)
	fmt.Printf("This is %T \n\n",username)


	var number int = 122
	fmt.Println(number)
	fmt.Printf("This is %T \n\n",number)


	var isBig bool = false
	fmt.Println(isBig)
	fmt.Printf("This is %T \n\n",isBig)


	var possitiveNumber uint = 8  // possitiveNumber >= 0  !!!!
	fmt.Println(possitiveNumber)
	fmt.Printf("This is %T \n\n",possitiveNumber)


	var floatNum float64 = 1.999999999999
	fmt.Println(floatNum)
	fmt.Printf("This is %T \n\n",floatNum)

	//defult varible
	var defultNumber int
	fmt.Println(defultNumber)
	
	var defultString string
	fmt.Println(defultString)//...

	// no var style 
	numberUser := 100
	fmt.Println(numberUser)
	fmt.Printf("This is %T \n\n",numberUser)//...

	
	nameUser := "John"
	fmt.Println(nameUser)
	fmt.Printf("This is %T \n\n",nameUser)//...
}