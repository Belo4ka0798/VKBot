package vk

import (
	"log"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
)

func (vk *VK) btnAnswerEvent(obj events.MessageEventObject) error {
	b := params.NewMessagesSendMessageEventAnswerBuilder()
	b.EventID(obj.EventID)
	b.UserID(obj.UserID)
	b.PeerID(obj.PeerID)
	_, err := vk.VKApi.MessagesSendMessageEventAnswer(b.Params)
	if err != nil {
		log.Printf("action button call1 error: %v", err)
		return err
	}
	return nil
}

func (vk *VK) actionBtnCallHome(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	vk.btnHome(b)
	b.Message("Home")
	b.RandomID(0)
	b.PeerID(obj.UserID)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnCall1(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	vk.btnHomeBtn1(b)
	b.Message("Hello menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Params.Confirm(true)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnCall2(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	vk.btnHomeBtn2(b)
	b.Message("Link menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Params.Confirm(true)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnCall3(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	vk.btnHomeBtn3(b)
	b.Message("Locate menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Params.Confirm(true)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnCall4(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	//vk.btnCall4(b)

	b.Message("Snak menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Params.Confirm(true)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionSayHello(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Hello VK!")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Params.Confirm(true)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}
