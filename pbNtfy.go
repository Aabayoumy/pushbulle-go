package main

import (
	"fmt"
	"net"
	"os"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

func main() {
	// Set the access token.
	token := "o.JC28TTeO5d9KxwA1WuIoRisoLOAL8WCL"
	ip := GetLocalIP()
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	// Create a client for Pushbullet.
	pb := pushbullet.New(token)

	// Create a push. The following codes create a note, which is one of push types.
	n := requests.NewNote()
	n.Title = name + " is up"
	n.Body = "IP : " + ip

	// Send the note via Pushbullet.
	if _, err := pb.PostPushesNote(n); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
