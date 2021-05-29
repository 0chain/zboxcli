# zbox - a CLI for 0Chain dStorage

zbox is a command line interface (CLI) tool to understand the capabilities of 0Chain dStorage and prototype your app. The utility is built using 0Chain's goSDK library written in Go. Check out a [video](https://youtu.be/TPrkRjdaHrY) on how to use the CLI to create an allocation (storage volume) and upload, download, update, delete, and share files and folders on the 0Chain dStorage platform.

![Storage](https://user-images.githubusercontent.com/65766301/120052450-0ab66700-c043-11eb-91ab-1f7aa69e133a.png)

[zbox](https://github.com/0chain/zboxcli#Command-with-no-arguments) supports the following features

1. [Register a Wallet](https://github.com/0chain/zboxcli#Register)
2. [Get detailed Allocation](https://github.com/0chain/zboxcli#Get)
3. [Create an allocation](https://github.com/0chain/zboxcli#Create-new-allocation)
4. [List allocations](https://github.com/0chain/zboxcli#List-allocations)
5. [Update an allocation](https://github.com/0chain/zboxcli#Update-allocation)
6. [Cancel allocation](https://github.com/0chain/zboxcli#Cancel-allocation)
7. [Finalize allocation](https://github.com/0chain/zboxcli#Finalize-allocation)
8. [List blobbers](https://github.com/0chain/zboxcli#List-blobbers)
9. [Detail blobber information](https://github.com/0chain/zboxcli#Detailed-blobber-information)
10. [Update blobber settings](https://github.com/0chain/zboxcli#Update-blobber-settings)
11. [Upload a file to dStorage](https://github.com/0chain/zboxcli#Upload)
12. [Download the uploaded file from dStorage](https://github.com/0chain/zboxcli#Download)
13. [Update the uploaded file on dStorage](https://github.com/0chain/zboxcli#Update)
14. [Delete the uploaded file on dStorage](https://github.com/0chain/zboxcli#Delete)
15. [Share the uploaded file on dStorage to the public](https://github.com/0chain/zboxcli#Share)
16. [List the uploaded files and folders](https://github.com/0chain/zboxcli#List)
17. [Copy uploaded files to another folder path on dStorage](https://github.com/0chain/zboxcli#Copy)
18. [Move uploaded files to another folder path on dStorage](https://github.com/0chain/zboxcli#Move)
19. [Get meta data of files](https://github.com/0chain/zboxcli#Get-metadata)
20. [Rename an object in allocation](https://github.com/0chain/zboxcli#Rename)
21. [Get file stats](https://github.com/0chain/zboxcli#Stats)
22. [Repair a file on dStorage](https://github.com/0chain/zboxcli#Repair)
23. [Sync your local folder to remote](https://github.com/0chain/zboxcli#Sync)
24. [Update file attributes](https://github.com/0chain/zboxcli#Update-file-attributes)
25. [Get wallet information](https://github.com/0chain/zboxcli#Get-wallet)
26. [Add Collaborator for a file](https://github.com/0chain/zboxcli#Add-collaborator)
27. [Remove Collaborator for a file](https://github.com/0chain/zboxcli#Delete-collaborator)
28. [Challenge pool information](https://github.com/0chain/zboxcli#Challenge-pool-information)
29. [Create read pool if not exists](https://github.com/0chain/zboxcli#Create-read-pool)
30. [Detailed read pool information](https://github.com/0chain/zboxcli#Read-pool-info)
31. [Lock tokens into read pool](https://github.com/0chain/zboxcli#Lock-tokens-into-read-pool)
32. [Unlock tokens from expired read pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-read-pool)
33. [Storage SC configurations](https://github.com/0chain/zboxcli#Storage-SC-configurations)
34. [Detailed stake pool information](https://github.com/0chain/zboxcli#Stake-pool-info)
35. [Lock tokens into stake pool](https://github.com/0chain/zboxcli#Lock-tokens-into-stake-pool)
36. [Unlock tokens from expired stake pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-stake-pool)
37. [Stake pools info of current user](https://github.com/0chain/zboxcli#Stake-pools-info-of-current-user)
38. [Pay interests](https://github.com/0chain/zboxcli#Pay-interests)
39. [Detailed write pool information](https://github.com/0chain/zboxcli#Write-pool-info)
40. [Lock tokens into write pool](https://github.com/0chain/zboxcli#Lock-tokens-into-write-pool)
41. [Unlock tokens from expired write pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-write-pool)
42. [Get download cost](https://github.com/0chain/zboxcli#Download-cost)
43. [Get upload cost](https://github.com/0chain/zboxcli#Upload-cost)

zbox CLI provides a self-explaining "help" option that lists commands and parameters they need to perform the intended action

## Getting started with zbox

## Pre-requisites

```
Go V1.12 or higher.
```

## Installation Guides

### [How to install on Linux](https://github.com/0chain/zboxcli/wiki/Build-Linux)

### [How to install on Windows](https://github.com/0chain/zboxcli/wiki/Build-Windows)

### [Other Platform Builds](https://github.com/0chain/zboxcli/wiki/Alternative-Platform-Builds)

### Use custom miner/sharder 

As mentioned in build guides, a ./zcn folder is created to store configuration files for zboxcli. Here is a sample network config file

```
  ---
  block_worker: http://localhost:9091
  signature_scheme: bls0chain
  min_submit: 50 # in percentage
  min_confirmation: 50 # in percentage
  confirmation_chain_length: 3
```

A blockWorker is used to connect to the network instead of giving network details directly, It will fetch the network details automatically from the blockWorker's network API. By default it will use the miner/sharder values which it will get using the `block_worker_url/network`. In case you want to override those values and give custom miner/sharder to use, You have to create a `network.yaml` in your ~/.zcn (config) folder and paste the miner/sharder values in below format.

```
miners:
  - http://localhost:7071
  - http://localhost:7072
  - http://localhost:7073
sharders:
  - http://localhost:7171
```

Note: This is helpful for the Mac OS users running local cluster and having trouble with docker internal IPs (block_worker return docker IPs in local)

## Running zbox commands

Note in this document, we will only show the commands for particular functionalities, the response will vary depending on your usage and may not be provided in all places. To get a more descriptive view of all the zbox functionalities check zbox cli documentation at docs.0chain.net.

### Command with no arguments

When you run `zbox` with no arguments, it will list all the supported commands.

```
./zbox 
zbox is a decentralized storage application written on the 0Chain platform.
                        Complete documentation is available at https://docs.0chain.net/0chain/

Usage:
  zbox [command]

Available Commands:
  add-collab        add collaborator for a file
  alloc-cancel      Cancel an allocation
  alloc-fini        Finalize an expired allocation
  bl-info           Get blobber info
  bl-update         Update blobber settings by its delegate_wallet owner
  commit            commit a file changes to chain
  copy              copy an object(file/folder) to another folder on blobbers
  cp-info           Challenge pool information.
  delete            delete file from blobbers
  delete-collab     delete collaborator for a file
  download          download file from blobbers
  get               Gets the allocation info
  get-diff          Get difference of local and allocation root
  get-download-cost Get downloading cost
  get-upload-cost   Get uploading cost
  getwallet         Get wallet information
  help              Help about any command
  list              list files from blobbers
  list-all          list all files from blobbers
  listallocations   List allocations for the client
  ls-blobbers       Show active blobbers in storage SC.
  meta              get meta data of files from blobbers
  move              move an object(file/folder) to another folder on blobbers
  newallocation     Creates a new allocation
  register          Registers the wallet with the blockchain
  rename            rename an object(file/folder) on blobbers
  rp-create         Create read pool if missing
  rp-info           Read pool information.
  rp-lock           Lock some tokens in read pool.
  rp-unlock         Unlock some expired tokens in a read pool.
  sc-config         Show storage SC configuration.
  share             share files from blobbers
  sign-data         Sign given data
  sp-info           Stake pool information.
  sp-lock           Lock tokens lacking in stake pool.
  sp-pay-interests  Pay interests not payed yet.
  sp-unlock         Unlock tokens in stake pool.
  sp-user-info      Stake pool information for a user.
  start-repair      start repair file to blobbers
  stats             stats for file from blobbers
  sync              Sync files to/from blobbers
  update            update file to blobbers
  update-attributes update object attributes on blobbers
  updateallocation  Updates allocation's expiry and size
  upload            upload file to blobbers
  version           Prints version information
  wp-info           Write pool information.
  wp-lock           Lock some tokens in write pool.
  wp-unlock         Unlock some expired tokens in a write pool.

Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
  -h, --help                       help for zbox
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key

Use "zbox [command] --help" for more information about a command.

```

### Register

Command `register`  creates a wallet if not created and registers it for use by the blockchain and blobbers. The wallet is created in the ~/.zcn directory and uses the keys stored in ~/.zcn/wallet.json. You can create multiple wallets with multiple allocations but make sure to specify a wallet file for every wallet.

#### Usage

```
./zbox register -h
Registers the wallet with the blockchain

Usage:
  zbox register [flags]

Flags:
  -h, --help   help for register

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox register
```

Response:

```
ZCN wallet created
Wallet registered
```

### Create new allocation

Command `newallocation` reserves hard disk space on the blobbers. Let's see the parameters it takes by using `--help`.

![allocation](https://user-images.githubusercontent.com/65766301/120052477-27529f00-c043-11eb-91bb-573558325b20.png)

#### Usage

```
./zbox newallocation --help

Creates a new allocation
Usage:
  zbox newallocation [flags]

Flags:
      --allocationFileName string   --allocationFileName allocation.txt (default "allocation.txt")
      --cost                        pass this option to only get the min lock demand
      --data int                    --data 2 (default 2)
      --expire duration             duration to allocation expiration (default 720h0m0s)
  -h, --help                        help for newallocation
      --lock float                  lock write pool with given number of tokens, required
      --mcct duration               max challenge completion time, optional, default 1h (default 1h0m0s)
      --parity int                  --parity 2 (default 2)
      --read_price string           select blobbers by provided read price range, use form 0.5-1.5, default is [0; inf)
      --size int                    --size 10000 (default 2147483648)
      --usd                         pass this option to give token value in USD
      --write_price string          select blobbers by provided write price range, use form 1.5-2.5, default is [0; inf)

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key


```

As you can see the `newallocation` command takes allocationFileName where the volume information is stored locally. All the parameters have default values. With more data shards, you can upload or download files faster. With more parity shards, you have higher availability.

#### Example

To create a new allocation with default values,use `newallocation` with a `--lock` flag to add some tokens to the write pool .On success a related write pool is created and the allocation information is stored under `$HOME/.zcn/allocation.txt`.

```
./zbox newallocation --lock 0.5
```

Response:

```
Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

### Update allocation

Command `updateallocation` updates hard disk space and expiry on the blobbers. Let's see the parameters it takes by using `--help` flag..

#### Usage

```
./zbox updateallocation -h
Updates allocation's expiry and size

Usage:
  zbox updateallocation [flags]

Flags:
      --allocation string   Allocation ID
      --expiry duration     adjust storage expiration time, duration
  -h, --help                help for updateallocation
      --lock float          lock write pool with given number of tokens, required
      --size int            adjust allocation size, bytes

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

Update an allocation for different storage expiration time, and  allocation size(in bytes).

```
./zbox updateallocation --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --expiry 48h --size 4096
```

Response:

```
Allocation updated with txId : fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d
```

You can see more txn details using above txID in block explorer [here](https://one.devnet-0chain.net/).

### Cancel allocation

Cancel allocation immediately return all tokens from challenge pool back to user (to write pool) and cancels the allocation. In this case blobber will not give their min lock demand. If blobbers already got some tokens, the tokens will not be returned.

#### Usage

```
./zbox alloc-cancel -h
Cancel allocation used to terminate an allocation where, because
of blobbers, it can't be used. Thus, the blobbers will not receive their
min_lock_demand. Other aspects of the cancellation follows the finalize
allocation flow.

Usage:
  zbox alloc-cancel [flags]

Flags:
      --allocation string   Allocation ID
  -h, --help                help for alloc-cancel

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox alloc-cancel --allocation <allocation_id>
```

### Finalize allocation

Finalize an expired allocation. When an allocation is expired, after its challenge completion time (after the expiration), it can be finalized by the owner or one of the allocation blobbers.

#### Usage

```
./zbox alloc-fini -h
Finalize an expired allocation by allocation owner or one of
blobbers of the allocation. It moves all tokens have to be moved between pools
and empties write pool moving left tokens to client.

Usage:
  zbox alloc-fini [flags]

Flags:
      --allocation string   Allocation ID
  -h, --help                help for alloc-fini

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox alloc-fini --allocation <allocation_id>
```

### List blobbers

Use `ls-blobbers` command to show active blobbers in storage SC.

#### Usage

```
./zbox ls-blobbers -h
Show active blobbers in storage SC.

Usage:
  zbox ls-blobbers [flags]

Flags:
  -h, --help   help for ls-blobbers
      --json   pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox ls-blobbers
- id:                    0ece681f6b00221c5567865b56040eaab23795a843ed629ce71fb340a5566ba3
  url:                   http://five.devnet-0chain.net:31302
  used / total capacity: 101.7 GiB / 1000.0 GiB
  terms:
    read_price:          0.01 tok / GB
    write_price:         0.01 tok / GB / time_unit
    min_lock_demand:     0.1
    cct:                 2m0s
    max_offer_duration:  744h0m0s
- id:                    788b1deced159f12d3810c61b4b8d381e80188c470e9798939f2e5036d964ffc
  url:                   http://five.devnet-0chain.net:31301
  used / total capacity: 102.7 GiB / 1000.0 GiB
  terms:
    read_price:          0.01 tok / GB
    write_price:         0.01 tok / GB / time_unit
    min_lock_demand:     0.1
    cct:                 2m0s
    max_offer_duration:  744h0m0s
```

### Detailed blobber information

Use `bl-info` command to get detailed blobber information.

#### Usage

```
./zbox bl-info --h
Get blobber info

Usage:
  zbox bl-info [flags]

Flags:
      --blobber_id string   blobber ID, required
  -h, --help                help for bl-info
      --json                pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox bl-info --blobber_id f65af5d64000c7cd2883f4910eb69086f9d6e6635c744e62afcfab58b938ee25
```

Response:

```
id:                f65af5d64000c7cd2883f4910eb69086f9d6e6635c744e62afcfab58b938ee25
url:               http://zerominer.xyz:5051
capacity:          1.0 GiB
last_health_check: 2021-04-08 22:54:50 +0700 +07
capacity_used:     0 B
terms:
  read_price:         0.01 tok / GB
  write_price:        0.1 tok / GB
  min_lock_demand:    10 %
  max_offer_duration: 744h0m0s
  cct:                2m0s
settings:
  delegate_wallet: 8b87739cd6c966c150a8a6e7b327435d4a581d9d9cc1d86a88c8a13ae1ad7a96
  min_stake:       1 tok
  max_stake:       100 tok
  num_delegates:   50
  service_charge:  30 %
```

### Update blobber settings

Use `./zbox bl-update --help` to get list of settings that can be updated.

#### Usage

```
./zbox bl-update -h
Update blobber settings by its delegate_wallet owner

Usage:
  zbox bl-update [flags]

Flags:
      --blobber_id string             blobber ID, required
      --capacity int                  update blobber capacity bid, optional
      --cct duration                  update challenge completion time (cct), optional
  -h, --help                          help for bl-update
      --max_offer_duration duration   update max_offer_duration, optional
      --max_stake float               update max_stake, optional
      --min_lock_demand float         update min_lock_demand, optional
      --min_stake float               update min_stake, optional
      --num_delegates int             update num_delegates, optional
      --read_price float              update read_price, optional
      --service_charge float          update service_charge, optional
      --write_price float             update write_price, optional

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

Update blobber read price

```
./zbox bl-update --blobber_id 0ece681f6b00221c5567865b56040eaab23795a843ed629ce71fb340a5566ba3 --read_price 0.1
```

### Upload

Use `upload` command to upload a file.

#### Usage

```
./zbox upload -h
upload file to blobbers

Usage:
  zbox upload [flags]

Flags:
      --allocation string                Allocation ID
      --attr-who-pays-for-reads string   Who pays for reads: owner or 3rd_party (default "owner")
      --commit                           pass this option to commit the metadata transaction
      --encrypt                          pass this option to encrypt and upload the file
  -h, --help                             help for upload
      --localpath string                 Local path of file to upload
      --remotepath string                Remote path to upload
      --thumbnailpath string             Local thumbnail path of file to upload

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

**Upload file with no encryption**

```
./zbox upload --localpath /absolute-path-to-local-file/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

Response:

```
12390 / 12390 [================================================================================] 100.00% 3s
Status completed callback. Type = application/octet-stream. Name = hello.txt
```

**Upload file with encryption**

Use upload command with optional encrypt parameter to upload a file in encrypted format. This can be downloaded as normal from same wallet/allocation or utilize Proxy Re-Encryption facility (see [download](https://github.com/0chain/zboxcli#Download) command).

```
./zbox upload --encrypt --localpath <absolute path to file>/sensitivedata.txt --remotepath /myfiles/sensitivedata.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

Response:

```
12390 / 12390 [================================================================================] 100.00% 3s
Status completed callback. Type = application/octet-stream. Name = sensitivedata.txt
```



### Download

Use `download` command to download your own or a shared file.

#### Usage

```
./zbox download -h
download file from blobbers

Usage:
  zbox download [flags]

Flags:
      --allocation string     Allocation ID
      --authticket string     Auth ticket fot the file to download if you dont own it
  -b, --blockspermarker int   pass this option to download multiple blocks per marker (default 10)
      --commit                pass this option to commit the metadata transaction
  -e, --endblock int          pass this option to download till specific block number
  -h, --help                  help for download
      --localpath string      Local path of file to download
      --lookuphash string     The remote lookuphash of the object retrieved from the list
      --remotepath string     Remote path to download
      --rx_pay                used to download by authticket; pass true to pay for download yourself
  -s, --startblock int        pass this option to download from specific block number
  -t, --thumbnail             pass this option to download only the thumbnail

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key

```

#### Example

```
./zbox download --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg --localpath ../horse.jpeg
```

Response:

```
4 / 4 [=======================================================================] 100.00% 3s
Status completed callback. Type = application/octet-stream. Name = horse.jpeg
```

Note: You can download by using only 1 on the below combination:

- `--remotepath`, `--allocation`
- `--authticket`

Downloaded file will be in the localpath specified.

### Update

Use `update` command to update content of an existing file in the remote path. Similar to [upload](https://github.com/0chain/zboxcli#Upload) command.

### Delete

Use `delete` command to delete your file on the allocation.

#### Usage

```
./zbox delete -h
delete file from blobbers

Usage:
  zbox delete [flags]

Flags:
      --allocation string   Allocation ID
      --commit              pass this option to commit the metadata transaction
  -h, --help                help for delete
      --remotepath string   Remote path of the object to delete

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox delete --allocation 3c0d32560ea18d9d0d76808216a9c634flist661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg
```

Response:

```
/myfiles/horse.jpeg deleted
```


File successfully deleted (Can be verified using [list](https://github.com/0chain/zboxcli#List))

### Share

Use `share` command to generate an authtoken that provides authorization to the holder to the specified file on the remotepath.

#### Usage

```
./zbox share -h
share files from blobbers

Usage:
  zbox share [flags]

Flags:
      --allocation string            Allocation ID
      --clientid string              ClientID of the user to share with. Leave blank for public share
      --encryptionpublickey string   Encryption public key of the client you want to share with. Can be retrieved by the getwallet command
  -h, --help                         help for share
      --remotepath string            Remote path to share

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

**Public share**

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt
```

Response:

```
Auth token :eyJjbGllbnRfaWQiOiIiLCJvd25lcl9pZCI6ImI2ZGU1NjJiNTdhMGI1OTNkMDQ4MDYyNGY3OWE1NWVkNDZkYmE1NDQ0MDQ1OTViZWUwMjczMTQ0ZTAxMDM0YWUiLCJhbGxvY2F0aW9uX2lkIjoiODY5NWI5ZTdmOTg2ZDRhNDQ3YjY0ZGUwMjBiYTg2ZjUzYjNiNWUyYzQ0MmFiY2ViNmNkNjU3NDI3MDIwNjdkYyIsImZpbGVfcGF0aF9oYXNoIjoiMjBkYzc5OGIwNGViYWIzMDE1ODE3Yzg1ZDIyYWVhNjRhNTIzMDViYWQ2Zjc0NDlhY2QzODI4YzhkNzBjNzZhMyIsImZpbGVfbmFtZSI6IjEudHh0IiwicmVmZXJlbmNlX3R5cGUiOiJmIiwiZXhwaXJhdGlvbiI6MTYyNjQyMDM1OSwidGltZXN0YW1wIjoxNjE4NjQ0MzU5LCJyZV9lbmNyeXB0aW9uX2tleSI6IiIsInNpZ25hdHVyZSI6ImFjNzIzZjdhMWQ0ZDBmMjc2ZmQ3Yzc2NWMxOTcyZTlhODc2OGI0MjU1ODkyMmMwNjEyZjMxNjBjMGZiODQ5MGMifQ==
```


**Encrypted share**

Use clientid and encryptionpublickey of the user to share with.
![Private File Sharing](https://user-images.githubusercontent.com/65766301/120052575-962ff800-c043-11eb-9cf7-433383d532a3.png)

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt --clientid b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae --encryptionpublickey 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=
```

Response:

```
Auth token :eyJjbGllbnRfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwib3duZXJfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwiYWxsb2NhdGlvbl9pZCI6Ijg2OTViOWU3Zjk4NmQ0YTQ0N2I2NGRlMDIwYmE4NmY1M2IzYjVlMmM0NDJhYmNlYjZjZDY1NzQyNzAyMDY3ZGMiLCJmaWxlX3BhdGhfaGFzaCI6IjIwZGM3OThiMDRlYmFiMzAxNTgxN2M4NWQyMmFlYTY0YTUyMzA1YmFkNmY3NDQ5YWNkMzgyOGM4ZDcwYzc2YTMiLCJmaWxlX25hbWUiOiIxLnR4dCIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MjY0MjA1NzQsInRpbWVzdGFtcCI6MTYxODY0NDU3NCwicmVfZW5jcnlwdGlvbl9rZXkiOiJ7XCJyMVwiOlwiOUpnci9aVDh6VnpyME1BcWFidlczdnhoWEZoVkdMSGpzcVZtVUQ1QTJEOD1cIixcInIyXCI6XCIrVEk2Z1pST3JCR3ZURG9BNFlicmNWNXpoSjJ4a0I4VU5SNTlRckwrNUhZPVwiLFwicjNcIjpcInhySjR3bENuMWhqK2Q3RXU5TXNJRzVhNnEzRXVzSlZ4a2N6YXN1K0VqQW89XCJ9Iiwic2lnbmF0dXJlIjoiZTk3NTYyOTAyODU4OTBhY2QwYTcyMzljNTFhZjc0YThmNjU2OTFjOTUwMzRjOWM0ZDJlMTFkMTQ0MTk0NmExYSJ9
```


Response contains an auth token- an encrypted string that can be shared.

### List

Use `list` command to list files in given remote path of the dStorage.

#### Usage

```
./zbox list -h
list files from blobbers

Usage:
  zbox list [flags]

Flags:
      --allocation string   Allocation ID
      --authticket string   Auth ticket fot the file to download if you dont own it
  -h, --help                help for list
      --json                pass this option to print response as json data
      --lookuphash string   The remote lookuphash of the object retrieved from the list
      --remotepath string   Remote path to list from

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox list --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /  
```

Response:

```
Auth token :eyJjbGllbnRfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwib3duZXJfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwiYWxsb2NhdGlvbl9pZCI6Ijg2OTViOWU3Zjk4NmQ0YTQ0N2I2NGRlMDIwYmE4NmY1M2IzYjVlMmM0NDJhYmNlYjZjZDY1NzQyNzAyMDY3ZGMiLCJmaWxlX3BhdGhfaGFzaCI6IjIwZGM3OThiMDRlYmFiMzAxNTgxN2M4NWQyMmFlYTY0YTUyMzA1YmFkNmY3NDQ5YWNkMzgyOGM4ZDcwYzc2YTMiLCJmaWxlX25hbWUiOiIxLnR4dCIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MjY0MjA1NzQsInRpbWVzdGFtcCI6MTYxODY0NDU3NCwicmVfZW5jcnlwdGlvbl9rZXkiOiJ7XCJyMVwiOlwiOUpnci9aVDh6VnpyME1BcWFidlczdnhoWEZoVkdMSGpzcVZtVUQ1QTJEOD1cIixcInIyXCI6XCIrVEk2Z1pST3JCR3ZURG9BNFlicmNWNXpoSjJ4a0I4VU5SNTlRckwrNUhZPVwiLFwicjNcIjpcInhySjR3bENuMWhqK2Q3RXU5TXNJRzVhNnEzRXVzSlZ4a2N6YXN1K0VqQW89XCJ9Iiwic2lnbmF0dXJlIjoiZTk3NTYyOTAyODU4OTBhY2QwYTcyMzljNTFhZjc0YThmNjU2OTFjOTUwMzRjOWM0ZDJlMTFkMTQ0MTk0NmExYSJ9
```

Response will be a list with information for each file/folder in the given path. The information includes lookuphash which is require for download via authticket.

### Copy

Use `copy` command to copy file to another folder path in dStorage.

#### Usage

```
./zbox copy -h
copy an object to another folder on blobbers

Usage:
  zbox copy [flags]

Flags:
      --allocation string   Allocation ID
      --commit              pass this option to commit the metadata transaction
      --destpath string     Destination path for the object. Existing directory the object should be copied to
  -h, --help                help for copy
      --remotepath string   Remote path of object to copy

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox copy --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath 
```

Response:

```
/file.txt --destpath /existingFolder
/file.txt copied
```

### Move

Use `move` command to move file to another remote folder path on dStorage.

```
./zbox move -h
move an object to another folder on blobbers

Usage:
  zbox move [flags]

Flags:
      --allocation string   Allocation ID
      --commit              pass this option to commit the metadata transaction
      --destpath string     Destination path for the object. Existing directory the object should be moved to
  -h, --help                help for move
      --remotepath string   Remote path of object to move

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox move --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath 
```

Response:

```
/file.txt --destpath /existingFolder
/file.txt moved
```

### List allocations

Use `listallocations` command to list all allocations for the client.

#### Usage

```
./zbox listallocations -h
List allocations for the client

Usage:
  zbox listallocations [flags]

Flags:
  -h, --help   help for listallocations
      --json   pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox listallocations
```

Response:

```
                                 ID                                |    SIZE    |          EXPIRATION           | DATASHARDS | PARITYSHARDS | FINALIZED | CANCELED |   R  PRICE   |   W  PRICE    
+------------------------------------------------------------------+------------+-------------------------------+------------+--------------+-----------+----------+--------------+--------------+
  8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc | 6442450944 | 2021-05-24 00:27:23 +0700 +07 |          4 |            2 | false     | false    | 0.0599999994 | 0.0599999994  
```

### Sync

sync command helps in syncing all files from the local folder recursively to the remote.

#### Usage

```
./zbox sync -h
Sync all files to/from blobbers from/to a localpath

Usage:
  zbox sync [flags]

Flags:
      --allocation string         Allocation ID
      --commit                    pass this option to commit the metadata transaction - only works with uploadonly
      --encryptpath string        Local dir path to upload as encrypted
      --excludepath stringArray   Remote folder paths exclude to sync
  -h, --help                      help for sync
      --localcache string         Local cache of remote snapshot.
                                  If file exists, this will be used for comparison with remote.
                                  After sync complete, remote snapshot will be updated to the same file for next use.
      --localpath string          Local dir path to sync
      --uploadonly                pass this option to only upload/update the files

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox sync --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --localpath /home/dung/Desktop/alloc --localcache /home/dung/Desktop/localcache.json
```

Response:

```
  OPERATION |      PATH       
+-----------+----------------+
  Download  | /1.txt          
  Download  | /afolder/1.txt  
  Download  | /d2.txt         

 4 / 4 [===========================================================================] 100.00% 0s
Status completed callback. Type = application/octet-stream. Name = 1.txt
 4 / 4 [===========================================================================] 100.00% 0s
Status completed callback. Type = application/octet-stream. Name = 1.txt
 7 / 7 [===========================================================================] 100.00% 0s
Status completed callback. Type = application/octet-stream. Name = d2.txt

Sync Complete
Local cache saved.
```

It will sync your localpath with the remote and do all the required CRUD operations.

### Update file attributes

Use `update-attributes` command to update file attributes. Only one attribute is currently supported: who-pays-for-reads that can be:

- `owner`, where allocation owner pays for own and 3rd_party reads
- `3rd_party`, where 3rd party readers pays for their downloads themselves

#### Usage

```
./zbox update-attributes -h
update object attributes on blobbers

Usage:
  zbox update-attributes [flags]

Flags:
      --allocation string           Allocation ID
      --commit                      pass this option to commit the metadata transaction
  -h, --help                        help for update-attributes
      --remotepath string           Remote path of object to rename
      --who-pays-for-reads string   Who pays for reads: owner or 3rd_party (default "owner")

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox update-attributes --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --who-pays-for-reads 3rd_party

```

Response:

```
attributes updated
```

### Get wallet

Use `getwallet` command to get additional wallet information including Encryption Public Key,Client ID which are required for Private File Sharing.

#### Usage

```
./zbox getwallet -h
Get wallet information

Usage:
  zbox getwallet [flags]

Flags:
  -h, --help   help for getwallet
      --json   pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox getwallet
```

Response:

```
                                                             PUBLIC KEY                                                            |                             CLIENTID                             |            ENCRYPTION PUBLIC KEY              
+----------------------------------------------------------------------------------------------------------------------------------+------------------------------------------------------------------+----------------------------------------------+
  3e7fa2dd6b924adfdf69e36cc61cb5d9012226dac619250ce5fc37eae25a05118008944d1727221b6d14b0998c5813acd13066040976598c5366e86519377001 | b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae | 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=  
```

Response will give details for current selected wallet (or wallet file specified by optional --wallet parameter)

### Get

Use `get` command to get the information about the allocation such as total size , used size, number of challenges and challenges passed/failed/open/redeemed.

#### Usage

```
./zbox get -h
Gets the allocation info

Usage:
  zbox get [flags]

Flags:
      --allocation string   Allocation ID
  -h, --help                help for get
      --json                pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox get --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc
```

Response:

```
allocation:
  id:              8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc
  tx:              026c9d331e9c93aee4f3247507c20bdd4b7429bd81d27845bfab83f9c9c082e6 (latest create/update allocation transaction hash)
  data_shards:     4
  parity_shards:   2
  size:            6.0 GiB
  expiration_date: 2021-05-24 00:27:23 +0700 +07
  blobbers:
    - blobber_id:       dea18e3f3c308666cb489877b9b2c7e2babf797d8b8c322fa9d074105787a9e9
      base URL:         http://five.devnet-0chain.net:31304
      size:             1.0 GiB
      min_lock_demand:  0.0012333333
      spent:            0.0000244839 (moved to challenge pool or to the blobber)
      penalty:          0 (blobber stake slash)
      read_reward:      0.000024414
      returned:         0.0000000233 (on challenge failed)
      challenge_reward: 0 (on challenge passed)
      final_reward:     0 (if finalized)
      terms: (allocation related terms)
        read_price:                0.0099999999 tok / GB (by 64KB chunks)
        write_price:               0.0099999999 tok / GB
        min_lock_demand:           10 %
        max_offer_duration:        744h0m0s
        challenge_completion_time: 2m0s
```

### Get metadata

Use `meta` command to get meta data for a given remote file. Use `-h` to know more about possible flags.

#### Usage

```
./zbox meta -h
get meta data of files from blobbers

Usage:
  zbox meta [flags]

Flags:
      --allocation string   Allocation ID
      --authticket string   Auth ticket fot the file to download if you dont own it
  -h, --help                help for meta
      --json                pass this option to print response as json data
      --lookuphash string   The remote lookuphash of the object retrieved from the list
      --remotepath string   Remote path to list from

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

**Without any authticket**

```
./zbox meta --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt
```

Response:

```
  TYPE | NAME  |  PATH  |                           LOOKUP HASH                            | SIZE |        MIME TYPE         |                   HASH                    
+------+-------+--------+------------------------------------------------------------------+------+--------------------------+------------------------------------------+
  f    | 1.txt | /1.txt | 20dc798b04ebab3015817c85d22aea64a52305bad6f7449acd3828c8d70c76a3 |    4 | application/octet-stream | 03cfd743661f07975fa2f1220c5194cbaff48451 
```

**With authticket **

```
./zbox meta --lookuphash 20dc798b04ebab3015817c85d22aea64a52305bad6f7449acd3828c8d70c76a3 --authticket eyJjbGllbnRfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwib3duZXJfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwiYWxsb2NhdGlvbl9pZCI6Ijg2OTViOWU3Zjk4NmQ0YTQ0N2I2NGRlMDIwYmE4NmY1M2IzYjVlMmM0NDJhYmNlYjZjZDY1NzQyNzAyMDY3ZGMiLCJmaWxlX3BhdGhfaGFzaCI6IjIwZGM3OThiMDRlYmFiMzAxNTgxN2M4NWQyMmFlYTY0YTUyMzA1YmFkNmY3NDQ5YWNkMzgyOGM4ZDcwYzc2YTMiLCJmaWxlX25hbWUiOiIxLnR4dCIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MjY0MjA1NzQsInRpbWVzdGFtcCI6MTYxODY0NDU3NCwicmVfZW5jcnlwdGlvbl9rZXkiOiJ7XCJyMVwiOlwiOUpnci9aVDh6VnpyME1BcWFidlczdnhoWEZoVkdMSGpzcVZtVUQ1QTJEOD1cIixcInIyXCI6XCIrVEk2Z1pST3JCR3ZURG9BNFlicmNWNXpoSjJ4a0I4VU5SNTlRckwrNUhZPVwiLFwicjNcIjpcInhySjR3bENuMWhqK2Q3RXU5TXNJRzVhNnEzRXVzSlZ4a2N6YXN1K0VqQW89XCJ9Iiwic2lnbmF0dXJlIjoiZTk3NTYyOTAyODU4OTBhY2QwYTcyMzljNTFhZjc0YThmNjU2OTFjOTUwMzRjOWM0ZDJlMTFkMTQ0MTk0NmExYSJ9
  
```

Response:

```
TYPE | NAME  |                           LOOKUP HASH                            | SIZE |        MIME TYPE         |                   HASH                    
+------+-------+------------------------------------------------------------------+------+--------------------------+------------------------------------------+
  f    | 1.txt | 20dc798b04ebab3015817c85d22aea64a52305bad6f7449acd3828c8d70c76a3 |    4 | application/octet-stream | 03cfd743661f07975fa2f1220c5194cbaff48451  
```


Response will be meta data for the given filepath/lookuphash (if using authTicket)

### Rename

`rename` command helps in renaming a file existing already on dStorage.

#### Usage

```
./zbox rename -h
rename an object on blobbers

Usage:
  zbox rename [flags]

Flags:
      --allocation string   Allocation ID
      --commit              pass this option to commit the metadata transaction
      --destname string     New Name for the object (Only the name and not the path). Include the file extension if applicable
  -h, --help                help for rename
      --remotepath string   Remote path of object to rename

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox rename --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --destname x.txt
```

Response:

```
/1.txt renamed
```

### Stats

`stats` command helps in getting upload, download and challenge information for a file.

```
./zbox stats -h
stats for file from blobbers

Usage:
  zbox stats [flags]

Flags:
      --allocation string   Allocation ID
  -h, --help                help for stats
      --json                pass this option to print response as json data
      --remotepath string   Remote path to list from

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox stats --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt
```

Response:

```
                              BLOBBER                              | NAME  |  PATH  | SIZE | UPLOADS | BLOCK DOWNLOADS | CHALLENGES | BLOCKCHAIN AWARE  
+------------------------------------------------------------------+-------+--------+------+---------+-----------------+------------+------------------+
  9c14598d5d39cb27177add6efabdadfb0a0478abe5d471ffe9080751dc89321c | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
  dea18e3f3c308666cb489877b9b2c7e2babf797d8b8c322fa9d074105787a9e9 | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
  0ece681f6b00221c5567865b56040eaab23795a843ed629ce71fb340a5566ba3 | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
  788b1deced159f12d3810c61b4b8d381e80188c470e9798939f2e5036d964ffc | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
  78a1a9db859cded21a5120a5bf808e97202a1fd7f94e51d2fd174edbdc4d7291 | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
  876b4cd610eb1aac63c53cdfd4d3a0ac91d94f2d6b858bb195f72b6dc0f33b55 | 1.txt | /1.txt | 2065 |       3 |               1 |          0 | true              
```

### Repair

Use `start-repair` command to repair a file on dStorage.
![repair](https://user-images.githubusercontent.com/65766301/120052600-b364c680-c043-11eb-9bf2-038ab244fed6.png)
\
#### Usage

```
./zbox start-repair -h
start repair file to blobbers

Usage:
  zbox start-repair [flags]

Flags:
      --allocation string   Allocation ID
  -h, --help                help for start-repair
      --repairpath string   Path to repair
      --rootpath string     File path for local files

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox start-repair --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --repairpath / --rootpath /home/dung/Desktop/alloc
```

Response:

```
Repair file completed, Total files repaired:  0
```

### Add collaborator

Use `add-collab` command to add a collaborator for a file on dStorage.
![collaboration](https://user-images.githubusercontent.com/65766301/120052678-0f2f4f80-c044-11eb-8ca6-1a032659eac3.png)
#### Usage

```
./zbox add-collab -h
add collaborator for a file

Usage:
  zbox add-collab [flags]

Flags:
      --allocation string   Allocation ID
      --collabid string     Collaborator's clientID
  -h, --help                help for add-collab
      --remotepath string   Remote path to list from

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox add-collab --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --collabid d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329
```

Response will be a confirmation that collaborator is added on all blobbers for the given file .

```
Collaborator d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329 added successfully for the file /1.txt
```

You can check all collaborators for a file in metadata json response.

### Delete collaborator

Use command delete-collab to remove a collaborator for a file

#### Usage

```
./zbox delete-collab -h
delete collaborator for a file

Usage:
  zbox delete-collab [flags]

Flags:
      --allocation string   Allocation ID
      --collabid string     Collaborator's clientID
  -h, --help                help for delete-collab
      --remotepath string   Remote path to list from

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox delete-collab --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --collabid d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329
```

Response will be a confirmation that collaborator is removed on all blobbers for the given file.

```
Collaborator d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329 removed successfully for the file /1.txt
```

### Challenge pool information

Use `cp-info` command to get the challenge pool brief information.

#### Usage

```
./zbox cp-info -h
Challenge pool information.

Usage:
  zbox cp-info [flags]

Flags:
      --allocation string   allocation identifier, required
  -h, --help                help for cp-info
      --json                pass this option to print response as json data

Global Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --verbose                    prints sdk log in stderr (default false)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
```

#### Example

```
./zbox cp-info --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc   
```

Response:

```
POOL ID: 6dba10422e368813802877a85039d3985d96760ed844092319743fb3a76712d7:challengepool:8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc
    BALANCE    |             START             |            EXPIRE             | FINIALIZED  
+--------------+-------------------------------+-------------------------------+------------+
  0.0000002796 | 2021-04-17 00:27:23 +0700 +07 | 2021-05-24 00:29:23 +0700 +07 | false    
```


 Balance is the current challenge pool balance. Start,Expire time and the finalization are allocations related.

### Create read pool

Use `rp-create` to create a read pool.

```
./zbox rp-create
```

### Read pool info

Use `rp-info` to get read pool information.

```
./zbox rp-info
```

### Lock tokens into read pool

Lock some tokens in read pool associated with an allocation. The tokens will be divided between allocation blobbers by their read price.

#### Usage

```
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

### Unlock tokens from read pool

Use `rp-unlock` to unlock tokens from an expired read pool by pool id. See `rp-info` for the POOL_ID and the expiration.

#### Usage

```
./zbox rp-unlock --pool_id <pool_id>
```

### Storage SC configurations

Show storage SC configuration.

```
./zbox sc-config
```

### Stake pool info

Use `sp-info` to get stake pool information and settings.

#### Usage

```
./zbox sp-info --blobber_id <blobber_id>
```

### Lock tokens into stake pool

Lock creates delegate pool for current client and given blobber. The tokens locked for the blobber stake can be unlocked any time, excluding where the tokens held by opened offers. The tokens collect interests.

#### Usage

```
./zbox sp-lock --blobber_id <blobber_id> --tokens 1.0
```

### Unlock tokens from stake pool

Unlock a stake pool by pool owner.

#### Usage

```
./zbox sp-unlock --blobber_id <blobber_id> --pool_id <pool_id>
```

### Stake pools info of current user

Get information about all stake pools of current user.

```
./zbox sp-user-info
```

### Pay interests

Changes in stake pool pays all pending rewards to calculate next rewards correctly and don't complicate stake pool. But if there are no changes interests will not be paid. To pay the interests  `sp-pay-interests`  command can be used to pays interest for all delegates. Use `sp-info` to check interests can be paid or not.

#### Usage

```
./zbox sp-pay-interests --blobber_id <blobber_id>
```

### Write pool info

Write pool information.

#### Usage

For all write pools.

```
./zbox wp-info
```

Filtering by allocation.

```
./zbox wp-info --allocation <allocation_id>
```

### Lock tokens into write pool

`wp-lock` can be used to lock tokens in a write pool associated with an allocation. All tokens will be divided between allocation blobbers depending on their write price.

#### Usage

```
./zbox wp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

### Unlock tokens from write pool

`wp-unlock` unlocks an expired write pool by its POOL_ID. See `wp-info` for the pool id and the expiration. 
An expired write pool, associated with an allocation, can be locked until allocation finalization even if it's expired. It possible in cases where related blobber doesn't give their min lock demands. The finalization will pay the demand and unlock the pool.

#### Usage

```
./zbox wp-unlock --pool_id <pool_id>
```

### Download cost

`get-download-cost` determines the cost for downloading the remote file from dStorage.

#### Usage

```
./zbox get-download-cost --allocation <allocation_id> --remotepath /path/file.ext
```

Also, there are `authticket` and `lookuphash` flags to get the cost for non allocation owners.

### Upload cost

`get-upload-cost` determines the cost for uploading a local file on dStorage.

#### Usage

```
./zbox get-upload-cost --allocation <allocation_id> --localpath ./path/file.ext
```

------

## Streaming

Streaming feature can be implemented together with player for each platforms (Android, IOS, Mac, Windows).

![streaming-android](https://user-images.githubusercontent.com/65766301/120052635-ce373b00-c043-11eb-94a5-a9711078ee54.png)
IOS documentation: https://github.com/0chain/0box-ios

Android documentation: https://github.com/0chain/0boxAndroid

Mac documentation: https://github.com/0chain/0BoxSyncMac

For platforms using zboxcli commands, implementation are similar to platforms above, i.e.:

1. Download file with downloadFileByBlocks method
2. Read chunked files to byte array (inputstream)
3. Add byte array to custom media source of player

Improvements was done in

downloadFileByBlocks - properly returns file-chunks with correct byte range, gosdk v1.2.4 and above only.

getFileMeta - returns actuaBlockNumbers and actualFileSize (exclude thumbnail size)

getFileMetaByAuth - same updates as getFileMeta

listAllocation - returns actuaBlockNumbers and actualFileSize (exclude thumbnail size)

# Troubleshooting

1. Both `rp-info` and `rp-lock` are not working.

```
./zbox rp-info
```

Response:

```
Failed to get read pool info: error requesting read pool info: consensus_failed: consensus failed on sharders
```

This can happen if read pool is not yet created for wallet. Read pool is usually created when new wallet is created by `zbox` or `zwallet`. However, if wallet is recovered through `zwallet recoverwallet`, read pool may not have been created. Simply run `zbox rp-create`  to create a read pool.

