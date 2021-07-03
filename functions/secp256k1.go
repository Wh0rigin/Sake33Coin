package functions

import (
	"bytes"
	"compress/gzip"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
)

const (
	version            = byte(0x00)
	addreddChechsumLen = 4
	privKeyBytesLen    = 32
)

type GKey struct {
	privateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
}

func MakeNewKey(randKey string) (*GKey, error) {
	var err error
	var gkey GKey
	var curve elliptic.Curve // 椭圆曲线参数

	lenth := len(randKey)
	if lenth < 224/8+8 {
		err = errors.New("RandKey is too short. It mast be longer than 36 bytes.")
		return &gkey, err
	} else if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}

	private, err := ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		log.Panic(err)
	}
	gkey = GKey{private, private.PublicKey}
	return &gkey, nil
}

/*
对text签名
返回加密结果，结果为数字证书r、s的序列化后拼接，然后用hex转换为string
*/
func (k GKey) Sign(text []byte) (string, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.privateKey, text)
	if err != nil {
		return "", err
	}
	rt, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	st, err := s.MarshalText()
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer w.Close()
	_, err = w.Write([]byte(string(rt) + "+" + string(st)))
	if err != nil {
		return "", err
	}
	w.Flush()
	return hex.EncodeToString(b.Bytes()), nil
}

/*
校验文本内容是否与签名一致
使用公钥校验签名和文本内容
*/
func Verify(text []byte, signature string, pubKey *ecdsa.PublicKey) (bool, error) {
	rint, sint, err := getSign(signature)
	if err != nil {
		return false, err
	}
	result := ecdsa.Verify(pubKey, text, &rint, &sint)
	return result, nil
}

/*
证书分解
通过hex解码，分割成数字证书r，s
*/
func getSign(signature string) (rint, sint big.Int, err error) {
	byterun, err := hex.DecodeString(signature)
	if err != nil {
		err = errors.New("decrypt error," + err.Error())
		return
	}
	r, err := gzip.NewReader(bytes.NewBuffer(byterun))
	if err != nil {
		err = errors.New("decode error," + err.Error())
		return
	}
	defer r.Close()
	buf := make([]byte, 1024)
	count, err := r.Read(buf)
	if err != nil {
		fmt.Println("decode = ", err)
		err = errors.New("decode read error," + err.Error())
		return
	}
	rs := strings.Split(string(buf[:count]), "+")
	if len(rs) != 2 {
		err = errors.New("decode fail")
		return
	}
	err = rint.UnmarshalText([]byte(rs[0]))
	if err != nil {
		err = errors.New("decrypt rint fail, " + err.Error())
		return
	}
	err = sint.UnmarshalText([]byte(rs[1]))
	if err != nil {
		err = errors.New("decrypt sint fail, " + err.Error())
		return
	}
	return
}

func (k GKey) GetPrivKey() []byte {
	d := k.privateKey.D.Bytes()
	b := make([]byte, 0, privKeyBytesLen)
	priKey := paddedAppend(privKeyBytesLen, b, d) // []bytes type
	// s := byteToString(priKey)
	return priKey
}

func (k GKey) GetPubKey() []byte {
	pubKey := append(k.PublicKey.X.Bytes(), k.privateKey.Y.Bytes()...) // []bytes type
	// s := byteToString(pubKey)
	return pubKey
}

func paddedAppend(size uint, dst, src []byte) []byte {
	/*
		把src数组转换成指定长度的数组，长度不够则添加0
			:param size: 要返回的数组长度
			:param dst: byte类型的切片，需要返回的切片
			:param src: 原byte数组
	*/
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
