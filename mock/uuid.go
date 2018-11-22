package mock

import (
	"fmt"

	"github.com/google/uuid"
)

// UUIDFactory is a factory that creates predictable UUIDs. This can be used in testing to create
// predictable UUIDs, instead of the random ones produced by uuid.New().
type UUIDFactory struct {
	uuids   []uuid.UUID
	nrCalls int
}

// NewUUIDFactory creates a new UUIDFactory that will generate the specified UUIDs. If any of the
// specified UUIDs is invalid, then this function panics.
func NewUUIDFactory(uuids []string) *UUIDFactory {
	u := &UUIDFactory{}
	for _, uuidStr := range uuids {
		newUUID := uuid.MustParse(uuidStr)
		u.uuids = append(u.uuids, newUUID)
	}
	return u
}

// New creates a new UUID according to the UUIDs passed to the NewUUIDFactory function. If all
// expected UUIDs have already been created, then New() panics.
func (u *UUIDFactory) New() uuid.UUID {
	if u.nrCalls >= len(u.uuids) {
		panic(fmt.Sprintf("Unexpected call to UUIDFactory.New(). Already had %d calls.", u.nrCalls))
	}
	newUUID := u.uuids[u.nrCalls]
	u.nrCalls++
	return newUUID
}

// AllCreated returns true when all the specifiec UUIDs have been created. Returns false otherwise.
func (u *UUIDFactory) AllCreated() bool {
	return len(u.uuids) == u.nrCalls
}
