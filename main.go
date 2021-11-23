package main


import (
	"net/http"
	"log"
	"net"
	"bytes"
)

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
func getMac(w http.ResponseWriter, r *http.Request) {
	log.Printf("getMac")
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	for _, a := range as {
		//fmt.Println(a)
		buffer.WriteString(a+",")
	}
	w.Write([]byte(buffer.String()))
}

func main(){
	http.HandleFunc("/mac",    getMac)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
