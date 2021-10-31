package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"godville/enc"
	"godville/structs"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MakeFeedPostRequest(c *http.Client, actionHash string, requestData map[string]interface{}) ([]byte, error) {
	var (
		r             *http.Response
		requestFields url.Values

		jsonString []byte
		a, b string

		responseBody []byte
		response     structs.GenericResponse

		err error
	)

	requestFields = url.Values{}

	a = enc.Vm(actionHash)

	// e.g. 9FwH2ahcM6oMrfS4DfuMyv1gcJksp
	requestFields.Set("a", a)

	if len(requestData) > 0 {

		jsonString, err = json.Marshal(requestData)

		if err != nil {
			return nil, errors.New(
				fmt.Sprintf("Error while encoding request data: %s\n", err.Error()),
			)
		}

		// e.g. W0vFCeyJhY3Rpb24iOiJwdW5pc2gifQ==GrS
		b = enc.Wm(jsonString)
		requestFields.Set("b", b)
	}

	r, _ = c.PostForm("https://godville.net/fbh/feed", requestFields)

	responseBody, err = ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &response)

	if err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, errors.New(fmt.Sprintf("%+v", response))
	}

	return responseBody, nil
}
