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

package main

import (
	"bufio"
	"github.com/OGFris/SluxDB/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetPrefix("[SluxDB] ")

	if _, err := os.Stat("./sluxdb_config.json"); err == os.ErrNotExist {
		// File doesn't exist so it'll start a new setup.
		scanner := bufio.NewScanner(os.Stdin)
		log.Println("Couldn't find a config file, creating a new one, please follow the setup instructions.")
		log.Print("Please set the username you want to use: ")
		scanner.Scan()
		username := scanner.Text()
		if len(username) == 0 {
			log.Fatalln("Username can't be empty, set it to default username.")
			username = "root"
		} else {
			log.Println() // to jump to the next line.
		}

		log.Print("Please set the password you want to use: ")
		scanner.Scan()
		password := scanner.Text()
		if len(password) == 0 {
			password = utils.GenerateToken()
			log.Fatalln("Password can't be empty, generated random passwword (Save it somewhere):", password)
		} else {
			log.Println()
		}

		hash, err := utils.Encrypt(password)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("./sluxdb_config.json")
		if err != nil {
			panic(err)
		}

		bytes, err := yaml.Marshal(struct {
			Username string
			Password string
		}{
			Username: username,
			Password: hash,
		})
		if err != nil {
			panic(err)
		}

		_, err = f.Write(bytes)
		if err != nil {
			panic(err)
		}

		err = f.Close()
		if err != nil {
			panic(err)
		}
	} else {
		bytes, err := ioutil.ReadFile("./sluxdb_config.json")
		if err != nil {
			panic(err)
		}

		out := struct {
			Username string
			Password string
		}{}

		err = yaml.Unmarshal(bytes, out)
		if err != nil {
			panic(err)
		}

		// TODO: Start the server with the out config.
	}
}
