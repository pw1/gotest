package uuidfactory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

// TestFactorySuite is the test function that runs the tests in the FactorySuite.
func TestFactorySuite(t *testing.T) {
	suite.Run(t, new(FactorySuite))
}

// TaskStorageSuite is the test suite for the TaskStorage object.
type FactorySuite struct {
	suite.Suite
}

func (t *FactorySuite) TestNew() {
	f := NewFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"06ef35bc-b8ae-4eff-ab81-2b2abd109fb5",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.Equal(uuid.MustParse("ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92"), f.New())
	t.Equal(uuid.MustParse("06ef35bc-b8ae-4eff-ab81-2b2abd109fb5"), f.New())
	t.Equal(uuid.MustParse("367dfdd2-bccb-41db-8364-cc8dd10e9f2a"), f.New())
}

// Tests that NewFactory panics when it receives an invalid UUID.
func (t *FactorySuite) TestNewFactoryPanics() {
	uuids := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"INVALID_UUID",
	}
	t.Panics(func() { NewFactory(uuids) })
}

// Check that New() panics when it is called too many times.
func (t *FactorySuite) TestNewPanic() {
	f := NewFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
	})
	f.New()
	t.Panics(func() { f.New() })
}

func (t *FactorySuite) TestAll() {
	f := NewFactory([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.False(f.AllCreated())
	f.New()
	t.False(f.AllCreated())
	f.New()
	t.True(f.AllCreated())
}
