package gphoto

/*
  Limited support to google api based on guide

	Reference: https://developers.google.com/photos/library/guides

*/
import (
	"net/http"
)

func isSupportedMediaItem() {
	/*
		Photos:	BMP, GIF, HEIC, ICO, JPG, PNG, TIFF, WEBP, some RAW files
		Videos: 3GP, 3G2, ASF, AVI, DIVX, M2T, M2TS, M4V, MKV, MMV, MOD, MOV, MP4, MPG, MTS, TOD, WMV
	*/

}

// UploadBytes to manage uploading raw media bytes to google photos
func UploadBytes(mediaPath string, httpClient *http.Client) (err error) {
	/*
		https://developers.google.com/photos/library/guides/upload-media#uploading-bytes

		a.k.a. simple upload
	*/
	return nil
}

// UploadResumableBytes is similar to UploadBytes but handle files (videos) larger than 50MB
func UploadResumableBytes() {
	/*
		https://developers.google.com/photos/library/guides/resumable-uploads

		This API is based on single request approach
	*/

}

// MediaItemsBatchCreate creates a ton of media items in google photos
func MediaItemsBatchCreate(mediaPath string, httpClient *http.Client) (err error) {
	return nil
}
