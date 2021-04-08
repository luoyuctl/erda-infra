// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"crypto/md5"
	"encoding/base64"
	"math/rand"
	"unsafe"

	uuid "github.com/satori/go.uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
	APPLICATION_CONFIG_PATH = "conf/application.yml"
	TEMPLATE_JSON_PATH      = "conf/template.json"
	IPDB_PATH               = "conf/ipdata.dat"
	TA_JS_PATH              = "static/ta.js"
	TRANSLATE_PATH          = "conf/translate.yml"
)

func UUID() (string, error) {
	u := uuid.NewV4()
	return u.String(), nil
}

func InString(s string, ss []string) bool {
	for _, item := range ss {
		if s == item {
			return true
		}
	}
	return false
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetMD5Base64(bytes []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	md5Value := md5Ctx.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(md5Value)
}

func GetMD5Base64WithLegth(bytes []byte, maxLength int) string {
	res := GetMD5Base64(bytes)
	if len(res) > maxLength {
		res = res[:maxLength]
	}
	return res
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
