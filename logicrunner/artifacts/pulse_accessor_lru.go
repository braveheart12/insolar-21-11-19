//
// Copyright 2019 Insolar Technologies GmbH
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
//

package artifacts

import (
	"context"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/pulse"
)

type CachedPulses interface {
	GetPulseForRequest(ctx context.Context, request insolar.Reference) (insolar.Pulse, error)
}

type pulseAccessorLRU struct {
	pulses pulse.Accessor
	client Client
	cache  *lru.Cache
}

func NewPulseAccessorLRU(pulses pulse.Accessor, client Client, size int) *pulseAccessorLRU {
	cache, err := lru.New(size)
	if err != nil {
		panic("failed to init pulse cache")
	}

	return &pulseAccessorLRU{
		pulses: pulses,
		client: client,
		cache:  cache,
	}
}

func (p pulseAccessorLRU) GetPulseForRequest(ctx context.Context, request insolar.Reference) (insolar.Pulse, error) {
	pn := request.GetLocal().Pulse()

	var (
		foundPulse insolar.Pulse
		err        error
	)

	val, ok := p.cache.Get(pn)
	if ok {
		return val.(insolar.Pulse), nil
	}

	foundPulse, err = p.pulses.ForPulseNumber(ctx, pn)
	if err == nil {
		p.cache.Add(pn, foundPulse)
		return foundPulse, nil
	}

	foundPulse, err = p.client.GetPulseForRequest(ctx, request)
	if err == nil {
		p.cache.Add(pn, foundPulse)
		return foundPulse, nil
	}

	return insolar.Pulse{}, errors.Wrapf(err, "couldn't get pulse for request: %s", request)
}
