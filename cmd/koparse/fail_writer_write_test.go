//ff:func feature=cli type=helper control=sequence
//ff:what Write — 늘 실패한다
package main

import "errors"

// Write implements io.Writer and always fails.
func (failWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }
