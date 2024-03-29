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

package payload

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/insolar/insolar/insolar/bus/meta"
)

func NewMessage(pl Payload) (*message.Message, error) {
	buf, err := Marshal(pl)
	if err != nil {
		return nil, err
	}
	return message.NewMessage(watermill.NewUUID(), buf), nil
}

func MustNewMessage(pl Payload) *message.Message {
	msg, err := NewMessage(pl)
	if err != nil {
		panic(err)
	}
	return msg
}

func NewResultMessage(pl Payload) (*message.Message, error) {
	msg, err := NewMessage(pl)
	if err != nil {
		return nil, err
	}
	msg.Metadata.Set(meta.Type, meta.TypeReturnResults)
	return msg, nil
}
