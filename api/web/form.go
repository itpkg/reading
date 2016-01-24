package web

type Form struct {
	Id     string   `json:"id"`
	Method string   `json:"method"`
	Action string   `json:"action"`
	Fields []*Field `json:"fields"`
}

func (p *Form) Text(id, val string) {
	p.Fields = append(p.Fields, &Field{Id: id, Value: val, Type: "text"})
}
func (p *Form) TextArea(id, val string) {
	p.Fields = append(p.Fields, &Field{Id: id, Value: val, Type: "textarea"})
}

func (p *Form) Password(id string) {
	p.Fields = append(p.Fields, &Field{Id: id, Type: "password"})
}

func (p *Form) Email(id string) {
	p.Fields = append(p.Fields, &Field{Id: id, Type: "email"})
}

func NewForm(id, action string) *Form {
	return &Form{Id: id, Method: "POST", Action: action, Fields: make([]*Field, 0)}
}

type Field struct {
	Id    string      `json:"id"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
