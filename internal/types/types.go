// Code generated by goctl. DO NOT EDIT.
package types

type CaptchaRequest struct {
	Phone string `json:"phone"` //手机号码
}

type CaptchaResponse struct {
	Captcha_key           string `json:"captcha_key"`
	Expired_at            string `json:"expired_at"`
	Captcha_image_content string `json:"captcha_image_content"`
}