package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"vcrlibraries/loggers"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

type Header struct {
	http.Header
}

// NewHeader Create customerized Header object
func NewHeader() *Header {
	return &Header{make(http.Header)}
}

func (self *Header) Set(key, value string) *Header {
	self.Header.Set(key, value)
	return self
}

func (self *Header) SetSN(sn string) *Header {
	self.Header.Set("X-Robot-SN", sn)
	return self
}

func (self *Header) SetSecretID(secretID string) *Header {
	self.Header.Set("X-Robot-SecretID", secretID)
	return self
}

func (self *Header) SetAccessToken(accessToken string) *Header {
	self.Header.Set("X-User-AccessToken", accessToken)
	return self
}

type URL struct {
	*url.URL
	query url.Values
}

// NewURL Create customerized URL object, save extra one import.
func NewURL(rawurl string) *URL {
	var u URL
	var err error
	u.URL, err = url.Parse(rawurl)
	if err != nil {
		loggers.Error.Println(err)
	} else {
		u.query = u.URL.Query()
	}
	return &u
}

func (self *URL) Set(key, value string) *URL {
	self.query.Set(key, value)
	return self
}

func (self *URL) String() string {
	self.URL.RawQuery = self.query.Encode()
	return self.URL.String()
}

// GetHost Get the HOST from `HOST:PORT` string
func GetHost(hostPort string) string {
	if index := strings.Index(hostPort, ":"); index != -1 {
		return hostPort[:index]
	} else {
		return hostPort
	}
}

// JsonToReader JSON object to io.Reader
func JsonToReader(v interface{}) io.Reader {
	data, err := json.Marshal(v)
	if err != nil {
		loggers.Error.Println(err)
		return nil
	}
	return bytes.NewReader(data)
}

func PostJson(url string, requestJson, responseJson interface{}) (*http.Response, error) {
	resp, err := http.Post(url, "application/json", JsonToReader(requestJson))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(responseJson)
	if err != nil {
		loggers.Error.Println(err)
		return resp, err
	}

	return resp, err
}

func GetJson(url string, responseJson interface{}) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(responseJson)
	if err != nil {
		loggers.Error.Println(err)
		return nil, err
	}

	return resp, nil
}

// Call A util api calling function, specific for Vincross backend API.
// Header is customrized header
func Call(method, url string, header *Header, requestJson interface{}) (*Response, error) {
	req, err := http.NewRequest(method, url, JsonToReader(requestJson))
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header.Header
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		loggers.Warn.Println(method, url, resp.StatusCode)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		loggers.Error.Println(resp.StatusCode, err)
		return nil, err
	}

	response.HTTPCode = resp.StatusCode
	return &response, nil
}

func CallReadCloser(method, url string, header *Header, readCloser io.ReadCloser) (*Response, error) {
	req, err := http.NewRequest(method, url, readCloser)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header.Header
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		loggers.Warn.Println(method, url, resp.StatusCode)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		loggers.Error.Println(resp.StatusCode, err)
		return nil, err
	}

	response.HTTPCode = resp.StatusCode
	return &response, nil
}
