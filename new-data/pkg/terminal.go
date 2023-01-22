package pkg

type Terminal struct {
	Elements []interface{}
}

func (t *Terminal) AddElement(elem interface{}) {
	t.Elements = append(t.Elements)
}
