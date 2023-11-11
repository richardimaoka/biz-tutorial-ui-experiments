package state2

type FromRowFields struct {
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`
}
