package vk

import (
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/object"
)

func (vk *VK) btnHome(b *params.MessagesSendBuilder) {
	keyboard := object.NewMessagesKeyboard(true)
	keyboard.AddRow()
	keyboard.AddCallbackButton("Hello", map[string]string{"type": "home_btn1"}, "primary")
	keyboard.AddCallbackButton("Dev VK", map[string]string{"type": "home_btn2"}, "primary")
	keyboard.AddCallbackButton("Locate", map[string]string{"type": "home_btn3"}, "primary")
	keyboard.AddCallbackButton("Snak", map[string]string{"type": "home_btn4"}, "primary")
	b.Keyboard(keyboard)
}

func (vk *VK) btnHomeBtn1(b *params.MessagesSendBuilder) {
	keyboard := object.NewMessagesKeyboard(true)
	keyboard.AddRow()
	keyboard.AddCallbackButton("Say Hello", map[string]string{"type": "say_hello"}, "primary")
	keyboard.AddCallbackButton("Go home", map[string]string{"type": "home_btn_back"}, "negative")
	b.Keyboard(keyboard)
}

func (vk *VK) btnHomeBtn2(b *params.MessagesSendBuilder) {
	keyboard := object.NewMessagesKeyboard(true)
	keyboard.AddRow()
	keyboard.AddOpenLinkButton("https://dev.vk.com", "Open VKApi", "")
	keyboard.AddCallbackButton("Go home", map[string]string{"type": "home_btn_back"}, "negative")
	b.Keyboard(keyboard)
}

func (vk *VK) btnHomeBtn3(b *params.MessagesSendBuilder) {
	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddLocationButton("")
	keyboard.AddCallbackButton("Go home", map[string]string{"type": "home_btn_back"}, "negative")
	b.Keyboard(keyboard)
}

//func (vk *VK) btnCall3(msg *params.MessagesSendBuilder) {
//	msg.Keyboard(object.NewMessagesKeyboard(true).
//		AddRow().
//		AddCallbackButton("Say Hello", map[string]string{"type": "call_btn1"}, "primary").
//		AddCallbackButton("Back", map[string]string{"type": "call_homebtn"}, "negative"))
//}
//
//func (vk *VK) btnCall4(msg *params.MessagesSendBuilder) {
//	msg.Keyboard(object.NewMessagesKeyboard(true).
//		AddRow().
//		AddCallbackButton("Say Hello", map[string]string{"type": "call_btn1"}, "primary").
//		AddCallbackButton("Back", map[string]string{"type": "call_homebtn"}, "negative"))
//}
