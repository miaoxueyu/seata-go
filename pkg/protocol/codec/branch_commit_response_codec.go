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
	"math"

	"github.com/seata/seata-go/pkg/common/bytes"
	serror "github.com/seata/seata-go/pkg/common/error"
	"github.com/seata/seata-go/pkg/protocol/branch"
	"github.com/seata/seata-go/pkg/protocol/message"
)

func init() {
	GetCodecManager().RegisterCodec(CodecTypeSeata, &BranchCommitResponseCodec{})
}

type BranchCommitResponseCodec struct {
}

func (g *BranchCommitResponseCodec) Decode(in []byte) interface{} {
	data := message.BranchCommitResponse{}
	buf := bytes.NewByteBuffer(in)

	data.ResultCode = message.ResultCode(bytes.ReadByte(buf))
	if data.ResultCode == message.ResultCodeFailed {
		data.Msg = bytes.ReadString8Length(buf)
	}
	data.TransactionExceptionCode = serror.TransactionExceptionCode(bytes.ReadByte(buf))
	data.Xid = bytes.ReadString16Length(buf)
	data.BranchId = int64(bytes.ReadUInt64(buf))
	data.BranchStatus = branch.BranchStatus(bytes.ReadByte(buf))

	return data
}

func (g *BranchCommitResponseCodec) Encode(in interface{}) []byte {
	data, _ := in.(message.BranchCommitResponse)
	buf := bytes.NewByteBuffer([]byte{})

	buf.WriteByte(byte(data.ResultCode))
	if data.ResultCode == message.ResultCodeFailed {
		msg := data.Msg
		if len(data.Msg) > math.MaxInt8 {
			msg = data.Msg[:math.MaxInt8]
		}
		bytes.WriteString8Length(msg, buf)
	}
	buf.WriteByte(byte(data.TransactionExceptionCode))
	bytes.WriteString16Length(data.Xid, buf)
	buf.WriteInt64(data.BranchId)
	buf.WriteByte(byte(data.BranchStatus))

	return buf.Bytes()
}

func (g *BranchCommitResponseCodec) GetMessageType() message.MessageType {
	return message.MessageTypeBranchCommitResult
}
