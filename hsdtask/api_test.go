package hsdtask_test

import (
	"testing"

	"github.com/smaTc/HSD_Task/hsdtask"
)

func TestApiCrudAbstraction(t *testing.T) {
	hsdtask.InitDB(testDbName)
	task := hsdtask.NewTask("Test Title", "Test Description", false)
	task2 := hsdtask.NewTask("Test Title 2", "Test Description 2", false)

	resp := hsdtask.CreateTask(&task)
	expectedResp := hsdtask.GenerateJsonResponse("message", "successfully created data with ID 1")
	if string(resp) != string(expectedResp) {
		t.Log(string(resp))
		t.Log(string(expectedResp))
		t.Error("error when creating task via crud abstraction")
	}

	resp2 := hsdtask.CreateTask(&task2)
	expectedResp2 := hsdtask.GenerateJsonResponse("message", "successfully created data with ID 2")
	if string(resp2) != string(expectedResp2) {
		t.Log(string(resp2))
		t.Log(string(expectedResp2))
		t.Error("error when creating task2 via crud abstraction")
	}

	taskListDb, err := hsdtask.JsonToObject[[]hsdtask.Task](hsdtask.GetTasks())
	if err != nil || len(taskListDb) != 2 {
		t.Error("errror when reading all tasks")
	}

	task.Description = "changed"
	resp3 := hsdtask.UpdateTask(int(taskListDb[0].ID), &task)
	expectedResp3 := hsdtask.GenerateJsonResponse("message", "successfully updated data with ID 1")
	if string(resp3) != string(expectedResp3) {
		t.Log(task)
		t.Log(string(resp3))
		t.Log(string(expectedResp3))
		t.Error("error when creating task via crud abstraction")
	}

	resp4 := hsdtask.DeleteTask(int(taskListDb[1].ID))
	expectedResp4 := hsdtask.GenerateJsonResponse("message", "successfully deleted data with ID 2")
	if string(resp4) != string(expectedResp4) {
		t.Log(string(resp4))
		t.Log(string(expectedResp4))
		t.Error("error when creating task via crud abstraction")
	}

	hsdtask.RemoveDb(testDbName)
}
