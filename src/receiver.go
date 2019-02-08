package main

import "fmt"

type Person struct {
  Name string
  MailAddress string
  Age int
  Gender string
}

type Setter interface {
  GetName() string
  GetMailAddress() string
  GetAge() int
  GetGender() string
  
  SetName(v string)
  SetMailAddress(v string)
  SetAge(v int)
  SetGender(v string)
}

func (v *Person) GetName() string{
  return v.Name
}

func (v *Person) GetMailAddress() string{
  return v.MailAddress
}

func (v *Person) GetAge() int{
  return v.Age
}

func (v *Person) GetGender() string{
  return v.Gender
}

func (v *Person) SetName(s string){
  v.Name = s
}
func (v *Person) SetMailAddress(s string){
  v.MailAddress = s
}
func (v *Person) SetAge(s int){
  v.Age = s
}
func (v *Person) SetGender(s string){
  v.Gender = s
}


func main() {
  p1 := &Person{}
  p1.SetName("Peter")
  p1.SetMailAddress("Peter@peter.com")
  p1.SetAge(20)
  p1.SetGender("m")
  
  fmt.Println(p1.GetName())
  fmt.Println(p1.GetMailAddress())
  
  p2 := p1
  p2.SetAge(30)
  fmt.Println(p1.GetAge())
  fmt.Println(p1.GetGender())
}