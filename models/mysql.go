package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	AesKey string = "$hejGRT^$*#@#12o"
)

const (
	_DB_Driver = "mysql"
)

func RegisterDB() {
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_schema := beego.AppConfig.String("db_schema")
	db_user := beego.AppConfig.String("db_user")
	db_passwd := beego.AppConfig.String("db_passwd")

	jdbcUrl := db_user + ":" + db_passwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_schema + "?charset=utf8" + "&loc=Local"
	beego.Info(fmt.Sprintf("connect to mysql server %v successfully !", db_host))
	orm.RegisterDriver(_DB_Driver, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_Driver, jdbcUrl, 30)

}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func Md5Encode(src []byte) string {
	m := md5.Sum(src)
	return hex.EncodeToString(m[:])
}

func AESEncode(msg, key string) (string, error) {
	if len(key) == 16 {
		var iv = []byte(key)[:aes.BlockSize]
		c := make([]byte, len(msg))
		be, err := aes.NewCipher([]byte(key))
		if err != nil {
			return "", err
		}
		e := cipher.NewCFBEncrypter(be, iv)
		e.XORKeyStream(c, []byte(msg))
		b64 := base64.StdEncoding.EncodeToString(c)
		b64 = strings.Replace(b64, "/", "-", -1)
		return b64, nil
	} else {
		return "", fmt.Errorf("%s", "Key length is not equal to 16.")
	}
}

func AESDecode(enmsg, key string) (string, error) {
	if len(key) == 16 {
		enmsg = strings.Replace(enmsg, "-", "/", -1)
		msg, err := base64.StdEncoding.DecodeString(enmsg)
		if nil != err {
			return "", err
		}
		var iv = []byte(key)[:aes.BlockSize]
		d := make([]byte, len(msg))
		var bd cipher.Block
		bd, err = aes.NewCipher([]byte(key))
		if err != nil {
			return "", err
		}
		e := cipher.NewCFBDecrypter(bd, iv)
		e.XORKeyStream(d, msg)
		return string(d), nil
	} else {
		return "", fmt.Errorf("%s", "Key length is not equal to 16.")
	}
}
