package main

import (
    "fmt"
    "temperature"
 
)


var  t float64
var u string
   

func main() {
   fmt.Println("Please enter your  temperature and its unit (C/F) : ")
   fmt.Scanln(&t,&u)
   if u == "C"{
   	fmt.Println	(t,"°C =",temperature.FAHRENHEIT(t),"°F")
   }else if u == "F"{
   	fmt.Println	(t,"°F =",temperature.CELSIUS(t),"°C")
   }else {
   	fmt.Println	("error")
   }
  


}
