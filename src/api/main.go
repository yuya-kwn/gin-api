package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
)

type(
	SampleHandler struct{
		Data string
	}
	SampleResponse struct{
		Status	string`json:"status"`
		Message  string`json:"message"`
		StatusCode string`json:"statuscode"`
	}
)

func main(){
	mux := &http.ServeMux{}
	mux.Handle("/view",NewSampleHandler())
	if err := http.ListenAndServe(":3000",mux);err != nil{
		log.Fatal(err)
	}
}

func NewSampleHandler()http.Handler{
	return &SampleHandler{"Hello,World!"}
}

func(h*SampleHandler)ServeHTTP(w http.ResponseWriter,r*http.Request){
	statuscode:=200

	res:= &SampleResponse{
		Status: "OK",
		Message: h.Data,
		StatusCode: strconv.Itoa(statuscode),
	}

	w.Header().Set("Content-Type","application/json; charset=UTF-8")

	w.WriteHeader(statuscode)

	buf,_:=json.MarshalIndent(res,""," ")
	_,_=w.Write(buf)
}

