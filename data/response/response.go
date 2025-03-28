package response

type Response struct {
 Code   int         `json:"code"`
 Status string      `json:"status"`
 Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
 Code    int         `json:"code"`
 Message string      `json:"message"`
}

type TaskResponse struct {
 Id     int    `json:"id"`
 Name   string `json:"name"`
 Description string `json:"description"`
}