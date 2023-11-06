package attempts

import (
	"github.com/google/uuid"
	"github.com/horockey/poller_backend/internal/model"
)

func FixID(a *model.Attempt) {
	if a.ID == "" {
		a.ID = uuid.NewString()
	}
}
