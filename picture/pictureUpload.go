package picture

import (
	"github.com/namjagbrawa/go_chat/exception"
	"io"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR      = "./files/uploads"
	IMAGETOTEXT_DIR = "./files/imageToText"

	URL_PICTURE_UPLOAD = "/picture/upload"
	URL_PICTURE_VIEW   = "/picture/view"

	UPDATE_PAGE = `<html>
			<head>
			<meta charset="utf-8">
			<title>Upload</title>
			</head>
			<body>
			<form method="POST" action="` + URL_PICTURE_UPLOAD + `" enctype="multipart/form-data">
			Choose an image to upload: <input name="image" type="file" />
			<input type="submit" value="Upload" />
			<div>Echo or not:
				<label><input type="radio" name="echo" value="true">yes</label>
				<label><input type="radio" name="echo" value="false">no</label>
			</div>
			</form>
			</body>
			</html>`
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		io.WriteString(w, UPDATE_PAGE)
	}

	// 处理图片上传
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		filename := h.Filename
		defer f.Close()

		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			// write to http client the internal error message, and return to exception handler user exception message, test.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return exception.UserError("create file error")
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			// write to http client internal error messgae.
			// may be log error message, and write to client the user error message by exception handler.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		if r.Form["echo"][0] == "true" {
			http.Redirect(w, r, URL_PICTURE_VIEW+"?filename="+filename, http.StatusFound)
		}
	}

	return nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request) error {
	filename := r.FormValue("filename")
	imagePath := UPLOAD_DIR + "/" + filename
	if exists := isExists(imagePath); !exists {
		return exception.UserError("file not found")
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
	return nil
}

func ImageToTextHandler(w http.ResponseWriter, r *http.Request) error {
	filename := r.FormValue("filename")
	textPath := IMAGETOTEXT_DIR + "/" + filename + ".txt"
	if exists := isExists(textPath); !exists {
		imagePath := UPLOAD_DIR + "/" + filename
		if exists := isExists(imagePath); !exists {
			http.NotFound(w, r)
			return exception.UserError("file not found")
		}
		imageToText()
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, textPath)
	return nil
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func imageToText() {

}
