package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var ips = make([]string, 0)
var hostname = ""

func init() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	hostname = name

	// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.String() != "127.0.0.1" {
				ips = append(ips, ip.String())
			}
		}
	}
}

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "7893"

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}

	return DefaultPort
}

// EchoHandler echos back the request as a response
func EchoHandler(writer http.ResponseWriter, request *http.Request) {

	log.Println("Echoing back request made to " + request.URL.Path + " to client (" + request.RemoteAddr + ")")

	writer.Write([]byte("Hostname: "))
	writer.Write([]byte(hostname))
	writer.Write([]byte("\n"))
	writer.Write([]byte("Server IPs: "))
	writer.Write([]byte(strings.Join(ips, ",")))
	writer.Write([]byte("\n"))

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	// allow pre-flight headers
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")
	request.Write(writer)
}

func KillMyself(w http.ResponseWriter, r *http.Request) {
	go panic("to be or not to be, it's a ...")
}

func main() {
	log.Println("starting server, listening on port " + getServerPort())

	http.HandleFunc("/", EchoHandler)
	http.HandleFunc("/kill", KillMyself)
	http.ListenAndServe(":"+getServerPort(), nil)
}
