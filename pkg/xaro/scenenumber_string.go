// Code generated by "stringer -type=SceneNumber"; DO NOT EDIT.

package xaro

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Game-0]
	_ = x[MainMenu-1]
}

const _SceneNumber_name = "GameMainMenu"

var _SceneNumber_index = [...]uint8{0, 4, 12}

func (i SceneNumber) String() string {
	if i < 0 || i >= SceneNumber(len(_SceneNumber_index)-1) {
		return "SceneNumber(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SceneNumber_name[_SceneNumber_index[i]:_SceneNumber_index[i+1]]
}
