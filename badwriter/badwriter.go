package badwriter

import "errors"

// BadWriter implements the io.Writer interface, but returns an error after ErrAfterBytes bytes are
// written. BadWriter can return a specified error, or create its own error message.
type BadWriter struct {
	// ErrAfterBytes indicates after how many successfully written bytes it must return an error.
	ErrAfterBytes int

	// ErrorValue is the value that will be returned when more than ErrAfterBytes bytes are written.
	// If the ErrorValue is not set, then a new error is generated. In that case ErrorValue is also
	// set to the new error.
	ErrorValue error

	// bytesWritten keeps track of how many bytes have been written to this writer.
	bytesWritten int
}

// Write simulates writing bytes to the underlying data stream. Implements io.Writer.
func (w *BadWriter) Write(p []byte) (n int, err error) {
	newBytesWritten := w.bytesWritten + len(p)

	if newBytesWritten <= w.ErrAfterBytes {
		w.bytesWritten = newBytesWritten
		return len(p), nil
	}

	actualBytesWritten := w.ErrAfterBytes - w.bytesWritten
	w.bytesWritten = w.ErrAfterBytes
	if w.ErrorValue == nil {
		w.ErrorValue = errors.New("Not good")
	}

	return actualBytesWritten, w.ErrorValue
}
