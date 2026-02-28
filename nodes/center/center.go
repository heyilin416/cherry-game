package center

import (
	"github.com/cherry-game/cherry"
	cherryCron "github.com/cherry-game/components/cron"
	"github.com/heyilin416/cherry-game/internal/data"
	"github.com/heyilin416/cherry-game/nodes/center/db"
	"github.com/heyilin416/cherry-game/nodes/center/module/account"
	"github.com/heyilin416/cherry-game/nodes/center/module/ops"
)

func Run(profileFilePath, nodeID string) {
	app := cherry.Configure(
		profileFilePath,
		nodeID,
		false,
		cherry.Cluster,
	)

	app.Register(cherryCron.New())
	app.Register(data.New())
	app.Register(db.New())

	app.AddActors(
		&account.ActorAccount{},
		&ops.ActorOps{},
	)

	app.Startup()
}
