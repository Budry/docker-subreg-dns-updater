package ip

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetPublicIp() string {
	url := "https://api.ipify.org?format=text"
	response, err := http.Get(url)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
