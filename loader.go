package main

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
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

func importProfiles(rds *redis.Conn, filename string) {

	defer timeTrack(time.Now(), "PROFILES IMPORT")

	//get columns
	var cols []string
	col_file, err := os.Open("data/columns.txt")
	PanicOnError(err)
	defer col_file.Close()

	scanner := bufio.NewScanner(col_file)
	for scanner.Scan() {
		cols = append(cols, scanner.Text())
	}

	fmt.Println("COLUMNS ARE:", cols)

	//get data
	data_file, err := os.Open(filename)
	defer data_file.Close()

	i := 0

	scanner = bufio.NewScanner(data_file)

	const maxCapacity = 256 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		rec := scanner.Text()

		arr := strings.Split(string(rec), "\t")
		key := "person:" + arr[0]

		var data []interface{}

		data = append(data, key)
		data = append(data, zip_arrays(cols, arr)...)

		msg, err := (*rds).Do("HMSET", data...)
		if err != nil {
			fmt.Println("MESSAGE:", msg)
			PanicOnError(err)
		}
		i++
		if i%100000 == 0 {
			fmt.Println("Processed so far :", i)
		}

	}

	fmt.Println("Processed all :", i)
}

func importRelations(rds *redis.Conn, filename string) {

	i := 0
	graph := "pokec"

	// Processed all : 30622564
	// RELATIONS IMPORT took 1h24m45.771996371s
	defer timeTrack(time.Now(), "RELATIONS IMPORT")

	//get data
	data_file, err := os.Open(filename)
	PanicOnError(err)
	defer data_file.Close()

	scanner := bufio.NewScanner(data_file)

	for scanner.Scan() {
		rec := scanner.Text()
		arr := strings.Split(string(rec), "\t")

		person := "person:" + arr[0]
		friend := "person:" + arr[1]

		_, err := (*rds).Do("GRAPH.ADDEDGE", graph, person, "friend", friend)
		if err != nil {
		}
		break
		i++
		if i%100000 == 0 {
			fmt.Println("Processed so far :", i)
		}

	}

	fmt.Println("Processed all :", i)
}

func main() {

	rds, err := redis.Dial("tcp", ":6389")
	PanicOnError(err)
	defer rds.Close()

	importProfiles(&rds, "data/soc-pokec-profiles.txt")
	importRelations(&rds, "data/soc-pokec-relationships.txt")

}

func zip_arrays(a, b []string) []interface{} {
	res := make([]interface{}, len(a)*2, len(a)*2)
	for i, x := range a {
		res[i*2] = x
		res[i*2+1] = b[i]
	}
	return res
}
