package main

import "fmt"
import "strconv"

var(
 firstName string="suresh"
 age int =30
 height float32=5.6
)

func main(){

  fmt.Println("Calculating two numbers: ")

  var firstNumber, secondNumber,sum int
  firstNumber=20
  secondNumber=10
  sum=firstNumber+secondNumber
  fmt.Println("Total =", +sum)

  // without declaring variable and datatype

  number:= 50
  fmt.Println(number)

  name:= "suresh"

  fmt.Println("")
  fmt.Println("")
  fmt.Println("")

  fmt.Printf("%T,%v",name,name)
  fmt.Println("")
  fmt.Println("")
  fmt.Println("")

  fmt.Println(name)

  fmt.Printf("%T",name)
  fmt.Println("")

  // printf for format of datatypes
  fmt.Printf("%v, %T",name,name)
  fmt.Println("")
  fmt.Printf("%v, %T", number , number)
  fmt.Println("")
  fmt.Println("")

    fmt.Println("My name is: "+firstName)
    fmt.Println("My age is: ",+age)
    fmt.Println("My height is: ",+height)




    i:=49
    var value string
    value=strconv.Itoa(i)
    fmt.Println("")
    fmt.Println("")
    fmt.Println(value)


    var iss bool=true

    fmt.Printf("%T,%v",iss,iss)


    var num int64 =98595959595
    fmt.Println(num)


//    s:="sandeep reddy"

  //  fmt.Printf("%v, %T",s[6],s[6])


    s:= "sandeep reddy"

    fmt.Printf("%v, %T\n",string(s[2]),s[2])

    alpha:="ABCDS"
    b:=[]byte(alpha)
    fmt.Printf("%v, %T\n\n",b,b)

    const myNum int=45

    fmt.Println(myNum)
    fmt.Println("\n\n")
    const myCompanyName string="sandeep corporation"


    fmt.Println(myCompanyName)



}
