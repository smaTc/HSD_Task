package hsdtask_test

import (
	"bytes"
	"testing"

	"github.com/smaTc/HSD_Task/hsdtask"
)

var json []byte = []byte("{\"id\":0,\"title\":\"test task\",\"description\":\"test description\",\"completed\":true}")

func TestObjectToJson(t *testing.T) {
	task := hsdtask.NewTask("test task", "test description", true)
	json2, err := hsdtask.ObjectToJson(task)
	if err != nil {
		t.Error("error when parsing object to json:", err)
		return
	}

	if !bytes.Equal(json, json2) {
		t.Log(string(json))
		t.Log(string(json2))
		t.Error("json parser creates wrong output")
	}

}

func TestJsonToObject(t *testing.T) {
	task, err := hsdtask.JsonToObject[hsdtask.Task](json)
	if err != nil {
		t.Error("error when parsing json to object:", err)
		return
	}

	if task.Title != "test task" || task.Description != "test description" || !task.Completed {
		t.Error("json parsed with wrong output")
	}

	_, err2 := hsdtask.JsonToObject[hsdtask.Task]([]byte("error"))
	if err2 == nil {
		t.Error("wrong json to object parsing not throwing error")
	}
}
