package controller

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Sha256Handler(w http.ResponseWriter, r *http.Request) {
	// to get the username, use the following:
	username := mux.Vars(r)["username"]

	sha256string := GetHash("Sha256", username)
	// to calculate sha256 hash of a string, use internal/hashes package and function GetHash

	// to send the response, use the following:
	fmt.Fprint(w, sha256string)

	w.WriteHeader(http.StatusOK)
}

func GithubUsernameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, os.Getenv("GITHUB_USERNAME"))
}

func GetHash(kind, text string) string {
	var (
		h      hash.Hash
		hashed string
	)
	switch kind {
	case "Sha1":
		h = sha1.New()
	case "Sha256":
		h = sha256.New()
	case "Sha384":
		h = sha512.New384()
	case "Sha512":
		h = sha512.New()
	case "Md5":
		h = md5.New()
	default:
		return ""
	}
	h.Write([]byte(text))
	hashed = hex.EncodeToString(h.Sum(nil))
	return hashed
}
