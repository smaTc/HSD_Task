package hsdtask

import (
	"log"
	"os"
	"strconv"
)

var Logger = log.New(os.Stderr, "[HSD-TASK] ", 0)

func GetTasks(id ...int) []byte {
	var data []Task
	if id != nil {
		ReadByIdData(&data, id)
	} else {
		ReadAllData(&data)
	}

	if len(data) == 0 {
		return GenerateJsonResponse("response", "task not found or no tasks stored in database")
	}

	var json []byte
	var err error
	if len(data) == 1 {
		json, err = ObjectToJson(data[0])
	} else {
		json, err = ObjectToJson(data)
	}

	if err != nil {
		Logger.Println("error when reading tasks from db")
		return GenerateJsonResponse("error", "reading tasks from db failed")
	}

	return json
}

func CreateTask(data *Task) []byte {
	task := *data
	task.ID = 0

	dbErr := CreateData(&task)
	if dbErr == nil {
		return GenerateJsonResponse("message", "successfully created data with ID "+formatTaskID(task.ID))
	}
	return GenerateJsonResponse("error", "data creation failed")
}

func UpdateTask(id int, newTask *Task) []byte {
	var task Task

	err := ReadByIdData(&task, []int{id})
	if err != nil {
		return GenerateJsonResponse("error", err.Error())
	}

	if task.ID == 0 {
		return GenerateJsonResponse("message", "task with ID "+formatTaskID(task.ID)+" was not found in database")
	}

	newTask.ID = task.ID
	dbErr := UpdateData(&newTask)
	if dbErr == nil {
		return GenerateJsonResponse("message", "successfully updated data with ID "+formatTaskID(task.ID))
	}
	return GenerateJsonResponse("error", "data update failed")
}

func DeleteTask(id int) []byte {
	var task Task

	err := ReadByIdData(&task, []int{id})
	if err != nil {
		return GenerateJsonResponse("error", err.Error())
	}

	if task.ID == 0 {
		return GenerateJsonResponse("message", "task with ID "+formatTaskID(task.ID)+" was not found in database")
	}

	dbErr := DeleteData(&task)
	if dbErr == nil {
		return GenerateJsonResponse("message", "successfully deleted data with ID "+formatTaskID(task.ID))
	}
	return GenerateJsonResponse("error", "data deletion failed")
}

func formatTaskID(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}
