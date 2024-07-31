package main

import (
	_ "github.com/lib/pq"

	_ "github.com/sirupsen/logrus"
	_ "go.uber.org/zap"
)

/*
type Country struct {
	ID       int
	Name     string
	Currency string
}

type Struct1 struct {
	x int
	y int
}

type MyStruct struct {
	Name    string `json:"age"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type contact struct {
	phone  int
	email  string
	street string
}
type user struct {
	firstname      string
	age            int
	isAwesome      bool
	height         float64
	contactDetails contact
}*/

type User struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	IsAwesom bool    `json:"is_awesome"`
	Height   float64 `json:"height"`
}

func main() {
	/*
				db, err := sql.Open("postgres", "host=localhost  port=5432 user=admin database=students password=1999 sslmode=disable")
				if err != nil {
					panic(err)
				}
		go get -u go.uber.org/zap

				if _, err := db.Exec(`delete from countries where id=$1`, id); err != nil {
					panic(err)
				}*/

	/*if _, err := db.Exec(`update countries set name=$1,currency=$2 where id=$3`, "malaziya", "malaz", 12); err != nil {
		panic(err)
	}

	fmt.Println("update data")*/
	/*
		countries := []Country{}

		rows, err := db.Query(`select *from countries`)
		if err != nil {
			panic(err)
		}

		for rows.Next() {
			country := Country{}
			if err := rows.Scan(&country.ID, &country.Name, &country.Currency); err != nil {
				panic(err)
			}

			countries = append(countries, country)
		}

		fmt.Println(countries)*/
	/*
		if _, err := db.Exec(`insert into countries(id,name,currency) values ($1,$2,$3)`, c.ID, c.Name, c.Currency); err != nil {
			panic(err)
		}
		fmt.Println(c)
	*/
	/*
	   ****************SELECT  ***************

	   	country := Country{}

	   	row := db.QueryRow(`select id,name ,currency from countries`)
	   	if err = row.Scan(&country.ID, &country.Name, &country.Currency); err != nil {
	   		panic(err)
	   	}

	   	fmt.Println(country)
	*/
	/*i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j*/

	/*
		fmt.Println(employees)
		fmt.Println(reflect.TypeOf(employees))*/

	/*if employees == nil {
		fmt.Println("this is map is not nil")
	} else {
		fmt.Println("this is map is nil")
	}

	employees["John Doe"] = 1

	fmt.Println(employees)*/
	/*employees := map[string]int{
		"jhon doe": 23,
		"mary doe": 19,
	}   contactDetails contact

	fmt.Println(employees)*/

	/*employees := make(map[string]int, 5)

	fmt.Println(employees)

	if employees == nil {
		fmt.Println("employees map is nil")
	} else {
		fmt.Println("employese map is not nil")
	}

	employees["jhon doe"] = 23
	employees["mary doe"] = 19

	fmt.Println(employees)*/
	/*
		fotbalclub := map[string]int{
			"ranaldo": 7,
			"Messi":   10,
			"neymar":  11,
			"mapae":   2,
		}

		fotbalclub["kristiano"] = 322

		delete(fotbalclub, "Messi")
		fmt.Println(fotbalclub)*/
	/*
		users := map[string]int{
			"jhon":  21,
			"david": 43,
			"paul":  54,
		}

		value, exist := users["john"]

		fmt.Println(value)
		fmt.Println(exist)*/

	/*
		for username, userage := range users {
			fmt.Println(username, userage)
		}
	*/
	/*
		users := map[int]string{
			100: "John",
			23:  "Mary",
			54:  "Paul",
			21:  "Josh",
		}
		var keys []int

		for key := range users {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		for _, key := range keys {
			fmt.Printf("Key : %d and Value : %s \n", key, users[key])
		}
	*/
	/*
		map1 := map[string]interface{}{
			"name":    "Amit Kumar",
			"age":     30,
			"address": "316 Some Main Road ,Creuasia",
		}

		jsonDate, _ := json.Marshal(map1)

		var structData MyStruct
		json.Unmarshal(jsonDate, &structData)

		fmt.Println(structData)*/

	/*logrus.Warn("Hello world")

	fmt.Println("*****************")

	logrus.Trace("Trace  log")
	logrus.Debug("debug log")
	logrus.Info("Info log")
	logrus.Warn("warm log")
	logrus.Error("error log")
	logrus.Fatal("fata; log")
	logrus.Panic("panic log")*/
	/*
		const (
			Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
			Ltime                         // the time in the local time zone: 01:23:23
			Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
			Llongfile                     // full file name and line number: /a/b/c/d.go:23
			Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
			LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
			Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
			LstdFlags     = Ldate | Ltime // initial values for the standard logger
		)*/

	/*log.Println("default log message") // default

	log.SetFlags(log.Ldate)

	log.SetFlags(log.Ldate | log.Ltime)*/
	/*
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		logger.Info("This is an info log with fields ", zap.String("key", "value"))
	*/

	/*user1 := User{firstname: "jhon", age: 24, isAwesome: true, height: 170}
	fmt.Println("firstname", user1.firstname)
	fmt.Println("Age ", user1.age)
	fmt.Println("is awesome", user1.isAwesome)
	fmt.Println("Height", user1.height)*/

	/*user1 := User{}
	var user2 User

	fmt.Println(user1)
	fmt.Println(user2)

	user1.firstname = "jhon"
	user1.age = 54
	user1.isAwesome = true
	user1.height = 4.3

	fmt.Println(user1)*/
	/*
		user1 := new(User)
		fmt.Println(user1)

		user1.firstname = "John"
		user1.age = 54
		user1.isAwesome = true
		user1.height = 4.3

		fmt.Println(*user1)
		fmt.Println(user1)*/
	/*
		contact1 := contact{
			phone:  234567,
			email:  "jhondoe@gmail.com",
			street: "4th go stree",
		}

		user1 := user{
			firstname:      "Jhon",
			age:            34,
			isAwesome:      true,
			height:         5.6,
			contactDetails: contact1,
		}

		fmt.Println(user1)
		fmt.Println(user1.contactDetails)
		fmt.Println(user1.contactDetails.phone)
		fmt.Println(user1.contactDetails.email)
		fmt.Println(user1.contactDetails.street)*/
	/*
		user1 := User{
			Username: "jhon doe",
			Email:    "jhondoe@gmail.com",
			IsAwesom: true,
			Height:   170,
		}

		user1.userInfo()
		user1.chaneEmail("changes@gmail.com")
		user1.userInfo()

		fmt.Println(user1)*/
}

/*
func (u User) userInfo() {
	response, err := json.Marshal(u)
	if err != nil {
		log.Fatal("Error : %v", err)
	}
	fmt.Println(string(response))
}

func (u *User) chaneEmail(email string) {
	u.Email = email
}*/
