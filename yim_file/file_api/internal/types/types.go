// Code generated by goctl. DO NOT EDIT.
package types

type ImageRequest struct {
}

type ImageResponse struct {
	Url string `json:"url"`
}

type ImageShowRequest struct {
	ImagesType string `json:"imagesType"`
	ImageName  string `json:"imageName"`
}