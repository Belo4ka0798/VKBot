package vk

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/object"
)

type Payload struct {
	TypeName string `json:"type,omitempty"`
}

func (vk *VK) messageNewObject(ctx context.Context, obj events.MessageNewObject) {
	text := strings.ToLower(obj.Message.Text)

	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	switch text {
	case "/menu":
		b.Message("Hello! Select a menu item for further work")
		vk.btnHome(b)
	case "/reset":
		b.Message("Reset Settings!")
		b.Keyboard(object.NewMessagesKeyboard(false))
	default:
		return
	}

	b.PeerID(obj.Message.PeerID)
	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		log.Fatal(err)
	}
}

func (vk *VK) joinGroupEvent(ctx context.Context, obj events.GroupJoinObject) {
	log.Printf("%d: %s", obj.UserID, obj.JoinType)
	b := params.NewMessagesSendBuilder()
	b.Message("Приветствую тебя в нашей группе!\n" +
		"Home menu - \"/home\"\n" +
		"Help - \"/help\"\n" +
		"Reset - \"/reset\"\n")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	vk.btnHome(b)
	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		log.Printf("cat't send message with join: %v", err)
		return
	}
}

func (vk *VK) messageEventObject(ctx context.Context, obj events.MessageEventObject) {
	err := vk.btnAnswerEvent(obj)
	if err != nil {
		return
	}

	p := &Payload{}
	err = json.Unmarshal(obj.Payload, &p)
	if err != nil {
		log.Printf("can't unmarshal json: %v", err)
		return
	}

	switch p.TypeName {
	case "home_btn_back":
		err = vk.actionBtnCallHome(obj)
		if err != nil {
			log.Printf("action button call3 error: %v", err)
			return
		}
	case "say_hello":
		err = vk.actionSayHello(obj)
		if err != nil {
			log.Printf("action button home error: %v", err)
			return
		}

	case "home_btn1":
		err = vk.actionBtnCall1(obj)
		if err != nil {
			log.Printf("action button call1 error: %v", err)
			return
		}

	case "home_btn2":
		err = vk.actionBtnCall2(obj)
		if err != nil {
			log.Printf("action button call2 error: %v", err)
			return
		}

	case "home_btn3":
		err = vk.actionBtnCall3(obj)
		if err != nil {
			log.Printf("action button call3 error: %v", err)
			return
		}

	case "home_btn4":
		err = vk.actionBtnCall4(obj)
		if err != nil {
			log.Printf("action button call3 error: %v", err)
			return
		}
	}
}
