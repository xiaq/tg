package tg

import (
	"unicode"
	"unicode/utf8"
)

func hanOrHangul(r rune) bool {
	return unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r)
}

func (u *User) DisplayName() string {
	if u.LastName == nil {
		return u.FirstName
	}
	lastName := *u.LastName
	r, _ := utf8.DecodeRuneInString(u.FirstName)
	s, _ := utf8.DecodeRuneInString(lastName)
	if hanOrHangul(r) && hanOrHangul(s) {
		return lastName + u.FirstName
	} else {
		return u.FirstName + " " + lastName
	}
}
