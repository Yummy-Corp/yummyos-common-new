package net

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yummy-corp/yummyos-common/common/config"
)

type Token struct {
	UpdatedAt   string `json:"updated_at" redis:"updated_at"`
	AccessToken string `json:"access_token" redis:"access_token"`
	Scope       string `json:"scope" redis:"scope"`
	ExpiresIn   int    `json:"expires_in" redis:"expires_in"`
	TokenType   string `json:"token_type" redis:"token_type"`
}

type Client struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type Credential struct {
	Uri         string `json:"uri"`
	ContentType string `json:"content_type"`
	Key         string `json:"key"`
	Client      Client `json:"client"`
	Token       Token  `json:"token"`
}

func (h *Credential) setToken() {
	rdb := config.InitRedis()
	ctx := context.Background()

	currentTime := time.Now()
	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, h.Key, "updated_at", currentTime.Format(time.RFC3339))
		rdb.HSet(ctx, h.Key, "access_token", h.Token.AccessToken)
		rdb.HSet(ctx, h.Key, "scope", h.Token.Scope)
		rdb.HSet(ctx, h.Key, "expires_in", h.Token.ExpiresIn)
		rdb.HSet(ctx, h.Key, "token_type", h.Token.TokenType)
		return nil
	}); err != nil {
		log.Printf("Error storing token: %s", err.Error())
	}
}

func (h *Credential) getToken() {
	rdb := config.InitRedis()
	ctx := context.Background()

	if err := rdb.HGetAll(ctx, h.Key).Scan(&h.Token); err != nil {
		panic(err)
	}
}

func (h *Credential) RequestToken() string {
	http := Http{
		Uri:         h.Uri,
		Method:      "POST",
		ContentType: h.ContentType,
	}
	h.getToken()

	var expire time.Time
	if h.Token.AccessToken != "" {
		updatedAt, _ := time.Parse(time.RFC3339, h.Token.UpdatedAt)
		expire = updatedAt.Add(time.Duration(h.Token.ExpiresIn) * time.Second)
	}

	if h.Token.AccessToken == "" || expire.Before(time.Now()) {
		http.Init()
		if err := http.SendPayload(h.Client, &h.Token); err != nil {
			log.Fatalf("Error requesting token: %s", err.Error())
		}
		h.setToken()
	}

	return h.Token.AccessToken
}

func (h *Credential) GenerateToken() Token {
	http := Http{
		Uri:    h.Uri,
		Method: "POST",
	}
	http.Init()
	if err := http.SendPayload(h.Client, &h.Token); err != nil {
		log.Fatalf("Error requesting token: %s", err.Error())
	}

	return h.Token
}
