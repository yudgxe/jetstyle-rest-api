package handler

type CreateInput struct {
	Owner int32  `json:"owner" minimum:"1"`
	Name  string `json:"name"  minimum:"4"`
}

type UpdateInput struct {
	CreateInput
	IsComplete bool `json:"is_complete" default:"false"`
}
