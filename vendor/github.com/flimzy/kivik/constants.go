package kivik

const (
	// KivikVersion is the version of the Kivik library.
	KivikVersion = "1.0.0-beta"
	// KivikVendor is the vendor string reported by this library.
	KivikVendor = "Kivik"
)

// SessionCookieName is the name of the CouchDB session cookie.
const SessionCookieName = "AuthSession"

// UserPrefix is the mandatory CouchDB user prefix.
// See http://docs.couchdb.org/en/2.0.0/intro/security.html#org-couchdb-user
const UserPrefix = "org.couchdb.user:"

// EndKeySuffix is a high Unicode character (0xfff0) useful for appending to an
// endkey argument, when doing a ranged search, as described here:
// http://couchdb.readthedocs.io/en/latest/ddocs/views/collation.html#string-ranges
//
// Example, to return all results with keys beginning with "foo":
//
//    rows, err := db.Query(context.TODO(), "ddoc", "view", map[string]interface{}{
//        "startkey": "foo",
//        "endkey":   "foo" + kivik.EndKeySuffix,
//    })
const EndKeySuffix = string(0xfff0)

// HTTP methods supported by CouchDB. This is almost an exact copy of the
// methods in the standard http package, with the addition of MethodCopy, and
// a few methods left out which are not used by CouchDB.
const (
	MethodGet    = "GET"
	MethodHead   = "HEAD"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	MethodCopy   = "COPY"
)

// HTTP response codes permitted by the CouchDB API.
// See http://docs.couchdb.org/en/1.6.1/api/basics.html#http-status-codes
const (
	StatusOK                           = 200
	StatusCreated                      = 201
	StatusAccepted                     = 202
	StatusFound                        = 302
	StatusNotModified                  = 304
	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusResourceNotAllowed           = 405
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusPreconditionFailed           = 412
	StatusBadContentType               = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusInternalServerError          = 500
	// StatusNotImplemented is not returned by CouchDB proper. It is used by
	// Kivik for optional features which are not implemented by some drivers.
	StatusNotImplemented = 501
)
