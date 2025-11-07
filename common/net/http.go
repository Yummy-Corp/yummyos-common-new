package net

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"reflect"
	"time"
)

type Http struct {
	Uri           string
	Method        string
	ContentType   string
	Authorization string
	Client        *http.Client
	Header        map[string]string
}

func (h *Http) Init() {
	if h.ContentType == "" {
		h.ContentType = "application/json"
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error initialize cookie jar: %s", err.Error())
	}
	h.Client = &http.Client{
		Timeout: 15 * time.Second,
		Jar:     jar}
}

func (h *Http) Send(data interface{}) error {
	req, err := http.NewRequest(h.Method, h.Uri, nil)
	if err != nil {
		log.Fatalf("Error initializing request: %s", err.Error())
	}

	if h.Authorization != "" {
		req.Header.Set("Authorization", h.Authorization)
	}

	for headerKey, headerValue := range h.Header {
		req.Header.Add(headerKey, headerValue)
	}

	res, err := h.Client.Do(req)
	if err != nil {
		log.Fatalf("Error consumimg request: %s", err.Error())
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(data)
}

func (h *Http) marshalRequest(payload interface{}) (*http.Request, error) {
	if h.ContentType == "application/x-www-form-urlencoded" {
		encodedPayload := url.Values{}
		values := reflect.ValueOf(payload)
		typesOf := values.Type()
		for i := 0; i < values.NumField(); i++ {
			encodedPayload.Set(typesOf.Field(i).Tag.Get("json"), values.Field(i).Interface().(string))
		}
		req, err := http.NewRequest(h.Method, h.Uri, bytes.NewBufferString(encodedPayload.Encode()))
		return req, err
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling payload: %s", err.Error())
	}
	req, err := http.NewRequest(h.Method, h.Uri, bytes.NewBuffer(jsonPayload))
	return req, err
}

func (h *Http) SendPayload(payload interface{}, data interface{}) error {
	req, err := h.marshalRequest(payload)

	if err != nil {
		log.Fatalf("Error initializing request: %s", err.Error())
	}

	req.Header.Set("Content-Type", h.ContentType)
	if h.Authorization != "" {
		req.Header.Set("Authorization", h.Authorization)
	}

	for headerKey, headerValue := range h.Header {
		req.Header.Add(headerKey, headerValue)
	}

	res, err := h.Client.Do(req)
	if err != nil {
		log.Fatalf("Error consumimg request: %s", err.Error())
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(data)
}
