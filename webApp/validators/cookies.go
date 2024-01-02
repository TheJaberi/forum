package forum

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
)

var (
	SecretKey       []byte
	ErrValueTooLong = errors.New("Cookie Value Too Long")
	ErrInvalidValue = errors.New("Cookie Value Invalid")
)

func SecretKeyInit() {
	var err error
	/*
	   Decode the random 64-character hex string to give us a slice containing
	   32 random bytes. For simplicity, I've hardcoded this hex string but in a
	   real application you should read it in at runtime from a command-line
	   flag or environment variable.
	*/
	SecretKey, err = hex.DecodeString("13d6b4dff8f84a10851021ec8608f814570d562c92fe6b5ec4c9f595bcb3234b")
	if err != nil {
		log.Fatalf("%d", err)
	}
}

func WriteSignedCookie(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)
	cookie.Value = string(signature) + cookie.Value
	return WriteCookie(w, cookie)
}

func WriteCookie(w http.ResponseWriter, cookie http.Cookie) error {
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))
	if len(cookie.String()) > 4096 { // web browsers maximum size limit for cookies
		return ErrValueTooLong
	}
	http.SetCookie(w, &cookie)
	return nil
}

func ReadSignedCookie(r *http.Request, name string, secretKey []byte) (string, error) {
	signedValue, err := ReadCookie(r, name)
	if err != nil {
		return "", err
	}
	if len(signedValue) < sha256.Size {
		return "", ErrInvalidValue
	}
	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	expectedSignature := mac.Sum(nil)

	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}
	return value, nil
}

func ReadCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}
	value, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", ErrInvalidValue
	}
	return string(value), nil
}
