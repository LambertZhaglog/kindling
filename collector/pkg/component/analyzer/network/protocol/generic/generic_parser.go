package generic

import (
	"github.com/Kindling-project/kindling/collector/pkg/component/analyzer/network/protocol"
	"github.com/Kindling-project/kindling/collector/pkg/model"
)

var empty_attributes = model.NewAttributeMap()

func NewGenericParser() *protocol.ProtocolParser {
	return protocol.NewProtocolParser(protocol.NOSUPPORT, false, parseHead, parsePayload)
}

func parseHead(data []byte, size int64, isRequest bool) (attributes protocol.ProtocolMessage) {
	return NewGenericAttributes(data, size, isRequest)
}

func parsePayload(attributes protocol.ProtocolMessage, isRequest bool) (ok bool) {
	return true
}

type GenericAttributes struct {
	*protocol.PayloadMessage
	request bool
}

func NewGenericAttributes(data []byte, size int64, isRequest bool) *GenericAttributes {
	return &GenericAttributes{
		PayloadMessage: protocol.NewPayloadMessage(data, 0, size, isRequest),
		request:        isRequest,
	}
}

func (generic *GenericAttributes) GetProtocol() string {
	return protocol.NOSUPPORT
}

func (generic *GenericAttributes) MergeRequest(request protocol.ProtocolMessage) bool {
	return true
}

func (generic *GenericAttributes) GetAttributes() *model.AttributeMap {
	return empty_attributes
}
