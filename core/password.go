package core

import (
	"github.com/pkg/errors"
	"math/rand"
	"time"
	"encoding/base64"
	"strings"
)

const PasswordLength = 256;

var ErrInvalidPassword = errors.New("不合法的密码");

type Password [PasswordLength]byte;

func init() {
	// 生成随机种子
	rand.Seed(time.Now().Unix());
}

func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:]);
}

// base64 解码获取密码
func ParsePassword(passwordString string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString));
	if err != nil || len(bs) != PasswordLength {
		return nil, ErrInvalidPassword
	}

	password := Password{}
	copy(password[:], bs);
	bs = nil;
	return &password, nil;
}

// 产生256位随机密码, 并且256位中不能出现重复的byte位
func RandPassword() *Password {
	intArr := rand.Perm(PasswordLength);
	password := &Password{}

	for i, v := range intArr {
		password[i] = byte(v);
		if i == v {
			return RandPassword();
		}
	}

	return password;
}
