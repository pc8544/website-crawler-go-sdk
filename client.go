package websitecrawler

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

const baseURL = "https://www.websitecrawler.org/api/crawl"

type Client struct {
    apiKey string
    token  string
}

func NewClient(apiKey string) *Client {
    return &Client{apiKey: apiKey}
}

func (c *Client) Authenticate() error {
    payload := map[string]string{"apiKey": c.apiKey}
    body, _ := json.Marshal(payload)

    resp, err := http.Post(baseURL+"/authenticate", "application/json", bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return c.handleError(resp)
    }

    var result AuthResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return err
    }
    c.token = result.Token
    return nil
}


func (c *Client) Start(url string, limit int) (*GenericResponse, error) {
    var resp GenericResponse
    if _, err := c.postWithRaw("/start", map[string]interface{}{"url": url, "limit": limit}, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}


func (c *Client) PollStatus(url string, interval time.Duration) (*StatusResponse, error) {
    for {
        var s StatusResponse
        _, err := c.postWithRaw("/status", map[string]string{"url": url}, &s)
        if err != nil {
            return nil, err
        }

        // Fetch current URL separately
        cur, _ := c.CurrentURL(url)

        if cur != nil {
            fmt.Printf("Status: %s | CurrentURL: %s\n", s.Status, cur.CurrentUrl)
        } else {
            fmt.Printf("Status: %s\n", s.Status)
        }

        if s.Status == "Completed!" {
            // When complete, fetch crawl data
            data, err := c.Data(url)
            if err != nil {
                return nil, err
            }
            fmt.Println("Crawl Data JSON:", data)
            return &s, nil
        }

        time.Sleep(interval)
    }
}


func (c *Client) CurrentURL(url string) (*CurrentURLResponse, error) {
    var resp CurrentURLResponse
    if _, err := c.postWithRaw("/currentURL", map[string]string{"url": url}, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}


func (c *Client) Data(url string) (*DataResponse, error) {
    var d DataResponse
    if _, err := c.postWithRaw("/cwdata", map[string]string{"url": url}, &d); err != nil {
        return nil, err
    }
    return &d, nil
}


func (c *Client) Clear(url string) (*GenericResponse, error) {
    var resp GenericResponse
    if _, err := c.postWithRaw("/clear", map[string]string{"url": url}, &resp); err != nil {
        return nil, err
    }
    return &resp, nil
}


func (c *Client) postWithRaw(endpoint string, payload interface{}, target interface{}) ([]byte, error) {
    body, _ := json.Marshal(payload)
    req, err := http.NewRequest("POST", baseURL+endpoint, bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer "+c.token)
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    raw, _ := ioutil.ReadAll(resp.Body)

    if resp.StatusCode == http.StatusBadRequest {
        return raw, errors.New(fmt.Sprintf("API endpoint returned this error: %s", string(raw)))
    }

    if err := json.Unmarshal(raw, target); err != nil {
        return raw, errors.New(fmt.Sprintf("JSON Decode error: %v\nRaw response: %s", err, string(raw)))
    }
    return raw, nil
}

func (c *Client) handleError(resp *http.Response) error {
    raw, _ := ioutil.ReadAll(resp.Body)
    return errors.New(fmt.Sprintf("API endpoint returned this error: %s", string(raw)))
}
