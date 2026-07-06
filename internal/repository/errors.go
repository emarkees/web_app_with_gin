package repository

import "errors"

// ErrNoRecord is returned when a requested record is not found.
var ErrNoRecord = errors.New("repository: no matching record found")
