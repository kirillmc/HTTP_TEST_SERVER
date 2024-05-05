package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillmc/http_test_server/internal/model"
)

const (
	newBaseUrl = "http://localhost:8080"
	aseUrl     = "https://4f48-93-100-98-132.ngrok-free.app"
	getPostfix = "/programs/%d"
)

func getNProgramsClient(n int64) (model.TrainPrograms, error) {
	resp, err := http.Get(fmt.Sprintf(newBaseUrl+getPostfix, n))
	if err != nil {
		log.Fatal("Failed to get programs:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.TrainPrograms{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.TrainPrograms{}, errors.New("failed to get user")
	}

	var programs model.TrainPrograms
	if err = json.NewDecoder(resp.Body).Decode(&programs); err != nil {
		return model.TrainPrograms{}, err
	}

	return programs, nil
}

func main() {
	start := time.Now()
	var n int64 = 1
	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}

	end := time.Now()
	numOfSets, err := json.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}
	log.Printf("|\t\t\tHTTP INFO: SIZE[%d]\t\t\t|\n", n)
	log.Printf("|\tTOTAL TIME TO GET PROGRAMS:\t%v\t\t|\n", end.Sub(start))
	log.Printf("|\tSIZE OF PROGRAMS:\t\t%s\t|\n", getSizeInFormattedString(int64(len(numOfSets))))
	//start := time.Now()
	//var user UserToGet
	//var err error
	//for i := 0; i < 101; i++ {
	//	user, err = getUserClient(1)
	//	if err != nil {
	//		log.Fatal("failed to get note:", err)
	//	}
	//}
	//end := time.Now()

	//var user UserToGet
	//var err error
	//start := time.Now()
	//for i := 0; i < 5; i++ {
	//	start := time.Now()
	//	user, err = getUserClient(1)
	//	if err != nil {
	//		log.Fatal("failed to get user:", err)
	//	}
	//	end := time.Now()
	//	log.Printf("time:%v\n", end.Sub(start))
	//}
	//end := time.Now()
	//
	//log.Printf("Last user info:%v\n", user)
	//log.Printf("time for 5 get requests: %v\n", end.Sub(start))

	//start := time.Now()
	//n := 101
	//wg.Add(n)
	//for i := 0; i < n; i++ {
	//	go testRequest(i)
	//}
	//wg.Wait()
	//end := time.Now()
	//var total time.Duration
	////	log.Printf("Last user info:%v\n", user)
	//
	//for i := range list {
	//	total += list[i]
	//}
	//
	//avg := total.Nanoseconds() / (int64(len(list)))
	//log.Printf("total requests time: %v\n", total)
	//log.Printf("time for %d get requests: %v\n", n, end.Sub(start))
	//log.Printf("avg time for %d get requests: %s\n", n, time.Duration(avg))
}

func getSizeInFormattedString(byteSize int64) string {
	if byteSize < 1024 {
		return fmt.Sprintf("%.3f байт\t", float64(byteSize))
	}
	if byteSize < 1024*1024 {
		return fmt.Sprintf("%.3f килобайт\t", float64(byteSize)/1024)
	} else {
		return fmt.Sprintf("%.3f мегабайт\t", float64(byteSize)/(1024*1024))
	}
}
