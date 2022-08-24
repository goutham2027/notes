package main

import "fmt"

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	companions []string
// }
// type Animal struct {
// 	Name   string
// 	Origin string
// }
type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

// type Bird struct {
// 	Animal
// 	SpeedKPH float32
// 	CanFly   bool
// }

func main() {
	// fmt.Println("Hello World!")

	// var i int
	// i = 42

	// var i int = 42

	// i := 42
	// fmt.Println(i)

	// var b bool
	// fmt.Printf("%v, %T\n", b, b)

	// n := 42
	// fmt.Printf("%v, %T\n", n, n)

	// names := [...]string{"Apple", "Bat"}
	// var names [2]string
	// names[0] = "Apple"
	// names[1] = "Bat"
	// fmt.Println(names)
	// fmt.Printf("Length of array is %d\n", len(names))

	// names := [3]string{"Apple", "Bat", "Cat"}
	// for i, name := range names {
	// 	fmt.Printf("%s at %d index\n", name, i)
	// }
	// // var scores = make(map[int]string)
	// // scores[0] = "zero"
	// // scores[1] = "one"
	// // elem, ok := scores[9]
	// // fmt.Println(elem)
	// fmt.Println(ok)
	// fmt.Println(scores)
	// for k, v in scores {
	// 	fmt.Println("Key %d Value %s", k, v)
	// }

	// name := "hello world This is a great world"
	// words := strings.Split(name, " ")
	// var count = make(map[string]int)
	// for _, word := range words {
	// 	count[word] += 1
	// }
	// fmt.Println(count)

	// nums := []int{1, 2, 3, 4, 5}
	// sub_nums := nums[3:]
	// fmt.Println(sub_nums)
	// fmt.Printf("length %d, capacity %d", len(sub_nums), cap(sub_nums))
	// sub_nums[0] = -4
	// fmt.Println(sub_nums)
	// fmt.Println(nums)

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))
	// a = append(a, 2, 3, 4, 5)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := make([]int, 10, 100)
	// fmt.Println(a)
	// a = append(a, 5)
	// fmt.Println(a)

	// slice := make([]int, 10, 100)
	// fmt.Printf("slice - %T\n", slice)
	// var arr []int
	// fmt.Printf("arr - %T", arr)

	// statePopulations := map[string]int{
	// 	"California": 1234,
	// 	"Texas":      8384,
	// }
	// fmt.Println(statePopulations)

	// statePopulations := make(map[string]int)
	// statePopulations["California"] = 123
	// fmt.Println(statePopulations)

	// struct
	// aDoctor := Doctor{
	// 	number:    3,
	// 	actorName: "Jon",
	// 	companions: []string{
	// 		"Liz", "Jo",
	// 	},
	// }

	// fmt.Println(aDoctor)
	// fmt.Println(aDoctor.actorName)

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b)

	// a := Animal{Name: "emu", Origin: "Australia"}
	// fmt.Println(a)

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)

	// switch 3 {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not one or two")
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// i := 0
	// for {
	// 	if i > 5 {
	// 		break
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// 	i++
	// }

	// Loop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	// s := []int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }

	// defer
	// fmt.Println("start")
	// defer fmt.Println("middle")
	// fmt.Println("end")

	// start - Defer example
	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()

	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)
	// end - Defer example

	// a := "start"
	// defer fmt.Println(a)
	// a = "end"

	// panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("Something bad happened")
	// fmt.Println("end")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// a := 42
	// b := a
	// fmt.Println(a, b)
	// a = 27
	// fmt.Println(a, b)

	// var a int = 42
	// var b *int = &a
	// fmt.Println(a, *b)
	// a = 27
	// fmt.Println(a, *b)

	// a := []int{1, 2, 3}
	// b := a
	// fmt.Println(a, b)
	// b[0] = 4
	// fmt.Println(a, b)

	// a := map[string]string{
	// 	"a": "apple",
	// }
	// b := a
	// fmt.Println(a, b)
	// a["a"] = "ant"
	// fmt.Println(a, b)

	sayMessage("hello go!")
}

func sayMessage(msg string) {
	fmt.Println(msg)
}
