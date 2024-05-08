package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kirillmc/http_test_server/internal/model"

	"github.com/kirillmc/data_filler/pkg/filler"
)

func main() {
	http.HandleFunc("/programs/", handleProgrmas)
	//http.HandleFunc("/users/", handleProgram)
	log.Println("Server is serving on: localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProgrmas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createProgram(w, r)
	case http.MethodGet:
		getNPrograms(w, r)
	case http.MethodPut:
		updatePrograms(w, r)
	case http.MethodDelete:
		deletePrograms(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//func handleProgram(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case http.MethodGet:
//		getUserById(w, r)
//	case http.MethodPut:
//		updatePrograms(w, r)
//	case http.MethodDelete:
//		deletePrograms(w, r)
//	default:
//		w.WriteHeader(http.StatusMethodNotAllowed)
//	}
//}

func getNPrograms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	n, err := getId(r)
	log.Print(n)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	programs := filler.CreateOwnSetOfPrograms(int(n))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(programs)
}

func createProgram(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON из тела запроса в структуру Programs
	var programs model.TrainPrograms
	err := json.NewDecoder(r.Body).Decode(&programs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := model.Response{Message: "Данные были добавлены"}

	// Отправляем успешный статус
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func updatePrograms(w http.ResponseWriter, r *http.Request) {

	// Декодируем JSON из тела запроса в структуру User
	var trainPrograms model.TrainPrograms
	err := json.NewDecoder(r.Body).Decode(&trainPrograms)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message := model.Response{Message: "Данные были обновлены"}

	// Отправляем успешный статус
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func deletePrograms(w http.ResponseWriter, r *http.Request) {
	_, err := getId(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	message := model.Response{Message: "Данные были удалены"}

	// Отправляем успешный статус
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)

}
func getId(r *http.Request) (int64, error) {
	id := r.URL.Path[len("/programs/"):]
	idRes, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, err
	}

	return idRes, nil
}
