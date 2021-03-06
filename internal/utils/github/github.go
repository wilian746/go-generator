package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetFileFromGithub(routerGithub string) ([]byte, error) {
	urlBase := "https://raw.githubusercontent.com/wilian746/go-generator/"
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", urlBase, routerGithub), nil)
	if err != nil {
		return []byte{}, err
	}
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
