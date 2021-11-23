package main


import (
	"net/http"
	"log"
	"fmt"
	"net"
	"bytes"
)
func mac() string {
	// 获取本机的MAC地址
	interfaces,err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	var buffer bytes.Buffer
	for _,inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		fmt.Println("MAC = ",mac)
		if "" != mac.String() {
			buffer.WriteString(mac.String()+",")
		}
	}

	return buffer.String()

}

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
	w.Write([]byte(string(mac())))
}

//func main(){
//	http.HandleFunc("/mac",    getMac)
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}

func main() {
	as, err := getMacAddr()
	if err != nil {
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	for _, a := range as {
		//fmt.Println(a)
		buffer.WriteString(a+",")
	}
	fmt.Println(buffer.String())
}
