package protocol

import "io/ioutil"

type CheckFileRequest struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}

func NewCheckFileRequest(filename string) (*CheckFileRequest, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &CheckFileRequest{Filename: filename, Content: string(bytes)}, nil
}

type CheckFileResponse struct {
	Filename string `json:"filename"`
	Display  string `json:"display"`
}
