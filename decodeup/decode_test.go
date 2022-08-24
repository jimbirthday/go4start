package decodeup

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

const base = `bnthPwD541YfEaBduYMdmg==eZh04vwhSUyR5PpBR+05EQ==ysPdkx3gAUiqJa05WapFKg==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`

func TestGCM(t *testing.T) {
	metadata := base[:72]
	//header, err := base64.StdEncoding.DecodeString(metadata[:24])
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	nonce, err := base64.StdEncoding.DecodeString(metadata[24:48])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//tag, err := base64.StdEncoding.DecodeString(metadata[48:])
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//ciphertest, err := base64.StdEncoding.DecodeString(base[72:])
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	key := []byte("qrvkuo6t4wG5mqGGVPpepPHKCr4Uznu1ng9QOQJur8s"[:32])

	//ciphertext := encrypt(plaintext, key, nonce)
	//fmt.Printf("Ciphertext: %x\n", ciphertext)

	recoveredPt := decrypt([]byte(metadata), key, nonce)
	fmt.Printf("Recovered plaintext: %s\n", recoveredPt)
}

func TestGCM1(t *testing.T) {
	fmt.Println("AES encryption with GCM")
	plaintext := []byte("加密")
	key := []byte("secretkey32bytessecretkey32bytes")
	nonce := make([]byte, 12)

	ciphertext := encrypt(plaintext, key, nonce)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	recoveredPt := decrypt(ciphertext, key, nonce)
	fmt.Printf("Recovered plaintext: %s\n", recoveredPt)
}

func encrypt(plaintext, key, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return ciphertext
}

func decrypt(ciphertext, key, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
