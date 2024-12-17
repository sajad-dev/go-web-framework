package exception


type CustomError struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
    Status    bool    `json:"status"`
}