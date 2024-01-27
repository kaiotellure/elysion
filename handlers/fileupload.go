package handlers

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/help"
	"github.com/ikaio/tailmplx/models"
	"github.com/ikaio/tailmplx/pages"
)

func NewFileUpload() *FileUploadHandler {
	return &FileUploadHandler{}
}

type FileUploadHandler struct {}

func (h *FileUploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Post(w, r)
	case http.MethodGet:
		h.Get(w, r)
	}
}

func (h *FileUploadHandler) Post(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		components.Error("file-too-big").Render(r.Context(), w)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		components.Error("unknown").Render(r.Context(), w)
		return
	}
	defer file.Close()

	upload := models.Upload{
		ID:     database.SF.Generate().String(),
		Title:  r.FormValue("title"),
		Author: r.FormValue("author"),
	}

	// My Image.png --> my-image.png
	sanitazed_name := strings.ToLower(strings.ReplaceAll(handler.Filename, " ", "-"))
	upload.Filename = upload.ID + "_" + sanitazed_name

	// Make sure upload folder exists
	upload_folder := help.Env(help.UPLOAD_FOLDER, "web/public/upload")
	os.MkdirAll(upload_folder, os.ModePerm)

	// Creating file on upload folder
	outfile_path := path.Join(upload_folder, upload.Filename)
	outfile, err := os.OpenFile(outfile_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		components.Error("creating-outfile").Render(r.Context(), w)
		return
	}
	defer outfile.Close()

	// Copying
	if _, err = io.Copy(outfile, file); err != nil {
		components.Error("copying-file").Render(r.Context(), w)
		return
	}

	// Saving upload metadata
	if err = upload.Save(); err != nil {
		components.Error("saving-upload").Render(r.Context(), w)
		return
	}

	components.DisplayUpload(&upload).Render(r.Context(), w)
}

func (h *FileUploadHandler) Get(w http.ResponseWriter, r *http.Request) {
	pages.Page(
		pages.Publish(), "Uploading New File",
		pages.DEFAULT_PROPS,
	).Render(r.Context(), w)
}
