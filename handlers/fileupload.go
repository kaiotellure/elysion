package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/help"
	"github.com/ikaio/tailmplx/models"
	"github.com/ikaio/tailmplx/pages"
)

func NewFileUpload() *FileUploadHandler {
	return &FileUploadHandler{}
}

type FileUploadHandler struct{}

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

	upload := models.Upload{
		ID:     database.SF.Generate().String(),
		Title:  help.OR(r.FormValue("title"), "Untitled Upload"),
		Author: help.OR(r.FormValue("author"), "Anonymous"),
		At:     time.Now(),
	}

	attachments := r.MultipartForm.File["attachments"]
	// Make sure upload folder exists
	upload_folder := help.Env(help.UPLOAD_FOLDER, "web/public/upload")
	os.MkdirAll(upload_folder, os.ModePerm)

	for index, attachment := range attachments {
		file, err := attachment.Open()
		if err != nil {
			fmt.Println("[CONTINUE-ERROR] Could not open attachment: " + err.Error())
			continue
		}

		// My Image.png --> my-image.png
		sanitazed_name := strings.ToLower(strings.ReplaceAll(attachment.Filename, " ", "-"))
		sanitazed_name = upload.ID + "_" + strconv.Itoa(index) + "_" + sanitazed_name
		upload.Files = append(upload.Files, sanitazed_name)

		// Creating file on upload folder
		outfile_path := path.Join(upload_folder, sanitazed_name)
		outfile, err := os.OpenFile(outfile_path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("[CONTINUE-ERROR] Could not open outfile: " + err.Error())
			continue
		}
		defer outfile.Close()

		// Copying
		if _, err = io.Copy(outfile, file); err != nil {
			components.Error("copying-file").Render(r.Context(), w)
			return
		}
	}

	// Saving upload metadata
	if err := upload.Save(); err != nil {
		components.Error("saving-upload").Render(r.Context(), w)
		return
	}

	components.DisplayUpload(&upload).Render(r.Context(), w)
}

func (h *FileUploadHandler) Get(w http.ResponseWriter, r *http.Request) {
	pages.Page(
		pages.FilestoreUpload(), "Upload to FILE-STORE",
		pages.DEFAULT_PROPS,
	).Render(r.Context(), w)
}
