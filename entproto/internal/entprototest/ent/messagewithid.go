// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithid"
	"entgo.io/ent/dialect/sql"
)

// MessageWithID is the model entity for the MessageWithID schema.
type MessageWithID struct {
	config
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithID) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithid.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MessageWithID", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithID fields.
func (mwi *MessageWithID) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithid.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mwi.ID = int32(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this MessageWithID.
// Note that you need to call MessageWithID.Unwrap() before calling this method if this MessageWithID
// was returned from a transaction, and the transaction was committed or rolled back.
func (mwi *MessageWithID) Update() *MessageWithIDUpdateOne {
	return NewMessageWithIDClient(mwi.config).UpdateOne(mwi)
}

// Unwrap unwraps the MessageWithID entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mwi *MessageWithID) Unwrap() *MessageWithID {
	_tx, ok := mwi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithID is not a transactional entity")
	}
	mwi.config.driver = _tx.drv
	return mwi
}

// String implements the fmt.Stringer.
func (mwi *MessageWithID) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithID(")
	builder.WriteString(fmt.Sprintf("id=%v", mwi.ID))
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithIDs is a parsable slice of MessageWithID.
type MessageWithIDs []*MessageWithID

func (mwi MessageWithIDs) config(cfg config) {
	for _i := range mwi {
		mwi[_i].config = cfg
	}
}
