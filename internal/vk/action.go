package vk

import (
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
		return err
	}
	return nil
}

func (vk *VK) actionBtnBackHome(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Home")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Keyboard(vk.btnHome())

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnHomeHello(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Hello menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Keyboard(vk.btnHomeHello())

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnHomeLink(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Link menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Keyboard(vk.btnHomeVKDevLink())

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnHomeLocate(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Locate menu")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Keyboard(vk.btnHomeLocation())

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnHomeSnak(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Snak menu")
	b.RandomID(0)
	b.PeerID(obj.PeerID)
	b.Keyboard(vk.btnHomeSnak())

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionBtnSayHello(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	b.Message("Hello Subscriber!")
	b.RandomID(0)
	b.PeerID(obj.UserID)

	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	return nil
}

func (vk *VK) actionShowSnak(obj events.MessageEventObject) error {
	b := params.NewMessagesSendMessageEventAnswerBuilder()
	b.PeerID(obj.UserID)
	b.UserID(obj.UserID)
	b.EventID(string(obj.EventID))
	b.EventData("{\n\"type\":\"show_snackbar\",\n\"text\":\"\nHi I am popup message!\"\n}")

	_, err := vk.VKApi.MessagesSendMessageEventAnswer(b.Params)
	if err != nil {
		return err
	}
	return nil
}
