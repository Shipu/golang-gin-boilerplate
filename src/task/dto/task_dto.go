package dto

// swagger:parameters CreateTaskRequest
type CreateTaskRequest struct {
	// required: true
	Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}

// swagger:parameters UpdateTaskRequest
type UpdateTaskRequest struct {
	// required: true
	Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}
