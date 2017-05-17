/*
 * Minio Cloud Storage, (C) 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// GOMAXPROCS=10 go test

package lsync_test

import (
	"testing"

	. "github.com/minio/lsync"
	"time"
)

func HammerLMutex(m *LMutex, loops int, cdone chan bool) {
	for i := 0; i < loops; i++ {
		m.Lock()
		m.Unlock()
	}
	cdone <- true
}

func TestLMutex(t *testing.T) {
	m := NewLMutex()
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go HammerLMutex(m, 1000, c)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}
