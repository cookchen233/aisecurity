package types

import "time"

type UploadedFile struct {
	Name       string    `json:"name"`
	URL        string    `json:"url"`
	Size       int64     `json:"size"`
	CreateTime time.Time `json:"create_time"`
}

type UploadedImage struct {
	UploadedFile
}

type UploadedVideo struct {
	UploadedFile
	Duration int64 `json:"duration"`
}

type AllowedMimeTypes struct {
	Types []string
}

func NewAllowedMimeTypes(allowTypes []string) *AllowedMimeTypes {
	return &AllowedMimeTypes{
		Types: allowTypes,
		//[]string{
		//"image/jpeg",
		//"image/png",
		// Add other image MIME types as needed
		//},
	}
}

func (a *AllowedMimeTypes) IsAllowed(mimeType string) bool {
	if len(a.Types) == 0 || a.Types == nil {
		return true
	}
	for _, t := range a.Types {
		if mimeType == t {
			return true
		}
	}
	return false
}
