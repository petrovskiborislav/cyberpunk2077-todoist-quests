package main

import (
	"log"

	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/powerpyx"
	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/services"
	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/todoist"
	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/todoist/models"
)

const (
	projectName = "Cyberpunk 2077"
)

func main() {
	todoistClient := todoist.NewClient()
	powerpyxClient := powerpyx.NewClient()

	projectResponse := models.CreateProjectResponseModel{}
	err := todoistClient.Create(models.CreateProjectRequestModel{Name: projectName, IsFavorite: true}, &projectResponse)
	if err != nil {
		log.Fatalf("Error while creating project: %v", err)
	}

	cyberpunkTodoistService := services.NewCyberpunkTodoistService(projectResponse.ID, powerpyxClient, todoistClient)
	err = cyberpunkTodoistService.CreateCyberpunk2077QuestsAsTodoistTasks()
	if err != nil {
		log.Fatalf("Error while creating Cyberpunk 2077 quests as Todoist tasks: %v", err)
	}
}
