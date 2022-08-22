package main

import (
	//"bufio"
	//"bufio"
	"encoding/json"
	"fmt"
	"os"

	//"io"
	"io/ioutil"
	"net/http"
	"strings"
)
func Auth(method string, path string) string {

	//POST https://sdkms.fortanix.com/sys/v1/session/auth
		client := &http.Client{}
		url := "https://sdkms.fortanix.com" + path
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		req.Header.Add("Authorization", "Basic MzRkYmYwNTItZmY5ZC00MTEwLWFkZTUtYWQ3MWRmNzU2YWQ1OmdQM1pSYWZrOFFkcUkyR1FrZG80SUp6ZU9kRGhjTVNqdDlhT3M5bVZ3VURuNXlWNWpyYzh4TEozeVZGUVE1NU5PcmdZVEltRGJSSE93T2ROS0NwT2RR")
	
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer res.Body.Close()
	
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	
		type Request struct {
			Type   string `json:"token_type"`
			Expiry int    `json:"expires_in"`
			Bearer string `json:"access_token"`
			Entity string `json:"entity_id"`
		}
		data := Request{}
		json.Unmarshal(body, &data)
		bearerToken := string(data.Bearer)
		return bearerToken
	}
	

func sign(bearer string, method string, path string, keyid string, data string) {
	//POST https://sdkms.fortanix.com/crypto/v1/sign
	// payload := strings.NewReader(`{"key": {"kid": "` + keyID + `"},"alg": "Aes","plain": "`+ plain +`","mode": "CBC"}`)
	url := "https://sdkms.fortanix.com" + path
	payload := strings.NewReader(`{
		"key": {
		  "kid": "` + keyid + `"
		},
		"hash_alg" : "SHA256",
		"data": "` + data + `"
	  }`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func verify(bearer string, method string, path string, keyid string, data string, signature string) {
	//POST https://sdkms.fortanix.com/crypto/v1/verify
	url := "https://sdkms.fortanix.com" + path
	payload := strings.NewReader(`{
		"key": {
			"kid": "` + keyid + `"
		},
		"hash_alg" : "SHA256",
		"data": "` + data + `",
		"signature": "` + signature + `"
		}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func main() {
	bearer := Auth("POST","/sys/v1/session/auth")
	sign(bearer, os.Args[1], os.Args[2], os.Args[3], os.Args[4])
	verify(bearer, os.Args[5], os.Args[6], os.Args[3], os.Args[4],
		os.Args[7])
}
