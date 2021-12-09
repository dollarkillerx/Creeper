package request

type DelIndexRequest struct {
	Index string `json:"index"`
}

type LogSlimmingRequest struct {
	Index         string `json:"index"`
	RetentionDays int64  `json:"retention_days"`
}

type LogRequest struct {
	Index   string `json:"index"`
	Message string `json:"message"`
}

type SearchRequest struct {
	Index   string `json:"index"`
	KeyWord string `json:"key_word"`

	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}
