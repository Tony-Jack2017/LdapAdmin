package util

import (
	"LdapAdmin/config"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"runtime"
)

//
// GenerateRsaKey
// @Description: This func is for generating RSA keys
// @param keySize: The private key size
// @param dirPath: The path to store the generated key
// @return error
//
func GenerateRsaKey(keySize int, keyDirPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, err.Error()))
	}
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}

	if !judgeFolderExist(keyDirPath) {
		if errCreate := os.MkdirAll(keyDirPath, os.ModePerm); errCreate != nil {
			_, file, line, _ := runtime.Caller(0)
			return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errCreate.Error()))
		}
	}

	if judgeFolderExist(fmt.Sprintf("%s/%s", keyDirPath, "private.pem")) {
		if errRemove := os.Remove(fmt.Sprintf("%s/%s", keyDirPath, "private.pem")); errRemove != nil {
			_, file, line, _ := runtime.Caller(0)
			return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errRemove.Error()))
		}
	}

	keyFile, errFile := os.Create(fmt.Sprintf("%s/%s", keyDirPath, "private.pem"))
	defer keyFile.Close()
	if errFile != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errFile.Error()))
	}
	errEncode := pem.Encode(keyFile, block)
	if errEncode != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errEncode.Error()))
	}
	publicKey := privateKey.PublicKey
	derStream, errStream := x509.MarshalPKIXPublicKey(&publicKey)
	if errStream != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errStream.Error()))
	}
	block = &pem.Block{
		Type:  "rsa public key",
		Bytes: derStream,
	}
	if judgeFolderExist(fmt.Sprintf("%s/%s", keyDirPath, "public.pem")) {
		if errRemove := os.Remove(fmt.Sprintf("%s/%s", keyDirPath, "private.pem")); errRemove != nil {
			_, file, line, _ := runtime.Caller(0)
			return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errRemove.Error()))
		}
	}
	keyFile, errFile = os.Create(fmt.Sprintf("%s/%s", keyDirPath, "public.pem"))
	if errFile != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errFile.Error()))
	}
	errEncode = pem.Encode(keyFile, block)
	if errEncode != nil {
		_, file, line, _ := runtime.Caller(0)
		return errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errEncode.Error()))
	}
	PrintlnSuccess("Generate rsa keys success >>>>")
	return nil
}

//
// DecryptPassword
// @Description: This func is for parsing the password that was encrypted by rsa public key
// @param password: The password
// @return []byte: The password's plaintext
// @return error
//
func DecryptPassword(password string) ([]byte, error) {
	block, err := readKeyFromFile(config.Conf.System.RsaKeyFolder + "private.pem")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, err.Error()))
	}
	privateKey, errParse := x509.ParsePKCS1PrivateKey(block)
	if errParse != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errParse.Error()))
	}
	defer func() {
		if err1 := recover(); err1 != nil {
			_, file, line, _ := runtime.Caller(0)
			err = errors.New(fmt.Sprintf("%s:%d\n%v", file, line, "rsa private key is nil"))
		}
	}()

	plaintext, errDecrypt := rsa.DecryptPKCS1v15(rand.Reader, privateKey, []byte(password))
	if errDecrypt != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, errors.New(fmt.Sprintf("%s:%d\n%v", file, line+1, errDecrypt.Error()))
	}
	return plaintext, err

}

//
// readKeyFromFile
// @Description: This func is for reading data from a file
// @param filename: The filename where the key is stored
// @return []byte: The data by reading from the file
//
func readKeyFromFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	var b []byte

	if err != nil {
		return b, err
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b = make([]byte, fileInfo.Size())
	_, err = f.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}

//
//  judgeFolderExist
//  @Description: This func is for checking if a folder exists
//  @param path: The path to the folder
//  @return bool
//
func judgeFolderExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
