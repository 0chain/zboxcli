# zbox Command-line Interface for 0Box Storage
zbox Command-line Interface is useful for quickly demonstrate and understand the capabilities of 0Box Storage. The utility is built using 0Chain's ClientSDK library written in Go V1.12.
##Features 
zbox supports following features
1. Register a Wallet
2. Create an allocation
3. Upload a file to 0Box
4. Download the uploaded file from 0Box
5. Update the uploaded file on 0Box
5. Delete the uploaded file on 0Box
6. Share the uploaded file on 0Box
7. List the uploaded files and folders

ZBox Command-line utility provides a self-explaining "help" option that lists out the commands it supports and the parameters each command needs to perform the intended action
## How to get it?
You can clone ZBox Command-line Interface from github repo [Here](https://github.com/0chain/zboxcli)
## Pre-requisites
* zbox Command-line Interface needs Go V1.12 or higher. You need to clone the gosdk from [here](https://github.com/0chain/gosdk)
## How to Build the code?
1. Make sure you've Go SDK 1.12 or higher and Go configurations are set and working on your system.
2. Clone [zboxcli](https://github.com/0chain/zboxcli)
3. Go to the root directory of the local repo
4. Run the following command:

        go build -tags bn256 -o zbox

5. zbox application is built in the local folder. 
## Getting started with zbox
### Before you start
Before you start playing with zbox, you need to know how to access the blockchain and blobbers and what encryption scheme is supported. Both of that information is stored in a configuration files under sample/clusters folder under repo. Choose the suitable one based on your needs.


### Setup
zbox Command-line Interface needs to know the configuration at runtime. By default, configuration files are assumed to be under $Home/.zcn folder. So, create $Home/.zcn folder and store the chosen yml files from clusters folder as nodes.yaml file there.
### Commands
To run the commands, cd to the folder where zbox is located.

Let's go over all the available commands and play with it. Note in this document, we will show only the commands, response will vary depending on your usage, so may not be provided in all places.

#### command with no arguments
When you run zbox with no arguments, it will list all the supported commands.

Command

    ./zbox

Response

    0Box is a decentralized storage application written on the 0Chain platform.
                Complete documentation is available at https://0chain.net

    Usage:

    zbox [command]

    Available Commands:
    delete        delete file from blobbers
    download      download file from blobbers
    get           Gets the allocation info
    help          Help about any command
    list          list files from blobbers
    newallocation Creates a new allocation
    register      Registers the wallet with the blockchain
    share         share files from blobbers
    stats         stats for file from blobbers
    update        update file to blobbers
    upload        upload file to blobbers

    Flags:
        --config string   config file (default is $HOME/.zcn/nodes.yaml)
    -h, --help            help for zbox
        --wallet string   wallet file (default is $HOME/.zcn/wallet.txt)

    Use "zbox [command] --help" for more information about a command.

#### register
Command register registers a wallet that will be used both by the blockchain and blobbers.

Command

     ./zbox register

Response

    ZCN wallet created
    Wallet registered


#### newallocation with help
Command newallocation reserves harddisk space on blobbers. Let's see the parameters it takes by using --help

Command

    ./zbox newallocation --help

Response

    Creates a new allocation

    Usage:
    zbox newallocation [flags]

    Flags:
        --allocationFileName string   --allocationFileName allocation.txt (default "allocation.txt")
        --data int                    --data 2 (default 2)
    -h, --help                        help for newallocation
        --parity int                  --parity 2 (default 2)
        --size int                    --size 10000 (default 2147483648)

    Global Flags:
        --config string   config file (default is $HOME/.zcn/nodes.yaml)
        --wallet string   wallet file (default is $HOME/.zcn/wallet.txt)
As you can see the newallocation command takes allocationFileName where the allocation information is stored locally, data and parity are used for redundancy factor, size is self-explanatory. All the parameters have default values. 

#### newallocation.
Create a new allocation with default values.
Command

    ./zbox newallocation

Response

    Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Also, the allocation information is stored under $Home/.zcn/allocation.txt

#### upload
Use upload command to upload a file. By using help for this command, you will see it takes parameters:
* --allocation -- the allocation id from the newallocation command
* --localpath -- absolute path to the file on your local system
* -- remote path -- remote path where you want to store. It should start with "/"

Command

    ./zbox upload --localpath < absolute path to file>/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type = application/octet-stream. Name = hello.txt

#### list
Use command "list" to list the files and detailed information of each file in the allocation.

Command

    ./zbox list --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath /myfiles

Response

Response will be a json string with array of filenames, it's path and more information about the folders and files. Amongst them lookuphash is useful later during download

#### get
Use command get to get the information about the allocation like  total size of the allocation, used size, number of challenges and the result of that, etc.

Command

    ./zbox get --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac 

Response

Response will have information about blobbers allocated and stats for the allocation. Stats contain important information about the size of the allocation, size used, number of write markers, and challenges passed/failed/open/redeemed

#### share
Use share command to generate an authtoken that provides authorization to the holder to the specified file on the remotepath.

Command

    ./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt

Response

Response contains auth token an encrypted string that can be shared.

#### download
Use download command to download your own or a shared file.

Command

    ./zbox download --authticket eyJjbGllbnRfaWQiOiIiLCJvd25lcl9pZCI6IjRiZjI4ODU5NzgzMjNiMmU0OGUyNGM0ZTNkODkwYTA1MzQwM2E3MDk3NDE3MDljMzA1YjAxZjE5ZDk2NDFhYTgiLCJhbGxvY2F0aW9uX2lkIjoiM2MwZDMyNTYwZWExOGQ5ZDBkNzY4MDgyMTZhOWM2MzRmNjYxOTc5ZDI5YmE1OWNjOGRhZmNjYjNlNWI5NTM0MSIsImZpbGVfcGF0aF9oYXNoIjoiNDE4NjVmMGM2YWFhNTcxM2VkMzkxZWJkZjgyMjU1MmZjNmNmYjU5YTg3YTI2MTY4MjgyNDJiYTNjYTBkY2U0OSIsImZpbGVfbmFtZSI6ImhvcnNlLmpwZyIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE1Njg3NTQ0ODQsInRpbWVzdGFtcCI6MTU2MDk3ODQ4NCwic2lnbmF0dXJlIjoiYjhkZWNhNzM4YjgyNGRiNmNlNzc0NDY1N2FlZmNiNzUzZTYxOWQ4MmJhODEzMjIzYWQ3MGI2NTlkOTQxNDM2YTVkMzQ0N2E5ZmUwNzE1NGYwMThmYjk5NDkyNDQ5ZDk5NmNjMmQ5M2RkMWM0NTJkYzgzNDEyYjVhZTNkMmFmMDEifQ== --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg --lookuphash 41865f0c6aaa5713ed391ebdf822552fc6cfb59a87a2616828242ba3ca0dce49 --localpath ../horse.jpeg

Response

Downloaded file will be in the localpath specified.

#### stats
stats command helps in getting upload, download and challenge information on a file.

Command

    ./zbox stats --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpg

#### update
Use update command to update content of an existing file in the remote path. See upload command



