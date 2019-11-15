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
	"sync"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
)

type descriptorsCache struct {
	Client Client `inject:""`

	codeCache  cache
	protoCache cache
	pulseCache cache
}

func NewDescriptorsCache() DescriptorsCache {
	return &descriptorsCache{
		codeCache:  newSingleFlightCache(),
		protoCache: newSingleFlightCache(),
		pulseCache: newSingleFlightCache(),
	}
}

func (c *descriptorsCache) ByPrototypeRef(
	ctx context.Context, protoRef insolar.Reference,
) (
	PrototypeDescriptor, CodeDescriptor, error,
) {
	protoDesc, err := c.GetPrototype(ctx, protoRef)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't get prototype descriptor")
	}

	codeRef := protoDesc.Code()
	codeDesc, err := c.GetCode(ctx, *codeRef)
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't get code descriptor")
	}

	return protoDesc, codeDesc, nil
}

func (c *descriptorsCache) ByObjectDescriptor(
	ctx context.Context, obj ObjectDescriptor,
) (
	PrototypeDescriptor, CodeDescriptor, error,
) {
	protoRef, err := obj.Prototype()
	if err != nil {
		return nil, nil, errors.Wrap(err, "couldn't get prototype reference")
	}

	if protoRef == nil {
		return nil, nil, errors.New("Empty prototype")
	}

	return c.ByPrototypeRef(ctx, *protoRef)
}

func (c *descriptorsCache) GetPrototype(
	ctx context.Context, ref insolar.Reference,
) (
	PrototypeDescriptor, error,
) {
	res, err := c.protoCache.get(ref, func() (interface{}, error) {
		return c.Client.GetPrototype(ctx, ref)
	})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get object")
	}

	return res.(PrototypeDescriptor), nil
}

func (c *descriptorsCache) GetCode(
	ctx context.Context, ref insolar.Reference,
) (
	CodeDescriptor, error,
) {
	res, err := c.codeCache.get(ref, func() (interface{}, error) {
		return c.Client.GetCode(ctx, ref)
	})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get code")
	}
	return res.(CodeDescriptor), nil
}

func (c *descriptorsCache) GetPulseForRequest(ctx context.Context, request insolar.Reference) (PulseDescriptor, error) {
	pn := request.GetLocal().Pulse()

	res, err := c.pulseCache.get(pn, func() (interface{}, error) {
		return c.Client.GetPulseForRequest(ctx, request)
	})
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't get pulse for request: %s", request)
	}

	return res.(PulseDescriptor), nil
}

//go:generate minimock -i github.com/insolar/insolar/logicrunner/artifacts.cache -o ./ -s _mock.go -g

type cache interface {
	get(key interface{}, getter func() (val interface{}, err error)) (val interface{}, err error)
}

type cacheEntry struct {
	mu    sync.Mutex
	value interface{}
}

type singleFlightCache struct {
	mu    sync.Mutex
	cache *lru.Cache
}

func newSingleFlightCache() cache {
	cache, err := lru.New(100)

	if err != nil {
		panic("failed to init cache")
	}

	return &singleFlightCache{cache: cache}
}

func (c *singleFlightCache) getEntry(key interface{}) *cacheEntry {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cache.Get(key); !ok {
		c.cache.Add(key, &cacheEntry{})
	}

	entry, _ := c.cache.Get(key)
	cEntry := entry.(*cacheEntry)

	return cEntry
}

func (c *singleFlightCache) get(
	key interface{},
	getter func() (value interface{}, err error),
) (
	interface{}, error,
) {
	e := c.getEntry(key)

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.value != nil {
		return e.value, nil
	}

	val, err := getter()
	if err != nil {
		return val, err
	}

	e.value = val
	return e.value, nil
}
