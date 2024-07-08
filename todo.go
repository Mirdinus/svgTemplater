package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var tasksEndpoint = "tasks?project_id=%d"

func fetchTodos(todoistProjectID int) TodoistResponse {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(Config.Todoist.BaseUrl, fmt.Sprintf(tasksEndpoint, todoistProjectID)), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Config.Todoist.ApiKey))
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error fetching todos: %v", err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading todos: %v", err)
	}

	todos := TodoistResponse{}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		log.Fatalf("Error unmarshalling todos: %v", err)
	}

	return todos
}

func ProcessTodo(template string) string {
	formated := template

	todos := fetchTodos(Config.Todoist.ProjectID)
	for i := 0; i < len(todos); i++ {
		formated = strings.ReplaceAll(formated, fmt.Sprintf("{notification_%d}", i+1), todos[i].Content)
	}

	for i := 0; i < 10; i++ {
		formated = strings.ReplaceAll(formated, fmt.Sprintf("{notification_%d}", i+1), "")
	}

	return formated
}
