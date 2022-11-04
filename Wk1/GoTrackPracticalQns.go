package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// basicdatatype1()
	// booleantype1()
	// intfloatcomplextypes1()
	// intfloatcomplextypes2()
	// intfloatcomplextypes3()
	// intfloatcomplextypes4()
	// string1()
	// ifstatements1()
	// ifstatements2()
	// ifstatements3()
	// ifstatements4()
	// ifstatements5()
	// forstatements1()
	// forstatements2()
	// forstatements3()
	// forstatements4()
	// forstatements5()
	// forstatements6()
	// forstatements7()
	// switchstatements1()
	// switchstatements2()
	// switchstatements3()
	// compositetypes1to4()
	// compositetypes5()
	// slices1to5()
	// struct1to2()
}
func basicdatatype1() {
	firstname := "Luke"
	lastname := "Skywalkter"
	age := 20
	weight := 73.0
	height := 1.72
	credits := 123.55
	accname := "admin"
	accpwd := "password"
	subscribe := true
	fmt.Printf("The following is the account information.\n")
	fmt.Printf("First Name : %s\n", firstname)
	fmt.Printf("Last Name : %s\n", lastname)
	fmt.Printf("Age : %d years old\n", age)
	fmt.Printf("Weight : %.2fkg\n", weight)
	fmt.Printf("height : %.2fm\n", height)
	fmt.Printf("Remaining Credits : $%.2f\n", credits)
	fmt.Printf("Account Name : %s\n", accname)
	fmt.Printf("Account Password : %s\n", accpwd)
	fmt.Printf("subscribed user : %t\n", subscribe)
}

