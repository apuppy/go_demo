package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	pb "./message"
)

func promptForAddress(r io.Reader) (*pb.Person, error)  {
	p := &pb.Person{}

	rd := bufio.NewReader(r)
	fmt.Print("Enter person ID number : ")

	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}

	fmt.Print("Enter name : ")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Name = strings.TrimSpace(name)

	fmt.Print("Enter email address (blank for none): ")
	email, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Email = strings.TrimSpace(email)

}