package main

import (
	"bufio"
	server "databaseServer"
	"encoding/json"
	"fmt"
	"net"
	"simpleDB/database"
)

func handleConnection(conn net.Conn, dbServer server.Server) {

	data, err := bufio.NewReader(conn).ReadBytes('\r')
	if err != nil {
		fmt.Println(err)
		conn.Close()
	}
	var requestPayload server.RequestPayload
	if err = json.Unmarshal(data, &requestPayload); err != nil {

		fmt.Println("failed to unmarshall payload" + err.Error())
		conn.Write(append([]byte(err.Error()), '\n'))
		conn.Close()
	}
	if val, err := dbServer.Protocol(requestPayload); err != nil {
		fmt.Println(err)
		conn.Write(append([]byte(err.Error()), '\n'))
		conn.Close()
	} else {
		conn.Write(append([]byte(val), '\n'))
	}
	defer conn.Close()
}
func main() {

	simpleDB := database.NewSimpleDB()

	// key := "key"

	// for i := 0; i < 1000001; i++ {
	// 	err := simpleDB.Write(fmt.Sprintf(key+"%v", i), fmt.Sprintf("%v", i))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// fmt.Println("sleeping...")
	// time.Sleep(time.Second * 2)
	// fmt.Println("Starting search...")
	// start := time.Now().UnixNano() / int64(time.Millisecond)
	// fmt.Println(simpleDB.FastGet("key1000000"))
	// end := time.Now().UnixNano() / int64(time.Millisecond)
	// fmt.Printf("It took %v UnixNano \n", end-start)
	// fmt.Println("Ending search...")

	// start = time.Now().UnixNano() / int64(time.Millisecond)
	// fmt.Println(simpleDB.Get("key1000000"))
	// end = time.Now().UnixNano() / int64(time.Millisecond)
	// fmt.Printf("It took %v UnixNano \n", end-start)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	fmt.Println("started server")
	dabaseServer := server.NewServer(*simpleDB)
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn, *dabaseServer)
	}
}
