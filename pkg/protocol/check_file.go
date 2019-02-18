package protocol

import "io/ioutil"

type CheckFileRequest struct {
	Filename string   `json:"filename"`
	Content  string   `json:"content"`
	Rules    []string `json:"rules"`
}

func NewCheckFileRequest(filename string, rules []string) (*CheckFileRequest, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &CheckFileRequest{
		Filename: filename,
		Rules:    rules,
		Content:  string(bytes),
	}, nil
}

type CheckFileResponse struct {
	Filename string `json:"filename"`
	Display  string `json:"display"`
}
