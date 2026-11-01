package knowledge

import "errors"

// ErrNotFound is returned by Entity when the adapter has no such entity.
var ErrNotFound = errors.New("knowledge: entity not found")
