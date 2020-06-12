package userrepository

import (
	entity "aprendendo/web_server/src/domain/entities"
	postResp "aprendendo/web_server/src/repository/responses"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePost(post *entity.Post) (*postResp.PostResponse, error) {
	bodyReq, err := json.Marshal(map[string]string{
		"title": post.Title,
		"body":  post.Body,
	})

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(bodyReq)
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", buffer)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bytesResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	var postResponse postResp.PostResponse
	errorUnmarshal := json.Unmarshal(bytesResponse, &postResponse)

	if errorUnmarshal != nil {
		return nil, errorUnmarshal
	}

	return &postResponse, nil
}
