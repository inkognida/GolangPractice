package main //making package for standalone executable

import (
	"fmt" // importing a package
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func	HelloWorld() string{	// function with Hello World example; return string;
	return "Hello, World!"
}

func	VariablesInit() (int, int){
	var a, b int // block with simple examples

	a = 10
	b = 5

	var _bool bool // block with specifics
	var _int = 123
	var __int = 456

	fmt.Println(_bool, _int, __int)

	_preferInt := 25

	fmt.Println("Case with := initialization", _preferInt)

	return a, b
}

func	TypeCasting(data float64) int{
	return int(data)
}

func	ConstantsAndEnums(){
	const PI = 3.14
	const STR string = "Hello, world!"

	fmt.Println(PI, STR)

	const (
		MALE = 0
		FEMALE = 1
		ANIMAL = 2
	)

	fmt.Printf("%d %d %d\n", MALE, FEMALE, ANIMAL)
}

type Celsius float32
type Fahrenheit float32

func	toFahrenheit(celsius Celsius) Fahrenheit{
	var tmp Fahrenheit

	tmp = (Fahrenheit(celsius) * 9/5)+ 32
	return tmp
}

// strings and strconv packages

func	stringMethods(){
	str := "Hey, babe. How u doing"

	fmt.Printf("\n")
	fmt.Println("Strings are here!")

	fmt.Println(strings.HasSuffix(str, "y")) // ends with "Hey"
	fmt.Println(strings.HasPrefix(str, "Hey")) // starts with "ggg"

	fmt.Println(strings.Index(str, "babe")) // find first occurrence
	fmt.Println(strings.LastIndex(str, "ing")) // find last occurrence
	fmt.Println(strings.Contains(str, "u")) // substr "u" in str

	fmt.Println(strings.Replace(str, "Hey", "Hi", 1)) // replace the first n occurrences of old in str by new
	fmt.Println(strings.Count(str, "i"))

	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.TrimSpace(str)) // remove all leading and trailing whitespaces
	fmt.Println(strings.Trim(str, "Hey")) // trim a specific string str from a string s

	fmt.Println(strings.Fields(str)) // splits the string s around each instance of one or more \
										// consecutive white space characters
	fmt.Println(strings.Split(str, "."))

	arr := []string{"a", "b", "c"}

	fmt.Println(strings.Join(arr, ":")) // join elements of arr into string with separator
}

func	strconvMethods(){
	fmt.Println(strconv.Itoa(123)) // default
	fmt.Println(strconv.Atoi("123")) // default <nil> is code error
	fmt.Println(strconv.FormatFloat(123.123, 'f', 2, 64)) // float to string
	fmt.Println(strconv.ParseFloat("123.123", 64)) // string to float <nil> is code error
}

// time package

func	timeMethods(){
	t := time.Now()

	fmt.Println(t.Year(), t.Month(), t.Day())
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("2006 01 02")
	fmt.Println(t, "=>", s)
}

// if, switch-case, select

//goland:noinspection ALL
func	ifElse(arg int) bool{

	if_ := rand.Intn(100)

	if if_ % 2 == 0{
		fmt.Println("fine")
	} else {
		fmt.Println("not fine")
	}

	if arg_ := arg; arg_ > 101{
		fmt.Println("more")
	} else {
		fmt.Println("not more")
	}

	if if_ % 2 == 0 {
		return true
	}
		return false
}

// error Tracking block

func	errorTracking(){
	v, _ := strconv.Atoi("101")
	fmt.Println(v)

	ev, error_ := strconv.Atoi("12")

	if error_ != nil{
		fmt.Println("Error inside original string")
		return
	}
	fmt.Printf("The value of ev is: %d\n", ev)

	file, error__ := os.Open("introOS.txt")
	if error__ != nil {
		os.Exit(1)
	}
	fmt.Println(file)
}

// switch case block

func	switchCase(){
	sc := 0

	switch sc {
	case 0: fallthrough // goes to case 3
	case 3: fmt.Println("Here we are")
	default:
		fmt.Println("default")
	}

	num := rand.Intn(100)

	switch {
	case num < 0: fmt.Println("Negative number")
	case num > 0: fmt.Println("Positive number")
	default:
		fmt.Println("Something unbelievable happened")
	}

	switch str, str_ := "hello", "world"; {
	default:
		fmt.Println(str, str_)
	}
}

func	Season(month int) string {
	switch month {
	case 1, 2, 12: return "Winter"
	case 3, 4, 5: return "Spring"
	case 6, 7, 8: return "Summer"
	case 9, 10, 11: return "Autumn"
	default:
		return "Season unknown"
	}
}

// loops

func	forLoop() {
	iter := 0

	for i := 0; i < 10; i++ {
		iter++
	}
	fmt.Println(iter)

	i := 0
	for i < 5 {
		i++
	}
	fmt.Println(i)

	for {
		fmt.Println(i, "endless loop")
		if i == 10 {
			break
		}
		i++
	}

	str := "Hello, world!"
	for index, value := range str {
		fmt.Println(index, " - ", string(value))
	}

	LABEL:
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, " - ", j)
			if j == 2 {
				continue LABEL
			}
		}
	}


	goto_ := false // goto control statement

	if goto_ {
		HERE:
		fmt.Println("goto identifier", i)
		if i == 12{
			return
		}
		i++
		goto HERE
	}
}

func main() { // entry point to my program
	fmt.Printf("%s\n", HelloWorld())
	fmt.Println(HelloWorld() + ":D")

	var a, b int
	a, b = VariablesInit()
	fmt.Println(a, b)

	fmt.Println(TypeCasting(10.9))

	ConstantsAndEnums()

	ra := rand.Int() // all scope of int
	rb := rand.Intn(10) // random value [0, n), my case is 10

	fmt.Println(ra, rb)

	fmt.Printf("%v\n", toFahrenheit(100))

	stringMethods()
	strconvMethods()
	timeMethods()
	fmt.Println(ifElse(10))
	errorTracking()
	switchCase()

	seasonCall := false
	if seasonCall {
		var seasonInput int
		fmt.Printf("Enter your number of month: ")
		if _, err := fmt.Scan(&seasonInput); err != nil {
			os.Exit(1)
		}
		fmt.Println(Season(seasonInput))
	}
	forLoop()
}
