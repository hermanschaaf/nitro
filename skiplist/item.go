// Copyright (c) 2016 Couchbase, Inc.
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
// except in compliance with the License. You may obtain a copy of the License at
//   http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the
// License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing permissions
// and limitations under the License.

package skiplist

import (
	"bytes"
	"fmt"
	"unsafe"
)

var (
	minItem unsafe.Pointer
	maxItem = unsafe.Pointer(^uintptr(0))
)

func compare(cmp CompareFn, this, that unsafe.Pointer) int {
	if this == minItem || that == maxItem {
		return -1
	}

	if this == maxItem || that == minItem {
		return 1
	}

	return cmp(this, that)
}

type byteKeyItem []byte

func (itm *byteKeyItem) String() string {
	return string(*itm)
}

func (itm byteKeyItem) Size() int {
	return len(itm)
}

// NewByteKeyItem creates a new item from bytes
func NewByteKeyItem(k []byte) unsafe.Pointer {
	itm := byteKeyItem(k)
	return unsafe.Pointer(&itm)
}

// CompareBytes is a byte item comparator
func CompareBytes(this, that unsafe.Pointer) int {
	thisItem := (*byteKeyItem)(this)
	thatItem := (*byteKeyItem)(that)
	return bytes.Compare([]byte(*thisItem), []byte(*thatItem))
}

type intKeyItem int

func (itm *intKeyItem) String() string {
	return fmt.Sprint(*itm)
}

func (itm intKeyItem) Size() int {
	return int(unsafe.Sizeof(itm))
}

// CompareInt is a helper integer item comparator
func CompareInt(this, that unsafe.Pointer) int {
	thisItem := (*intKeyItem)(this)
	thatItem := (*intKeyItem)(that)
	return int(*thisItem - *thatItem)
}
