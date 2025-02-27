package hsdtask_test

import (
	"testing"

	"github.com/smaTc/HSD_Task/hsdtask"
)

func TestNewTask(t *testing.T) {
	task := hsdtask.NewTask("Test Task", "Short testing task", false)
	task2 := hsdtask.NewTask("Test Task 2", "Another short testing task", true)

	if task.Title != "Test Task" || task.Description != "Short testing task" || task.Completed {
		t.Error("Task creation error")
	}

	if task2.Title != "Test Task 2" || task2.Description != "Another short testing task" || !task2.Completed {
		t.Error("Task creation error")
	}

}
