module github.com/0chain/zboxcli

go 1.13

require (
	github.com/0chain/gosdk v1.2.7-0.20210528185355-76efc0601709
	github.com/aws/aws-sdk-go v1.38.33
	github.com/fatih/color v1.7.0 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.1
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	gopkg.in/cheggaaa/pb.v1 v1.0.28
)

// temporary, for development
//replace github.com/0chain/gosdk => ../gosdk
