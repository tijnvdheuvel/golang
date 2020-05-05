package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
var(
	ipadress string
input string
)
func main(){
	fmt.Println("Welkom bij webinfo requester")
	fmt.Println("Van welke website wilt u de info opzoeken?(e.g. google.com)")
	//User input verwerking
	input=UserInput()
	WebInfo(input)
	//Press enter einde
	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func UserInput() string{
	a := bufio.NewScanner(os.Stdout)
	a.Scan()
	b :=a.Text()
	fmt.Println("U heeft ingevoerd: ", b)
	return b
}

func WebInfo(u string) {
	//Input handling
	if u==""||u=="\n"||strings.Contains(u,"www.")||strings.Contains(u,"https://"){
		fmt.Println("[Invalid input, and/or format]")
		os.Exit(3)
	}
	fmt.Println("**************************")
	//Host input
	fmt.Println("----------------\nIP adresses:")
	hosts, err := net.LookupHost(u)
	if err != nil {
		panic(err)
	}
	ipadress = hosts[0]
	for b, h := range hosts {

		fmt.Print("|[", b, "] ")
		if len(h) < 17 {
			fmt.Println("IPv4: ", h)
		} else {
			fmt.Println("IPV6: ", h)
		}
	}
	fmt.Println("Final common name:")
	cname,err:=net.LookupCNAME(u)
	if err !=nil{
		panic(err)
	}
	fmt.Println("|",cname)
	//Reverse lookup related names
	relname, err := net.LookupAddr(ipadress)
	if err != nil {
		panic(err)
	}
	fmt.Println("----------------\nReverse lookup:")
	if len(relname)==0{
		fmt.Println("no record")
	}
	for a, n := range relname {
		fmt.Println("|[", a, "]: ", n)
	}
//Check DNS servers
fmt.Println("-------------\nDNS NS records:")
	DNS,err :=net.LookupNS(u)
	if err!=nil {
		p := err.Error()
		if !strings.Contains(p,"no such host"){
			panic(err)
		}

	}
	if len(DNS)==0{
		fmt.Println("|No records found")
	}else {
		for i,s:=range DNS{
			fmt.Println("|[",i,"]: ",s.Host)
		}
	}
	//Check DNS TXT records
	fmt.Println("----------------\nDNS TXT records:")
	TXT, err :=net.LookupTXT(u)
	if len(TXT) ==0||TXT==nil{
		fmt.Println("|No records found.")
	} else if err !=nil && TXT!=nil{
		panic(err)
	}else {

		for i, d := range (TXT) {
			fmt.Println("|[", i, "]: ", d)
		}
	}
	fmt.Println("----------------")
	fmt.Println("**************************")

}
