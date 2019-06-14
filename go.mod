module github.com/0chain/zboxcmd

replace github.com/0chain/zboxcmd => ../zboxcmd

replace github.com/0chain/gosdk => ../gosdk

go 1.12

require (
	github.com/0chain/gosdk v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)
