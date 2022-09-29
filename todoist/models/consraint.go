package models

// TodoistRequestModelConstraint is a unifier for the todist
// struct models and is used as a generics constraint.
type TodoistRequestModelConstraint interface {
	Endpoint() string
}

// TodoistResponseModelConstraint is a unifier for the todist
// struct models and is used as a generics constraint.
type TodoistResponseModelConstraint interface {
	ResponseModel()
}
