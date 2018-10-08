package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var baseURL = "goooooogle"

func main() {

}

func DeleteApp(platform string, hostname string, header string) error {

	if platform == "" || hostname == "" {
		fmt.Printf("[DELETE APP] Platform and hostname can not be empty strings")
		return errors.New("[DELETE APP] Platform and hostname can not be empty strings")
	}
	url := fmt.Sprintf(baseURL+"delete-app/%s/%s", platform, hostname)

	// Create client & set timeout
	client := &http.Client{}
	client.Timeout = time.Second * 15

	// Create request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("[DELETE APP] Could not create request : %v", err)
		return err
	}

	//check for optional header
	if header != "" {
		req.Header.Add("X-Fields", header)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[DELETE APP] Could not fetch request : %v", err)
		return err
	}
	defer resp.Body.Close()

	//Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[DELETE APP] Could not read response body : %v", err)
		return err
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))

	return nil
}
