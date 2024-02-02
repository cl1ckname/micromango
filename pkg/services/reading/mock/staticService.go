package mock

import "micromango/pkg/common"

type StaticServiceMock struct{}

func (s *StaticServiceMock) UploadPage(_ string, _ string, _ common.File) (string, error) {
	return "path/to/file", nil
}
