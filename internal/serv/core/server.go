package core

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"homework-6/internal/app/producer"
	"homework-6/internal/serv/server"
	"io"
	"net/http"
	"strconv"
)

const queryParamKey = "id"

type FacadeServer struct {
	Server        server.IServer
	KafkaProducer *producer.Service
}

func (fc *FacadeServer) Get(w http.ResponseWriter, req *http.Request) {
	fc.KafkaProducer.Send("GET", "")
	id, codeReq := GetId(req)
	if codeReq == http.StatusBadRequest {
		toBodyErr(codeReq, w)
		return
	}
	codeReq, artReq := fc.Server.Get(req.Context(), id)
	if codeReq != http.StatusOK {
		toBodyErr(codeReq, w)
		return
	}

	articleJson, err := json.Marshal(artReq)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	_, err = w.Write(articleJson)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(codeReq)
}

func (fc *FacadeServer) DeleteArticle(w http.ResponseWriter, req *http.Request) {
	fc.KafkaProducer.Send("DELETE", "")
	id, codeReq := GetId(req)
	if codeReq == http.StatusBadRequest {
		toBodyErr(codeReq, w)
		return
	}
	codeReq = fc.Server.DeleteArticle(req.Context(), id)

	if codeReq != http.StatusOK {
		toBodyErr(codeReq, w)
		return
	}
	w.WriteHeader(codeReq)
}

func (fc *FacadeServer) CreatePost(w http.ResponseWriter, req *http.Request) {
	id, codeReq := GetId(req)
	if codeReq == http.StatusBadRequest {
		toBodyErr(codeReq, w)
		return
	}

	body, err := io.ReadAll(req.Body)
	fc.KafkaProducer.Send("POST", string(body))
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unm server.PostRequest
	if err := json.Unmarshal(body, &unm); err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codeReq, artReq := fc.Server.CreatePost(req.Context(), id, &unm)

	if codeReq == http.StatusBadRequest {
		toBodyErr(codeReq, w)
		return
	}

	articleJson, err := json.Marshal(artReq)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}

	_, err = w.Write(articleJson)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(codeReq)
}

func (fc *FacadeServer) CreateArticle(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fc.KafkaProducer.Send("POST", string(body))

	var unm server.ArticleRequest
	if err := json.Unmarshal(body, &unm); err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
	}

	codeReq, artReq := fc.Server.CreateArticle(req.Context(), &unm)

	if codeReq != http.StatusOK {
		toBodyErr(codeReq, w)
		return
	}

	articleJson, err := json.Marshal(artReq)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	_, err = w.Write(articleJson)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(codeReq)
}

func (fc *FacadeServer) UpdateArticle(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fc.KafkaProducer.Send("PUT", string(body))

	var unm server.ArticleRequest
	if err := json.Unmarshal(body, &unm); err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	codeReq, artReq := fc.Server.UpdateArticle(req.Context(), &unm)

	if codeReq != http.StatusOK {
		toBodyErr(codeReq, w)
		return
	}

	articleJson, err := json.Marshal(artReq)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	_, err = w.Write(articleJson)
	if err != nil {
		toBodyErr(http.StatusInternalServerError, w)
		return
	}
	w.WriteHeader(codeReq)
}

func GetId(req *http.Request) (int64, int) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		return 0, http.StatusBadRequest
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return keyInt, http.StatusOK
}

func CreateRouter(implementation FacadeServer) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/article/{id:[0-9]+}", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			implementation.Get(w, req)
		}
		if req.Method == http.MethodDelete {
			implementation.DeleteArticle(w, req)
		}
		if req.Method != http.MethodGet && req.Method != http.MethodDelete {
			toBodyErr(http.StatusBadRequest, w)
		}
	})

	router.HandleFunc("/article/{id:[0-9]+}/post", func(w http.ResponseWriter, req *http.Request) {

		if req.Method == http.MethodPost {
			implementation.CreatePost(w, req)
		} else {
			toBodyErr(http.StatusBadRequest, w)
		}
	})

	router.HandleFunc("/article", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			implementation.CreateArticle(w, req)
		}
		if req.Method == http.MethodPut {
			implementation.UpdateArticle(w, req)
		}
		if req.Method != http.MethodPost && req.Method != http.MethodPut {
			toBodyErr(http.StatusBadRequest, w)
		}
	})
	return router
}

type Error struct {
	Key string `json:"error_msg"`
}

func toBodyErr(key int, w http.ResponseWriter) {
	valBody := ""
	switch key {
	case 404:
		valBody = "not found"
	case 400:
		valBody = "invalid query"
	case 500:
		valBody = "Server error"
	}
	mr := Error{valBody}
	if val, err := json.Marshal(&mr); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		_, err = w.Write(val)
		w.WriteHeader(key)
	}
}
