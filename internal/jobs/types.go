package jobs

type Status string

const (
	StatusPending    Status = "pending"
	StatusProcessing Status = "processing"
	StatusDone       Status = "done"
	StatusFailed     Status = "failed"
)

type Job struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
	Result string `json:"result, omitempty"`
	Error  string `json:"error, omitempty"`
}
