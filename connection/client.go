package connection

import (
	"bytes"
	"fmt"
	"github.com/ggmolly/belfast/logger"
	"github.com/ggmolly/belfast/orm"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"net"
)

type Client struct {
	IP          net.IP
	Port        int
	State       int
	PacketIndex int
	Hash        uint32
	Connection  *net.Conn
	Commander   *orm.Commander
	Buffer      bytes.Buffer
	Server      *Server
}

func (client *Client) CreateCommander(uname string) (uint32, error) {
	guser := orm.GatewayUser{
		Username: uname,
	}
	if err := orm.GormDB.Create(&guser).Error; err != nil {
		logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("failed to create account: %v", err), logger.LOG_LEVEL_ERROR)
		return 0, err
	}
	// Create a new commander for the account
	commander := orm.Commander{
		AccountID: guser.AccountID,
		Name:      fmt.Sprintf("Unnamed commander #%d", guser.AccountID),
	}
	if err := orm.GormDB.Create(&commander).Error; err != nil {
		logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("failed to create commander for account %d: %v", guser.AccountID, err), logger.LOG_LEVEL_ERROR)
		return 0, err
	}

	// Since we have no tutorial / first login, we'll also give a secretary to the new commander
	if err := orm.GormDB.Create(&orm.OwnedShip{
		OwnerID:           commander.CommanderID,
		ShipID:            202124, // Belfast (6 stars)
		IsSecretary:       true,
		SecretaryPosition: proto.Uint32(0),
	}).Error; err != nil {
		logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("failed to give Belfast to account %d: %v", guser.AccountID, err), logger.LOG_LEVEL_ERROR)
		return 0, err
	}

	// Give default items to commander
	if err := orm.GormDB.Create(&([]orm.CommanderItem{{
		// Wisdom Cube
		CommanderID: commander.CommanderID,
		ItemID:      20001,
		Count:       0,
	}})).Error; err != nil {
		logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("failed to give default items to account %d: %v", guser.AccountID, err), logger.LOG_LEVEL_ERROR)
		return 0, err
	}
	// Give default resources to commander
	if err := orm.GormDB.Create(&([]orm.OwnedResource{{
		// Gold
		CommanderID: commander.CommanderID,
		ResourceID:  1,
		Amount:      0,
	}, {
		// Oil
		CommanderID: commander.CommanderID,
		ResourceID:  2,
		Amount:      0,
	}, {
		// Gem
		CommanderID: commander.CommanderID,
		ResourceID:  4,
		Amount:      0,
	}})).Error; err != nil {
		logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("failed to give default resources to account %d: %v", guser.AccountID, err), logger.LOG_LEVEL_ERROR)
		return 0, err
	}

	logger.LogEvent("Client", "CreateCommander", fmt.Sprintf("created new commander for account %d", guser.AccountID), logger.LOG_LEVEL_INFO)
	return guser.AccountID, nil
}

func (client *Client) GetCommander(accountId uint32) error {
	err := orm.GormDB.Where("account_id = ?", accountId).First(&client.Commander).Error
	return err
}

// Sends SC_10999 (disconnected from server) message to the Client, reasons are defined in consts/disconnect_reasons.go
func (client *Client) Disconnect(reason uint8) error {
	_, _, err := SendProtoMessage(10999, client, &protobuf.SC_10999{
		Reason: proto.Uint32(uint32(reason)),
	})
	return err
}

// Sends the content of the buffer to the client via TCP
func (client *Client) Flush() {
	_, err := (*client.Connection).Write(client.Buffer.Bytes())
	if err != nil {
		logger.LogEvent("Client", "Flush()", fmt.Sprintf("%s:%d -> %v", client.IP, client.Port, err), logger.LOG_LEVEL_ERROR)
	}
	// logger.LogEvent("Client", "Flush()", fmt.Sprintf("%s:%d -> %d bytes flushed", client.IP, client.Port, n), logger.LOG_LEVEL_INFO)
	client.Buffer.Reset()
}

func (client *Client) SendMessage(packetId int, message any) (int, int, error) {
	return SendProtoMessage(packetId, client, message)
}
