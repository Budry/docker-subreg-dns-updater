package ip

import (
	"io/ioutil"
	"net/http"
)

func GetPublicIp() (string, error) {
	url := "https://api.ipify.org?format=text"
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
