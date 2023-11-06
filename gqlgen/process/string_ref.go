package process

// TODO retire this file
func stringRef(s string) *string {
	if s == "" {
		return nil
	} else {
		return &s
	}
}
