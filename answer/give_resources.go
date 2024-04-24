package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func GiveResources(buffer *[]byte, client *connection.Client) (int, int, error) {
	// We won't verify the packet, just send whatever the client asked for
	var data protobuf.CS_11013
	err := proto.Unmarshal(*buffer, &data)
	if err != nil {
		return 0, 11014, err
	}

	var number uint32
	var fieldId uint32
	switch data.GetType() {
	case 1: // Number of resource #7
		number = client.Commander.GetResourceCount(7)
		fieldId = 7
	case 2: // Number of resource #5
		number = client.Commander.GetResourceCount(5)
		fieldId = 5
	}

	res := proto.Uint32(0)

	// Add the requested resource
	if err := client.Commander.AddResource(data.GetType(), number); err != nil {
		res = proto.Uint32(1)
	}
	// Remove the field resource
	if err := client.Commander.SetResource(fieldId, 0); err != nil {
		res = proto.Uint32(1)
	}
	// Send the response
	response := protobuf.SC_11014{
		Result: res,
	}
	if number == 0 {
		response.Result = proto.Uint32(1)
	}
	return client.SendMessage(11014, &response)
}
