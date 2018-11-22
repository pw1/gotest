package uuidfactory

import (
	"fmt"

	"github.com/google/uuid"
)

// Factory is a factory that creates predictable UUIDs. This can be used in testing to create
// predictable UUIDs, instead of the random ones produced by uuid.New().
type Factory struct {
	uuids   []uuid.UUID
	nrCalls int
}

// NewFactory creates a new Factory that will generate the specified UUIDs. If any of the
// specified UUIDs is invalid, then this function panics.
func NewFactory(uuids []string) *Factory {
	u := &Factory{}
	for _, uuidStr := range uuids {
		newUUID := uuid.MustParse(uuidStr)
		u.uuids = append(u.uuids, newUUID)
	}
	return u
}

// New creates a new UUID according to the UUIDs passed to the NewFactory function. If all
// expected UUIDs have already been created, then New() panics.
func (u *Factory) New() uuid.UUID {
	if u.nrCalls >= len(u.uuids) {
		panic(fmt.Sprintf("Unexpected call to Factory.New(). Already had %d calls.", u.nrCalls))
	}
	newUUID := u.uuids[u.nrCalls]
	u.nrCalls++
	return newUUID
}

// AllCreated returns true when all the specifiec UUIDs have been created. Returns false otherwise.
func (u *Factory) AllCreated() bool {
	return len(u.uuids) == u.nrCalls
}
