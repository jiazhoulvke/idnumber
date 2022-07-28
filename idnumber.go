package idnumber

import (
	"strconv"

	"golang.org/x/xerrors"
)

// 从 0 开始算：
// 0-5 位是地区行政代码
// 6-13 位是出生年月日
// 14-16 位是对地区内同一天出生的人编定的顺序号，奇数为男，偶数为女
// 17 位是校验码

type IDNumberInfo struct {
	Region Region `json:"region"`
	Year   int64  `json:"year"`
	Month  int64  `json:"month"`
	Day    int64  `json:"day"`
	//Gender 性别 0:女 1:男
	Gender uint8 `json:"gender"`
}

var ErrIDNumberInvalid = xerrors.New("idnumber is not a valid")

//checkCode 校验身份证号码
func checkCode(s string) (uint8, error) {
	var wn = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var n int
	for i := 0; i < 17; i++ {
		n1, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return 0, err
		}
		n += wn[i] * n1
	}
	modNumber := n % 11
	if modNumber == 10 {
		return 'X', nil
	}
	codes := []uint8{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	return codes[modNumber], nil
}

//Parse 解析身份证号码，如果不是正确的号码则会报错
func Parse(s string) (IDNumberInfo, error) {
	var idNumber IDNumberInfo
	if len(s) != 18 {
		return idNumber, ErrIDNumberInvalid
	}
	checkCode, err := checkCode(s)
	if err != nil {
		return idNumber, err
	}
	if checkCode != s[17] {
		return idNumber, ErrIDNumberInvalid
	}
	found := false
	for _, region := range Regions() {
		if region.Code == s[0:6] {
			found = true
			idNumber.Region = region
			break
		}
	}
	if !found {
		return idNumber, ErrIDNumberInvalid
	}
	idNumber.Year, err = strconv.ParseInt(s[6:10], 10, 64)
	if err != nil {
		return idNumber, ErrIDNumberInvalid
	}
	if idNumber.Year < 1900 {
		return idNumber, ErrIDNumberInvalid
	}
	idNumber.Month, err = strconv.ParseInt(s[10:12], 10, 64)
	if err != nil {
		return idNumber, ErrIDNumberInvalid
	}
	if idNumber.Month < 1 || idNumber.Month > 12 {
		return idNumber, ErrIDNumberInvalid
	}
	idNumber.Day, err = strconv.ParseInt(s[12:14], 10, 64)
	if err != nil {
		return idNumber, ErrIDNumberInvalid
	}
	if idNumber.Day < 1 || idNumber.Day > 31 {
		return idNumber, ErrIDNumberInvalid
	}
	seq, err := strconv.ParseInt(s[14:17], 10, 64)
	if err != nil {
		return idNumber, ErrIDNumberInvalid
	}
	idNumber.Gender = uint8(seq % 2)
	return idNumber, nil
}
