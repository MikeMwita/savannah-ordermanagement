package service

//
//import (
//	"context"
//	"fmt"
//	"github.com/AndroidStudyOpenSource/africastalking-go/sms"
//	"github.com/MikeMwita/savannah-ordermanagement/internal/core/models"
//)
//
//type AtSms struct {
//	senderInfo models.Resource
//}
//
////SendSMS sends an SMS to the recipient
//
//func (a *AtSms) SendSMS(ctx context.Context, recipient string, message string) (sms.SendMessageResponse, error) {
//
//	//Call the Gateway, and pass the constants here!
//	smsService := sms.NewService(a.senderInfo.Label, a.senderInfo.MetaStringValue0, a.senderInfo.MetaKey0)
//
//	//Send SMS
//	recipients, err := smsService.Send(a.senderInfo.Name, recipient, message) //Leave blank, "", if you don't have one)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(recipients)
//
//	return *recipients, err
//
//}
//
//func NewAtSms(sender models.Resource) *AtSms {
//	return &AtSms{
//		senderInfo: sender,
//	}
//}
