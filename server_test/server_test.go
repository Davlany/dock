package servertest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lab6/entities"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	testData := entities.TestData{
		DecodeString: "test string",
	}
	testExpected := "d5579c46dfcc7f18207013e65b44e4cb4e2c2298f4ac457ba8f82743f31e930b"
	for {
		time.Sleep(time.Second)
		requestBody, err := json.Marshal(testData)
		if err != nil {
			fmt.Println("Ошибка в создании JSON тела:", err)
			return
		}
		resp, err := http.Post("http://127.0.0.1:8080/sha256", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Println("Not connect, retrying...")
			continue
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			var responseData entities.ResponseData
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				fmt.Println("Ошибка в разборе JSON ответа:", err)
				return
			}
			if responseData.Sha256String != testExpected {
				t.Errorf("Test not complete, excpected: %s\nget: %s", testExpected, responseData.Sha256String)
			} else {
				t.Log("Test complete")
			}
			resp.Body.Close()
			break
		}

	}
}
