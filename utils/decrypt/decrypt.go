package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
)

func DecryptAES(key string, text string) (string, error) {
	ciphertext, _ := hex.DecodeString(text)

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s, nil
}

func DecryptTripleDES(triplekey string, text string) (string, error) {
	plaintext := []byte("Hello Wo") // Hello Wo = 8 bytes.

	block, err := des.NewTripleDESCipher([]byte(triplekey))

	if err != nil {
		return "", err
	}

	ciphertext := []byte(text)
	iv := ciphertext[:des.BlockSize] // const BlockSize = 8

	//decrypt
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(plaintext))
	decrypter.CryptBlocks(decrypted, []byte(text))
	return string(decrypted), nil
}
