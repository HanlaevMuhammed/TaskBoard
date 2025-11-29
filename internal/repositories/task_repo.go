package storage

import (
	"errors"
	"taskBoard_API/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetByID(id uint) (models.Task, error) {

	var task models.Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetByTitle(title string) (models.Task, error) {
	var task models.Task
	if err := r.db.Where("LOWER(title) = LOWER(?)", title).First(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) Create(task models.Task) error {
	return r.db.Create(&task).Error
}

func (r *TaskRepository) DeleteByID(id uint) error {
	if result := r.db.Where("id = ?", id).Delete(&models.Task{}); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("NOT FOUND!")
	}
	return nil
}

func (r *TaskRepository) DeleteByTitle(title string) error {
	if result := r.db.Where("LOWER(title) = LOWER(?)", title).Delete(&models.Task{}); result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("NOT FOUND")
	}
	return nil
}

func (r *TaskRepository) UpdateTask(id uint, updatedTask models.Task) error {

	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return errors.New("NOT FOUND")
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status
	return r.db.Save(&task).Error

}
