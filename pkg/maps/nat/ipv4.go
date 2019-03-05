// Copyright 2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nat

import (
	"fmt"
	"unsafe"

	"github.com/cilium/cilium/common/types"
	"github.com/cilium/cilium/pkg/byteorder"
	"github.com/cilium/cilium/pkg/tuple"
)

// NatEntry represents an entry in the NAT table.
type NatEntry4 struct {
	Created uint64     `align:"created"`
	Addr    types.IPv4 `align:"to_saddr"`
	Port    uint16     `align:"to_sport"`
}

// GetValuePtr returns the unsafe.Pointer for n.
func (n *NatEntry4) GetValuePtr() unsafe.Pointer { return unsafe.Pointer(n) }

// String returns the readable format.
func (n *NatEntry4) String() string {
	return fmt.Sprintf("Addr=%s Port=%d Created=%d\n",
		n.Addr,
		n.Port,
		n.Created)
}

// Dumps nat entry to string.
func (n *NatEntry4) Dump(key tuple.TupleKey, start uint64) string {
	var which string

	if key.GetFlags()&tuple.TUPLE_F_IN != 0 {
		which = "DST"
	} else {
		which = "SRC"
	}
	return fmt.Sprintf("XLATE_%s %s:%d Created=%s\n",
		which,
		n.Addr,
		n.Port,
		NatDumpCreated(start, n.Created))
}

// ToHost converts NatEntry4 ports to host byte order.
func (n *NatEntry4) ToHost() NatEntry {
	x := *n
	x.Port = byteorder.NetworkToHost(n.Port).(uint16)
	return &x
}
