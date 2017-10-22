package channel

import "os/user"

type Channel struct {
	name	string
	owner	user.User
	set		ChannelSetting
	users	[]user.User
}