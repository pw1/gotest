package mock

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// TestUUIDFactorySuite is the test function that runs the tests in the UUIDFactorySuite.
func TestUUIDFactorySuite(t *testing.T) {
	suite.Run(t, new(UUIDFactorySuite))
}

// TaskStorageSuite is the test suite for the TaskStorage object.
type UUIDFactorySuite struct {
	suite.Suite
}

func (t *UUIDFactorySuite) TestNew() {
	f := NewUUIDFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"06ef35bc-b8ae-4eff-ab81-2b2abd109fb5",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.Equal(uuid.MustParse("ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92"), f.New())
	t.Equal(uuid.MustParse("06ef35bc-b8ae-4eff-ab81-2b2abd109fb5"), f.New())
	t.Equal(uuid.MustParse("367dfdd2-bccb-41db-8364-cc8dd10e9f2a"), f.New())
}

// Tests that NewUUIDFactory panics when it receives an invalid UUID.
func (t *UUIDFactorySuite) TestNewUUIDFactoryPanics() {
	uuids := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"INVALID_UUID",
	}
	t.Panics(func() { NewUUIDFactory(uuids) })
}

// Check that New() panics when it is called too many times.
func (t *UUIDFactorySuite) TestNewPanic() {
	f := NewUUIDFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
	})
	f.New()
	t.Panics(func() { f.New() })
}

func (t *UUIDFactorySuite) TestAll() {
	f := NewUUIDFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.False(f.AllCreated())
	f.New()
	t.False(f.AllCreated())
	f.New()
	t.True(f.AllCreated())
}
