package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kirillmc/data_filler/pkg/filler"
	fil "github.com/kirillmc/data_filler/pkg/model"
	"github.com/kirillmc/http_test_server/internal/model"
)

const (
	newBaseUrl  = "http://localhost:8080"
	aseUrl      = "https://4f48-93-100-98-132.ngrok-free.app"
	getPostfix  = "/programs/%d"
	postPostfix = "/programs/"
	avg         = 5
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

func postNProgramsClient(programs fil.TrainPrograms) (model.Response, float64, error) {

	dataToSend, err := json.Marshal(programs)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	resp, err := http.Post(fmt.Sprintf(newBaseUrl+postPostfix), "application/json", bytes.NewBuffer(dataToSend))
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {

		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err

	}

	if resp.StatusCode != http.StatusOK {

		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err

	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToSend)), err
	}

	return mess, float64(len(dataToSend)), nil
}

func updateNProgramsClient(programs fil.TrainPrograms) (model.Response, float64, error) {

	dataToUpdate, err := json.Marshal(programs)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(newBaseUrl+postPostfix), bytes.NewBuffer(dataToUpdate))
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err

	}

	if resp.StatusCode != http.StatusOK {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err

	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		return model.Response{Message: err.Error()}, float64(len(dataToUpdate)), err
	}

	return mess, float64(len(dataToUpdate)), nil
}

func deleteNProgramsClient(n int64) (model.Response, error) {

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(newBaseUrl+getPostfix, n), nil)
	if err != nil {
		return model.Response{Message: err.Error()}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Response{Message: err.Error()}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return model.Response{Message: err.Error()}, err
	}

	if resp.StatusCode != http.StatusOK {
		return model.Response{Message: err.Error()}, err
	}

	var mess model.Response
	if err = json.NewDecoder(resp.Body).Decode(&mess); err != nil {
		return model.Response{Message: err.Error()}, err
	}

	return mess, nil
}

func main() {
	//var trainPorgrams, wq = oneToPost(1)
	//log.Printf("POST: %v\t%v", trainPorgrams, wq)
	//
	//getMethodFrom0ToNWithAVG(n)
	methodFrom0ToNWithAVG(505, oneToDelete)
	//	printAvgOfConst(1, oneToGet)
}

func methodFrom0ToNWithAVG(n int64, fun func(int64) (float64, float64)) {
	for i := int64(0); i <= n; i++ {
		printAvgOfConst(i, fun)
	}
}

func printAvgOfConst(n int64, fun func(int64) (float64, float64)) {
	var avgTime float64
	var avgSize float64
	for i := 1; i <= avg; i++ {
		avgTempTime, avgTempSize := fun(n)
		avgTime += avgTempTime
		avgSize += avgTempSize
	}

	log.Printf("\t%d;\t%f;\t%f;\n", n, avgTime/avg, avgSize/avg)
}

func oneToGet(n int64) (float64, float64) {
	start := time.Now()

	programs, err := getNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	numOfSets, err := json.Marshal(programs)
	if err != nil {
		fmt.Errorf("fail to get json: %v", err)
	}

	return float64(end.Sub(start).Nanoseconds()), float64(len(numOfSets))
}

func oneToPost(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))
	_, postMessSize, err := postNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()

	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToUpdate(n int64) (float64, float64) {
	start := time.Now()
	programs := filler.CreateOwnSetOfPrograms(int(n))
	_, postMessSize, err := updateNProgramsClient(programs)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()), postMessSize
}

func oneToDelete(n int64) (float64, float64) {
	start := time.Now()
	_, err := deleteNProgramsClient(n)
	if err != nil {
		log.Println("ERROR")
	}
	end := time.Now()
	return float64(end.Sub(start).Nanoseconds()), 0.0
}

func oldPrint(n int64) {
	start := time.Now()
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

//func main() {
//
//	var n int64 = 21
//	//log.Printf("\tCOUNT;\tTIME(nanoS);\tSIZE;\n")
//	//for i := int64(1); i <= n; i++ {
//	//	printAvgOfConst(i)
//	//}
//	oldPrint(n)
//	//start := time.Now()
//	//var user UserToGet
//	//var err error
//	//for i := 0; i < 101; i++ {
//	//	user, err = getUserClient(1)
//	//	if err != nil {
//	//		log.Fatal("failed to get note:", err)
//	//	}
//	//}
//	//end := time.Now()
//
//	//var user UserToGet
//	//var err error
//	//start := time.Now()
//	//for i := 0; i < 5; i++ {
//	//	start := time.Now()
//	//	user, err = getUserClient(1)
//	//	if err != nil {
//	//		log.Fatal("failed to get user:", err)
//	//	}
//	//	end := time.Now()
//	//	log.Printf("time:%v\n", end.Sub(start))
//	//}
//	//end := time.Now()
//	//
//	//log.Printf("Last user info:%v\n", user)
//	//log.Printf("time for 5 get requests: %v\n", end.Sub(start))
//
//	//start := time.Now()
//	//n := 101
//	//wg.Add(n)
//	//for i := 0; i < n; i++ {
//	//	go testRequest(i)
//	//}
//	//wg.Wait()
//	//end := time.Now()
//	//var total time.Duration
//	////	log.Printf("Last user info:%v\n", user)
//	//
//	//for i := range list {
//	//	total += list[i]
//	//}
//	//
//	//avg := total.Nanoseconds() / (int64(len(list)))
//	//log.Printf("total requests time: %v\n", total)
//	//log.Printf("time for %d get requests: %v\n", n, end.Sub(start))
//	//log.Printf("avg time for %d get requests: %s\n", n, time.Duration(avg))
//}
