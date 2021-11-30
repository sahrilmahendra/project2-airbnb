package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// function untuk generate latitude, longitude menggunakan api geocode (google)
func GetGeocodeLocations(s string) (float64, float64, error) {
	google_api_key := "AIzaSyDVzOsxD_6zAX2dG6jqLFjBBw3jh9lGLbI"
	locationString := strings.ReplaceAll(s, " ", "+")

	url := ("https://maps.googleapis.com/maps/api/geocode/json?address=" + locationString + "&key=" + google_api_key)
	response, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return 0, 0, err2
	}

	var longitude float64
	var latitude float64
	var values map[string]interface{}

	json.Unmarshal(body, &values)
	for _, v := range values["results"].([]interface{}) {
		for i2, v2 := range v.(map[string]interface{}) {
			if i2 == "geometry" {
				latitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lat"].(float64)
				longitude = v2.(map[string]interface{})["location"].(map[string]interface{})["lng"].(float64)
				break
			}
		}
	}
	return latitude, longitude, nil
}
