//ff:type feature=noun type=model
//ff:what segState — 복합어 나누기 동적 계획의 한 칸(가능 여부, 조각 수, 접사 수, 앞 칸, 이 조각을 아는지)
package noun

// segState is the best segmentation of a stem prefix.
type segState struct {
	ok      bool
	pieces  int
	affixes int
	prev    int
	known   bool
}
