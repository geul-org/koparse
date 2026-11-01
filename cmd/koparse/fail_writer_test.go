//ff:type feature=cli type=model
//ff:what failWriter — 시험용: 쓰기마다 실패하는 Writer
package main

// failWriter fails every write.
type failWriter struct{}
