package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	Id      int
	Name    string
	Address string
	Jobs    string
	Reason  string
}

func main() {
  args := os.Args[1]
  id, _ := strconv.Atoi(args)
  students := []student{
    {
      Id:      1,
      Name:    "Daffa Haryadi",
      Address: "Depok",
      Jobs:    "Backend Devs",
      Reason:  "Melanjutkan Pembelajaran golang dan project",
    },
    {
      Id:      2,
      Name:    "M yoga irvandi",
      Address: "Sumatra Selatan",
      Jobs:    "Mahasiswa",
      Reason:  "karena Golang merupakan bahasa baru,yg menarik untuk di pelajari bagi saya.",
    },
  }
  id--
  res := students[id].printStudent(id)
  fmt.Println(res)
}

func (s student) printStudent(i int) string {
  return fmt.Sprintf(" Nama %s dengan id %d dari %s alasan %s", s.Name, s.Id, s.Address, s.Reason)
}
