package websitecrawler

import (
    "errors"
    "fmt"
    "net/http"
    "io/ioutil"
)

func handleAPIError(resp *http.Response) error {
    body, _ := ioutil.ReadAll(resp.Body)
    return errors.New(fmt.Sprintf("API Error: %s", string(body)))
}
