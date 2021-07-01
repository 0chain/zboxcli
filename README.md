# zbox - a CLI for 0Chain dStorage

zbox is a command line interface (CLI) tool to understand the capabilities of 0Chain dStorage and prototype your app. The utility is built using 0Chain's goSDK library written in Go. Check out a [video](https://youtu.be/TPrkRjdaHrY) on how to use the CLI to create an allocation (storage volume) and upload, download, update, delete, and share files and folders on the 0Chain dStorage platform.

![Storage](https://user-images.githubusercontent.com/65766301/120052450-0ab66700-c043-11eb-91ab-1f7aa69e133a.png)

- [Getting Started](#getting-started)
- [Running zbox](https://github.com/0chain/zboxcli#Command-with-no-arguments) 
- [Global Flags](#global-flags)
- [Commands](#commands)
  - Creating and Managing Allocations
    - [Register a Wallet](https://github.com/0chain/zboxcli#Register)
    - [Get detailed Allocation](https://github.com/0chain/zboxcli#Get)
    - [Create an allocation](https://github.com/0chain/zboxcli#Create-new-allocation)
    - [List allocations](https://github.com/0chain/zboxcli#List-allocations)
    - [Update an allocation](https://github.com/0chain/zboxcli#Update-allocation)
    - [Cancel allocation](https://github.com/0chain/zboxcli#Cancel-allocation)
    - [Finalize allocation](https://github.com/0chain/zboxcli#Finalize-allocation)
  - Uploading and Managing Files
    - [Upload a file to dStorage](https://github.com/0chain/zboxcli#Upload)
    - [Download the uploaded file from dStorage](https://github.com/0chain/zboxcli#Download)
    - [Update the uploaded file on dStorage](https://github.com/0chain/zboxcli#Update)
    - [Delete the uploaded file on dStorage](https://github.com/0chain/zboxcli#Delete)
    - [List the uploaded files and folders](https://github.com/0chain/zboxcli#List)
    - [Copy uploaded files to another folder path on dStorage](https://github.com/0chain/zboxcli#Copy)
    - [Move uploaded files to another folder path on dStorage](https://github.com/0chain/zboxcli#Move)
    - [Rename a file on dStorage](https://github.com/0chain/zboxcli#Rename)
    - [Get meta data of files](https://github.com/0chain/zboxcli#Get-metadata)
    - [Get file stats](https://github.com/0chain/zboxcli#Stats)
    - [Update file attributes](https://github.com/0chain/zboxcli#Update-file-attributes)
    - [Get download cost](https://github.com/0chain/zboxcli#Download-cost)
    - [Get upload cost](https://github.com/0chain/zboxcli#Upload-cost)
  - Advanced Features
    - [Repair a file on dStorage](https://github.com/0chain/zboxcli#Repair)
    - [Sync your local folder to remote](https://github.com/0chain/zboxcli#Sync)
    - [Share the uploaded file on dStorage](https://github.com/0chain/zboxcli#Share)
    - [Add Collaborator for a file](https://github.com/0chain/zboxcli#Add-collaborator)
    - [Remove Collaborator for a file](https://github.com/0chain/zboxcli#Delete-collaborator)
    - [Video Streaming](https://github.com/0chain/zboxcli#Streaming)
  - Locking and unlocking tokens    
    - [Get wallet information](https://github.com/0chain/zboxcli#Get-wallet)
    - [Challenge pool information](https://github.com/0chain/zboxcli#Challenge-pool-information)
    - [Create read pool if not exists](https://github.com/0chain/zboxcli#Create-read-pool)
    - [Detailed read pool information](https://github.com/0chain/zboxcli#Read-pool-info)
    - [Lock tokens into read pool](https://github.com/0chain/zboxcli#Lock-tokens-into-read-pool)
    - [Unlock tokens from expired read pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-read-pool)
    - [Detailed write pool information](https://github.com/0chain/zboxcli#Write-pool-info)
    - [Lock tokens into write pool](https://github.com/0chain/zboxcli#Lock-tokens-into-write-pool)
    - [Unlock tokens from expired write pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-write-pool)
    - [Detailed stake pool information](https://github.com/0chain/zboxcli#Stake-pool-info)
    - [Lock tokens into stake pool](https://github.com/0chain/zboxcli#Lock-tokens-into-stake-pool)
    - [Unlock tokens from expired stake pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-stake-pool)
    - [Stake pools info of current user](https://github.com/0chain/zboxcli#Stake-pools-info-of-current-user)
    - [Pay interests](https://github.com/0chain/zboxcli#Pay-interests)
  - zbox Configuration info
    - [Storage SC configurations](https://github.com/0chain/zboxcli#Storage-SC-configurations)
    - [List blobbers](https://github.com/0chain/zboxcli#List-blobbers)
    - [Detail blobber information](https://github.com/0chain/zboxcli#Detailed-blobber-information)
    - [Update blobber settings](https://github.com/0chain/zboxcli#Update-blobber-settings)
    

   - [Troubleshooting](#troubleshooting)


## Getting Started

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

## Running zbox

When you run the `zbox` command in terminal with no arguments, it will list all the available commands and the global flags.

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
### Global Flags

 Global Flags are parameters in zbox that can be used with any command to override the default configuration.zbox supports the following global parameters.

| Flags                      | Description                                                  | Usage                                             |
| -------------------------- | ------------------------------------------------------------ | ------------------------------------------------- |
| --config string            | Specify a zbox configuration file (default is [$HOME/.zcn/config.yaml](#zcnconfigyaml)) | zbox [command] --config config1.yaml              |
| --configDir string         | Specify a zbox configuration directory (default is $HOME/.zcn) | zbox [command] --configDir /$HOME/.zcn2           |
| -h, --help                 | Gives more information about a particular command.           | zbox [command] --help                             |
| --network string           | Specify a network file to overwrite the network details(default is [$HOME/.zcn/network.yaml](#zcnnetworkyaml)) | zbox [command] --network network1.yaml            |
| --verbose                  | Provides additional details as to what the particular command is doing. | zbox [command] --verbose                          |
| --wallet string            | Specify a wallet file or 2nd wallet (default is $HOME/.zcn/wallet.json) | zbox [command] --wallet wallet2.json              |
| --wallet_client_id string  | Specify a wallet client id (By default client_id specified in $HOME/.zcn/wallet.json is used) | zbox [command] --wallet_client_id <client_id>     |
| --wallet_client_key string | Specify a wallet client_key (By default client_key specified in $HOME/.zcn/wallet.json is used) | zbox [command] --wallet_client_key  < client_key> |

 
# Commands

Note in this document, we will only show the commands for particular functionalities, the response will vary depending on your usage and may not be provided in all places. To get a more descriptive view of all the zbox functionalities check zbox cli documentation at docs.0chain.net.


## `register` Register wallet

`register` is used when needed to register a given wallet to the blockchain. This could be that the blockchain network is reset and you wished to register the same wallet at `~/.zcn/wallet.json`.

![image](https://user-images.githubusercontent.com/6240686/124104251-fb6b7480-da59-11eb-9397-9151b04de363.png)

Sample command

```sh
./zwallet register
```

Sample output

```
Wallet registered
```

## newallocation Create new allocation

Command `newallocation` reserves hard disk space on the blobbers. Later `upload`
can be used to save files to the blobber. `newallocation` has three modes triggered by the presence or absence of the `cost` 
and `free_storage` parameters.
* `cost` Converts `newallocation` into a query that returns the cost of the allocation
  determined by the remaining parameters.
* `free_storage` Creates an allocation using a free storage marker. All other
  parameters except `cost` will be ignored. The allocation settings will be set
  automatically by `0chain`, from preconfigured values.  
* `otherwise` Creates an allocation applying the settings indicated by the
  remaining parameters.  


| Parameter          | Description                                               | Default        | Valid Values |
|--------------------|-----------------------------------------------------------|----------------|--------------|
| allocationFileName | local file to store allocation information                | allocation.txt | file path    |
| cost               | returns the cost of the allocation, no allocation created |                | flag         |
| data               | number of data shards, effects upload and download speeds | 2              | int          |
| expire             | duration to allocation expiration                         | 720h           | duration     |
| free_storage       | free storage marker file.                                 |                | file path    |
| lock               | lock write pool with given number of tokens               |                | float        |
| mcct               | max challenge completion time                             | 1h             | duration     |
| parity             | number of parity shards, effects availability             | 2              | int          |
| read_price         | filter blobbers by read price range                       | 0-inf          | range        |
| size               | size of space reserved on blobbers                        | 2147483648     | bytes          |
| usd                | give token value in USD                                   |                | flag         |
| write_price        | filter blobbers by write price range                      | 0-inf          | range        |


<details>
  <summary> New allocation </summary>

![allocation](https://user-images.githubusercontent.com/65766301/120052477-27529f00-c043-11eb-91bb-573558325b20.png)

![image](https://user-images.githubusercontent.com/6240686/124010595-e8fc2700-d9d6-11eb-83b9-dc10cbeb75e0.png)

</details>

<details>
  <summary> Free storage new allocation </summary>

![image](https://user-images.githubusercontent.com/6240686/124010041-3926b980-d9d6-11eb-80f1-f062c92751ed.png)

</details>

#### Free storage allocation

Entities give free `0chain` storage in the form of markers. A marker takes the 
form of a json file
```json
{
  "assigner": "my_corporation",
  "recipient": "f174cdda7e24aeac0288afc2e8d8b20eda06b18333efd447725581dc80552977",
  "free_tokens": 2.1,
  "timestamp": 2000000,
  "signature": "9edb86c8710d5e3ee4fde247c638fd6b81af67e7bb3f9d60700aec8e310c1f06"
}
```
* `assigner` A label for the entity gifting the free storage.
* `recipient` The marker has to be run by the recipient to be valid.
* `free_tokens` The amount free tokens; equivalent to the `size` filed.
* `timestamp` Used to prevent multiple applications of the same marker.
* `signature` Signed by the assigner, validated using public key on the blockchain.
All allocation settings, other than `lock`, will be set automatically by 0chain. 
  Once created, an allocation funded by a free storage marker becomes identical to
  any other allocation; Its history forgotten.
  
```shell
./zbox newallocation --free_allocation markers/referal_marker.json
```
```shell
Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```  

#### Example

To create a new allocation with default values,use `newallocation` with a `--lock` flag to add 
some tokens to the write pool .On success a related write pool is created and the allocation 
information is stored under `$HOME/.zcn/allocation.txt`.
```shell
./zbox newallocation --lock 0.5
```

```shell
./zbox newallocation --lock 0.5 --free_storage markers/my_marker.json
```

Response:

```
Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

## updateallocation Update allocation

`updateallocation` updates allocation settings. It has two modes depending on 
the presence of the `free_storage` field. 
* `free_storage` Uses a free storage marker to fund this allocation update; settings
  predefined by `0chain`. See [newallocation](#free-storage-allocation) for further details.
* `otherwise` Update an allocation applying the settings indicated by the
  remaining parameters.  
  
| Parameter     | Required | Description                                            | Valid Values |
|---------------|----------|--------------------------------------------------------|--------------|
| allocation    | yes      | allocation id                                          | string       |
| expiry        |          | adjust storage expiration time                         | duration     |
| free_storage  |          | free storage marker file                               | string       |
| lock          | yes*     | lock additional tokens in write pool                    | int          |
| set_immutable |          | sets allocation so that data can no longer be modified | boolean      |
| size          |          | adjust allocation size                                 | bytes        |
* only required if free_storage not set.

<details>
  <summary>Update allocation </summary>

![image](https://user-images.githubusercontent.com/6240686/124003064-65d6d300-d9ce-11eb-808d-2d59340b00e7.png)

</details>

<details>
  <summary>Free storage update allocation </summary>

![image](https://user-images.githubusercontent.com/6240686/124003924-602dbd00-d9cf-11eb-910c-1d286c2a173c.png)

</details>

```
./zbox updateallocation --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --expiry 48h --size 4096
```


```shell
./zbox updateallocation --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --free_storage "markers/my_marker.json"
```

Output:

```
Allocation updated with txId : fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d
```

You can see more txn details using above txID in block explorer [here](https://one.devnet-0chain.net/).

### alloc-cancel Cancel allocation

Cancel allocation immediately return all tokens from challenge pool back to user
(to write pool) and cancels the allocation. In this case blobber will
not give their min lock demand. If blobbers already got some tokens, 
the tokens will not be returned. Cancelling an allocation can only occur
if the amount of failed challenges exceed a preset threshold.

| Parameter  | Required | Description   | Valid Values |
|------------|----------|---------------|--------------|
| allocation | yes      | allocation id | string       |

<details>
  <summary>Cancel allocation</summary>

![image](https://user-images.githubusercontent.com/6240686/124147442-9aa66100-da86-11eb-8b88-cd20306bfde1.png)

</details>

#### Example

```
./zbox alloc-cancel --allocation <allocation_id>
```

## alloc-fini Finalize allocation

Finalize an expired allocation. When an allocation expires, 
after its challenge completion time (after the expiration), 
it can be finalized by the owner or one of the allocation blobbers.

| Parameter  | Required | Description   | Valid Values |
|------------|----------|---------------|--------------|
| allocation | yes      | allocation id | string       |

<details>
  <summary>Cancel allocation</summary>

![image](https://user-images.githubusercontent.com/6240686/124149297-5c11a600-da88-11eb-9274-1fb756d93358.png)

</details>

#### Example

```
./zbox alloc-fini --allocation <allocation_id>
```

## ls-blobbers List blobbers

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
<details>
  <summary>rp-create sequence diagram</summary>

![image](https://user-images.githubusercontent.com/6240686/123973204-77f74800-d9b3-11eb-8165-96741cc0b291.png)

</details>

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

* Uses two different formats, you can either define a specific blobber
to lock all tokens, or spread across all the allocations blobbers automatically.
* If the user does not have a pre-existing read pool, then the smart-contract
creates one.

<details>
  <summary>rp-lock with a specific blobber</summary>

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1 --blobber f65af5d64000c7cd2883f4910eb69086f9d6e6635c744e62afcfab58b938ee25 
```
![image](https://user-images.githubusercontent.com/6240686/123973055-5302d500-d9b3-11eb-844a-90aa4b43f56b.png)

</details>

<details>
  <summary>rp-lock spread across all blobbers</summary>

Tokens are spread between the blobber pools weighted by 
each blobber's Terms.ReadPrice.

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

![image](https://user-images.githubusercontent.com/6240686/123973442-abd26d80-d9b3-11eb-9e37-c8e6551ed48c.png)

</details>

### Unlock tokens from read pool

Use `rp-unlock` to unlock tokens from an expired read pool by pool id. See `rp-info` for the POOL_ID and the expiration.

#### Usage

```
./zbox rp-unlock --pool_id <pool_id>
```
<details>
  <summary>rp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/123973808-ff44bb80-d9b3-11eb-9b70-c952fd95858f.png)

</details>

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

* Uses two different formats, you can either define a specific blobber
  to lock all tokens, or spread across all the allocations blobbers automatically.
* If the user does not have a pre-existing read pool, then the smart-contract
  creates one.

<details>
  <summary>rp-lock with a specific blobber</summary>

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1 --blobber f65af5d64000c7cd2883f4910eb69086f9d6e6635c744e62afcfab58b938ee25 
```
![image](https://user-images.githubusercontent.com/6240686/123988183-b4c93c00-d9bf-11eb-825c-9a5849fedbbf.png)

</details>

<details>
  <summary>rp-lock spread across all blobbers</summary>

Tokens are spread between the blobber pools weighted by
each blobber's Terms.ReadPrice.

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

![image](https://user-images.githubusercontent.com/6240686/123979735-e5f23e00-d9b8-11eb-8232-339a4a3374d0.png)

</details>

### Unlock tokens from write pool

`wp-unlock` unlocks an expired write pool by its POOL_ID. See `wp-info` for the pool id and the expiration. 
An expired write pool, associated with an allocation, can be locked until allocation finalization even if it's expired. It possible in cases where related blobber doesn't give their min lock demands. The finalization will pay the demand and unlock the pool.

#### Usage

```
./zbox wp-unlock --pool_id <pool_id>
```

<details>
  <summary>rp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/123980742-b09a2000-d9b9-11eb-8987-c18ff90ee705.png)

</details>

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

### Streaming

Video streaming with Zbox CLI can be implemented with players for different operating platforms(iOS, Android Mac).Zbox CLI does not have a player itself and use the the downloadFileByBlocks helper function to properly returns file-chunks with correct byte range.

![streaming-android](https://user-images.githubusercontent.com/65766301/120052635-ce373b00-c043-11eb-94a5-a9711078ee54.png)

#### How it works:

When the user starts the video player (ExoPlayer for Android or AVPlayer for iOS), A ZChainDataSource starts chunked download and requests chunks of video from the buffer(a Middleman between streaming player and Zbox).

After the arrival of the first chunk, the player starts requesting more chunks from the buffer, which requests the Zbox SDK. Zbox SDK, which is built using GO, makes use of the downloadFileByBlocks method to reliably download large files by chunking them into a sequence of parts that can be downloaded individually. Once the blocks are downloaded, they are read into input streams and added to the media source of the streaming player.

The task of downloading files and writing them to buffer using Zbox SDK happens constantly, and If players request random bits of video, they are delivered instantly by a buffer.

In a case, if the player didn't receive chunks (for example, it's still not downloaded), then the player switches to STALE state, and the video stream will pause. During the STALE state, a player tries to make multiple requests for chunks; if didn't receive a response, the video stream stops.


#### Usage

To understand how Zbox CLI provides downloading of files by blocks. Let's consider an allocation that has `audio. mp3 ` file stored on dStorage. Make sure the file has a large size(more than 64 kB(64000 bytes)) to download the file by blocks. The size and other attributes of the sample `audio. mp3` file can be viewed using

```
./zbox list --allocation $ALLOC --remotepath /myfiles
```

Response:

```
  TYPE |   NAME    |        PATH        |  SIZE   | NUM BLOCKS |LOOKUP HASH      | IS ENCRYPTED | DOWNLOADS PAYER
+------+-----------+--------------------+---------+------------+----------------
  f    | audio.mp3 | /myfiles/audio.mp3 | 5992396 |         92 | 3cea39505cc30fb9f6fc5c6045284188feb14eac8ff3a19577701c4f6d973239 | NO           | owner

```

Here we can see the `audio.mp3` file of size (5993286) bytes having 92 blocks.If we want to download a certain number of blocks for the `audio.mp3` file we can use the `--endblock` or `--startblock` flag with `./zbox download` command. Other flags for download can be viewed using `./zbox download --help`


```
Flags:
    
  -b, --blockspermarker int   pass this option to download multiple blocks per marker (default 10)       
  -e, --endblock int          pass this option to download till specific block number                     
  -h, --help                  help for download
  --localpath string          Local path of file to download                           
  --remotepath string         Remote path to download
   -s, --startblock int       Pass this option to download from specific block number                                                                                                         
```

For only downloading three blocks of `audio.mp3` file, we specify `--startblock` and`--endblock` with integer value of 1 and 3. `--blockspermarker` flag can also be specified to download multiple blocks at a time(default is 10).

Sample command for downloading till 3rd block of the `audio.mp3` file would be:

```
./zbox download --localpath /root --remotepath /myfiles/audio.mp3 --allocation $ALLOC --startblock 1 --endblock 3 
```

Response:

```
 393216 / 2996198 [====================>-----------------------------------------------------------------------------------------------------------------------------------------]  13.12% 1s
Status completed callback. Type = audio/mpeg. Name = audio.mp3

```

As we can see, the downloaded file size(393216) is less than the original(2996198), which means zbox has downloaded some blocks of the file.



## Troubleshooting

1. Both `rp-info` and `rp-lock` are not working.

```
./zbox rp-info
```

Response:

```
Failed to get read pool info: error requesting read pool info: consensus_failed: consensus failed on sharders
```

This can happen if read pool is not yet created for wallet. Read pool is usually created when new wallet is created by `zbox` or `zwallet`. However, if wallet is recovered through `zwallet recoverwallet`, read pool may not have been created. Simply run `zbox rp-create`  to create a read pool.

