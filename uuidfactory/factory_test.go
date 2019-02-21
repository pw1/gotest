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

// FactorySuite is the test suite for the Factory object.
type FactorySuite struct {
	suite.Suite
}

// Tests that NewFactory panics when it receives an invalid UUID.
func (t *FactorySuite) TestNewFactoryPanics() {
	uuids := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"INVALID_UUID",
	}
	t.Panics(func() { New(uuids) })
}

func (t *FactorySuite) TestNew() {
	f := New([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"06ef35bc-b8ae-4eff-ab81-2b2abd109fb5",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.Equal(uuid.MustParse("ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92"), f.New())
	t.Equal(uuid.MustParse("06ef35bc-b8ae-4eff-ab81-2b2abd109fb5"), f.New())
	t.Equal(uuid.MustParse("367dfdd2-bccb-41db-8364-cc8dd10e9f2a"), f.New())
}

func (t *FactorySuite) TestNewWithError() {
	strings := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"ERROR",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	}
	f := New(strings)
	for _, str := range strings {
		if str == "ERROR" {
			t.Panics(func() { f.New() })
		} else {
			t.Equal(uuid.MustParse(str), f.New())
		}
	}
}

func (t *FactorySuite) TestNewRandom() {
	strings := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"06ef35bc-b8ae-4eff-ab81-2b2abd109fb5",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	}
	f := New(strings)
	for _, str := range strings {
		expectedUUID := uuid.MustParse(str)
		actualUUID, actualErr := f.NewRandom()
		t.Equal(expectedUUID, actualUUID)
		t.Nil(actualErr)
	}
}

func (t *FactorySuite) TestNewRandomWithError() {
	strings := []string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"ERROR",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	}
	f := New(strings)
	for _, str := range strings {
		actualUUID, actualErr := f.NewRandom()
		if str == "ERROR" {
			t.Equal(uuid.Nil, actualUUID)
			t.NotNil(actualErr)
		} else {
			t.Equal(uuid.MustParse(str), actualUUID)
			t.Nil(actualErr)
		}
	}
}

// Check that New() panics when it is called too many times.
func (t *FactorySuite) TestNewPanic() {
	f := New([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
	})
	f.New()
	t.Panics(func() { f.New() })
}

func (t *FactorySuite) TestAllCreated() {
	f := New([]string{
		"ce8631ce-3b9d-4ac9-a9fa-d536d2e11a92",
		"367dfdd2-bccb-41db-8364-cc8dd10e9f2a",
	})
	t.False(f.AllCreated())
	f.New()
	t.False(f.AllCreated())
	f.New()
	t.True(f.AllCreated())
}
