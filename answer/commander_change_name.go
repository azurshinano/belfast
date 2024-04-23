package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/orm"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// TODO: add change name cd time
func CommanderChangeName(buffer *[]byte, client *connection.Client) (int, int, error) {
	var data protobuf.CS_11007
	if err := proto.Unmarshal(*buffer, &data); err != nil {
		return 0, 11008, err
	}
	response := protobuf.SC_11008{
		Result: proto.Uint32(0),
	}

	client.Commander.Name = data.GetName()
	if err := orm.GormDB.Save(&client.Commander).Error; err != nil {
		response.Result = proto.Uint32(1)
	}
	return client.SendMessage(11008, &response)
}
