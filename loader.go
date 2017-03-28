package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"os"
	"strings"
	"time"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	r, err := redis.Dial("tcp", ":6389")
	defer r.Close()

	str, err := r.Do("GET", "ololo")
	PanicOnError(err)
	fmt.Printf("Value: %s", str)

	defer timeTrack(time.Now(), "iMPORT")

	f, err := os.Open("data/soc-pokec-profiles.txt")

	rdr := bufio.NewReader(f)

	i := 0
	rec, isPref, err := rdr.ReadLine()

	if isPref {
		// 	s, isPref, err := rdr.ReadLine()

		// 	rec = append(rec, s)
	}

	for err != io.EOF {
		i++
		if i%100000 == 0 {
			fmt.Println("Processed so far :", i)
		}

		// if rec == nil {}
		arr := strings.Split(string(rec), "\t")
		_, err := r.Do("SETEX", arr[0], 3600, arr)

		// fmt.Printf("%s\n", strings.Split(string(rec), "\t"))

		rec, isPref, err = rdr.ReadLine()
		PanicOnError(err)
	}
	fmt.Println("Processed all :", i)

}
