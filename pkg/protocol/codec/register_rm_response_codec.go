/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package codec

import (
	"github.com/seata/seata-go/pkg/protocol/message"
)

func init() {
	GetCodecManager().RegisterCodec(CodecTypeSeata, &RegisterRMResponseCodec{})
}

type RegisterRMResponseCodec struct {
	AbstractIdentifyResponseCodec
}

func (g *RegisterRMResponseCodec) Decode(in []byte) interface{} {
	req := g.AbstractIdentifyResponseCodec.Decode(in)
	abstractIdentifyResponse := req.(message.AbstractIdentifyResponse)
	return message.RegisterRMResponse{
		AbstractIdentifyResponse: abstractIdentifyResponse,
	}
}

func (c *RegisterRMResponseCodec) Encode(in interface{}) []byte {
	resp := in.(message.RegisterRMResponse)
	return c.AbstractIdentifyResponseCodec.Encode(resp.AbstractIdentifyResponse)
}

func (g *RegisterRMResponseCodec) GetMessageType() message.MessageType {
	return message.MessageTypeRegRmResult
}
