// Code generated by goctl. DO NOT EDIT.
package types

type GenerateRequest struct {
	Lurl          string `json:"lurl"`
	UpdateBy      string `json:"updateBy"`
	EffectiveDays int32  `json:"effectiveDays"`
}

type GenerateResponse struct {
	Surl string `json:"surl"`
}

type SurlRequest struct {
	Surl string `path:"surl"`
}
