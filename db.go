package main

import "fmt"



type myNumber int

func(num myNumber) square() myNumber  {
	if num==0{
		return 1
	}
	return num*num
}

func main(){
	num:=myNumber(25)
	sq:=num.square()
	fmt.Printf("the square of %d is %d \n", num, sq)
}

/*
type User struct {
	Name  string
	Email string
}

func (u User) userDetails() string {
	return fmt.Sprintf("User name is  : %s and email is : %s", u.Name, u.Email)
}

func main() {
	user1 := User{Name: "Johon doe", Email: "jhondoe@gmail.com"}
	fmt.Println(user1.userDetails())
}*/

/*

func sendMessage(msg chan string) {
	msg <- "Hello From GoLinux Academy"
}

func main() {

	res := make(chan string)

	go sendMessage(res)

	fmt.Println("Rsponse of chabnel %v", <-res)
}*/

/*
func main() {
	resualts := make(map[string]int, 0)

	resualts["Google"] = 300000
	resualts["Chrome"] = 4000
	resualts["GoLinux"] = 800000
	resualts["AWS"] = 60000

	fmt.Println("original values : %v \n", resualts)
	fmt.Println("Pass-values   to  function : %v \n", resualts)
	fmt.Println("original values after function exution : %v \n, ", resualts)

}

func UPdateCompanyNet(data map[string]int) map[string]int {
	data["Google"] = 9000
	data["AWS"] = 2000
	return data
}*/

/*

type Company struct {
	Name        string
	Description string
}

func UPdateCompanyDetails(c Company) Company {
	c.Name = "GoLinux"
	c.Description = "Online Study Gurus"

	return c
}

func main() {
	res := Company{
		Name:        "JetBrains",
		Description: "Proficient in IDEs",
	}

	fmt.Println("original  struct set values : %v \n", res)
	fmt.Println("Pass values to function update:", UPdateCompanyDetails(res))
	fmt.Println("Original  struct set values after excution: %v", res)
}*/

/*
func multiply(n int) int {
	return n * n
}

func main() {
	val := 3

	fmt.Println("Orginal :", val)
}
*/
/*
func init() {
   fmt.Println("<<< First >>>")
}
func init() {
   fmt.Println("<<< Second >>>")
}
func init() {
   fmt.Println("<<< Third >>>")
}

func main() {
   fmt.Println("I execute after init() functions")
}*/
