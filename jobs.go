package model

const (
	JobTypePendingTripRequest = "pending_trip_request"
)
const (
	JobStatusPending  = "pending"
	JobStatusCanceled = "canceled"
	JobStatusError    = "error"
	JobStatusSuccess  = "success"
)

type Job struct {
	ID            string         `json:"id"`
	JobType       string         `json:"jobType"`
	CreatedAt     int64          `json:"createdAt"`
	ExecutionTime int64          `json:"executionTime"`
	JobStatus     string         `json:"JobStatus"`
	JobData       map[string]any `json:"jobData"`
}

type JobRequest struct {
	JobType       string         `json:"jobType"`
	ExecutionTime int64          `json:"executionTime"`
	JobData       map[string]any `json:"jobData"`
}
