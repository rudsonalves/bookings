package forms

type errors map[string][]string

// Add adds an errors message for a givien form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return the frist error message
func (e errors) Get(field string) string {
	messageList, ok := e[field]
	if !ok {
		return ""
	}
	return messageList[0]
}
