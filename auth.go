/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// ErrMissingPEM happens when the bytes cannot be decoded as a PEM block.
var ErrMissingPEM = errors.New("no PEM blob found")

// ErrInvalidPrivateKey happens when a key cannot be parsed as a ECDSA PKCS8 private key.
var ErrInvalidPrivateKey = errors.New("key could not be parsed as a valid ecdsa.PrivateKey")

// AuthTransport is an http.RoundTripper implementation that stores the JWT created.
// If the token expires, the Rotate function should be called to update the stored token.
type AuthTransport struct {
	Transport    http.RoundTripper
	jwtGenerator jwtGenerator
}

type jwtGenerator interface {
	Token() (string, error)
	IsValid() bool
}

type standardJWTGenerator struct {
	keyID          string
	issuerID       string
	expireDuration time.Duration
	privateKey     *ecdsa.PrivateKey

	token string
}

// NewTokenConfig returns a new AuthTransport instance that customizes the Authentication header of the request during transport.
// It can be customized further by supplying a custom http.RoundTripper instance to the Transport field.
func NewTokenConfig(keyID string, issuerID string, expireDuration time.Duration, privateKey []byte) (*AuthTransport, error) {
	return NewTokenConfigWithProxy(keyID, issuerID, expireDuration, privateKey, nil)
}

func NewTokenConfigWithProxy(keyID string, issuerID string, expireDuration time.Duration, privateKey []byte, proxyURL *url.URL) (*AuthTransport, error) {
	key, err := parsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	gen := &standardJWTGenerator{
		keyID:          keyID,
		issuerID:       issuerID,
		privateKey:     key,
		expireDuration: expireDuration,
	}
	_, err = gen.Token()

	return &AuthTransport{
		Transport:    newTransport(proxyURL),
		jwtGenerator: gen,
	}, err
}

func parsePrivateKey(blob []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(blob)
	if block == nil {
		return nil, ErrMissingPEM
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	if key, ok := parsedKey.(*ecdsa.PrivateKey); ok {
		return key, nil
	}

	return nil, ErrInvalidPrivateKey
}

// RoundTrip implements the http.RoundTripper interface to set the Authorization header.
func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.jwtGenerator.Token()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	return t.transport().RoundTrip(req)
}

// Client returns a new http.Client instance for use with asc.Client.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport == nil {
		t.Transport = newTransport(nil)
	}

	return t.Transport
}

func (g *standardJWTGenerator) Token() (string, error) {
	if g.IsValid() {
		return g.token, nil
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, g.claims())
	t.Header["kid"] = g.keyID

	token, err := t.SignedString(g.privateKey)
	if err != nil {
		return "", err
	}

	g.token = token

	return token, nil
}

func (g *standardJWTGenerator) IsValid() bool {
	if g.token == "" {
		return false
	}

	parsed, err := jwt.Parse(
		g.token,
		jwt.KnownKeyfunc(jwt.SigningMethodES256, g.privateKey),
		jwt.WithAudience("appstoreconnect-v1"),
		jwt.WithIssuer(g.issuerID),
	)
	if err != nil {
		return false
	}

	return parsed.Valid
}

func (g *standardJWTGenerator) claims() jwt.Claims {
	// 当前时间减去1分钟
	adjustedTime := time.Now().Add(-1 * time.Minute)
	// 基于调整后的时间设置过期时间
	expiry := adjustedTime.Add(g.expireDuration)

	return jwt.StandardClaims{
		Audience:  jwt.ClaimStrings{"appstoreconnect-v1"},
		Issuer:    g.issuerID,
		ExpiresAt: jwt.At(expiry),
	}
}

func newTransport(proxyURL *url.URL) http.RoundTripper {
	// 配置代理
	// 客户端 → 代理服务器 → 目标服务器
	// 超时参数	                   作用范围	                           代理环境下的具体含义
	// DialContext.Timeout	  TCP 连接建立阶段                  客户端到代理服务器的连接超时
	// TLSHandshakeTimeout	  TLS 握手阶段                     客户端与代理服务器之间的 TLS 握手超时
	// ResponseHeaderTimeout  从请求发送到收到响应头的总时间	   客户端→代理→目标服务器的全链路响应头超时
	// ExpectContinueTimeout  等待 "100 Continue" 响应的时间	客户端等待代理返回 "100 Continue" 的时间
	// IdleConnTimeout	      空闲连接保持时间	                 客户端与代理服务器之间的空闲连接保持时间

	return &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout:   15 * time.Second, // 到代理的TCP连接超时
			KeepAlive: 30 * time.Second, // 保持活跃探测间隔
		}).DialContext,
		// 安全控制
		TLSHandshakeTimeout: 10 * time.Second, // 到代理的TLS握手超时

		// 响应控制
		ResponseHeaderTimeout: 50 * time.Second, // 全链路响应头超时
		ExpectContinueTimeout: 2 * time.Second,  // // 等待代理的100 Continue

		// 连接池控制
		MaxIdleConns:        50,             // 最大空闲连接数
		MaxIdleConnsPerHost: 10,             // 每个主机最大空闲连接
		IdleConnTimeout:     defaultTimeout, // // 到代理的空闲连接超时
		DisableKeepAlives:   true,           // false: 默认值, 启用 Keep-Alive，复用 TCP 连接（同一主机多次请求可共用连接）/ 禁用 Keep-Alive，每次请求完成后强制关闭连接（短连接模式）
		ForceAttemptHTTP2:   false,          // 启用HTTP/2
	}
}
