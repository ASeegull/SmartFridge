package protoStruct

import (
	"github.com/golang/protobuf/proto"
)

//SetParameters sets parameter to struct and executes token
func (str *Setup) SetParameters(agentID, userID, productID string, heartbeat int, time string) {
	str.UserID = userID
	str.ProductID = productID
	str.Heartbeat = int32(heartbeat)
	str.StateExpires = time
}

//MarshalStruct marshals this struct
func (str *Agentstate) MarshalStruct() ([]byte, error) {
	return proto.Marshal(str)
}

//UnmarshalToStruct unmarshals to struct
func (str *Agentstate) UnmarshalToStruct(data []byte) error {
	return proto.Unmarshal(data, str)
}

//MarshalStruct marshals this struct
func (str *Setup) MarshalStruct() ([]byte, error) {
	return proto.Marshal(str)
}

//UnmarshalToStruct unmarshals to struct
func (str *Setup) UnmarshalToStruct(data []byte) error {
	return proto.Unmarshal(data, str)
}

//UnmarshalToStruct unmarshals to struct
func (str *Request) UnmarshalToStruct(data []byte) error {
	return proto.Unmarshal(data, str)
}
