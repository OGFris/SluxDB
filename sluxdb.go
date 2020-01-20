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

package SluxDB

import (
	"github.com/OGFris/SluxDB/routes"
	"github.com/OGFris/SluxDB/storage"
	"github.com/OGFris/SluxDB/utils"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	Password string
	Storage  *storage.Storage
)

func Start(port string) {
	if _, err := os.Stat("./sluxdb_config.yaml"); err == os.ErrNotExist {
		Password = utils.GenerateToken()
		f, err := os.Create("./sluxdb_config.yaml")
		utils.PanicErr(err)

		bytes, err := yaml.Marshal(struct {
			Password string
		}{
			Password: Password,
		})
		utils.PanicErr(err)

		_, err = f.Write(bytes)
		utils.PanicErr(err)

		err = f.Close()
		utils.PanicErr(err)
	} else {
		bytes, err := ioutil.ReadFile("./sluxdb_config.yaml")
		utils.PanicErr(err)

		out := struct {
			Password string
		}{}

		err = yaml.Unmarshal(bytes, out)
		utils.PanicErr(err)

		Password = out.Password
	}

	var err error
	Storage, err = storage.NewStorage("./data.db")
	utils.PanicErr(err)

	router := mux.NewRouter()

	router.HandleFunc("/query", routes.Query).Methods("POST")
	router.HandleFunc("/query/put", routes.QueryPut).Methods("POST")
	router.HandleFunc("/query/batch", routes.QueryBatch).Methods("POST")
	router.HandleFunc("/bucket", routes.Bucket).Methods("POST")

	log.Fatalln(http.ListenAndServe(":"+port, router))
}
