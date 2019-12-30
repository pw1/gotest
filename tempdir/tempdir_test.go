package tempdir

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestTempDirSuite runs the test suite.
func TestTempDirSuite(t *testing.T) {
	suite.Run(t, new(TempDirSuite))
}

// TempDirSuite contains the tests.
type TempDirSuite struct {
	suite.Suite

	TempDir TempDir
}

func (s *TempDirSuite) SetupSuite() {
	s.TempDir.Init(s)
}

func (s *TempDirSuite) TearDownSuite() {
	s.TempDir.Cleanup()
}

func (s *TempDirSuite) TestBase() {
	s.NotEmpty(s.TempDir.Base())
}

func (s *TempDirSuite) TestSubDir() {
	tmpDir := s.TempDir.SubDir()
	s.NotEmpty(tmpDir)
	s.DirExists(tmpDir)
	relPath, err := filepath.Rel(s.TempDir.Base(), tmpDir)
	s.Nil(err)
	s.False(strings.Contains(relPath, ".."))
}

func (s *TempDirSuite) TestFilePath() {
	tmpFilePath := s.TempDir.FilePath()
	s.NotEmpty(tmpFilePath)
	_, err := os.Stat(tmpFilePath)
	s.True(os.IsNotExist(err))
	relPath, err := filepath.Rel(s.TempDir.Base(), tmpFilePath)
	s.Nil(err)
	s.False(strings.Contains(relPath, ".."))
}

func (s *TempDirSuite) TestFile() {
	tmpFile := s.TempDir.File()
	s.NotNil(tmpFile)
	s.FileExists(tmpFile.Name())
	relPath, err := filepath.Rel(s.TempDir.Base(), tmpFile.Name())
	s.Nil(err)
	s.False(strings.Contains(relPath, ".."))
}

func (s *TempDirSuite) TestNotInitialized() {
	tempDir := &TempDir{}
	s.Panics(func() { tempDir.Base() })
}
