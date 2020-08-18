/*
@Time : 2020/4/17 13:41
@Author : FB
@File : TokenResult.go
@Software: GoLand
*/
package entities

type TokenResp struct {
	Code             int     `json:"error"`
	ErrorDescription string  `json:"error_description"`
	ExpiresIn        float64 `json:"expires_in"`
	AccessToken      string  `json:"access_token"`

	// 以下参数请忽略，暂时使用
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}
