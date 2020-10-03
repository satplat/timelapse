package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type farm struct {
	FarmTitle string `json:"farm_title"`
	Polygon string `json:"polygon"`
	CultivationStartDate string `json:"cultivation_start_date"`
	CultivationEndDate string `json:"cultivation_end_date"`
	Services []int `json:"services"`
	ProductType string `json:"product_type"`
	SeedWeight float32 `json:"seed_weight"`
	SeedType string `json:"seed_type"`
	WaterySeedType string `json:"watery_seed_type"`
}

func RegisterLand(coordinates, token string) (string, error) {
	url := "https://api.apieco.ir/farm-info/api/v1/write/legalCustomer/farm/information"
	method := "POST"
	randomName := randStringBytes(10) + string(time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
	farm := farm{
		FarmTitle: randomName,
		Polygon: fmt.Sprintf("POLYGON((%s))", coordinates),
		CultivationStartDate: "2019-04-01",
		CultivationEndDate: "2019-05-01",
		Services: []int{2},
		ProductType: "GANDOM",
		SeedWeight: 1,
		SeedType: "HOMEBRED",
		WaterySeedType: "",
	}
	payload, err := json.Marshal(farm)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, strings.NewReader(string(payload)))

	req.Header.Add("apieco-key", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", nil
	}

	return strings.Split(fmt.Sprintf("%f", data["id"].(float64)), ".")[0], nil
}

func randStringBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}