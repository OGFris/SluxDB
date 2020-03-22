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

package storage

import (
	"fmt"
	"github.com/OGFris/SluxDB/utils"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

type Storage struct {
	Engine *bolt.DB
	Local  map[string]map[string]int
}

func NewStorage(path string) (*Storage, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return &Storage{}, err
	}

	s := Storage{Engine: db}
	utils.PanicErr(s.parseAll())

	return &s, nil
}

func (s *Storage) CreateBucket(bucket string) error {
	return s.Engine.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err == nil {
			s.Local[bucket] = make(map[string]int)
		}

		return err
	})
}

func (s *Storage) DeleteBucket(bucket string) error {
	return s.Engine.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucket))
		if err == nil {
			delete(s.Local, bucket)
		}

		return err
	})
}

func (s *Storage) PutKey(bucket, key string, value int) error {
	return s.Engine.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(fmt.Sprint(value)))
		if err == nil {
			s.Local[bucket][key] = value
		}

		return err
	})
}

func (s *Storage) DeleteKey(bucket, key string) error {
	return s.Engine.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Delete([]byte(key))
		if err == nil {
			delete(s.Local[bucket], key)
		}

		return err
	})
}

func (s *Storage) parseAll() error {
	s.Local = make(map[string]map[string]int)
	return s.Engine.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, b *bolt.Bucket) error {
			s.Local[string(name)] = make(map[string]int)
			return s.Engine.View(func(tx *bolt.Tx) error {
				b := tx.Bucket(name)
				return b.ForEach(func(k, v []byte) error {
					n, err := strconv.Atoi(string(v))
					if err != nil {
						return err
					}
					s.Local[string(name)][string(k)] = n

					return nil
				})
			})
		})
	})
}
