package handler

const EchoServiceName = "handler/EchoService"

type NewEchoService struct{}

func (s *NewEchoService) Echo(request string, reply *string) error {
	*reply = request
	return nil
}
