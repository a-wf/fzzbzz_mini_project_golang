package main

import(
	"testing"
	"github.com/akdj/fizzbuzz/internal/services"
	"net/http"
	"net/http/httptest"
	"github.com/akdj/fizzbuzz/internal/routes"

	"github.com/gavv/httpexpect"
)

func TestFizzBuzzServiceValidAnswer(t * testing.T){

	answer := services.FizzBuzzAnswer(2, 3, 14, "salut", "yes");
	t.Log("answer")
	
	if answer != "1,salut,yes,salut,5,salutyes,7,salut,yes,salut,11,salutyes,13,salut" {
		t.Errorf("Failed")
	 }
}


func TestFizzBuzzAPIValid(t * testing.T){

	server := httptest.NewServer(routes.Init())

	e := httpexpect.New(t, server.URL)
	defer server.Close()


	e.GET("/v1/fizzbuzz").
		WithQuery("int1", 2).
		WithQuery("int2", 3).
		WithQuery("limit", 14).
		WithQuery("str1", "salut").
		WithQuery("str2", "yes").
		Expect().
		Status(http.StatusOK).
		JSON().Object().ContainsKey("answer").ValueEqual("answer", "1,salut,yes,salut,5,salutyes,7,salut,yes,salut,11,salutyes,13,salut")
}




func TestFizzBuzzAPIFailed1(t * testing.T){

	server := httptest.NewServer(routes.Init())

	e := httpexpect.New(t, server.URL)
	defer server.Close()

	e.GET("/v1/fizzbuzz").
		WithQuery("int2", 3).
		WithQuery("limit", 14).
		WithQuery("str1", "salut").
		WithQuery("str2", "yes").
		Expect().
		Status(400)
}

func TestFizzBuzzAPIFailed2(t * testing.T){

	server := httptest.NewServer(routes.Init())

	e := httpexpect.New(t, server.URL)
	defer server.Close()

	e.GET("/v1/fizzbuzz").
		WithQuery("int1", 2).
		WithQuery("int2", 3).
		WithQuery("limit", 14).
		WithQuery("str1", "salut").
		Expect().
		Status(400)
}