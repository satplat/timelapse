package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type imageLink struct {
	ImageLink string `json:"image_link"`
}

type imagesData struct {
	ImagesInfo []imageData `json:"images_info"`
}

type imageData struct {
	Date     string
	Filename string
}

func GetImagesData(landID, token string) error {
	// Retrieving list of images for a year
	for i := 1; i < 13; i++ {
		url := fmt.Sprintf(
			"https://api.apieco.ir/farm-info/api/v1/read/legalCustomer/satelliteImageInfo/%s/%d",
			landID, i)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, strings.NewReader("{}"))
		if err != nil {
			return err
		}

		req.Header.Add("apieco-key", token)
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			return err
		}

		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		var data imagesData
		err = json.Unmarshal(body, &data)
		if err != nil {
			return err
		}
		for _, v := range data.ImagesInfo {
			saveImage(v.Date, v.Filename, token, landID)
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func saveImage(date, filename, token, landID string) error {
	dateSlice := strings.Split(date, "/")
	fileName := fmt.Sprintf("%s-%s-%s.jpg", dateSlice[0], addZero(dateSlice[1]), addZero(dateSlice[2]))
	month := dateSlice[1]
	url := fmt.Sprintf(
		"https://api.apieco.ir/farm-info/api/v1/read/legalCustomer/RGB/%s/%s/%s",
		landID, month, filename)
	method := "GET"
	payload := strings.NewReader("{}")
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("apieco-key", token)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var data []imageLink
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	url2 := data[0].ImageLink
	resp, _ := http.Get(url2)

	os.Mkdir(".images", 0777)
	file, _ := os.Create(".images/" + fileName)
	// //Write the bytes to the fiel
	io.Copy(file, resp.Body)
	fmt.Println(fileName + " Saved!")
	return nil
}

func addZero(a string) string {
	if len(a) == 2 {
		return a
	}
	return "0" + a
}
