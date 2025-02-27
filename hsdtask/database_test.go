package hsdtask_test

import (
	"testing"

	"github.com/smaTc/HSD_Task/hsdtask"
)

var testDbName = "test.db"

func TestDatabase(t *testing.T) {
	hsdtask.InitDB(testDbName)
	task := hsdtask.NewTask("title 1", "description 1", true)
	task2 := hsdtask.NewTask("title 2", "description 2", false)
	task3 := hsdtask.NewTask("title 3", "description 3", false)

	hsdtask.CreateData(&task)
	hsdtask.CreateData(&task2)

	var tasksFromDb []hsdtask.Task
	hsdtask.ReadAllData(&tasksFromDb)

	if len(tasksFromDb) != 2 {
		t.Error("Wrong number of objects in db")
	}

	hsdtask.CreateData(&task3)

	updTask := tasksFromDb[0]
	updTask.Title = "UPDATED"
	hsdtask.UpdateData(&updTask)

	var idRead hsdtask.Task
	hsdtask.ReadByIdData(&idRead, []int{1})

	if idRead.Title != "UPDATED" {
		t.Error("error when updating task in db")
	}

	hsdtask.DeleteData(tasksFromDb[1])

	tasksFromDb = []hsdtask.Task{}
	hsdtask.ReadAllData(&tasksFromDb)

	if len(tasksFromDb) != 2 {
		t.Error("error when deleting data from db")
	}

	hsdtask.RemoveDb(testDbName)
}
