package message

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMessageFromString(t *testing.T) {
	tests := []struct {
		CQString string
		Expected Message
	}{
		{``, Message{}},
		{
			`Gorilla[CQ:text]`,
			Message{Text("Gorilla"), MessageSegment{Type: "text", Data: map[string]string{}}},
		},
		{
			`[CQ:face,id=123][CQ:face,id=1234]  `,
			Message{Face(123), Face(1234), Text("  ")},
		},
		{
			`ȐĉņþƦȻƝƃ[CQ:rcnb][CQ:ɌćƞßɌĆnƅŕĉ,a=b]`,
			Message{
				Text("ȐĉņþƦȻƝƃ"),
				MessageSegment{Type: "rcnb", Data: map[string]string{}},
				MessageSegment{Type: "ɌćƞßɌĆnƅŕĉ", Data: map[string]string{"a": "b"}},
			},
		},
		{
			`[CQ:face,id=123]🍟🍟🍟[CQ:face,id=1234]  `,
			Message{Face(123), Text(`🍟🍟🍟`), Face(1234), Text("  ")},
		},
		{
			`[CQ:face,id=123,id=123,id=123,id=123][CQ:face,id=1234]  [CQ:]`,
			Message{Face(123), Face(1234), Text("  "), MessageSegment{Type: "", Data: map[string]string{}}},
		},
		{
			`[CQ:image,file=file:///C:\path\to\my\img-123\###.png]`, // https://github.com/Mrs4s/go-cqhttp/issues/169
			Message{MessageSegment{Type: "image", Data: map[string]string{"file": "file:///C:\\path\\to\\my\\img-123\\###.png"}}},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := ParseMessageFromString(test.CQString)
			assert.Equal(t, test.Expected, got)
		})
	}
}

const bench = `rcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQ[CQ:face,id=123][CQ:face,id=1234][CQ:face,id=123][CQ:face,id=1234]ȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞ,a=b][CQ:rcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞrcnbCQȓČņÞ`

func BenchmarkParseMessageFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessageFromString(bench)
	}
}

func BenchmarkParseMessageFromStringWithUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessageFromStringWithUnsafe(bench)
	}
}
