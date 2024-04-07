package liaison

type StatusEnum string

const (
	Success      StatusEnum = "success"
	Created      StatusEnum = "created"
	MultiStatus  StatusEnum = "multiStatus"
	UserError    StatusEnum = "userError"
	Unauthorized StatusEnum = "unauthorized"
	InfraError   StatusEnum = "infraError"
)

type LiaisonOutput struct {
	Status StatusEnum  `json:"status"`
	Body   interface{} `json:"body"`
}

func NewLiaisonOutput(status StatusEnum, body interface{}) LiaisonOutput {
	return LiaisonOutput{
		Status: status,
		Body:   body,
	}
}
