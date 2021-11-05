module github.com/0chain/zboxcli

go 1.13

require (
	github.com/0chain/errors v1.0.3
	github.com/0chain/gosdk v1.3.0-beta.3.0.20211105104758-a142151bcd19 //
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)

// temporary, for development
//replace github.com/0chain/gosdk => ../gosdk
