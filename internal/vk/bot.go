package vk

import (
	"VKBot/config"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
)

type VK struct {
	config   *config.Config
	VKApi    *api.VK
	LongPoll *longpoll.LongPoll
}

func NewVK(cfg *config.Config) (*VK, error) {

	vk := api.NewVK(cfg.VKBot.Token)

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		return nil, err
	}

	bot := &VK{
		config:   cfg,
		VKApi:    vk,
		LongPoll: lp,
	}
	return bot, nil
}

// RegisterEvents new message event
func (vk *VK) RegisterEvents() {
	vk.LongPoll.MessageNew(vk.messageNewObject)
	vk.LongPoll.GroupJoin(vk.joinGroupEvent)
	vk.LongPoll.MessageEvent(vk.messageEventObject)
}

func (vk *VK) Run() error {
	return vk.LongPoll.Run()
}
