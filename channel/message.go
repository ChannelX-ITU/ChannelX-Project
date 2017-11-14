package channel

import (
)

type Message struct {
	to 			*User
	sub			string
	msg			string
}