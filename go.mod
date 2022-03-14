module github.com/0chain/zboxcli

go 1.17

require (
	github.com/0chain/errors v1.0.3
	github.com/0chain/gosdk v1.7.7-0.20220314114910-0c923e538d35
	github.com/icza/bitio v1.1.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)

// temporary, for development
//replace github.com/0chain/gosdk => ../gosdk
