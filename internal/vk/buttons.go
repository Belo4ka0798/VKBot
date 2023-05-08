package vk

import (
	"github.com/SevereCloud/vksdk/v2/object"
)

// TODO  Убрать b.Keyboard() из кнопок!
func (vk *VK) btnHome() *object.MessagesKeyboard {
	return object.NewMessagesKeyboard(true).
		AddRow().
		AddCallbackButton("Hello", map[string]string{"type": "btn_home_hello"}, "primary").
		AddCallbackButton("Dev VK", map[string]string{"type": "btn_home_link"}, "primary").
		AddCallbackButton("Locate", map[string]string{"type": "btn_home_locate"}, "primary").
		AddCallbackButton("Snak", map[string]string{"type": "btn_home_snak"}, "primary")
}

func (vk *VK) btnHomeHello() *object.MessagesKeyboard {
	return object.NewMessagesKeyboard(true).
		AddRow().
		AddCallbackButton("Say Hello", map[string]string{"type": "btn_say_hello"}, "primary").
		AddCallbackButton("Go home", map[string]string{"type": "btn_back_home"}, "negative")
}

func (vk *VK) btnHomeVKDevLink() *object.MessagesKeyboard {
	return object.NewMessagesKeyboard(true).
		AddRow().
		AddOpenLinkButton("https://dev.vk.com", "Open VKApi", "").
		AddCallbackButton("Go home", map[string]string{"type": "btn_back_home"}, "negative")
}

func (vk *VK) btnHomeLocation() *object.MessagesKeyboard {
	return object.NewMessagesKeyboard(false).
		AddRow().
		AddLocationButton("").
		AddCallbackButton("Go home", map[string]string{"type": "btn_back_home"}, "negative")
}

func (vk *VK) btnHomeSnak() *object.MessagesKeyboard {
	return object.NewMessagesKeyboard(false).
		AddRow().
		AddCallbackButton("Snak", map[string]string{"type": "btn_snak_push"}, "primary").
		AddCallbackButton("Go home", map[string]string{"type": "btn_back_home"}, "negative")
}
