package main

import "fmt"
import "regexp"

func main(){
  var nama = "joni, amran, manurung"
  var regex, err = regexp.Compile(`[a-z]+`)

  if err != nil{
    fmt.Println(err.Error())
  }

  var hasil = regex.FindAllString(nama, 3)
  fmt.Println(hasil)

  var index1 = regex.FindStringIndex(nama)
  fmt.Println(index1)
}
