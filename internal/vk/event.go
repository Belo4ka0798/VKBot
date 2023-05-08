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
		b.Message("Select a menu item for further work")
		b.Keyboard(vk.btnHome())
	case "/reset":
		b.Message("Reset Settings!")
		b.Keyboard(object.NewMessagesKeyboard(false))
	case "/help":
		b.Message("\"Home menu - \"/menu\"\n" +
			"Reset - \"/reset\"\n")
	default:
		b.Keyboard(vk.btnHome())
		return
	}
	b.Keyboard(vk.btnHome())

	b.PeerID(obj.Message.PeerID)
	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		log.Fatal(err)
	}
}

func (vk *VK) joinGroupEvent(ctx context.Context, obj events.GroupJoinObject) {
	log.Printf("%d: %s", obj.UserID, obj.JoinType)
	b := params.NewMessagesSendBuilder()
	b.Message("Welcome to our group!\n" +
		"Home menu - \"/menu\"\n" +
		"Help - \"/help\"\n" +
		"Reset - \"/reset\"\n")
	b.RandomID(0)
	b.PeerID(obj.UserID)
	b.Keyboard(vk.btnHome())
	_, err := vk.VKApi.MessagesSend(b.Params)
	if err != nil {
		log.Printf("cat't send message with join: %v", err)
		return
	}
}

func (vk *VK) messageEventObject(ctx context.Context, obj events.MessageEventObject) {
	err := vk.btnAnswerEvent(obj)
	if err != nil {
		log.Printf("can't AnswerEvent: %v", err)
		return
	}

	p := &Payload{}
	err = json.Unmarshal(obj.Payload, &p)
	if err != nil {
		log.Printf("can't unmarshal json: %v", err)
		return
	}

	switch p.TypeName {
	case "btn_back_home":
		err = vk.actionBtnBackHome(obj)
		if err != nil {
			log.Printf("action button back home error: %v", err)
			return
		}
	case "btn_say_hello":
		err = vk.actionBtnSayHello(obj)
		if err != nil {
			log.Printf("action button say hello error: %v", err)
			return
		}

	case "btn_home_hello":
		err = vk.actionBtnHomeHello(obj)
		if err != nil {
			log.Printf("action button home hello error: %v", err)
			return
		}

	case "btn_home_link":
		err = vk.actionBtnHomeLink(obj)
		if err != nil {
			log.Printf("action button home link error: %v", err)
			return
		}

	case "btn_home_locate":
		err = vk.actionBtnHomeLocate(obj)
		if err != nil {
			log.Printf("action button home locate error: %v", err)
			return
		}

	case "btn_home_snak":
		err = vk.actionBtnHomeSnak(obj)
		if err != nil {
			log.Printf("action button home snak error: %v", err)
			return
		}
	}
}
