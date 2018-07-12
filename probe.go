package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
	"strconv"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func loadDataFromSource(source string) (result string, err error) {

	res, err := http.Get(source)

	if err != nil {
		return result, err
	}

	resultBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return result, err
	}

	return string(resultBody), err

}

func SocketClient(ip string, port int) (conection net.Conn, err error) {

	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Println("Error connecting to socket server: ", err)
	}

	return conn, err
}

func writeToSocket(conn net.Conn, message string) (err error) {

	_, err = conn.Write([]byte("#BOM#"))
	_, err = conn.Write([]byte(message))
	_, err = conn.Write([]byte("#EOM#"))

	return err
}

func loadDataAndSendToServerPeriodically(parameters Parameters) {

	conn, socketErr := SocketClient(parameters.socketServerUrl, parameters.serverPort)

	ticker := time.NewTicker(time.Duration(parameters.statusCheckInterval) * time.Second)
	go func() {
		for t := range ticker.C {

			status, err := loadDataFromSource(parameters.source)

			if err != nil {
				log.Println("Error getting actual state: ", err)
				continue
			}

			fmt.Println("Got data (%v)", t)

			if conn != nil {
				socketErr = writeToSocket(conn, status)
			}

			if socketErr != nil {
				log.Println("Reconnecting to socket ... ")
				conn, socketErr = SocketClient(parameters.socketServerUrl, parameters.serverPort)
			}
		}
	}()
}

func main() {

	parameters := parseFlags()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	loadDataAndSendToServerPeriodically(parameters)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
