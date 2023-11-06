package state

func stringRef(s string) *string {
	if s == "" {
		return nil
	} else {
		return &s
	}
}
