package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T){
	record:=httptest.NewRecorder()
	req,err:=http.NewRequest(http.MethodGet,"/",nil)
	if err!=nil {
		t.Fatal(err)
	}
	mainHandler(record,req)
	if record.Body.String()!="hello world" {
		t.Fatalf("Response body: expected %s got %s",record.Body.String(),"hello world")
	}
}