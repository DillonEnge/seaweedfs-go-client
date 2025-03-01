package seaweedfs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadFileResp struct {
	Size int `json:"size"`
}

func (c *Client) UploadFile(f *os.File) (UploadFileResp, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	ffw, err := w.CreateFormFile("file", f.Name())
	if err != nil {
		return UploadFileResp{}, err
	}

	io.Copy(ffw, f)

	w.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		c.config.VolumesURL,
		&b,
	)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return UploadFileResp{}, err
	}

	var uploadFileResp UploadFileResp
	err = json.NewDecoder(resp.Body).Decode(&uploadFileResp)
	if err != nil {
		return UploadFileResp{}, err
	}

	return uploadFileResp, nil
}
