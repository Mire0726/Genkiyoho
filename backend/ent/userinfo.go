// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Genkiyoho/ent/user"
	"Genkiyoho/ent/userinfo"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserInfo is the model entity for the UserInfo schema.
type UserInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserInfoQuery when eager-loading is set.
	Edges        UserInfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserInfoEdges holds the relations/edges for other nodes in the graph.
type UserInfoEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserInfoEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userinfo.FieldID, userinfo.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userinfo.FieldTitle, userinfo.FieldBody:
			values[i] = new(sql.NullString)
		case userinfo.FieldCreatedAt, userinfo.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserInfo fields.
func (ui *UserInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ui.ID = int(value.Int64)
		case userinfo.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ui.UserID = int(value.Int64)
			}
		case userinfo.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				ui.Title = value.String
			}
		case userinfo.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				ui.Body = value.String
			}
		case userinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ui.CreatedAt = value.Time
			}
		case userinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ui.UpdatedAt = value.Time
			}
		default:
			ui.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserInfo.
// This includes values selected through modifiers, order, etc.
func (ui *UserInfo) Value(name string) (ent.Value, error) {
	return ui.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserInfo entity.
func (ui *UserInfo) QueryUser() *UserQuery {
	return NewUserInfoClient(ui.config).QueryUser(ui)
}

// Update returns a builder for updating this UserInfo.
// Note that you need to call UserInfo.Unwrap() before calling this method if this UserInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ui *UserInfo) Update() *UserInfoUpdateOne {
	return NewUserInfoClient(ui.config).UpdateOne(ui)
}

// Unwrap unwraps the UserInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ui *UserInfo) Unwrap() *UserInfo {
	_tx, ok := ui.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserInfo is not a transactional entity")
	}
	ui.config.driver = _tx.drv
	return ui
}

// String implements the fmt.Stringer.
func (ui *UserInfo) String() string {
	var builder strings.Builder
	builder.WriteString("UserInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ui.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ui.UserID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(ui.Title)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(ui.Body)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ui.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ui.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserInfos is a parsable slice of UserInfo.
type UserInfos []*UserInfo
