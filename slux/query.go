// MIT License
//
// Copyright (c) 2019 Ilyes Cherfaoui
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

package slux

import (
	"fmt"
	"github.com/OGFris/SluxDB/utils"
	"net/http"
)

func Query(w http.ResponseWriter, r *http.Request) {
	password := r.PostFormValue("password")
	if Password != password {
		utils.WriteErr(w, "Wrong password", http.StatusUnauthorized)

		return
	}
	bucket := r.PostFormValue("bucket")
	key := r.PostFormValue("key")
	if bucket == "" {
		utils.WriteErr(w, "Couldn't be found", http.StatusNotFound)

		return
	}

	if key == "" {
		// returns all
		if v, exist := Storage.Local[bucket]; exist {
			utils.WriteJson(w, v)
		} else {
			utils.WriteErr(w, "Couldn't be found", http.StatusNotFound)

			return
		}
	}

	if _, exist := Storage.Local[bucket]; exist {
		if v, exist := Storage.Local[bucket][key]; exist {
			_, err := w.Write([]byte(fmt.Sprint(v)))
			utils.PanicErr(err)
		} else {
			utils.WriteErr(w, "Couldn't be found", http.StatusNotFound)
		}
	} else {
		utils.WriteErr(w, "Bucket couldn't be found", http.StatusNotFound)
	}
}
