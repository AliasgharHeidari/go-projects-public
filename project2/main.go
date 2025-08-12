package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Students struct {
	Name       string
	Id         int
	Age        string
	SignUpTime time.Time
}

func memUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc : %.2f MiB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Totall Alloc : %.2f Mib\n", float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("Sys : %.2f Mib\n", float64(m.Sys)/1024/1024)
	fmt.Printf("NumGC : %v\n", m.NumGC)
}

func main() {
	fmt.Println("--------------------Student sign up------------------")
	rand.Seed(time.Now().UnixNano())
	var s Students

	fmt.Print("Enter your name : ")
	fmt.Scan(&s.Name)
	var int_age int
	var err error
	for {

		fmt.Print("Enter your age : ")
		fmt.Scan(&s.Age)

		bufio.NewReader(os.Stdin).ReadBytes('\n')

		int_age, err = strconv.Atoi(s.Age)

		if err != nil {
			fmt.Println("Invalid age. Please tap Enter then enter a number.")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}
		break
	}

	s.Id = rand.Intn(90000) + 10000
	s.SignUpTime = time.Now()

	fmt.Println("---------------------------you're succesfully regestered------------------------")
	fmt.Println("Name : ", s.Name)
	fmt.Println("Age : ", int_age)
	fmt.Println("Student id : ", s.Id)
	fmt.Println("regestration time : ", s.SignUpTime)
	fmt.Println("---------------------------memory usage-------------------------------")
	memUsage()
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Press Enter to exit...")
	var exit string
	fmt.Scanln(&exit)
}
