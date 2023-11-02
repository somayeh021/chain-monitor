// Code generated by "stringer -type TokenType"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TokenTypeUnknown-0]
	_ = x[TokenTypeETH-1]
	_ = x[TokenTypeERC20-2]
	_ = x[TokenTypeERC721-3]
	_ = x[TokenTypeERC1155-4]
}

const _TokenType_name = "TokenTypeUnknownTokenTypeETHTokenTypeERC20TokenTypeERC721TokenTypeERC1155"

var _TokenType_index = [...]uint8{0, 16, 28, 42, 57, 73}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
