# zbox - a CLI for 0Chain dStorage
zbox is a command line interface (CLI) tool to understand the capabilities of 0Chain dStorage and prototype your app. The utility is built using 0Chain's ClientSDK library written in Go. Check out a [video](https://youtu.be/TPrkRjdaHrY) on how to use the CLI to create an allocation (storage volume) and upload, download, update, delete, and share files and folders to dStor on the 0Chain dStorage platform.
##Features
zbox supports the following features
1. Register a Wallet
2. Create an allocation
3. Upload a file to dStorage
4. Download the uploaded file from dStorage
5. Update the uploaded file on dStorage
6. Delete the uploaded file on dStorage
7. Share the uploaded file on dStorage to the public
8. List the uploaded files and folders
9. Copy uploaded files to another folder path on dStorage
10. Upload encrypted files to dStorage
11. Share an encrypted file using proxy re-encryption (PRE) with your friend

zbox CLI provides a self-explaining "help" option that lists commands and parameters they need to perform the intended action
## How to get it?
    git clone https://github.com/0chain/zboxcli.git
## Pre-requisites
    Go V1.12 or higher.
### How to install Go on Linux
    All build requirements here for 64-bit only.
#### Using wget
    wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
    rm go1.13.linux-amd64.tar.gz
Add Go to path (alternatively, add to .profile)
    export PATH=$PATH:/usr/local/go/bin
#### Using apt on Ubuntu18+ e.g. Ubuntu 18, Mint 19
    sudo apt update
    sudo apt install build-essentials
    sudo apt install git
#### Using yum on RHEL7+ e.g. Centos 7
    yum update -y
    yum install -y openssl-devel
    yum groupinstall -y "Development Tools"
    yum install -y git
    yum install -y wget
    yum install -y make
    yum install -y g++
### How to build zboxcli on Linux
    git clone https://github.com/0chain/zboxcli.git
    cd zboxcli
    make install
Make a .zcn folder and copy a sample yaml file that represents the dStorage network. Currently it is devb.yaml
    mkdir $HOME/.zcn
    cp sample/config/devb.yaml $HOME/.zcn/nodes.yaml
Type ./zbox and you should see the help list appear!

### How to build on Windows
Windows 64bit (tested with Windows 10)
#### Make (e.g. gnuwin32 on sourceforge)
* Install executables make3.8.1 binary & make3.8.1 dependencies

        http://gnuwin32.sourceforge.net/packages/make.htm

* The 3 files from the bin folder of above two packages (make.exe and 2xDLLs) need to be copied into a windows folder such as Windows/system, the zboxcli folder, or a folder of your choosing and the path added to windows system path

        Windows System > Control Panel > System and Security > System > Advanced System Settings > Environment Variables > (System Variables - Path > Edit ) > New [add your chosen path]

#### Compiler tools via MinGW-W64
    https://sourceforge.net/projects/mingw-w64/

* Will install required compiler tools for you;
* Select Architecture x86_64, leave other options as is;
* Once again, you may need to manually add the mingw64 binaries path to your system path
* for whatever path has been installed on your system e.g.

        Windows System > Control Panel > System and Security > System > Advanced System Settings > Environment Variables > (System Variables - Path > Edit ) > New C:\Program Files\mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\bin

#### Git for Windows (Optional)
* Installer should handle defaults for you

        gitforwindows.org


#### Golang for Windows
* Via Installer (Windows Amd64 msi) should handle defaults and perform windows path updates for you

        https://golang.org/dl/

* Any Open Command (Terminal) Windows may need to be restarted to 'see' any paths updated above

#### zboxcli
        md go
        cd go
        git clone https://github.com/0chain/zboxcli.git
        cd zboxcli
        make install

(As an alternative to the git clone command, you can manually download and extract from the github website, 0chain/zboxcli)
If you get as far as this;

        process_begin: CreateProcess(NULL, cp -f zwallet /sample/zwallet, ...) failed.
        make (e=2): The system cannot find the file specified.
        make: *** [install] Error 2

Dont worry, its built fine, just rename the file to an .exe like

        rename zbox zbox.exe

* Locate your home folder, like \Users\<username>
* Make a folder called .zcn
* copy devb.yml file from zboxcli\samples\config to the new .zcn folder
* Also move zbox.exe to that folder if desired.
* You should now be able to run the "zbox" command and get the help menu


## Getting started with zbox
### Before you start
Before you start playing with zbox, you need to access the blockchain. Go to sample/clusters folder in the repo, and choose a  network. Copy it to your ~/.zcn folder and then rename it as config.yaml file.

    mkdir ~/.zcn
    cp ~/.zcn/devb.yaml ~/.zcn/config.yaml

Sample config.yaml

      miners:
      - http://virb.devb.testnet-0chain.net:7071
      - http://vira.devb.testnet-0chain.net:7071
      - http://cala.devb.testnet-0chain.net:7071
      - http://calb.devb.testnet-0chain.net:7071
      sharders:
      - http://cala.devb.testnet-0chain.net:7171
      - http://vira.devb.testnet-0chain.net:7171
      signature_scheme: bls0chain
      min_submit: 50 # in percentage
      min_confirmation: 50 # in percentage
      confirmation_chain_length: 6

### Setup
The zbox command line uses the ~/.zcn/nodes.yaml file at runtime to point to the network specified in that file.

