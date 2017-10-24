// Package couchdb is a driver for connecting with a CouchDB server over HTTP.
// This version of the package is deprecated, and no longer receiving updates.
// Please use github.com/go-kivik/couchdb instead.
package couchdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/driver"
	"github.com/flimzy/kivik/driver/couchdb/chttp"
)

const (
	typeJSON  = "application/json"
	typeText  = "text/plain"
	typeMixed = "multipart/mixed"
)

// Couch represents the parent driver instance.
type Couch struct{}

var _ driver.Driver = &Couch{}

func init() {
	kivik.Register("couch", &Couch{})
}

// CompatMode is a flag indicating the compatibility mode of the driver.
type CompatMode int

// Compatibility modes
const (
	CompatUnknown = iota
	CompatCouch16
	CompatCouch20
)

// Known vendor strings
const (
	VendorCouchDB  = "The Apache Software Foundation"
	VendorCloudant = "IBM Cloudant"
)

type client struct {
	*chttp.Client
	Compat CompatMode
}

var _ driver.Client = &client{}

// NewClient establishes a new connection to a CouchDB server instance. If
// auth credentials are included in the URL, they are used to authenticate using
// CookieAuth (or BasicAuth if compiled with GopherJS). If you wish to use a
// different auth mechanism, do not specify credentials here, and instead call
// Authenticate() later.
func (d *Couch) NewClient(ctx context.Context, dsn string) (driver.Client, error) {
	chttpClient, err := chttp.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	c := &client{
		Client: chttpClient,
	}
	c.setCompatMode(ctx)
	return c, nil
}

func (c *client) setCompatMode(ctx context.Context) {
	info, err := c.Version(ctx)
	if err != nil {
		// We don't want to error here, in case the / endpoint is just blocked
		// for security reasons or something; but then we also can't infer the
		// compat mode, so just return, defaulting to CompatUnknown.
		return
	}
	switch info.Vendor {
	case VendorCouchDB, VendorCloudant:
		switch {
		case strings.HasPrefix(info.Version, "2.0."):
			c.Compat = CompatCouch20
		case strings.HasPrefix(info.Version, "1.6"):
			c.Compat = CompatCouch16
		}
	}
}

func (c *client) DB(_ context.Context, dbName string, options map[string]interface{}) (driver.DB, error) {
	forceCommit, err := forceCommit(options)
	if err != nil {
		return nil, err
	}
	if key, exists := getAnyKey(options); exists {
		return nil, fmt.Errorf("kivik: unrecognized option '%s'", key)
	}
	return &db{
		client:      c,
		dbName:      dbName,
		forceCommit: forceCommit,
	}, nil
}

type putResponse struct {
	ID  string `json:"id"`
	OK  bool   `json:"ok"`
	Rev string `json:"rev"`
}
