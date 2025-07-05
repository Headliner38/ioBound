package worker

import (
	"math/rand/v2"
	"time"

	"github.com/Headliner38/go-project4.git/Desktop/ioBound/pkg/models"
)

func ExecuteTask(task *models.Task) { // функция симулирующая работу I/O операции
	time.Sleep(15 * time.Second) // задержка перед запуском операции статус pending
	task.UpdateTaskStatus(models.StatusRunning)
	time.Sleep(1 * time.Minute)
	if rand.Float64() < 0.5 { // иммитация ошибки во время выполнения задачи, шанс = 50%
		task.UpdateTaskStatus(models.StatusFailed)
		task.Result = "Task failed by unexpected error"
		return
	}
	task.UpdateTaskStatus(models.StatusCompleted)
	task.Result = "Task completed successfully"
}
