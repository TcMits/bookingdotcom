package bookingdotcom

import (
	"github.com/TcMits/bookingdotcom/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingDotCom interface {
	DB() (*mongo.Database, error)
  DaoHook() *hook.Hook[*core.DaoEvent]
}
