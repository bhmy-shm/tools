package tools

import "github.com/bhmy-shm/tools/gtop/common"

var (
	Cmd       = common.NewCommand("tools", common.WithRunE(nil))
	pingCmd   = common.NewCommand("ping", common.WithRunE(pingServiceCommand))
	sshCmd    = common.NewCommand("ssh", common.WithRunE(sshServiceCommand))
	scpCmd    = common.NewCommand("scp", common.WithRunE(scpServiceCommand))
	packetCmd = common.NewCommand("packet", common.WithRunE(packetServiceCommand))
)

func init() {
	var (
		pingCmdFlags   = pingCmd.Flags()
		sshCmdFlags    = sshCmd.Flags()
		scpCmdFlags    = scpCmd.Flags()
		packetCmdFlags = packetCmd.Flags()
	)

	pingCmdFlags.IntVarP(&VarPingCountInt, "count", "c", 5, "ping 次数")
	pingCmdFlags.IntVarP(&VarPingInterval, "interval", "i", 1, "ping 间隔，单位：秒")
	pingCmdFlags.IntVarP(&VarPingTimeout, "timeout", "t", 30, "ping 超时，单位：秒")
	pingCmdFlags.BoolVar(&VarPingPrivileged, "privileged")

	sshCmdFlags.StringVar(&VarSSHUser, "user")
	sshCmdFlags.StringVar(&VarSSHAddress, "host")

	scpCmdFlags.StringVarP(&VarScpUser, "user", "u")
	scpCmdFlags.StringVarP(&VarScpAddress, "host", "h")
	scpCmdFlags.StringVarP(&VarScpSource, "source", "s")
	scpCmdFlags.StringVarP(&VarScpDirect, "direct", "d")
	scpCmdFlags.BoolVarP(&VarScpRemoteToLocal, "remote", "r")

	packetCmdFlags.StringVar(&VarPacketDeviceName, "net")

	Cmd.AddCommand(pingCmd, sshCmd, scpCmd, packetCmd)
}