### Commands
Note in this document, we will show only the commands, response will vary depending on your usage, so may not be provided in all places.

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
    copy            copy an object(file/folder) to another folder on blobbers
    delete          delete file from blobbers
    download        download file from blobbers
    get             Gets the allocation info
    getwallet       Get wallet information
    help            Help about any command
    list            list files from blobbers
    listallocations List allocations for the client
    meta            get meta data of files from blobbers
    newallocation   Creates a new allocation
    register        Registers the wallet with the blockchain
    rename          rename an object(file/folder) on blobbers
    share           share files from blobbers
    stats           stats for file from blobbers
    sync            Sync files to/from blobbers
    update          update file to blobbers
    upload          upload file to blobbers
    version         Prints version information

    Flags:
        --config string   config file (default is $HOME/.zcn/nodes.yaml)
    -h, --help            help for zbox
        --verbose         prints sdk log in stdio (default false)
        --wallet string   wallet file (default is $HOME/.zcn/wallet.txt)

    Use "zbox [command] --help" for more information about a command.

#### register
Command register registers a wallet that will be used both by the blockchain and blobbers, and is created in the ~/.zcn directory. If you have created a wallet with another network, you will need to remove and recreate it. If you want to create multiple wallets with multiple allocations, make sure you store the wallet information. zbox uses the keys in ~/.zcn/wallet.txt when it executes the commands.

Command

     ./zbox register

Response

    ZCN wallet created
    Wallet registered


#### newallocation with help
Command newallocation reserves hard disk space on the blobbers. Let's see the parameters it takes by using --help

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
As you can see the newallocation command takes allocationFileName where the volume information is stored locally. All the parameters have default values. With more data shards, you can upload or download files faster. With more parity shards, you have higher availability.

#### newallocation.
Create a new allocation with default values. If you have not registered a wallet, it will automatically create a wallet.
Command

    ./zbox newallocation

Response

    Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Also, the allocation information is stored under $Home/.zcn/allocation.txt

#### upload
Use upload command to upload a file. By using help for this command, you will see it takes parameters:
* --allocation -- the allocation id from the newallocation command
* --localpath -- absolute path to the file on your local system
* --remote path -- remote path where you want to store. It should start with "/"
* --thumbnailpath -- Local thumbnail path of file to upload
* --encrypt -- [OPTIONAL] pass this option to encrypt and upload the file


Command

    ./zbox upload --localpath < absolute path to file>/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type = application/octet-stream. Name = hello.txt


#### upload --encrypt
Use upload command with optional encrypt parameter to upload a file in encrypted format. This can be downloaded as normal from same wallet/allocation or utilize Proxy Re-Encryption facility (see download command).


Command

    ./zbox upload --encrypt --localpath <absolute path to file>/sensitivedata.txt --remotepath /myfiles/sensitivedata.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type =

#### list
Use list command to list files in given path. By using help for this command, you will see it takes parameters:
      --allocation string   Allocation ID
      --remotepath string   Remote path to list from (Required for --allocation)
      --authticket string   Auth ticket fot the file to download if you dont own it
      --lookuphash string   The remote lookuphash of the object retrieved from the list


Command

    ./zbox list --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath /myfiles

Response

Response will be a list with information for each file/folder in the given path. The information includes lookuphash which is require for download via authticket
(Optional file list in json format)

#### get
Use command get to get the information about the allocation like  total size of the allocation, used size, number of challenges and the result of that, etc.

Command

    ./zbox get --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

Response will have information about blobbers allocated and stats for the allocation. Stats contain important information about the size of the allocation, size used, number of write markers, and challenges passed/failed/open/redeemed

#### getwallet
Use command get to get additional wallet information including Encryption Public Key required for Proxy Re-Encryption.

Command

    ./zbox getwallet

Response

Response will give details for current selected wallet (or wallet file specified by optional --wallet parameter)

#### share
Use share command to generate an authtoken that provides authorization to the holder to the specified file on the remotepath.
      --allocation string            Allocation ID
      --clientid string              ClientID of the user to share with. Leave blank for public share
      --encryptionpublickey string   Encryption public key of the client you want to share with (from getwallet command)


Command

    ./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt

Response

Response contains auth token an encrypted string that can be shared.

#### download
Use download command to download your own or a shared file.
      --allocation string     Allocation ID
      --authticket string     Auth ticket fot the file to download if you dont own it
      --localpath string      Local path of file to download
      --lookuphash string     The remote lookuphash of the object retrieved from the list
      --remotepath string     Remote path to download
  -t, --thumbnail             pass this option to download only the thumbnail

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

#### sync
sync command helps to sync all files in localfolder recursively to remote.

Command help

    Sync all files to/from blobbers from/to a localpath

    Usage:
    zbox sync [flags]

    Flags:
        --allocation string         Allocation ID
        --excludepath stringArray   Remote folder paths exclude to sync
    -h, --help                      help for sync
        --localcache string         Local cache of remote snapshot.
                                    If file exists, this will be used for comparison with remote.
                                    After sync complete, remote snapshot will be updated to the same file for next use.
        --localpath string          Local dir path to sync

    Global Flags:
        --config string   config file (default is $HOME/.zcn/nodes.yaml)
        --verbose         prints sdk log in stdio (default false)
        --wallet string   wallet file (default is $HOME/.zcn/wallet.txt)

