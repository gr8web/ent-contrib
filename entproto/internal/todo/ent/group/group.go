// Code generated by ent, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "user_id"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_group"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
