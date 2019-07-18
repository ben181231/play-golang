package leak

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type request struct {
	First  int `json:"first"`
	Second int `json:"second"`
	Third  int `json:"third"`
}

type response struct {
	Result int `json:"result"`
}

func simpleSum(first, second, third int, result chan<- int) {
	result <- first + second + third
}

func loopSum(input [3]int, result chan<- int) {
	sum := 0
	for _, elem := range input {
		sum += elem
	}

	result <- sum
}

// GetSumHandler returns a handler sum up 3 integers
func GetSumHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("Error during reading body: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		req := request{}
		if err := json.Unmarshal(body, &req); err != nil {
			log.Printf("Error during unmarshal body: %s", err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		writeResult := func(rw http.ResponseWriter, result int) {
			resp := response{
				Result: result,
			}

			respBytes, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Error during marshaling response: %s", err.Error())
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			rw.WriteHeader(http.StatusOK)
			rw.Write(respBytes)
		}

		result1 := make(chan int)
		result2 := make(chan int)
		go simpleSum(req.First, req.Second, req.Third, result1)
		go loopSum(
			[3]int{req.First, req.Second, req.Third},
			result2,
		)

		select {
		case result := <-result1:
			writeResult(rw, result)
		case result := <-result2:
			writeResult(rw, result)
		}
	})
}
