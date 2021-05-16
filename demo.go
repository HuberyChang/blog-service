/*
* @Author: HuberyChang
* @Date: 2021/5/15 15:52
 */

package main

import (
	"encoding/base64"
	"log"
)

func main() {
	payload, _ := base64.StdEncoding.DecodeString("eyJhcHBfa2V5IjoiMjc1NjY4YmE2NTUwNDljZDczOWQxZDllNmIzMWNjZjEiLCJhcHBfc2VjcmV0IjoiYmY1ZmIxMTM5ZjQ0Zjc1YmZkYmE2OGIxM2I5YjIxNTYiLCJleHAiOjE2MjEwNzIzNDgsImlzcyI6ImJsb2dfc2VydmljZSJ9")
	log.Println(string(payload))
}
