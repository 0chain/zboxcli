# zbox - a CLI for 0Chain dStorage
zbox is a command line interface (CLI) tool to understand the capabilities of 0Chain dStorage and prototype your app. The utility is built using 0Chain's goSDK library written in Go. Check out a [video](https://youtu.be/TPrkRjdaHrY) on how to use the CLI to create an allocation (storage volume) and upload, download, update, delete, and share files and folders to dStor on the 0Chain dStorage platform.

## Features

[zbox](#Command-with-no-arguments) supports the following features
1. [Register a Wallet](#Register)
2. [Create an allocation](#Create-new-allocation)
3. [Update an allocation](#Update-allocation)
4. [Upload a file to dStorage](#Upload)
5. [Download the uploaded file from dStorage](#Download)
6. [Update the uploaded file on dStorage](#Update)
7. [Delete the uploaded file on dStorage](#Delete)
8. [Share the uploaded file on dStorage to the public](#Share)
9. [List the uploaded files and folders](#List)
10. [Copy uploaded files to another folder path on dStorage](#Copy)
11. [Upload encrypted files to dStorage](#with---encrypt)
12. [Share an encrypted file using proxy re-encryption (PRE) with your friend](#share-encrypted)
13. [List allocations](#List-allocations)
14. [Sync your local folder to remote](#Sync)
15. [Get wallet information](#Get-wallet)
16. [Get Allocation info](#Get)
17. [Get meta data of files](#Get-metadata)
18. [Rename an object in allocation](#Rename)
19. [Get file stats](#Stats)
20. [Repair a file on dStorage](#Repair)
21. [Cancel allocation](#Cancel-allocation)
22. [Finalize allocation](#Finalize-allocation)
23. [Challenge pool information](#Challenge-pool-information)
24. [List blobbers](#List-blobbers)
25. [Read pool](#Read-pool)
26. [Storage SC configurations](#Storage-SC-configurations)
27. [Stake pool](#Stake-pool)
28. [Write pool](#Write-pool)
29. [Costs](#Costs)

zbox CLI provides a self-explaining "help" option that lists commands and parameters they need to perform the intended action
## How to get it?
    git clone https://github.com/0chain/zboxcli.git
## Pre-requisites
    Go V1.12 or higher.

### [How to build on Linux](https://github.com/0chain/zboxcli/wiki/Build-Linux)
### [How to build on Windows](https://github.com/0chain/zboxcli/wiki/Build-Windows)

## Getting started with zbox
### Before you start
Before you start playing with zbox, you need to access the blockchain. Go to network folder in the repo, and choose a network. Copy it to your ~/.zcn folder and then rename it as config.yaml file.

    mkdir ~/.zcn
    cp network/one.yaml ~/.zcn/config.yaml

Sample config.yaml

      miners:
      - http://one.devnet-0chain.net:31071
      - http://one.devnet-0chain.net:31072
      - http://one.devnet-0chain.net:31073
      - http://one.devnet-0chain.net:31074                  
      sharders:
      - http://one.devnet-0chain.net:31171
      - http://one.devnet-0chain.net:31172
      preferred_blobbers:
      - http://one.devnet-0chain.net:31051
      - http://one.devnet-0chain.net:31052
      signature_scheme: bls0chain
      min_submit: 50 # in percentage
      min_confirmation: 50 # in percentage
      confirmation_chain_length: 3

### Setup
The zbox command line uses the ~/.zcn/config.yaml file at runtime to point to the network specified in that file.

## Commands
Note in this document, we will show only the commands, response will vary depending on your usage, so may not be provided in all places.

### Command with no arguments
When you run zbox with no arguments, it will list all the supported commands.

Command

    ./zbox

Response

    0Box is a decentralized storage application written on the 0Chain platform.
    			Complete documentation is available at https://0chain.net
    
    Usage:
      zbox [command]
    
    Available Commands:
      alloc-cancel      Cancel an allocation
      alloc-fini        Finalize an expired allocation
      copy              copy an object(file/folder) to another folder on blobbers
      cp-info           Challenge pool information.
      delete            delete file from blobbers
      download          download file from blobbers
      get               Gets the allocation info
      get-download-cost Get downloading cost
      get-upload-cost   Get uploading cost
      getwallet         Get wallet information
      help              Help about any command
      list              list files from blobbers
      listallocations   List allocations for the client
      ls-blobbers       Show active blobbers in storage SC.
      meta              get meta data of files from blobbers
      newallocation     Creates a new allocation
      register          Registers the wallet with the blockchain
      rename            rename an object(file/folder) on blobbers
      rp-create         Create read pool if missing
      rp-info           Read pool information.
      rp-lock           Lock some tokens in read pool.
      rp-unlock         Unlock some expired tokens in a read pool.
      sc-config         Show storage SC configuration.
      share             share files from blobbers
      sp-info           Stake pool information.
      sp-lock           Lock tokens lacking in stake pool.
      sp-pay-interests  Pay interests not payed yet.
      sp-unlock         Unlock tokens in stake pool.
      start-repair      start repair file to blobbers
      stats             stats for file from blobbers
      sync              Sync files to/from blobbers
      update            update file to blobbers
      updateallocation  Updates allocation's expiry and size
      upload            upload file to blobbers
      version           Prints version information
      wp-info           Write pool information.
      wp-lock           Lock some tokens in write pool.
      wp-unlock         Unlock some expired tokens in a write pool.

    Flags:
          --config string      config file (default is config.yaml)
          --configDir string   configuration directory (default is $HOME/.zcn)
      -h, --help               help for zbox
          --verbose            prints sdk log in stderr (default false)
          --wallet string      wallet file (default is wallet.json)
    
    Use "zbox [command] --help" for more information about a command.

### Register
Command register registers a wallet that will be used both by the blockchain and blobbers, and is created in the ~/.zcn directory. If you have created a wallet with another network, you will need to remove and recreate it. If you want to create multiple wallets with multiple allocations, make sure you store the wallet information. zbox uses the keys in ~/.zcn/wallet.json when it executes the commands.

Command

     ./zbox register

Response

    ZCN wallet created
    Wallet registered


### Create new allocation
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
          --expire duration             duration to allocation expiration (default 720h)
      -h, --help                        help for newallocation
          --lock string                 lock write pool with given number of tokens, or use 'auto' (default "0")
          --mcct duration               max challenge completion time, optional, default 1 hour
          --parity int                  --parity 2 (default 2)
          --read_price string           select blobbers by provided read price range, use form 0.5-1.5, default is [0; inf)
          --size int                    --size 10000 (default 2147483648)
          --write_price string          select blobbers by provided write price range, use form 1.5-2.5, default is [0; inf)

    Global Flags:
          --config string      config file (default is config.yaml)
          --configDir string   configuration directory (default is $HOME/.zcn)
          --verbose            prints sdk log in stderr (default false)
          --wallet string      wallet file (default is wallet.json)

As you can see the newallocation command takes allocationFileName where the volume information is stored locally. All the parameters have default values. With more data shards, you can upload or download files faster. With more parity shards, you have higher availability.

#### Usage

Create a new allocation with default values. If you have not registered a wallet,
it will automatically create a wallet. The newallocation also, on success,
creates related write pool. Use `--lock` flag to add some tokens to the pool
in the transaction.

Command

    ./zbox newallocation

Response

    Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Also, the allocation information is stored under $Home/.zcn/allocation.txt

### Update allocation
Command updateallocation updates hard disk space and expiry on the blobbers. Let's see the parameters it takes by using --help

Command

    ./zbox updateallocation --help

Response

    Updates allocation's expiry and size

    Usage:
      zbox updateallocation [flags]

    Flags:
          --allocation_id string   allocation ID
          --expiry duration        adjust storage expiration time, duration (default 720h)
      -h, --help                   help for updateallocation
          --size int               adjust allocation size, bytes (default 2147483648)

    Global Flags:
          --config string      config file (default is config.yaml)
          --configDir string   configuration directory (default is $HOME/.zcn)
          --verbose            prints sdk log in stderr (default false)
          --wallet string      wallet file (default is wallet.json)

#### Usage

Create a new allocation with default values. If you have not registered a wallet,
it will automatically create a wallet.

Command

    ./zbox updateallocation --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Allocation updated with txId : fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d

You can see more txn details usin above txID in block explorer.

### Finalize and expired allocation

Finalize an expired allocation. It makes sure all blobbers has got min_lock_demand,
unlocks all related pools. Moves tokens from write pool to allocation owner.

     ./zbox finialloc --allocaiton <allocaiton_id>


### Upload
Use upload command to upload a file. By using help for this command, you will see it takes parameters:
* --allocation -- the allocation id from the newallocation command
* --localpath -- absolute path to the file on your local system
* --remote path -- remote path where you want to store. It should start with "/"
* --thumbnailpath -- Local thumbnail path of file to upload
* --encrypt -- [OPTIONAL] pass this option to encrypt and upload the file
* --commit -- [OPTIONAL] pass this option to commit the metadata transaction

Command

    ./zbox upload --localpath <absolute path to file>/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type = application/octet-stream. Name = hello.txt


#### with --encrypt
Use upload command with optional encrypt parameter to upload a file in encrypted format. This can be downloaded as normal from same wallet/allocation or utilize Proxy Re-Encryption facility (see [download](#Download) command).

Command

    ./zbox upload --encrypt --localpath <absolute path to file>/sensitivedata.txt --remotepath /myfiles/sensitivedata.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type =

### Download
Use download command to download your own or a shared file.
* --allocation string     Allocation ID
* --commit                Pass this option to commit the metadata transaction
* --authticket string     Auth ticket fot the file to download if you dont own it
* --localpath string      Local path of file to download
* --lookuphash string [OPTIONAL]     The remote lookuphash of the object retrieved from the list
* --remotepath string     Remote path to download
* -t, --thumbnail             pass this option to download only the thumbnail

Command

    ./zbox download --authticket eyJjbGllbnRfaWQiOiIiLCJvd25lcl9pZCI6IjRiZjI4ODU5NzgzMjNiMmU0OGUyNGM0ZTNkODkwYTA1MzQwM2E3MDk3NDE3MDljMzA1YjAxZjE5ZDk2NDFhYTgiLCJhbGxvY2F0aW9uX2lkIjoiM2MwZDMyNTYwZWExOGQ5ZDBkNzY4MDgyMTZhOWM2MzRmNjYxOTc5ZDI5YmE1OWNjOGRhZmNjYjNlNWI5NTM0MSIsImZpbGVfcGF0aF9oYXNoIjoiNDE4NjVmMGM2YWFhNTcxM2VkMzkxZWJkZjgyMjU1MmZjNmNmYjU5YTg3YTI2MTY4MjgyNDJiYTNjYTBkY2U0OSIsImZpbGVfbmFtZSI6ImhvcnNlLmpwZyIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE1Njg3NTQ0ODQsInRpbWVzdGFtcCI6MTU2MDk3ODQ4NCwic2lnbmF0dXJlIjoiYjhkZWNhNzM4YjgyNGRiNmNlNzc0NDY1N2FlZmNiNzUzZTYxOWQ4MmJhODEzMjIzYWQ3MGI2NTlkOTQxNDM2YTVkMzQ0N2E5ZmUwNzE1NGYwMThmYjk5NDkyNDQ5ZDk5NmNjMmQ5M2RkMWM0NTJkYzgzNDEyYjVhZTNkMmFmMDEifQ== --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg --lookuphash 41865f0c6aaa5713ed391ebdf822552fc6cfb59a87a2616828242ba3ca0dce49 --localpath ../horse.jpeg

Note: You can download by using only 1 on the below combination
* --remotepath , --allocation 
* --authticket

Response

Downloaded file will be in the localpath specified.

### Update
Use update command to update content of an existing file in the remote path. Similar to [upload](#Upload) command.

### Delete
Use delete command to delete your file on allocation.
      --allocation string     Allocation ID
      --commit                Pass this option to commit the metadata transaction
      --remotepath string     Remote path to download
Command

    ./zbox delete --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg

Response

File successfully deleted (Can be verified using [list](#List))

### Share
Use share command to generate an authtoken that provides authorization to the holder to the specified file on the remotepath.
      --allocation string            Allocation ID
      --clientid string              ClientID of the user to share with. Leave blank for public share
      --encryptionpublickey string   Encryption public key of the client you want to share with (from [getwallet](#Get-wallet) command )
      --remotepath string            Remote path to share

Command             

    ./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt

#### share-encrypted
Use clientid and encryptionpublickey of the user to share with.

Command             

    ./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt --clientid d52d82133177ec18505145e784bc87a0fb811d7ac82aa84ae6b013f96b93cfaa --encryptionpublickey +BV37Ip05OdHo+Sz4N8xgCrACCUOVLnQICU1IiJq8uU=

Response

Response contains auth token an encrypted string that can be shared.

### List
Use list command to list files in given path. By using help for this command, you will see it takes parameters:
* --allocation string --Allocation ID
* --remotepath string --Remote path to list from (Required for --allocation)
* --authticket string --Auth ticket fot the file to download if you dont own it
* --lookuphash string --The remote lookuphash of the object retrieved from the list


Command

    ./zbox list --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath /myfiles

Response

Response will be a list with information for each file/folder in the given path. The information includes lookuphash which is require for download via authticket
(Optional file list in json format)

### Copy
Use copy command to copy file to another folder path in dStorage. By using help for this command, you will see it takes parameters:
* --allocation string --Allocation ID
* --destpath string     Destination path for the object. Existing directory the object should be copied to.
* --remotepath string   Remote path of object to copy

Command

    ./zbox copy --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath /file.txt --destpath /existingFolder/file.txt

### List allocations
Use list allocations command to list all allocations for a client.

Command

    ./zbox listallocations

Response

    Response will be table format output in which all of your allocations will be listed with their ID, SIZE, EXPIRATION, DATASHARDS and PARITYSHARDS

### Sync
Sync command helps to sync all files in localfolder recursively to remote.

Command

    ./zbox sync --help

Response

    Sync all files to/from blobbers from/to a localpath

    Usage:
      zbox sync [flags]

    Flags:
          --allocation string         Allocation ID
          --commit                    pass this option to commit the metadata transaction - only works with     uploadonly
          --excludepath stringArray   Remote folder paths exclude to sync
      -h, --help                      help for sync
          --localcache string         Local cache of remote snapshot.
                                      If file exists, this will be used for comparison with remote.
                                      After sync complete, remote snapshot will be updated to the same file for next    use.
          --localpath string          Local dir path to sync
          --uploadonly                pass this option to only upload/update the files

    Global Flags:
          --config string      config file (default is config.yaml)
          --configDir string   configuration directory (default is $HOME/.zcn)
          --verbose            prints sdk log in stderr (default false)
          --wallet string      wallet file (default is wallet.json)

#### Usage
Command

    ./zbox sync --localpath /myLocalFolder --remotepath / --allocation 06e227bc283cb9ebc98c164f8f3fd4f02e70f76ec93cb8b41b276146db3f329f --localcache /someFolder/Localcache.txt

Response

It will sync your localpath with the remote and do all the requried CRUD operations.

### Get wallet
Use command get to get additional wallet information including Encryption Public Key required for Proxy Re-Encryption.

Command

    ./zbox getwallet

Response

Response will give details for current selected wallet (or wallet file specified by optional --wallet parameter)

### Get
Use command get to get the information about the allocation like  total size of the allocation, used size, number of challenges and the result of that, etc.

Command

    ./zbox get --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

Response will have information about blobbers allocated and stats for the allocation. Stats contain important information about the size of the allocation, size used, number of write markers, and challenges passed/failed/open/redeemed



### Get metadata
Use command meta to get meta data for the given remote file. Use with help to know more about possible flags.
* --allocation string -- Allocation ID
* --authticket string -- Auth ticket fot the file to download if you dont own it
* --lookuphash string -- The remote lookuphash of the object retrieved from the list
* --remotepath string -- Remote path to list from

Command

    ./zbox meta --allocation 0834e2848e3006714423f485cfc1c22266368d655026152bd762cd0cdd0a4aeb --remotepath /file.txt

With authticket

    ./zbox meta --authticket <authticket in encoded hash string> --lookuphash <optional>

Response

Response will be meta data for the given filepath/lookuphash (if using authTicket)


### Rename
rename command helps in renaming a file already exists in dStorage.
* --allocation string -- Allocation ID
* --destname string -- New Name for the object (Only the name and not the path). Include the file extensio*if   applicable
* --remotepath string -- Remote path of object to rename

Command

    ./zbox rename -- allocation 0834e2848e3006714423f485cfc1c22266368d655026152bd762cd0cdd0a4aeb --remotepath /file1.txt --destname file2.txt 

### Stats
stats command helps in getting upload, download and challenge information on a file.

Command

    ./zbox stats --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpg

### Repair
Use repair command to repair a file. By using help for this command, you will see it takes parameters:
* --allocation -- the allocation id from the newallocation command
* --localpath -- absolute path to the file on your local system
* --remote path -- remote path where you want to store. It should start with "/"

Command

    ./zbox repair --localpath <absolute path to file>/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac

Response

    Status completed callback. Type = application/octet-stream. Name = hello.txt

# Cancel allocation

Cancel allocation immediately and return all tokens from challenge pool
back to user (to write pool). This case blobber will not give their min
lock demand. If blobbers already got some tokens, the tokens will not be
returned.

It returns tokens to user's write pool, creating one expired if necessary.

    ./zbox alloc-cancel --allocation ALLOCATION_ID

# Finalize allocation

Finalize an expired allocation. When allocation expired, after its
challenge completion time (after the expiration) an allocation can be
finalized by owner or one of allocation blobbers.

      ./zbox alloc-fini --allocation ALLOCATION_ID

# Challenge pool information

Challenge pool brief information.

      ./zbox cp-info --allocation ALLOCATION_ID

Example

```
    BALANCE    |             START             |            EXPIRE             | FINIALIZED  
+--------------+-------------------------------+-------------------------------+------------+
  0.0039020457 | 2020-04-24 03:54:35 +0400 +04 | 2020-04-26 03:56:35 +0400 +04 | false       
```

Where time bounds and the finalization are allocations related. Balance is
current challenge pool balance.

# List blobbers

Show active blobbers in storage SC.

      ./zbox ls-blobbers

# Read pool

## Create

Create read pool if missing.

      ./zbox rp-create

## Info

Read pool information.

      ./zbox rp-info

Filtering result for an allocation.

      ./zbox rp-info --allocation ALLOCATION_ID

## Lock

Lock some tokens in read pool associated with an allocation. The tokens
will be divided between allocation blobbers by their read price.

      ./zbox rp-lock --allocation ALLOCATION_ID --duration 40m --tokens 1

## Unlock

Unlock an expired read pool by pool id. See `rp-info` for the POOL_ID and
the expiration.

      ./zbox rp-unlock --pool_id POOL_ID

# Storage SC configurations

Show storage SC configuration.

      ./zbox sc-config

# Stake pool

Stake pool related commands.


## Info

Stake pool information and settings.

      ./zbox sp-info --blobber_id BLOBBER_ID

## Lock

Lock crates delegate pool for current client and given blobber. The tokens
locked for the blobber stake and can be unlocked any time, excluding where
the tokens held by opened offers. The tokens collect interests.

      ./zbox sp-lock --blobber_id BLOBBER_ID --tokens 1.0

## Unlock

Unlock a stake pool by pool owner.

      ./zbox sp-unlock --blobber_id BLOBBER_ID --pool_id POOL_ID

## Pay interests

Some changes in stake pool pays all pending rewards to calculate next rewards
correctly and don't complicate stake pool. But if there is no changes interests
will not be payed. The `sp-pay-interests` can be used to pay the interests. The
command never fails. It pays interests for all delegates. Use `sp-info` to check
interests can be payed.

      ./zbox sp-pay-interests --blobber_id BLOBBER_ID

# Write pool

## Info

Write pool information.

For all write pools.

      ./zbox wp-info

Filtering by allocation.

      ./zbox wp-info --allocation ALLOCATION_ID       

## Lock

Lock tokens in a write pool associated with an allocation. All tokens will
be divided between allocation blobbers depending on their write price.

      ./zbox wp-lock --allocation ALLOCATION_ID --duration 40m --tokens 1

## Unlock

Unlock an expired write pool by its POOL_ID. See wp-info  for the id and the
expiration. An expired write pool, associated with an allocation, can be
locked until allocation finalization even if it's expired. It possible in
cases where related blobber doesn't give their min lock demands. The
finalization will pay the demand an unlock the pool.

      ./zbox wp-unlock --pool_id POOL_ID

# Costs

## Download cost.

How much a remote file downloading.

      ./zbox get-download-cost --allocation ALLOC_ID --remotepath /path/file.ext

Also, there are `authticket` and `lookuphash` flags to get the cost for
non allocation owners.

## Upload cost.

How much a local file uploading.

    ./zbox get-upload-cost --allocation ALLOC_ID --localpath ./path/file.ext

---
