package main
import "testing"

/*
this is how the program would look like
type Message struct{
	// ....data here
}

func (m *Message) Send(email, subject string, body []byte) error{
	//...logic
	return nil
}
Lets write a test case for this
*/

type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	return m.Send("someone@example.com", "some subject", problem)
}

// You have defined interface above. Now define a MockTestCase that Interface



type MockMessage struct {
	email, subject string
	body []byte
}

func (m *MockMessage) Send(email, subject string, body []byte) error {
	m.email = email
	m.subject = subject
	m.body = body

	return nil
}

func TestAlert(t *testing.T){
	msgr := new(MockMessage)
	body := []byte("some subject")

	Alert(msgr, body)
	if msgr.subject != "some subject"{
		t.Errorf("expected some subject but got %s", msgr.subject)
	}
}