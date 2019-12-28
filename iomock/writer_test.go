package iomock

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"
)

// TestWriterSuite is the test function that runs the tests in the WriterSuite.
func TestWriterSuite(t *testing.T) {
	suite.Run(t, new(WriterSuite))
}

// WriterSuite is the test suite for the Writer object.
type WriterSuite struct {
	suite.Suite
}

// TestImplementsIoWriter verifies that Writer implements io.Writer. This won't compile if it
// doesn't implement it correctly. The test itself doesn't do anything.
func (s *WriterSuite) TestImplementsIoWriter() {
	var writer io.Writer
	writer = &Writer{}
	s.NotNil(writer)
}

// TestDefaultErrOnFirstByte verifies that when ErrAfterBytes is not specified explicitely it will
// return an error on the first byte.
func (s *WriterSuite) TestDefaultErrOnFirstByte() {
	writer := &Writer{}
	n, err := writer.Write([]byte{0x00})
	s.Equal(0, n)
	s.NotNil(err)
	s.Same(writer.ErrorValue, err)
}

// TestErrOnSecondByte verifies that one byte can be written, but an error is returned on the second
// byte.
func (s *WriterSuite) TestErrOnSecondByte() {
	writer := &Writer{ErrAfterBytes: 1}
	n, err := writer.Write([]byte{0x00, 0x01})
	s.Equal(1, n)
	s.NotNil(err)
	s.Same(writer.ErrorValue, err)
}

// TestErrOnSecondWrite verifies that an error is returned the second time Write() is called, not
// the first time.
func (s *WriterSuite) TestErrOnSecondWrite() {
	writer := &Writer{ErrAfterBytes: 5}
	n, err := writer.Write([]byte{0x00, 0x01, 0x02})
	s.Equal(3, n)
	s.Nil(err)
	n, err = writer.Write([]byte{0x03, 0x04, 0x05})
	s.Equal(2, n)
	s.NotNil(err)
	s.Same(writer.ErrorValue, err)
}

// TestWithMyError verifies that the user-specified (by setting ErrorValue) error is returned,
// instead of a new one.
func (s *WriterSuite) TestWithMyError() {
	myErr := errors.New("This is my error")
	writer := &Writer{ErrorValue: myErr}
	n, err := writer.Write([]byte{0x00, 0x01, 0x02})
	s.Equal(0, n)
	s.Same(myErr, err)
	s.Same(myErr, writer.ErrorValue)
}

// TestErrOnMultipleCalls verifies that consecutive calls to Writes also return errors.
func (s *WriterSuite) TestErrOnMultipleCalls() {
	writer := &Writer{}
	n1, err1 := writer.Write([]byte{0x00, 0x01})
	n2, err2 := writer.Write([]byte{0x02, 0x03})
	s.Equal(0, n1)
	s.Equal(0, n2)
	s.NotNil(err1)
	s.Same(err1, err2)
}
