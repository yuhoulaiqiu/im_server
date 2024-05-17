// Code generated by goctl. DO NOT EDIT.
package types

type AuthenticationRequest struct {
	Token     string `header:"Token,optional"`
	ValidPath string `header:"ValidPath,optional"`
}

type AuthenticationResponse struct {
	UserID uint `json:"userId"`
	Role   int  `json:"role"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OpenLoginRequest struct {
	Code string `json:"code"`
	Flag string `json:"flag"` //标识
}

type OpenLoginResponse struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"`
}
