package production

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ikaio/tailmplx/internal/database"
	"github.com/ikaio/tailmplx/internal/help"
)

func HandlePut(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body reading failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var production database.Production
	err = json.Unmarshal(body, &production)
	if err != nil {
		http.Error(w, "json parsing failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	primary, err := help.GetImagePrimaryColorFromURL(production.Images.Cover)
	if err != nil {
		http.Error(w, "cover image primary color detection failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	production.Properties.PrimaryColor = primary

	err = production.Save()
	if err != nil {
		http.Error(w, "database saving failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Production was saved successfully."))
}
