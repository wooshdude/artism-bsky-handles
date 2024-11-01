package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	handle string
	did    string
}

func fetch_did(handle string) string {
	var usr string = fmt.Sprintf("https://%s.bsky.social/.well-known/atproto-did", handle)
	r, _ := http.Get(usr)

	data, _ := io.ReadAll(r.Body)

	return string(data)
}

func get_usr(handle string) User {
	var usr_path string = "accounts.json"
	content, _ := os.ReadFile(usr_path)
	var data map[string]string
	json.Unmarshal(content, &data)

	var user User
	user.handle = handle
	user.did = data[handle]

	return user
}
