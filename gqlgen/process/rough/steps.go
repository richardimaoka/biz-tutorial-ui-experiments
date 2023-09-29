package rough

type RoughStep struct {
	Step         string `json:"step"`
	Phase        string `json:"phase"`
	Type         string `json:"type"`
	Instruction  string `json:"instruction"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`
	ModalText    string `json:"modalText"`
	Commit       string `json:"commit"`
	Comment      string `json:"comment"`
}
