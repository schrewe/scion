// Copyright 2019 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"github.com/scionproto/scion/go/lib/prom"
)

const (

	// Namespace is the metrics namespace for the beacon server.
	Namespace = "bs"

	// DstBR indicates the destination to be Border Router.
	DstBR string = "br"
	// DstPS indicates the destination to be Path Server.
	DstPS string = "ps"

	// ErrProcess indicates an error during processing.
	ErrProcess string = prom.ErrProcess

	// RevNew indicates a new issued revocation.
	RevNew string = "new"
	// RevRenew indicates a renew of an already issued revocation.
	RevRenew string = "renew"
	// RevFromCtrl indicates that revocation was sent control payload.
	RevFromCtrl string = "ctrl"

	// Success indicates a successful result.
	Success string = prom.Success
)

var (
	// Ifstate is the single-instance struct to get prometheus metrics or counters.
	Ifstate = newIfstate()
	// Keepalive is the single-instance struct to get keepalive prometheus counters.
	Keepalive = newKeepalive()
	// Revocation is the single-instance struct to get prometheus counters.
	Revocation = newRevocation()
)