func booleantype1() {
	var x = int(39)
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scan(&guess)
	if guess < x {
		fmt.Println("Too low, try again next time!")
	} else if guess > x {
		fmt.Println("Too high, try again next time!")
	} else {
		fmt.Println("Well Done! Your guess is correct")
	}
}
func intfloatcomplextypes1() {
	var i int
	var j int
	var k int
	fmt.Println("Enter first value from 0 to 9: ")
	fmt.Scanln(&i)
	fmt.Println("Select one of the arithmetic functions:\n1. add\n2. Subtract\n3.multiply\n4.divide")
	fmt.Scanln(&j)
	fmt.Println("Enter second value from 0 to 9: ")
	fmt.Scanln(&k)
	if j == 1 {
		fmt.Println(i + k)
	} else if j == 2 {
		fmt.Println(i - k)
	} else if j == 3 {
		fmt.Println(i * k)
	} else {
		fmt.Println(i / k)
	}
}
func intfloatcomplextypes2() {
	var x int
	fmt.Println("Select the format input temperature:\n1. Kelvin\n2. Celsius\n3. Fahrenheit")
	fmt.Scanln(&x)
	var y float64
	fmt.Println("Please input the current temperature: ")
	_, err := fmt.Scanf("%f", &y)

	if err != nil {
		fmt.Println(err)
	}

	f, c, k := float64(0), float64(0), float64(0)

	if x == 3 {
		f = y
		c = 5 * (f - 32) / 9
		k = 5 * (f + 459.67) / 9
	} else if x == 2 {
		c = y
		f = 9*c/5 + 32
		k = 5 * (f + 459.67) / 9
	} else {
		k = y
		f = 9*k/5 - 459.67
		c = 5 * (f - 32) / 9
	}

	fmt.Println(fmt.Sprintf("%.2f Fahrenheit, %.2f Celcius, %.2f Kelvin", f, c, k))
}
func intfloatcomplextypes3() {
	var a float32
	fmt.Println("Enter the number of 1-dollar coins: ")
	fmt.Scanln(&a)
	var b float32
	fmt.Println("Enter the number of 50-cent coins: ")
	fmt.Scanln(&b)
	var c float32
	fmt.Println("Enter the number of 20-cent coins: ")
	fmt.Scanln(&c)
	var d float32
	fmt.Println("Enter the number of 10-cent coins: ")
	fmt.Scanln(&d)
	total := a*1 + b*0.5 + c*0.2 + d*0.1
	fmt.Println(fmt.Sprintf("Total Amount: $%.2f ", total))
	notes := int16(total / 2)
	fmt.Println("Number of 2-dollar notes can be exchange: ", notes)
	remaining := math.Remainder(float64(total), float64(2))
	fmt.Println(fmt.Sprintf("Remaining change to return after exchanging: %.2f", remaining))
}
func intfloatcomplextypes4() {
	var x float64
	fmt.Println("Please enter the first length: ")
	fmt.Scanln(&x)
	var y float64
	fmt.Println("Please enter the second length: ")
	fmt.Scanln(&y)
	var z float64
	fmt.Println("Please enter the angle in radian: ")
	fmt.Scanln(&z)
	resultant := math.Sqrt(x*x + y*y - 2*x*y*math.Cos(z))
	fmt.Println("Resultant length: ", resultant)
}
func string1() {
	list1 := [4]string{"ans", "wer", "is", "of"}
	list2 := [4]string{"-", "+", "*", "/"}
	list3 := [4]string{"5", "10", "4", "0"}

	x, err := strconv.Atoi(list3[0])
	if err != nil {
		// ... handle error
		panic(err)
	}
	y, err := strconv.Atoi(list3[2])
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println(list1[0]+list1[1]+" "+list1[3]+" "+list3[0]+" "+list2[1]+" "+list3[2]+" "+list1[2], x+y)
}
func ifstatements1() {
	x := 17
	y := 24
	if x > y {
		fmt.Println(x, "is bigger than ", y, "by ", x-y)
	} else {
		fmt.Println(y, "is bigger than ", x, "by ", y-x)
	}
}
func ifstatements2() {

	x := 17
	if len(fmt.Sprint(x)) > 1 {
		fmt.Println(x, "have more than 1 digit")
	} else {
		fmt.Println(x, "have 1 digit only")
	}
	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else if x%2 == 1 {
		fmt.Println(x, "is odd")
	}
}
func ifstatements3() {
	var x int
	fmt.Println("Please enter an integer: ")
	fmt.Scanln(&x)
	if x%5 > 0 || x%6 > 0 {
		fmt.Println(x, "is not divisible by either 5 or 6")
	} else {
		fmt.Println(x, "is divisible by both 5 and 6")
	}
}
func ifstatements4() {
	var x string
	fmt.Println("Please input your username: ")
	fmt.Scanln(&x)
	if x == "Admin" {
		fmt.Println("Welcome, Admin!")
	} else if x == "Robin" || x == "John" {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("Merry Men")
	}
}
func ifstatements5() {
	var x int
	fmt.Println("Please input a year: ")
	fmt.Scanln(&x)
	if x%100 == 0 {
		if x%400 == 0 {
			fmt.Println(x, "is a leap year!")
		} else {
			fmt.Println(x, "is not a leap year!")
		}
	} else if x%4 == 0 {
		fmt.Println(x, "is a leap year!")
	} else {
		fmt.Println(x, "is not a leap year!")
	}
}
func forstatements1() {
	// print 0-1000
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}
	// print even number
	for i := 0; i <= 1000; i = i + 2 {
		fmt.Println(i)
	}
	// print odd number
	for i := 1; i <= 1000; i = i + 2 {
		fmt.Println(i)
	}
}
func forstatements2() {
	var i int
	for {
		fmt.Println("To exit, enter -1\nPlease enter a number: ")
		fmt.Scanln(&i)
		if i == -1 {
			break
		}
		if i%2 == 0 {
			fmt.Println(i, "is an even number")
		} else {
			fmt.Println(i, "is an odd number")
		}
	}
}
func forstatements3() {
	var x, y int
	fmt.Println("Please enter first number: ")
	fmt.Scanln(&x)
	fmt.Println("Please enter second number: ")
	fmt.Scanln(&y)
	// if x is even print all even and print odd
	// if x is odd print all odd and print even after
	for i := x; i <= y; i = i + 2 {
		fmt.Println(i)
	}
	for i := x + 1; i <= y; i = i + 2 {
		fmt.Println(i)
	}
}
func forstatements4() {
	var x int
	for {
		fmt.Println("Enter a value:")
		fmt.Scanln(&x)
		if x%2 == 0 {
			fmt.Println(x, "is an even number")
		} else {
			fmt.Println(x, "is an odd number")
		}
		if x > 9 {
			fmt.Println(x, "has more than 1 digit")
		} else {
			fmt.Println(x, "does not have more than 1 digit")
		}
	}
}
func forstatements5() {
	var x int
	var y int
	fmt.Println("Please input first number: ")
	fmt.Scanln(&x)
	fmt.Println("Please input second number: ")
	fmt.Scanln(&y)
	for i := x; i < y; i++ {
		fmt.Println(i)
	}
	for i := y; i >= x; i-- {
		fmt.Println(i)
	}
}
func forstatements6() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter Name: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	x := rand.Intn(99) + 1
	var y int
	for i := 0; i < 5; i++ {
		fmt.Println("Make a guess: ")
		fmt.Scanln(&y)
		if y > x {
			fmt.Println("lower please")
		} else if y < x {
			fmt.Println("higher please")
		} else {
			fmt.Println("Well Done you are right!")
		}
	}
}
func forstatements7() {
	var x string
	fmt.Println("Please enter a vheicle plate number: ")
	fmt.Scanln(&x)
	const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := float64(0)
	y := strings.ToUpper(x[len(x)-1:])
	z := strings.ToUpper(string(x[0]))

	//check first and last char if it is a letter
	if !strings.Contains(alpha, y) || !strings.Contains(alpha, z) {
		fmt.Println("The vehicle plate given is not correct!")
		os.Exit(0)
	}

	//check the len must not be more than 8
	if len(x) > 8 {
		fmt.Println("The vehicle plate given is not correct!")
		os.Exit(0)
	}
	list1 := []int{2, 3, 4, 5}
	list2 := []int{4, 9}
	a := 0
	b := 0
	c := 0
	for i := len(x) - 2; i >= 0; i-- {
		if n, err := strconv.Atoi(string(x[i])); err == nil {
			if b <= 0 {
				total += float64(n * list1[a])
			} else {
				fmt.Println("The vehicle plate given is not correct!")
			}
			a += 1
		} else {
			b += 1
			total += float64((strings.Index(alpha, strings.ToUpper(string(x[i]))) + 1) * list2[c])
			c += 1
			if c > 1 {
				break
			}
		}
	}
	remaining := math.Mod(total, 19)
	checksum := "AZYXUTSRPMLKJHGEDCB"
	//check if checksum is the same as the last char
	if string(checksum[int64(remaining)]) == y {
		fmt.Println("The vehicle plate given is correct!")
	} else {
		fmt.Println("The vehicle plate given is not correct!")
	}
}
func switchstatements1() {
	x := time.Now().Weekday()
	fmt.Println(x)
	switch {
	case x%2 == 0:
		fmt.Println("today is even")
	case x%2 > 0:
		fmt.Println("today is odd")
	}
}
func switchstatements2() {
	var x float64
	var y float64
	fmt.Println("Please input your height in metres: ")
	fmt.Scanln(&x)
	fmt.Println("Please input your weight in kg: ")
	fmt.Scanln(&y)
	bmi := y / (x * x)
	switch {
	case bmi < 18.5:
		fmt.Println("Underweight")
	case bmi >= 18.5 && bmi <= 24.9:
		fmt.Println("Healthy Weight")
	case bmi >= 25 && bmi <= 29.9:
		fmt.Println("Overweight")
	case bmi >= 30 && bmi <= 34.9:
		fmt.Println("Obese")
	case bmi >= 35 && bmi <= 39.9:
		fmt.Println("Severly Obese")
	case bmi >= 40:
		fmt.Println("Morbidly Obese")
	}
}
func switchstatements3() {
	// output will be 1>10
	switch {
	case 20 > 40:
		fmt.Println("20 > 40")
	case 10 > 1:
		fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 1")
	}

	//no output
	switch {
	case 10 > 11:
		fmt.Println("10 > 11")
	case 5 > 1:
		// fallthrough
	case 1 > 10:
		fmt.Println("1 > 10")
	default:
		fmt.Println("None 2")
	}

	//no output
	switch {
	case 20 > 21:
		fmt.Println("20 > 21")
	case 10 > 1:
		break
	case 20 > 10:
		fmt.Println("20 > 10")
	default:
		fmt.Println("None")
	}
}
func compositetypes1to4() {
	array1 := []string{"Windows XP", "Linux 1.0", "Raspbian Teensy", "MacOS", "IOS", "Google Android"}
	//Activity 1
	// for i, _ := range array1 {
	// 	fmt.Println(len(array1[i]))
	// }
	//Activity 2
	mapping := make(map[string]string)
	for _, s := range array1 {
		mapping[s] = ""
	}
	mapping["Windows XP"] = "Windows 10"
	mapping["Linux 1.0"] = "Linux 16.04"
	mapping["Raspbian Teensy"] = "Raspbian Buster"

	//Activity 3
	mapping["Ubuntu"] = ""
	mapping["MS-Dos"] = ""
	mapping["Solaris"] = ""

	//Activity 4
	first3elements := make(map[string]string)
	next3elements := make(map[string]string)
	last3elements := make(map[string]string)
	i := 0
	for k, v := range mapping {
		if i < 3 {
			first3elements[k] = v
		} else if i < 6 {
			next3elements[k] = v
		} else {
			last3elements[k] = v
		}
		i++
	}
	fmt.Println("first 3 elements are: ", first3elements)
	fmt.Println("next 3 elements are: ", next3elements)
	fmt.Println("last 3 elements are: ", last3elements)
}
func compositetypes5() {
	total := float64(0)
	array1 := []float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}
	for i := 0; i < len(array1); i++ {
		total += array1[i]
	}
	fmt.Println(fmt.Sprintf("Average temperature for the day: %.2f ", total/24))
}
func slices1to5() {
	//Activity 1
	slice := make([]float64, 0, 5)
	slice = append(slice, 9.5, 8, 10.2, 7.4)
	fmt.Println("Length of the slice: ", len(slice))
	fmt.Println("Capacity of the slice: ", cap(slice))

	//Activity 2
	slice[2] = 9.8
	fmt.Println(slice)

	//Activity 3
	slice = append(slice, 8.4, 9.4, 7.2)
	fmt.Println("Length of the slice: ", len(slice))
	fmt.Println("Capacity of the slice: ", cap(slice))

	//Activity 4
	subslice := slice[3:]
	fmt.Println(subslice)

	//Activity 5
	dictionary := make(map[string][]int)
	dictionary["Room 1"] = []int{20, 21, 23, 25, 22}
	dictionary["Room 2"] = []int{27, 23, 25, 20, 30, 24}
	dictionary["Room 3"] = []int{22, 23, 24, 22}
	for k, v := range dictionary {
		total := int(0)
		for i := 0; i < len(v); i++ {
			total += v[i]
		}
		fmt.Println(fmt.Sprintf("Average Temperature for %v is: %.2f", k, float64(total)/float64(len(v))))
	}
}
func struct1to2() {
	//Activity 1
	type customer struct {
		fname            string
		lname            string
		Age              int
		Subscriber       bool
		HomeAddress      string
		Phone            int
		CreditAvailable  float32
		CurrentCartCost  float32
		CurrentOrderCost float32
	}

	//Activity 2
	customer1 := customer{
		fname:            "Annakin",
		lname:            "Skywalker",
		Age:              45,
		Subscriber:       true,
		HomeAddress:      "Death Star",
		Phone:            1234567,
		CreditAvailable:  10000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00,
	}
	customer2 := customer{
		fname:            "Han",
		lname:            "Solo",
		Age:              50,
		Subscriber:       false,
		HomeAddress:      "Tatooine",
		Phone:            4321765,
		CreditAvailable:  20000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00,
	}
	fmt.Println(customer1)
	fmt.Println(customer2)
}

//Channel Activity 1
//Output: Received 0-9 true
//Channel Activity 2
//Output: 1,2
