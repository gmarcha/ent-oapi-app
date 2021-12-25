// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/gmarcha/goswagger-ent-app/ent/event"
	"github.com/gmarcha/goswagger-ent-app/ent/schema"
	"github.com/gmarcha/goswagger-ent-app/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescDescription is the schema descriptor for description field.
	eventDescDescription := eventFields[2].Descriptor()
	// event.DefaultDescription holds the default value on creation for the description field.
	event.DefaultDescription = eventDescDescription.Default.(string)
	// eventDescDateCreation is the schema descriptor for date_creation field.
	eventDescDateCreation := eventFields[3].Descriptor()
	// event.DefaultDateCreation holds the default value on creation for the date_creation field.
	event.DefaultDateCreation = eventDescDateCreation.Default.(func() time.Time)
	// eventDescDateUpdate is the schema descriptor for date_update field.
	eventDescDateUpdate := eventFields[4].Descriptor()
	// event.DefaultDateUpdate holds the default value on creation for the date_update field.
	event.DefaultDateUpdate = eventDescDateUpdate.Default.(func() time.Time)
	// event.UpdateDefaultDateUpdate holds the default value on update for the date_update field.
	event.UpdateDefaultDateUpdate = eventDescDateUpdate.UpdateDefault.(func() time.Time)
	// eventDescID is the schema descriptor for id field.
	eventDescID := eventFields[0].Descriptor()
	// event.DefaultID holds the default value on creation for the id field.
	event.DefaultID = eventDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescTutor is the schema descriptor for tutor field.
	userDescTutor := userFields[5].Descriptor()
	// user.DefaultTutor holds the default value on creation for the tutor field.
	user.DefaultTutor = userDescTutor.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}