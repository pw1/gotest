package uuidfactory

import (
	"fmt"

	"github.com/google/uuid"
)

// Factory is a factory that creates predictable UUIDs. This can be used in testing to create
// predictable UUIDs, instead of the random ones produced by uuid.New().
type Factory struct {
	uuids   []*uuid.UUID
	nrCalls int
}

// NewFactory creates a new Factory that will generate the specified UUIDs when New() of NewRandom()
// are called. If a specified UUID is "ERROR", then the corresponding call to New() will panic, or
// the corresponding call to NewRandom() wil return a nil UUID and an error.
// If any of the specified UUIDs is not "ERROR" and is invalid, then this function panics.
func NewFactory(uuids []string) *Factory {
	u := &Factory{}
	for _, uuidStr := range uuids {
		if uuidStr == "ERROR" {
			u.uuids = append(u.uuids, nil)
		} else {
			newUUID := uuid.MustParse(uuidStr)
			u.uuids = append(u.uuids, &newUUID)
		}
	}
	return u
}

// New creates a new UUID according to the UUIDs passed to the NewFactory function. If the
// corresponding UUID was specified as "ERROR", then New() will panic. If all expected UUIDs have
// already been created, then New() panics.
func (u *Factory) New() uuid.UUID {
	newUUID, err := u.NewRandom()
	if err != nil {
		panic(err)
	}
	return newUUID
}

// NewRandom creates a new UUID according to the UUIDs passed to the NewFactory function. If the
// corresponding UUID was specified as "ERROR", then NewRandom() returns error.  If all expected
// UUIDs have already been created, then NewRandom() panics.
func (u *Factory) NewRandom() (uuid.UUID, error) {
	if u.nrCalls >= len(u.uuids) {
		panic(fmt.Sprintf("Unexpected call to Factory.New(). Already had %d calls.", u.nrCalls))
	}
	newUUID := u.uuids[u.nrCalls]
	u.nrCalls++

	if newUUID == nil {
		return uuid.Nil, fmt.Errorf("ERROR you asked for an ERROR")
	}
	return *newUUID, nil
}

// AllCreated returns true when all the specifiec UUIDs have been created. Returns false otherwise.
func (u *Factory) AllCreated() bool {
	return len(u.uuids) == u.nrCalls
}
