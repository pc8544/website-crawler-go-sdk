package websitecrawler

type AuthResponse struct {
    Token string `json:"token"`
}

type GenericResponse struct {
    Message string `json:"message"`
}

type StatusResponse struct {
    Status string `json:"status"`
}

type CurrentURLResponse struct {
    CurrentUrl string `json:"currentUrl"`
}

type DataResponse struct {
    Pages []string `json:"pages"`
}
