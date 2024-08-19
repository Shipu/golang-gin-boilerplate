package dto

// swagger:parameters CreateTodoRequest
type CreateTodoRequest struct {
	// required: true
	Todo string `form:"todo" json:"todo" xml:"todo"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}

// swagger:parameters UpdateTodoRequest
type UpdateTodoRequest struct {
	// required: true
	Todo string `form:"todo" json:"todo" xml:"todo"  binding:"required,min=1,max=300"`
	// required: true
	Status string `form:"status" json:"status" xml:"status"  binding:"required,oneof=active inactive"`
}
