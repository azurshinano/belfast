package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func SendMailList(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_30003

	response.MailList = make([]*protobuf.MAIL_SIGLE, len(client.Commander.Mails))
	for i, mail := range client.Commander.Mails {
		response.MailList[i] = &protobuf.MAIL_SIGLE{
			Id:         proto.Uint32(mail.ID),
			ReadFlag:   proto.Uint32(boolToUint32(mail.Read) + boolToUint32(mail.AttachmentsCollected)),
			Date:       proto.Uint32(uint32(mail.Date.Unix())),
			AttachFlag: proto.Uint32(boolToUint32(len(mail.Attachments) > 0) + boolToUint32(mail.AttachmentsCollected)),
			ImpFlag:    proto.Uint32(boolToUint32(mail.IsImportant)),
		}
		if mail.CustomSender != nil {
			response.MailList[i].Title = proto.String(mail.Title + "||" + *mail.CustomSender)
		} else {
			response.MailList[i].Title = proto.String(mail.Title)
		}
		if len(mail.Attachments) > 0 { // only if the mail has attachments, otherwise we set as nil (default)
			response.MailList[i].AttachmentList = make([]*protobuf.ATTACHMENT, len(mail.Attachments))
			for j, attachment := range mail.Attachments {
				response.MailList[i].AttachmentList[j] = &protobuf.ATTACHMENT{
					Type:   proto.Uint32(attachment.Type),
					Id:     proto.Uint32(attachment.ItemID),
					Number: proto.Uint32(0),
				}
			}
		}
	}
	return client.SendMessage(30003, &response)
}
