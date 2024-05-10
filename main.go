package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/gin-gonic/gin"
)

type TestData struct {
	DecodeString string `json:"decodeString"`
}

func main() {
	r := gin.Default()
	r.GET("/status", func(ctx *gin.Context) {
		ctx.String(200, "Work")
	})
	r.POST("/sha256", func(ctx *gin.Context) {
		var testData TestData

		err := ctx.ShouldBindJSON(&testData)
		if err != nil {
			log.Fatalln(err)
		}
		hash := sha256.Sum256([]byte(testData.DecodeString))
		ctx.JSON(200, gin.H{
			"sha256": hex.EncodeToString(hash[:]),
		})
	})
	r.Run()
}

// func testData(data string) {
// 	jsonData := TestData{DecodeString: data}
// 	jsonBuffer := new(bytes.Buffer)
// 	json.NewEncoder(jsonBuffer).Encode(jsonData)
// 	resp, err := http.Post("http://127.0.0.1:8080/sha256", "application/json", jsonBuffer)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var responseData TestData
// 	if err := json.Unmarshal(body, &responseData); err != nil {
// 		fmt.Println("Ошибка при декодировании JSON ответа:", err)
// 		return
// 	}
// 	fmt.Println("Декодированный JSON ответ:", responseData)

// }
