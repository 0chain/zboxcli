module github.com/0chain/zboxcli

go 1.13

require (
	github.com/0chain/errors v1.0.3
	github.com/0chain/gosdk v1.5.1-0.20220211100741-875ec57b7709
	github.com/olekukonko/tablewriter v0.0.5
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)

// temporary, for development
//replace github.com/0chain/gosdk => ../gosdk
