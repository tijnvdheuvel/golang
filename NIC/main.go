package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	input string
AdaptIndex int
	)
func main(){
	fmt.Println("Welkom bij NIC-view. \nU gaat zometeen uw NIC's zien.")
	//Command alle interfaces
AlleInterfaces()
fmt.Println("Specificeer het ID van de interface waarvan u de specificaties wilt zien. [eg. 1 (int)]" )
	scanner:= bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	if input ==""{
		AdaptIndex =1
		fmt.Println("U heeft niks ingevuld, uw Index is 1.")
	}else {
		n, err := strconv.Atoi(input)
		{
			if err != nil {
				n=1
				AdaptIndex = n
				fmt.Println("U heeft niks ingevuld, uw index is 1.")
			}else {
				fmt.Println("U heeft gekozen voor index: ", input)
				AdaptIndex=n
			}
		}
	}
	//Info geven over NIC
SpecifiekeInterface(AdaptIndex)

		fmt.Print("Press 'Enter' to exit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}

func AlleInterfaces(){
	fmt.Println("*************************")
	fmt.Println("ID:  Name:")
	ifaces, err := net.Interfaces()
	var samenvoegsel string
	if err != nil{
		panic(err)
}
for _,i := range ifaces{
	index := i.Index
	name := i.Name
	//Lengte kloppend maken
	samenvoegsel = (strconv.Itoa(index))
	if len(samenvoegsel) == 4{
		samenvoegsel += ": " + name
	}
	if len(samenvoegsel) == 3{
		samenvoegsel+=":  "+ name
	}
	if len(samenvoegsel) < 2 {
		samenvoegsel +=":   "+ name
	}else {samenvoegsel += ":   " + name
	}
	fmt.Println(samenvoegsel)
	samenvoegsel=""

}
	fmt.Println("*************************")
}
func SpecifiekeInterface(index int){
	fmt.Println("*************************")
	i, err := net.InterfaceByIndex(index)
	if err != nil{
		panic(err)
	}
	fmt.Println(" De naam van uw interface is: ", i.Name)
	fmt.Println(" Interface flags: ",i.Flags)
	fmt.Println(" De maximale transmissie limiet is: ", i.MTU)

	//Check voor hardware addr
	if i.HardwareAddr !=nil{
		fmt.Println(" Het hardware adres is:", i.HardwareAddr)
	}else{
		fmt.Println(" Het hardware adres is helaas niet gevonden.")
	}

	//Print alle IP'adressen
	fmt.Println("----------------")
	fmt.Println("|Unicast adressen:")
	uniaddrs, err := i.Addrs()
	if err!= nil{
		panic(err)
	}
	for b,p := range uniaddrs{
		fmt.Println("|",b,": ",p)
	}
	fmt.Println("----------------")
	fmt.Println("----------------")
	fmt.Println("|Multicast adressen:")
	multiaddrs, err := i.MulticastAddrs()
	if err!= nil{
		panic(err)
	}
	for b,p := range multiaddrs{
		fmt.Println("|",b,": ",p)
	}
	fmt.Println("----------------")

	fmt.Println("*************************")
}
