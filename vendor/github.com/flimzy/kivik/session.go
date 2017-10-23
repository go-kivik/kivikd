package kivik

import (
	"context"
	"encoding/json"

	"github.com/flimzy/kivik/driver"
	"github.com/flimzy/kivik/errors"
)

// Session represents an authentication session.
type Session struct {
	// Name is the name of the authenticated user.
	Name string
	// Roles is a list of roles the user belongs to.
	Roles []string
	// AuthenticationMethod is the authentication method that was used for this
	// session.
	AuthenticationMethod string
	// AuthenticationDB is the user database against which authentication was
	// performed.
	AuthenticationDB string
	// AuthenticationHandlers is a list of authentication handlers configured on
	// the server.
	AuthenticationHandlers []string
	// RawResponse is the raw JSON response sent by the server, useful for
	// custom backends which may provide additional fields.
	RawResponse json.RawMessage
}

// Session returns information about the currently authenticated user.
func (c *Client) Session(ctx context.Context) (*Session, error) {
	if sessioner, ok := c.driverClient.(driver.Sessioner); ok {
		session, err := sessioner.Session(ctx)
		if err != nil {
			return nil, err
		}
		var ses Session = Session(*session)
		return &ses, nil
	}
	return nil, errors.Status(StatusNotImplemented, "kivik: driver does not support sessions")
}
