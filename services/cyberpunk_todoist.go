package services

import (
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/powerpyx"
	"github.com/petrovskiborislav/cyberpunk2077-todoist-quests/todoist"

	powerpyxModels "github.com/petrovskiborislav/cyberpunk2077-todoist-quests/powerpyx/models"
	todoistModels "github.com/petrovskiborislav/cyberpunk2077-todoist-quests/todoist/models"
)

// CyberpunkTodoistService is the interface of the service used for creating Cyberpunk 2077 quests as Todoist tasks.
type CyberpunkTodoistService interface {
	CreateCyberpunk2077QuestsAsTodoistTasks() error
}

// NewCyberpunkTodoistService returns a new instance of the CyberpunkTodoistService.
func NewCyberpunkTodoistService(projectID string, powerpyxClient powerpyx.Client, todoistClient todoist.Client) CyberpunkTodoistService {
	return &cyberpunkTodoistService{
		projectID:      projectID,
		powerpyxClient: powerpyxClient,
		todoistClient:  todoistClient,
	}
}

type cyberpunkTodoistService struct {
	projectID      string
	powerpyxClient powerpyx.Client
	todoistClient  todoist.Client
}

// CreateCyberpunk2077QuestsAsTodoistTasks creates Cyberpunk 2077 quests as Todoist tasks.
func (c *cyberpunkTodoistService) CreateCyberpunk2077QuestsAsTodoistTasks() error {
	questsByCategory, err := c.powerpyxClient.GetCyberpunk2077Quests()
	if err != nil {
		return err
	}

	for category, quests := range questsByCategory {
		questCategoryTask, err := c.createTask("", category)
		if err != nil {
			return err
		}

		questsWithSubCategories := lo.GroupBy(quests, func(quest powerpyxModels.Cyberpunk2077Quest) string {
			return quest.SubCategory
		})

		err = c.createSubCategoriesAndQuestsTasks(questCategoryTask.ID, questsWithSubCategories)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cyberpunkTodoistService) createSubCategoriesAndQuestsTasks(parentTaskID string, questsWithSubCategories map[string][]powerpyxModels.Cyberpunk2077Quest) error {
	var err error
	subCategoryTask := todoistModels.CreateTaskResponseModel{}
	parentID := parentTaskID

	for subCategory, quests := range questsWithSubCategories {
		if subCategory != "" {
			subCategoryTask, err = c.createTask(parentTaskID, subCategory)
			if err != nil {
				return err
			}
			parentID = subCategoryTask.ID
		}

		for _, quest := range quests {
			taskNameWithLink := make([]string, 0)
			for name, link := range quest.NameLinkPair {
				taskNameWithLink = append(taskNameWithLink, fmt.Sprintf("[%s](%s)", name, link))
			}
			questContent := strings.Join(taskNameWithLink, " / ")
			_, err = c.createTask(parentID, questContent)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *cyberpunkTodoistService) createTask(parentID, content string) (todoistModels.CreateTaskResponseModel, error) {
	questSubCategoryTask := todoistModels.CreateTaskResponseModel{}
	questSubCategoryTaskRequest := todoistModels.CreateTaskRequestModel{
		Content:   content,
		ParentID:  parentID,
		ProjectID: c.projectID,
	}
	err := c.todoistClient.Create(questSubCategoryTaskRequest, &questSubCategoryTask)
	if err != nil {
		return todoistModels.CreateTaskResponseModel{}, err
	}

	return questSubCategoryTask, nil
}
