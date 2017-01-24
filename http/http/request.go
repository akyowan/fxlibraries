package http

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"vcrlibraries/domain"
	"vcrlibraries/loggers"

	"strconv"

	"bytes"

	"github.com/gorilla/mux"
)

type Request struct {
	Method      string            `bson:"-"`
	Body        *bytes.Buffer     `bson:"-"`
	RemoteAddr  string            `bson:"client_ip,omitempty"`
	QueryParams url.Values        `bson:"query_params,omitempty"`
	UrlParams   map[string]string `bson:"url_params,omitempty"`
	URL         *url.URL          `bson:"-"`

	AccessToken string `bson:"access_token,omitempty"`
	Userhash    string `bson:"userhash,omitempty"`
	SN          string `bson:"sn,omitempty"`
	SecretID    string `bson:"secret_id,omitempty"`

	DefaultLanguage string   `bson:"default_language,omitempty"`
	Languages       []string `bson:"languages,omitempty"`

	UserAgent     string `bson:"user_agent,omitempty"`
	DeviceType    string `bson:"device_type,omitempty"`
	DeviceVersion string `bson:"device_version,omitempty"`

	User  *domain.User  `bson:"user,omitempty"`
	Robot *domain.Robot `bson:"robot,omitempty"`

	ManagerToken string             `bson:"manager_token,omitempty"`
	ManagerID    string             `bson:"manager_id,omitempty"`
	ManagerType  domain.ManagerType `bson:"manager_type,omitempty"`
}

func NewRequest(r *http.Request) (*Request, error) {
	var request Request
	request.Method = r.Method
	request.URL = r.URL
	request.Body = new(bytes.Buffer)
	request.Body.ReadFrom(r.Body)
	// TODO: get from  x-forword-for
	request.RemoteAddr = r.Header.Get("X-Forwarded-For")
	if request.RemoteAddr == "" {
		request.RemoteAddr = r.RemoteAddr
	}
	err := request.parseUrl(r)
	if err != nil {
		return nil, err
	}
	err = request.parseHeader(r)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (self *Request) parseUrl(r *http.Request) error {
	self.UrlParams = mux.Vars(r)
	self.QueryParams = r.URL.Query()
	return nil
}

func (self *Request) parseHeader(r *http.Request) error {
	self.AccessToken = r.Header.Get("X-User-AccessToken")
	self.Userhash = r.Header.Get("X-User-Userhash")
	self.SN = r.Header.Get("X-Robot-SN")
	self.SecretID = r.Header.Get("X-Robot-SecretID")
	language := r.Header.Get("Accept-Language")
	if language != "" {
		languages := strings.Split(language, ",")
		self.Languages = []string{}
		for i, v := range languages {
			if i == 0 {
				self.DefaultLanguage = strings.TrimSpace(v)
			}
			self.Languages = append(self.Languages, strings.TrimSpace(v))
		}
	}
	self.UserAgent = r.UserAgent()
	uaParts := strings.Split(self.UserAgent, " ")
	uaHexaParts := strings.Split(uaParts[0], "/")

	if len(uaHexaParts) == 3 && uaHexaParts[0] == "HEXA" {
		self.DeviceType = uaHexaParts[1]
		self.DeviceVersion = uaHexaParts[2]
	}

	self.ManagerToken = r.Header.Get("X-Manager-Token")
	managerTypeStr := r.Header.Get("X-Manager-Type")
	if managerTypeStr != "" {
		if manageTypeInt, err := strconv.Atoi(managerTypeStr); err == nil {
			self.ManagerType = domain.ManagerType(manageTypeInt)
			self.ManagerID = r.Header.Get("X-Manager-ID")
		}
	}

	return nil
}

func (self *Request) Parse(v interface{}) error {
	return json.Unmarshal(self.Body.Bytes(), v)
}

func (self *Request) ParseJson() map[string]interface{} {
	res := make(map[string]interface{})
	err := self.Parse(&res)
	if err != nil {
		loggers.Error.Println(err)
		return nil
	}
	return res
}
