package main

import (
 "fmt"
 "math/rand"
 "time"
)

type PersonInfo struct {
 Info  string
 Value int
}

type Product struct {
 Price  float64
 Weight float64
}

func generateCode(length int) string {
 const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
 code := make([]byte, length)
 for i := range code {
  code[i] = charset[rand.Intn(len(charset))]
 }
 return string(code)
}

func generateProduct() Product {
 return Product{
  Price:  rand.Float64()*100 + 10, 
  Weight: rand.Float64()*5 + 0.5,  
 }
}

// second commit


func task1() {
 arr := [10]int{41, 20, 26, 11, 46, 48, 5, 6, 29, 12}

 a := arr[:5]
 b := arr[3:7]
 c := arr[8:]

 fmt.Println("Початковий масив:", arr)
 fmt.Println("Зріз a:", a)
 fmt.Println("Зріз b:", b)
 fmt.Println("Зріз c:", c)


 if len(a) > 1 {
  a[1] = 0
 }
 if len(b) > 0 {
  b[0] = 0
 }
 if len(c) > 0 {
  c[len(c)-1] = 0
 }

 fmt.Println("Масив після змін:", arr)
}

func task2() {
 arr := [10]int{41, 20, 26, 11, 46, 48, 5, 6, 29, 12}

 students := []PersonInfo{
  {Info: "Кулик Дарʼя Олександрівна, група 73, варіант 12,", Value: arr[0]},
  {Info: "Кулик Дарʼя Олександрівна, група 73, варіант 12,", Value: arr[1]},
  {Info: "Кулик Дарʼя Олександрівна, група 73, варіант 12,", Value: arr[2]},
  {Info: "Кулик Дарʼя Олександрівна, група 73, варіант 12,", Value: arr[3]},
 }

 fmt.Println("Зріз структур (students):")
 for i, s := range students {
  fmt.Printf("Елемент %d:\n", i+1)
  fmt.Printf("  Інформація: %s\n", s.Info)
  fmt.Printf("  Значення: %d\n", s.Value)
 }
}

func task3() {
 slice1 := make([]int, 0, 3)
 slice2 := make([]int, 0, 5)
 slice3 := make([]int, 0, 7)

 slice1 = append(slice1, 12)

 slice2 = append(slice2, 18, 29, 50)

 slice3 = append(slice3, 94, 91, 79, 80, 93)

 fmt.Printf("slice1: %v len=%d cap=%d\n", slice1, len(slice1), cap(slice1))
 fmt.Printf("slice2: %v len=%d cap=%d\n", slice2, len(slice2), cap(slice2))
 fmt.Printf("slice3: %v len=%d cap=%d\n", slice3, len(slice3), cap(slice3))

 newCap := cap(slice3) - 2
 if newCap < len(slice3) {
  newCap = len(slice3)
 }
 reducedSlice3 := make([]int, len(slice3), newCap)
 copy(reducedSlice3, slice3)

 fmt.Println("\nПісля зменшення ємності slice3:")
 fmt.Printf("slice3: %v len=%d cap=%d\n", reducedSlice3, len(reducedSlice3), cap(reducedSlice3))

 
 allSlices := [][]int{slice1, slice2, reducedSlice3}
 fmt.Println("\nЗріз зрізів:")
 for i, s := range allSlices {
  fmt.Printf("Зріз #%d: %v (len=%d, cap=%d)\n", i+1, s, len(s), cap(s))
 }
}


func task4() {
 rand.Seed(time.Now().UnixNano())

 products := make(map[string]Product)

 for i := 0; i < 8; i++ {
  code := generateCode(5)
  products[code] = generateProduct()
 }

 fmt.Println("Наявні коди продуктів:")
 for code := range products {
  fmt.Println("-", code)
 }

 var input string
 fmt.Print("\nВведіть код продукту, щоб побачити інформацію: ")
 fmt.Scanln(&input)

 if product, ok := products[input]; ok {
  fmt.Printf("\nІнформація про продукт %s:\n", input)
  fmt.Printf("Ціна: %.2f грн\n", product.Price)
  fmt.Printf("Вага: %.2f кг\n", product.Weight)
 } else {
  fmt.Println("\n Продукт із таким кодом не знайдено.")
 }
}


// ====== Головне меню ======
func main() {
 var choice int
 fmt.Println("Оберіть завдання для виконання:")
 fmt.Println("1 - Завдання 1 (масив і зрізи)")
 fmt.Println("2 - Завдання 2 (зріз структур)")
 fmt.Println("3 - Завдання 3 (зріз зрізів)")
 fmt.Println("4 - Завдання 4 (структура продуктів)")
 fmt.Print("Введіть номер: ")
 fmt.Scan(&choice)

 fmt.Println()

 switch choice {
 case 1:
  task1()
 case 2:
  task2()
 case 3:
  task3()
 case 4:
  task4()
 default:
  fmt.Println(" Невірний вибір")
 }
}