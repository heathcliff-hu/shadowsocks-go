package core

import (
	"testing"
	"sort"
	"reflect"
)

func (password *Password) Len() int {
	return PasswordLength;
}

func (password *Password) Less(i, j int) bool {
	return password[i] < password[j];
}

func (password *Password) Swap(i, j int) {
	password[i], password[j] = password[j], password[i];
}

func TestRandPassword(t *testing.T) {
	password := RandPassword();
	t.Log(password);
	sort.Sort(password);
	for i := 0; i < PasswordLength; i++ {
		if password[i] != byte(i) {
			t.Error("不能出现一个重复的比特位，并且必须全部包含");
		}
	}
}

func TestRandPasswordString(t *testing.T) {
	password := RandPassword();
	passwordStr := password.String();
	decodePassword, err := ParsePassword(passwordStr);
	if err != nil {
		t.Error(err);
	} else {
		if (!reflect.DeepEqual(password, decodePassword)) {
			t.Error("密码转化为字符串与反解码之后的数据不对应");
		}
	}
}
