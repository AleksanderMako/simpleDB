package main

import (
	"fmt"
	"simpleDB/database"
	"time"
)

func main() {

	simpleDB := database.NewSimpleDB()

	key := "key"

	for i := 0; i < 1000001; i++ {
		err := simpleDB.Write(fmt.Sprintf(key+"%v", i), fmt.Sprintf("%v", i))
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("sleeping...")
	time.Sleep(time.Second * 2)
	fmt.Println("Starting search...")
	start := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(simpleDB.FastGet("key1000000"))
	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("It took %v UnixNano \n", end-start)
	fmt.Println("Ending search...")

	start = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(simpleDB.Get("key1000000"))
	end = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("It took %v UnixNano \n", end-start)
}
