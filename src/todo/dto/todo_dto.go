package dto

// swagger:parameters CreateTodoRequest
type CreateTodoRequest struct {
	// required: true
	Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}

// swagger:parameters UpdateTodoRequest
type UpdateTodoRequest struct {
	// required: true
    Task string `form:"task" json:"task" xml:"task"  binding:"required,min=1,max=300"`
    // required: true
    Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}
