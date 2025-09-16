package flow

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) NewProject(projectName string, moduleName string) error {
	return nil
}
