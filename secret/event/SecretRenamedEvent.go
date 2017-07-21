package event

import (
	"github.com/function61/pi-security-module/state"
	"github.com/function61/pi-security-module/util/eventbase"
)

type SecretRenamed struct {
	eventbase.Event
	Id    string
	Title string
}

func (e *SecretRenamed) Apply() {
	for idx, s := range state.Inst.State.Secrets {
		if s.Id == e.Id {
			s.Title = e.Title
			state.Inst.State.Secrets[idx] = s
			return
		}
	}
}
