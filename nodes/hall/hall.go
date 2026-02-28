package hall

import (
	"github.com/cherry-game/cherry"
	cherrySnowflake "github.com/cherry-game/cherry/extend/snowflake"
	cstring "github.com/cherry-game/cherry/extend/string"
	cherryCron "github.com/cherry-game/components/cron"
	cherryGops "github.com/cherry-game/components/gops"
	checkCenter "github.com/heyilin416/cherry-game/internal/component/check_center"
	"github.com/heyilin416/cherry-game/internal/data"
	"github.com/heyilin416/cherry-game/nodes/hall/db"
	"github.com/heyilin416/cherry-game/nodes/hall/module/player"
)

func Run(profileFilePath, nodeID string) {
	// 从 nodeID 中提取数字部分作为 serverId
	serverIdStr := ""
	for _, c := range nodeID {
		if c >= '0' && c <= '9' {
			serverIdStr += string(c)
		}
	}

	if serverIdStr == "" {
		panic("node parameter must contain number.")
	}

	// snowflake global id
	serverId, _ := cstring.ToInt64(serverIdStr)
	cherrySnowflake.SetDefaultNode(serverId)

	// 配置cherry引擎
	app := cherry.Configure(profileFilePath, nodeID, false, cherry.Cluster)

	// diagnose
	app.Register(cherryGops.New())
	// 注册调度组件
	app.Register(cherryCron.New())
	// 注册数据配置组件
	app.Register(data.New())
	// 注册检测中心节点组件，确认中心节点启动后，再启动当前节点
	app.Register(checkCenter.New())
	// 注册db组件
	app.Register(db.New())

	app.AddActors(
		&player.ActorPlayers{},
	)

	app.Startup()
}
