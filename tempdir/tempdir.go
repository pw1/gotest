package tempdir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

// TempDir manages a temp directory. This can be used for test suites, but other application are
// also possible. When TempDir is initialzized it creates a base temp directory. Typically, there is
// one base directory per test suite. The name of the base temp diriectory is based on the suite's
// name. Individual tests can request their own temp subdirectory within this base temp dir. Also
// temp files within the base temp dir can be requested. When all test are finished, then the whole
// base temp dir is removed.
type TempDir struct {
	// If set, then the temp dir is not deleted by Cleanup(). Usefull for troubleshooting.
	DontDeleteTempDir bool

	initialized bool   // Flag to track whether the TempDir has been initialized, or not.
	tempDir     string // Created base temp dir
	dirName     string // Name of the base temp dir.
}

// Init creates the base temp dir. Pass a pointer to your suite as suite argument. The name of the
// temp directory is based on the name of your suite. You can also pass any other object from which
// you want to derive the temp dir name. This function is typically called from your SetupSuite().
func (t *TempDir) Init(suite interface{}) {
	t.dirName = t.buildDirName(suite)
	fmt.Printf("Setup TempDir %s\n", t.dirName)

	// Find the system's temp dir
	tmpDir := os.TempDir()
	if tmpDir == "" {
		panic("unable to determine system's temp dir")
	}
	fi, err := os.Stat(tmpDir)
	if err != nil {
		panic(fmt.Sprintf("unable to check temp dir: %s", err))
	}
	if !fi.IsDir() {
		panic(fmt.Sprintf("temp dir %s is not a directory", tmpDir))
	}

	// Create the temp dir for this TempDit instance.
	t.tempDir = filepath.Join(tmpDir, t.dirName)
	err = os.Mkdir(t.tempDir, 0700)
	if err != nil {
		if os.IsExist(err) {
			msg := fmt.Sprintf("WARNING\nWARNING temp dir %s existed already", t.tempDir)
			msg += ", did you forget to call Cleanup(), typically from your own TearDownSuite()?"
			msg += "\nWARNING\n"
			fmt.Print(msg)
		} else {
			panic(fmt.Sprintf("unable to create temp dir: %s", err))
		}
	}

	fmt.Printf("Created TempDir %s\n", t.tempDir)
	t.initialized = true
}

// Cleanup removes the base temp directory, if DontDeleteTempDir is not set. This function is
// typically called from your TearDownSuite().
func (t *TempDir) Cleanup() {
	fmt.Println("TempDir cleanup")
	if (t.tempDir != "") && !t.DontDeleteTempDir {
		fmt.Printf("Removing TempDir %s\n", t.tempDir)
		err := os.RemoveAll(t.tempDir)
		if err != nil {
			panic(err)
		}
	}
	t.initialized = false
	t.tempDir = ""
	t.dirName = ""
}

func (t *TempDir) ensureInitialized() {
	if !t.initialized {
		panic("TempDir is not initialized, you must call Init() first")
	}
}

// buildDirName constructs the directory name of the base temp directory.
func (t *TempDir) buildDirName(suite interface{}) string {
	dirName := fmt.Sprintf("%s-%s", reflect.TypeOf(t).String(), reflect.TypeOf(suite).String())
	dirName = strings.ReplaceAll(dirName, "*", "")
	dirName = strings.ReplaceAll(dirName, ".", "-")
	return dirName
}

// Base returns the full path to the base temp dir.
func (t *TempDir) Base() string {
	t.ensureInitialized()
	return t.tempDir
}

// SubDir creates a new sub temp dir. This new temp dir is a subdirectory of the base temp dir.
func (t *TempDir) SubDir() string {
	dir, err := ioutil.TempDir(t.Base(), "")
	if err != nil {
		panic(err)
	}
	return dir
}

// FilePath constructs a new unique file path within the base temp dir.
func (t *TempDir) FilePath() string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return filepath.Join(t.Base(), u.String())
}

// File creates a new temporary file within the base temp dir and opens it for reading and
// writing.
func (t *TempDir) File() *os.File {
	f, err := ioutil.TempFile(t.Base(), "")
	if err != nil {
		panic(err)
	}
	return f
}
