package osmock

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// TestFactorySuite is the test function that runs the tests in the FactorySuite.
func TestFileInfoSuite(t *testing.T) {
	suite.Run(t, new(FileInfoSuite))
}

// FileInfoSuite is the test suite for the Factory object.
type FileInfoSuite struct {
	suite.Suite
}

// TestFileInfoImplementOsFileInfo verifies that FileInfo implements the os.FileInfo interface.
func (s *FileInfoSuite) TestFileInfoImplementOsFileInfo() {
	// This won't compile if FileInfo doesn't implemen the os.FileInfo interface
	var fileInfo os.FileInfo = &FileInfo{}
	s.NotNil(fileInfo)
}

// TestReturnsValues verifies that all functions return the corresponding value.
func (s *FileInfoSuite) TestReturnsValues() {
	fileInfo := &FileInfo{
		NameV:    "the name",
		SizeV:    123456,
		ModeV:    os.ModeAppend | os.ModeTemporary,
		ModTimeV: time.Date(2019, time.January, 2, 3, 4, 5, 6, time.UTC),
		IsDirV:   true,
		SysV:     "hi there",
	}
	s.Equal("the name", fileInfo.Name())
	s.Equal(int64(123456), fileInfo.Size())
	s.Equal(os.ModeAppend|os.ModeTemporary, fileInfo.Mode())
	s.Equal(time.Date(2019, time.January, 2, 3, 4, 5, 6, time.UTC), fileInfo.ModTime())
	s.Equal(true, fileInfo.IsDir())
	s.Equal("hi there", fileInfo.Sys())
}
