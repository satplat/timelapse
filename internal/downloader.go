package internal

import (
	"encoding/csv"
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

type IndexData struct {
	ImageInfo []IndexValue `json:"image_info"`
}

type IndexValue struct {
	Value string  `json:"value"`
	X float32 `json:"x_coor"`
	Y float32 `json:"y_coor"`
}

func GetImagesData(landID, token string) error {
	// Retrieving list of images for a year
	os.Mkdir(".images", 0777)
	os.Mkdir(".csv", 0777)
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
			err = saveImage(v.Date, v.Filename, token, landID)
			if err != nil {
				continue
			}
			err = saveIndex(v.Date, v.Filename, landID, "NDVI", token)
			if err != nil {
				continue
			}
			time.Sleep(1 * time.Second)
		}
	}
	return nil
}

func saveIndex(date, filename, landID, index, token string) error {
	dateSlice := strings.Split(date, "/")
	fileName := fmt.Sprintf("%s-%s-%s.csv", dateSlice[0], addZero(dateSlice[1]), addZero(dateSlice[2]))
	month := dateSlice[1]
	url := fmt.Sprintf(
		"https://api.apieco.ir/farm-info/api/v1/read/legalCustomer/Index/%s/%s/%s/%s",
		landID, index, month, filename)
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
	csvFile, err := os.Create(".csv/" + fileName)
	csvWriter := csv.NewWriter(csvFile)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var data IndexData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	for _, item := range data.ImageInfo {
		row := []string{item.Value, fmt.Sprintf("%f", item.X), fmt.Sprintf("%f", item.Y)}
		_ = csvWriter.Write(row)
	}
	csvWriter.Flush()
	csvFile.Close()
	fmt.Println("New CSV created.")
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

	file, _ := os.Create(".images/" + fileName)
	// //Write the bytes to the file
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
