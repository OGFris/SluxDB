// MIT License
//
// Copyright (c) 2019 Fris
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"testing"
	"time"
)

// Encrypt encrypts passwords and return the hash with an error if there's one.
func Encrypt(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passwordBytes), err
}

// Compare checks if a hashed password matches the password provided.
func Compare(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// GenerateToken returns a 16 chars long string of random characters.
func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	tokenBytes := make([]rune, 16)
	for i := range tokenBytes {
		tokenBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(tokenBytes)
}

// AssertEqual checks if the two given values are the same otherwise it fails the test
func AssertEq(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatal(fmt.Sprintf("%v != %v", a, b))
	}
}
