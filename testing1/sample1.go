package main


type Message struct{
	// ....data here
}

func (m *Message) Send(email, subject string, body []byte) error{
	//...logic
	return nil
}
