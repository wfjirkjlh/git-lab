package main

import (
 "fmt"
 "math"
)


type Student struct {
 Name    string
 Surname string
 Group   int
 Variant int
}


type ContactInfo struct {
 Email string
 Phone string
}

type Address struct {
 City   string
 Street string
}

type Employee struct {
 FirstName string
 LastName  string
 Position  string
 Contact   ContactInfo
 Address   
}


func Calculations(g *int, v *int, c *float64) {
 origG := *g
 origV := *v
 origCint := int(*c)

 sum := origG + origV + origCint
 product := origG * origV * origCint

 var trig float64
 if origV%2 == 0 {
  trig = math.Sin(float64(sum))
 } else {
  trig = math.Cos(float64(product))
 }

 *g = sum
 *v = product
 *c = trig
}

func main() {
 
 initialGroup := 73
 initialVariant := 12
 initialCourse := 4.0

 group := initialGroup
 variant := initialVariant
 course := initialCourse

 fmt.Println("Початкові значення:")
 fmt.Printf("група = %d, варіант = %d, курс = %v\n\n", group, variant, course)

 Calculations(&group, &variant, &course)

 fmt.Println("Нові значення після Calculations:")
 fmt.Printf("група = %d, варіант = %d, курс = %v\n\n", group, variant, course)


 
 student := Student{
  Name:    "Дарʼя",
  Surname: "Кулик",
  Group:   initialGroup,
  Variant: initialVariant,
 }

 fmt.Println("Структура до:")
 fmt.Printf("Ім'я: %s, Прізвище: %s, Група: %d, Варіант: %d\n\n",
  student.Name, student.Surname, student.Group, student.Variant)

 rewriteVariantSum(&student)

 fmt.Println("Структура після:")
 fmt.Printf("Ім'я: %s, Прізвище: %s, Група: %d, Варіант: %d\n\n",
  student.Name, student.Surname, student.Group, student.Variant)

 
 employee := Employee{
  FirstName: "Дарʼя",
  LastName:  "Кулик",
  Position:  "Студентка",
  Contact: ContactInfo{
   Email: "kylikdasha22@gmail.com",
   Phone: "+380957663831", 
  },
  Address: Address{
   City:   "Дніпро",
   Street: "вулиця Відпочинку, 68",
  },
 }

 fmt.Println("--- Інформація про співробітника ---")
 fmt.Printf("Повне ім'я: %s %s\n", employee.FirstName, employee.LastName)
 fmt.Printf("Посада: %s\n", employee.Position)
 fmt.Printf("Email: %s\n", employee.Contact.Email)
 fmt.Printf("Місто: %s\n\n", employee.Address.City)

 
 arr := []float64{
  1,
  2,
  float64(group),
  float64(variant),
  course,
 }

 fmt.Println("Масив до змін:", arr)

 arr[0] = float64(initialGroup)
 arr[1] = float64(int(initialCourse))

 for i := range arr {
  arr[i] += float64(initialVariant)
 }

 fmt.Println("Масив після змін:")
 fmt.Println(arr)
}

func rewriteVariantSum(s *Student) {
 s.Variant = s.Group + s.Variant
}