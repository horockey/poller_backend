package polls

import (
	"github.com/google/uuid"
	"github.com/horockey/poller_backend/internal/model"
)

func FixID(p *model.Poll) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
}
