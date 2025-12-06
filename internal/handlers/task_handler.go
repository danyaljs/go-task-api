package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"task-api/internal/models"
	"task-api/internal/repository"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{
		repo: repo,
	}
}

type ErrorResponse struct {
	Error string `json:"error" example:"task not found"`
}

type SuccessResponse struct {
	Message string       `json:"message" example:"task created successfully"`
	Task    *models.Task `json:"task,omitempty"`
}

type ListResponse struct {
	Message string         `json:"message" example:"tasks retrieved successfully"`
	Tasks   []*models.Task `json:"tasks"`
	Count   int            `json:"count" example:"5"`
}

func (h *TaskHandler) respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *TaskHandler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, ErrorResponse{Error: message})
}

func (h *TaskHandler) extractID(path string) (int, error) {

	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path: no ID found")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, models.ErrInvalidID
	}

	if id <= 0 {
		return 0, models.ErrInvalidID
	}

	return id, nil
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleGetTasks(w, r)

	case http.MethodPost:
		h.handleCreateTasks(w, r)

	default:
		h.respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {

	id, err := h.extractID(r.URL.Path)
	if err != nil {
		h.respondError(w, http.StatusBadRequest, models.ErrInvalidID.Error())
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.handleGetTask(w, r, id)
	case http.MethodPut:
		h.handleUpdateTask(w, r, id)
	case http.MethodDelete:
		h.handleDeleteTask(w, r, id)
	default:
		h.respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *TaskHandler) handleGetTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.repo.GetAll()

	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to retrieve tasks")
		return
	}

	h.respondJSON(w, http.StatusOK, ListResponse{
		Message: "Tasks retrieved successfully",
		Tasks:   tasks,
		Count:   len(tasks),
	})
}

func (h *TaskHandler) handleCreateTasks(w http.ResponseWriter, r *http.Request) {

	var req models.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		h.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.repo.Create(req.Title, req.Description)

	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to create task")
		return
	}

	h.respondJSON(w, http.StatusCreated, SuccessResponse{
		Message: "task created successfully",
		Task:    task,
	})
}

func (h *TaskHandler) handleGetTask(w http.ResponseWriter, r *http.Request, id int) {

	task, err := h.repo.GetByID(id)
	if err == models.ErrTaskNotFound {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to retrieve task")
		return
	}

	h.respondJSON(w, http.StatusOK, SuccessResponse{
		Message: "task retrieved successfully",
		Task:    task,
	})
}

func (h *TaskHandler) handleUpdateTask(w http.ResponseWriter, r *http.Request, id int) {

	var req models.UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := req.Validate(); err != nil {
		h.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.repo.Update(id, &req)
	if err == models.ErrTaskNotFound {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to update task")
		return
	}

	h.respondJSON(w, http.StatusOK, SuccessResponse{
		Message: "task updated successfully",
		Task:    task,
	})
}

func (h *TaskHandler) handleDeleteTask(w http.ResponseWriter, r *http.Request, id int) {

	err := h.repo.Delete(id)

	if err == models.ErrTaskNotFound {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}

	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to delete task")
		return
	}

	h.respondJSON(w, http.StatusOK, SuccessResponse{
		Message: "task deleted successfully",
	})
}
