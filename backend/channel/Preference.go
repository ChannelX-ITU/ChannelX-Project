package channel

import "time"

type Preference struct {
	prefId		int
	startDate	time.Time	//will be optimized
	duratDays	int
	userId		int
	chanId		int
}
