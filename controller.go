package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ListFamily []Family
var ListPeoples []People
var ListTown []Town

func (p People) Input() {
	var consoleReader = bufio.NewReader(os.Stdin)

	fmt.Println("Nhap so can cuoc cong dan\n")
	iden, _ := consoleReader.ReadString('\n')
	iden = strings.Replace(iden, "\n", "", -1)
	idenInt, _ := strconv.ParseInt(iden, 10, 64)
	p.IdentityNumber = idenInt

	fmt.Println("Nhap ten cua cong dan\n")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	p.Name = name

	fmt.Println("Nhap tuoi cua cong dan\n")
	age, _ := consoleReader.ReadString('\n')
	age = strings.Replace(age, "\n", "", -1)
	ageInt, _ := strconv.ParseInt(age, 10, 64)
	p.Age = ageInt

	fmt.Println("Nhap nghe nghiep cua cong dan\n")
	profes, _ := consoleReader.ReadString('\n')
	profes = strings.Replace(profes, "\n", "", -1)
	p.Profession = profes
	ListPeoples = append(ListPeoples, p)
}
func (f Family) AddFamily() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("so thanh vien trong gia dinh\n")
	num, _ := consoleReader.ReadString('\n')
	num = strings.Replace(num, "\n", "", -1)
	numInt, _ := strconv.ParseInt(num, 10, 64)
	f.NumberMem = numInt

	fmt.Println("Nhap dia chi cua gia dinh\n")
	add, _ := consoleReader.ReadString('\n')
	add = strings.Replace(add, "\n", "", -1)
	f.Address = add

	a := make([]People, numInt)
	for i := 0; i < int(numInt); i++ {
		fmt.Scan(&f.Members[i])
	}
	f.Members = a
}

func (t Town) AddTown() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("so ten khu pho\n")
	name, _ := consoleReader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	t.NameTown = name

	fmt.Println("so gia dinh trong khu pho\n")
	num, _ := consoleReader.ReadString('\n')
	num = strings.Replace(num, "\n", "", -1)
	numInt, _ := strconv.ParseInt(num, 10, 64)
	t.NumFamily = numInt

	familyT := make([]Family, numInt)
	for i := 0; i < int(numInt); i++ {
		fmt.Scan(&t.Familys[i])
	}
	t.Familys = familyT
}
