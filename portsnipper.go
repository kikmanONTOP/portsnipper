package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const asciiArt = `
MADE BY https://github.com/kikmanONTOP
dMMMMb  .aMMMb  dMMMMb dMMMMMMP .dMMMb  dMMMMb  dMP dMMMMb  dMMMMb  dMMMMMP dMMMMb 
dMP.dMP dMP"dMP dMP.dMP   dMP   dMP" VP dMP dMP amr dMP.dMP dMP.dMP dMP     dMP.dMP 
dMMMMP" dMP dMP dMMMMK"   dMP    VMMMb  dMP dMP dMP dMMMMP" dMMMMP" dMMMP   dMMMMK"  
dMP     dMP.aMP dMP"AMF   dMP   dP .dMP dMP dMP dMP dMP     dMP     dMP     dMP"AMF   
dMP      VMMMP" dMP dMP   dMP    VMMMP" dMP dMP dMP dMP     dMP     dMMMMMP dMP dMP    
																					
`

func isOpen(ip string, port int) string {
	timeout := time.Duration(2 * time.Second)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(ip, strconv.Itoa(port)), timeout)
	if err != nil {
		return fmt.Sprintf("Port %d: Closed or Unreachable", port)
	}
	defer conn.Close()
	return fmt.Sprintf("Port %d: Open", port)
}

func getIP(host string) string {
	ips, err := net.LookupIP(host)
	check(err)
	return ips[0].String()
}

func scanPorts(ip string, ports []int) []string {
	var portStatus []string
	for _, port := range ports {
		status := isOpen(ip, port)
		portStatus = append(portStatus, status)
	}
	return portStatus
}

func writeToFile(filename string, data []string) {
	err := ioutil.WriteFile(filename, []byte(strings.Join(data, "\n")), 0644)
	check(err)
}

func main() {
	fmt.Println(asciiArt)

	fmt.Print("Enter the IP or hostname: ")
	var target string
	_, err := fmt.Scanln(&target)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	ip := getIP(target)

	fmt.Printf("Starting scan on target: %s (%s)\n", target, ip)

	ports := make([]int, 1200)
	for i := range ports {
		ports[i] = i + 1
	}

	portStatus := scanPorts(ip, ports)

	for _, status := range portStatus {
		fmt.Println(status)
	}

	writeToFile("ports.txt", portStatus)

	fmt.Println("Port scan complete. Results saved to 'ports.txt'.")
}
