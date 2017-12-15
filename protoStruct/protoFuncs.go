package protoStruct

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"
)

//CheckToken checks tokens conformity
func (str *Agentstate) CheckToken() bool {
	return str.Token == getHash(fmt.Sprintf("%s%s%s", str.AgentID, str.UserID, str.ProductID))
}

//SetParameters sets parameter to struct and executes token
func (str *Setup) SetParameters(agentID, userID, productID string, heartbeat int) {
	str.UserID = userID
	str.ProductID = productID
	str.Heartbeat = int32(heartbeat)
	str.Token = getHash(fmt.Sprintf("%s%s%s", agentID, userID, productID))
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

//MarshalStruct marshals this struct
func (str *Request) MarshalStruct() ([]byte, error) {
	return proto.Marshal(str)
}

//UnmarshalToStruct unmarshals to struct
func (str *Request) UnmarshalToStruct(data []byte) error {
	return proto.Unmarshal(data, str)
}

func getHash(stringToHash string) string {
	hash := md5.Sum([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}
