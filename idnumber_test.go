package idnumber

import (
	"testing"
)

func TestIDNumber(t *testing.T) {
	// 这里的身份证号码由身份证号码生成器生成，非真实信息
	idnubmers := []string{"11010519430606229X", "110105194306066397", "110105194306062396"}
	for _, idNumber := range idnubmers {
		info, err := Parse(idNumber)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v\n", info)
	}
}
