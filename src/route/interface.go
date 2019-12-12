//author: richard
package route

type HttpError struct {
	Code  int 	`json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseVersion struct {
	Info string `json:"info,omitempty"`
}


