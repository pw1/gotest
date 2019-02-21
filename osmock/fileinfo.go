package osmock

import (
	"os"
	"time"
)

// FileInfo is an os.FileInfo implementation for testing purposes. It is a simple data struct
// that specifies which values must be returned by methods of the os.FileInfo interface.
type FileInfo struct {
	// NameV is the Value that Name() will return
	NameV string

	// SizeV is the Value that Size() will return
	SizeV int64

	// ModeV is the Value that Mode() will return
	ModeV os.FileMode

	// ModTimeV is the Value that ModTime() will return
	ModTimeV time.Time

	// IsDirV is the Value that IsDir() will return
	IsDirV bool

	// SysV is the Value that Sys() will return
	SysV interface{}
}

// Name returns the value of the NameV field.
func (f *FileInfo) Name() string {
	return f.NameV
}

// Size returns the value of the SizeV field.
func (f *FileInfo) Size() int64 {
	return f.SizeV
}

// Mode returns the value of the ModeV field.
func (f *FileInfo) Mode() os.FileMode {
	return f.ModeV
}

// ModTime returns the value of the ModTimeV field.
func (f *FileInfo) ModTime() time.Time {
	return f.ModTimeV
}

// IsDir returns the value of the IsDirV field.
func (f *FileInfo) IsDir() bool {
	return f.IsDirV
}

// Sys returns the value of the SysV field.
func (f *FileInfo) Sys() interface{} {
	return f.SysV
}
