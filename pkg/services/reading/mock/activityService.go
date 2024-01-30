package mock

// Mock for activity service
type ActivityServiceMock struct{}

func (a *ActivityServiceMock) GetReadChapters(_ string, _ string) ([]string, error) {
	return []string{"1", "2"}, nil
}
