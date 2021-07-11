package handlefunc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jimmiepr/file-system/model"
	"github.com/jimmiepr/file-system/service"
)

func GetFileList(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet == r.Method {
		service.ResponseIP(r)
		Fl := service.GetList("./mockup")
		var fileList []model.FileList
		for i := 0; i < len(Fl); i++ {
			fileList = append(fileList, model.FileList{
				Id:       i,
				Name:     Fl[i],
				DateTime: service.GetBirthTime(fmt.Sprintf("./mockup/%s", Fl[i])),
			})
		}
		json, err := json.Marshal(fileList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
