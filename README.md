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
    - [List allocations](https://github.com/0chain/zboxcli#List-allocations)     
    - [Create new allocation](#Create-new-allocation)
    - [Update allocation](#update-allocation)
    - [Cancel allocation](#cancel-allocation)
    - [Finalise allocation](#finalise-allocation)
    - [Add curator](#add-curator)
    - [Transfer allocation ownership](#transfer-allocation-ownership)
  - Uploading and Managing Files
    - [Upload a file to dStorage](#upload)
    - [Download the uploaded file from dStorage](#download)
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
    - [Detailed stake pool information](#stake-pool-info)
    - [Lock tokens into stake pool](https://github.com/0chain/zboxcli#Lock-tokens-into-stake-pool)
    - [Unlock tokens from expired stake pool](https://github.com/0chain/zboxcli#Unlock-tokens-from-stake-pool)
    - [Stake pools info of current user](#stake-pools-info-of-user)
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

**Command**|**Description**
:-----:|:-----:
add|Adds free storage assigner
[add-collab](#add-collaborator)|add collaborator for a file
[addcurator](#add-curator)|Adds a curator to an allocation
[alloc-cancel](#cancel-allocation)|Cancel an allocation
[alloc-fini](#finalise-allocation)|Finalize an expired allocation
[bl-info](#detailed-blobber-information)|Get blobber info
[bl-update](#update-blobber-settings)|Update blobber settings by its delegate\_wallet owner
commit|commit a file changes to chain
[copy](#copy)|copy an object(file/folder) to another folder on blobbers
[cp-info](#challenge-pool-information)|Challenge pool information.
[delete](#delete)|delete file from blobbers
[delete-collab](#delete-collaborator)|delete collaborator for a file
[download](#download)|download file from blobbers
[get](#get)|Gets the allocation info
[get-diff](#get-differences)|Get difference of local and allocation root
[get-download-cost](#download-cost)|Get downloading cost
[get-upload-cost](#upload-cost)|Get uploading cost
[getwallet](#get-wallet)|Get wallet information
help|Help about any command
[list](#list)|list files from blobbers
[list-all](#list-all-allocations)|list all files from blobbers
[listallocations](#list-allocations)|List allocations for the client
[ls-blobbers](#list-blobbers)|Show active blobbers in storage SC.
[meta](#get-metadata)|get meta data of files from blobbers
[move](#move)|move an object(file/folder) to another folder on blobbers
[newallocation](#create-new-allocation)|Creates a new allocation
[register](#register-wallet)|Registers the wallet with the blockchain
[rename](#rename)|rename an object(file/folder) on blobbers
[rp-create](#create-read-pool)|Create read pool if missing
[rp-info](#read-pool-info)|Read pool information.
[rp-lock](#lock-tokens-into-read-pool)|Lock some tokens in read pool.
[rp-unlock](#unlock-tokens-from-read-pool)|Unlock some expired tokens in a read pool.
[sc-config](#storage-sc-configurations)|Show storage SC configuration.
[share](#share)|share files from blobbers
sign-data|Sign given data
[sp-info](#stake-pool-info)|Stake pool information.
[sp-lock](#lock-tokens-into-stake-pool)|Lock tokens lacking in stake pool.
[sp-pay-interests](#pay-interests)|Pay interests not payed yet.
[sp-unlock](#unlock-tokens-from-stake-pool)|Unlock tokens in stake pool.
[sp-user-info](#stake-pools-info-of-user)|Stake pool information for a user.
[start-repair](#repair)|start repair file to blobbers
[stats](#stats)|stats for file from blobbers
[sync](#sync)|Sync files to/from blobbers
[transferallocation](#transfer-allocation-ownership)|Transfer an allocation between owners
[update](#update)|update file to blobbers
update-attributes|depreciated
[updateallocation](#update-allocation)|Updates allocation's expiry and size
[upload](#upload)|upload file to blobbers
version|Prints version information
[wp-info](#write-pool-info)|Write pool information.
[wp-lock](#lock-tokens-into-write-pool)|Lock some tokens in write pool.
[wp-unlock](#unlock-tokens-from-write-pool)|Unlock some expired tokens in a write pool.

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


## Register wallet

`register` is used when needed to register a given wallet to the blockchain. This could be that the blockchain network is reset and you wished to register the same wallet at `~/.zcn/wallet.json`.

Sample command

```sh
./zwallet register
```

Sample output

```
Wallet registered
```

## Create new allocation

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
  <summary>newallocation </summary>

![allocation](https://user-images.githubusercontent.com/65766301/120052477-27529f00-c043-11eb-91bb-573558325b20.png)

![image](https://user-images.githubusercontent.com/6240686/125315476-15953480-e32f-11eb-8a11-b069079911d3.png)

</details>

<details>
  <summary>Free storage newallocation </summary>

![image](https://user-images.githubusercontent.com/6240686/125315943-876d7e00-e32f-11eb-95b8-908d5bf456bc.png)

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
* `assigner` A label for the entity giving the free storage.
* `recipient` The marker; has to be run by the recipient to be valid.
* `free_tokens` The number of free tokens; equivalent to the `size` filed.
* `timestamp` A unique timestamp. Used to prevent multiple applications of the same marker.
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

Example

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

## Update allocation

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
`*` only required if free_storage not set.

<details>
  <summary>updateallocation </summary>


![image](https://user-images.githubusercontent.com/6240686/125335948-0b7e3080-e345-11eb-82af-20fd1e4501df.png)
</details>

<details>
  <summary>Free storage updateallocation</summary>

![image](https://user-images.githubusercontent.com/6240686/125335821-e984ae00-e344-11eb-9960-648a76550bc3.png)

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

### Cancel allocation

`alloc-cancel` immediately return all tokens from challenge pool back to the 
allocation's owner and cancels the allocation. If blobbers already got some tokens, 
the tokens will not be returned. Cancelling an allocation can only occur
if the amount of failed challenges exceed a preset threshold.

| Parameter  | Required | Description   | Valid Values |
|------------|----------|---------------|--------------|
| allocation | yes      | allocation id | string       |

<details>
  <summary>alloc-cancel</summary>

![image](https://user-images.githubusercontent.com/6240686/125453211-a65caad7-3d46-4ea9-84ab-1ed8cd5f5820.png)


</details>

Example

```
./zbox alloc-cancel --allocation <allocation_id>
```

## Finalise allocation
 
`alloc-fini` finalises an expired allocation. When an allocation expires, 
after its challenge completion time (after the expiration), 
it can be finalised by the owner or one of the allocation blobbers.

| Parameter  | Required | Description   | Valid Values |
|------------|----------|---------------|--------------|
| allocation | yes      | allocation id | string       |

<details>
  <summary>alloc-fini</summary>

![image](https://user-images.githubusercontent.com/6240686/125453928-5d881535-0426-4c93-96fd-aed3bf70ee17.png)


</details>

Example

```
./zbox alloc-fini --allocation <allocation_id>
```

## Add curator

`addcurator` adds a curator to an allocation.
A curator can transfer ownership of an allocation. Each allocation 
maintains a list of these curators.

| Parameter  | Required | Description                           | Valid Values |
|------------|----------|---------------------------------------|--------------|
| allocation | yes      | allocation id                         | string       |
| curator    | yes      | id of new curator to add to allocation | string       |

<details>
  <summary>addcurator</summary>

![image](https://user-images.githubusercontent.com/6240686/125454687-9de43b37-62ac-45c6-8b1d-a5c883672d56.png)

</details>

```shell
./zbox addcurator --allocation fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d --curator  e49458a13f8a000b5959d03f8f7b6fa397b578643940ba50d3470c201d333429
```

```shell
e49458a13f8a000b5959d03f8f7b6fa397b578643940ba50d3470c201d333429 added as a curator to allocation fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d
```

## Transfer allocation ownership

`transferallocation` changes the owner of an allocation. Only a curator, 
previously added by an [addcurator](#add-curator) command can change an 
allocation's ownership.

`transferallocation` does not move any funds, only changes the owner, 
and the owner's public key.

| Parameter     | Required | Description             | Valid Values |
|---------------|----------|-------------------------|--------------|
| allocation    | yes      | allocatino id           | string       |
| new_owner     | yes      | id of the new owner     | string       |
| new_owner_key | yes      | public key of new owner | string       |

<details>
  <summary>transferallocation</summary>

![image](https://user-images.githubusercontent.com/6240686/125456952-115b5468-c4f8-4564-8160-86f8164c5ce6.png)

</details>

```shell
./zbox trnasferallocation --allocation fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d \
    --new_owner 8b87739cd6c966c150a8a6e7b327435d4a581d9d9cc1d86a88c8a13ae1ad7a96
    --new_owner_key a2df94e69954e51999f768aeca40bf0678e168fa9eb21ee5c82c32a9c25fb71fe9a340b726456d6e557f92854975ef04270291cdc1853e56000b0a6b48312d13
```

Output

```shell
transferred ownership of fb84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d to 8b87739cd6c966c150a8a6e7b327435d4a581d9d9cc1d86a88c8a13ae1ad7a96
```

## List blobbers

Use `ls-blobbers` command to show active blobbers.

| Parameter | Required | Description                          | Valid Values |
|-----------|----------|--------------------------------------|--------------|
| all       | no       | shows active and non active blobbers | flag         |
| json      | no       | display result in .json format       | flag         |

Example

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

## Detailed blobber information

Use `bl-info` command to get detailed blobber information.

| Parameter  | Required | Description                         | default | Valid values |
|------------|----------|-------------------------------------|---------|--------------|
| blobber id | yes      | blobber on which to get information |         | string       |
| json       | no       | print result in json format         | false   | boolean      |

<details>
  <summary>bl-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124609407-7fad6580-de67-11eb-896c-1f7be1faf7c0.png)

</details>

Example

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

## Lost all files

`list-all` lists al the files stored with an allocation


| Parameter               | Required | Description                            | Default | Valid values                            |
|-------------------------|----------|----------------------------------------|---------|-----------------------------------------|
| allocation              | yes      | allocation id, sender must be allocation owner                   |         | string                                  |

```shell
./zbox list-all --allocation 4ebeb69feeaeb3cd308570321981d61beea55db65cbeba4ba3b75c173c0f141b
```

## List all allocations

`listallocations` provides a list of all allocations owned by the user.

| Parameter          | Required | Description                               | default | Valid values |
|--------------------|----------|-------------------------------------------|---------|--------------|
| json         | no     | print output in json format |         | boolean       |

```shell
./zbox listallocations
ZED | CANCELED | R  PRICE |   W  PRICE    
+------------------------------------------------------------------+-----------+-------------------------------+------------+--------------+-----------+----------+----------+--------------+
  4ebeb69feeaeb3cd308570321981d61beea55db65cbeba4ba3b75c173c0f141b | 104857600 | 2021-07-16 13:34:29 +0100 BST |          1 |            1 | false     | false    |     0.02 | 0.1999999998 
```


## Update blobber settings

Use `./zbox bl-update to update a blobber's configuration settings. This updates the settings
on the blockchain not the blobber.

| Parameter          | Required | Description                               | default | Valid values |
|--------------------|----------|-------------------------------------------|---------|--------------|
| blobber_id         | yes      | id of blobber of which to update settings |         | string       |
| capacity           | no       | update blobber capacity                   |         | int          |
| cct                | no       | update challenge completion time          |         | duration     |
| max_offer_duration | no       | update max offer duration                 |         | duration     |
| max_stake          | no       | update maximum stake                      |         | float        |
| min_lock_demand    | no       | update minimum lock demand                |         | float        |
| min_stake          | no       | update minimum stake                        |         | float        |
| num_delegates      | no       | update maximum number of delegates          |         | int          |
| read_price         | no       | update read price                        |         | float        |
| service_charge     | no       | update service charge                     |         | float        |
| write_price        | no       | update write price                        |         | float        |

<details>
  <summary>bl-update</summary>

![image](https://user-images.githubusercontent.com/6240686/124616924-6825ab00-de6e-11eb-80a7-13e8061dd20b.png)

</details>

Example

Update blobber read price

```
./zbox bl-update --blobber_id 0ece681f6b00221c5567865b56040eaab23795a843ed629ce71fb340a5566ba3 --read_price 0.1
```

## Upload

Use `upload` command to upload a file. The user must be the owner of the allocation.
You can request the file be encrypted before upload, and can send thumbnails 
with the file. 

| Parameter               | Required | Description                            | Default | Valid values                            |
|-------------------------|----------|----------------------------------------|---------|-----------------------------------------|
| allocation              | yes      | allocation id, sender must be allocation owner                   |         | string                                  |
| commit                  | no       | save metadata to blockchain                                      | false   | boolean                                 |
| encrypt                 | no       | encrypt file before upload                                       | false   | boolean                                 |
| localpath               | yes      | local path of the file to upload                                 |         | file path                               |
| remotepath              | yes      | remote path to upload file to, use to access file later          |         | string                              |
| thumbnailpath           | no       | local path of thumbnaSil                                         |         | file path                               |


<details>
  <summary>upload</summary>

![image](https://user-images.githubusercontent.com/6240686/124287350-cf2e2180-db47-11eb-8079-40f069a5e0c2.png)

</details>

Example

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

Use upload command with optional encrypt parameter to upload a file in encrypted 
format. This can be downloaded as normal from same wallet/allocation or utilize 
Proxy Re-Encryption facility (see [download](https://github.com/0chain/zboxcli#Download) command).

```
./zbox upload --encrypt --localpath <absolute path to file>/sensitivedata.txt --remotepath /myfiles/sensitivedata.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

Response:

```
12390 / 12390 [================================================================================] 100.00% 3s
Status completed callback. Type = application/octet-stream. Name = sensitivedata.txt
```

## Download

Use `download` command to download your own or a shared file. 
* `owner` The owner of the allocation can always download files, in this case the owner pays for the download.
* `collaborator` A collaborator can download files, the owner pays. To add collaborators to an allocation, use
  [add-collab](#add-collaborator).
* `authticket` To download a file using `authticket`, you must have previous be given an auth
  ticket using the [share](#share) command. Use rx_pay to indicate who pays, `rx_pay = true` you pay,
  `rx_pay = false` the allocation owner pays.
Use `startblock` and `endblock` to only download part of the file.   

| Parameter       | Required | Description                                                              | Default | Valid values |
|-----------------|----------|--------------------------------------------------------------------------|---------|--------------|
| allocation      | yes      | allocation id                                                            |         | string       |
| authticket      | no       | auth ticked if not owner of the allocation, use share to get auth ticket |         | string       |
| blockspermarker | no       | download multiple blocks per marker                                      | 10      | int          |
| commit          | no       | save metadata to blockchain                                              | false   | boolean      |
| endblock        | no       | download until specified block number                                    |         | int          |
| localpath       | yes      | local path to which to download the file to                              |         | file path    |
| remotepath      | yes      | remote path to which the file was uploaded                               |         | string       |
| rx_pay          | no       | `authticket` must be valid, true = sender pays, false = allocation owner pays                                      | false   | boolean      |
| startblock      | no       | start download from specified block                                      |         | int          |
| thumbail        | no       | only download the thumbnail                                              | false   | boolean      |

<details>
  <summary>download</summary>

![image](https://user-images.githubusercontent.com/6240686/124352957-79b34c80-dbfb-11eb-883f-4bb583b9a618.png)

</details>

Example

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

## Update

Use `update` command to update content of an existing file in the remote path. 
Like [upload](#upload) command. Only the owner of the allocation or a collaborator
can update a file.  To add collaborators to an allocation, use
[add-collab](#add-collaborator).

| Parameter     | Required | Description                   | Default | Valid values |
|---------------|----------|-------------------------------|---------|--------------|
| allocation    | yes      | allocation id                 |         | string       |
| encrypt       | no       | encrypt file before upload    | false   | boolean      |
| localpath     | yes      | local file to upload          |         | file path    |
| remotepath    | yes      | remote file to upload         |         | string       |
| thumbnailpath | no       | local fumbnail file to upload |         | file path    |
| commit        | no       | save meta data to blockchain  | false   | boolean      |

<details>
  <summary>update</summary>

![image](https://user-images.githubusercontent.com/6240686/124354473-14b02480-dc04-11eb-9463-5a91d4f6f02d.png)

</details>

## Delete

Use `delete` command to delete your file on the allocation. Only the owner
of the application can delete a file.

| Parameter  | Required | Description                   | Default | Valid values |
|------------|----------|-------------------------------|---------|--------------|
| allocation | yes      | allocation id                 |         | string       |
| remotepath | yes      | remote path of file to delete |         | string       |
| commit     | no       | save meta data to blockchain  | false   | boolean      |

<details>
  <summary>=delete</summary>

![image](https://user-images.githubusercontent.com/6240686/124353872-0f050f80-dc01-11eb-9e45-ddf2c888223b.png)

</details>

Example

```
./zbox delete --allocation 3c0d32560ea18d9d0d76808216a9c634flist661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/horse.jpeg
```

Response:

```
/myfiles/horse.jpeg deleted
```

File successfully deleted (Can be verified using [list](https://github.com/0chain/zboxcli#List))

## Share

Use `share` command to generate an authorisation ticket that provides authorisation to the
holder to the specified file on the `remotepath`. Use the returned auth ticket with the
`--authticket` parameter.

`auth ticket` can be used with  [download](#download), [commit](#commit) and [list](#list), 
[meta](#get-metadata) and [get_download_cost](#download-cost), but only for files in 
the pre-defined remote path.

| Parameter           | Required | Description                                                       | Valid values |
|---------------------|----------|-------------------------------------------------------------------|--------------|
| allocation          | yes      | allocation id                                                     | string       |
| clientid            | no       | id of user to share file with, leave blank for public share       | string       |
| encryptionpublickey | no       | public key of the client to share file with, required if clientId | string       |
| remotepath          | yes      | remote path of file to share                                      | string       |

<details>
  <summary>share</summary>

![image](https://user-images.githubusercontent.com/6240686/124355532-876fce80-dc09-11eb-8166-bc7018480404.png)

</details>

Example

**Public share**

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt
```

Response:

```
auth ticket :eyJjbGllbnRfaWQiOiIiLCJvd25lcl9pZCI6ImI2ZGU1NjJiNTdhMGI1OTNkMDQ4MDYyNGY3OWE1NWVkNDZkYmE1NDQ0MDQ1OTViZWUwMjczMTQ0ZTAxMDM0YWUiLCJhbGxvY2F0aW9uX2lkIjoiODY5NWI5ZTdmOTg2ZDRhNDQ3YjY0ZGUwMjBiYTg2ZjUzYjNiNWUyYzQ0MmFiY2ViNmNkNjU3NDI3MDIwNjdkYyIsImZpbGVfcGF0aF9oYXNoIjoiMjBkYzc5OGIwNGViYWIzMDE1ODE3Yzg1ZDIyYWVhNjRhNTIzMDViYWQ2Zjc0NDlhY2QzODI4YzhkNzBjNzZhMyIsImZpbGVfbmFtZSI6IjEudHh0IiwicmVmZXJlbmNlX3R5cGUiOiJmIiwiZXhwaXJhdGlvbiI6MTYyNjQyMDM1OSwidGltZXN0YW1wIjoxNjE4NjQ0MzU5LCJyZV9lbmNyeXB0aW9uX2tleSI6IiIsInNpZ25hdHVyZSI6ImFjNzIzZjdhMWQ0ZDBmMjc2ZmQ3Yzc2NWMxOTcyZTlhODc2OGI0MjU1ODkyMmMwNjEyZjMxNjBjMGZiODQ5MGMifQ==
```


**Encrypted share**

Use clientid and encryptionpublickey of the user to share with.

![Private File Sharing](https://user-images.githubusercontent.com/65766301/120052575-962ff800-c043-11eb-9cf7-433383d532a3.png)

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt --clientid b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae --encryptionpublickey 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=
```

Response:

```
auth ticket :eyJjbGllbnRfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwib3duZXJfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwiYWxsb2NhdGlvbl9pZCI6Ijg2OTViOWU3Zjk4NmQ0YTQ0N2I2NGRlMDIwYmE4NmY1M2IzYjVlMmM0NDJhYmNlYjZjZDY1NzQyNzAyMDY3ZGMiLCJmaWxlX3BhdGhfaGFzaCI6IjIwZGM3OThiMDRlYmFiMzAxNTgxN2M4NWQyMmFlYTY0YTUyMzA1YmFkNmY3NDQ5YWNkMzgyOGM4ZDcwYzc2YTMiLCJmaWxlX25hbWUiOiIxLnR4dCIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MjY0MjA1NzQsInRpbWVzdGFtcCI6MTYxODY0NDU3NCwicmVfZW5jcnlwdGlvbl9rZXkiOiJ7XCJyMVwiOlwiOUpnci9aVDh6VnpyME1BcWFidlczdnhoWEZoVkdMSGpzcVZtVUQ1QTJEOD1cIixcInIyXCI6XCIrVEk2Z1pST3JCR3ZURG9BNFlicmNWNXpoSjJ4a0I4VU5SNTlRckwrNUhZPVwiLFwicjNcIjpcInhySjR3bENuMWhqK2Q3RXU5TXNJRzVhNnEzRXVzSlZ4a2N6YXN1K0VqQW89XCJ9Iiwic2lnbmF0dXJlIjoiZTk3NTYyOTAyODU4OTBhY2QwYTcyMzljNTFhZjc0YThmNjU2OTFjOTUwMzRjOWM0ZDJlMTFkMTQ0MTk0NmExYSJ9
```


Response contains an auth ticket- an encrypted string that can be shared.

## List

Use `list` command to list files in given remote path of the dStorage. An auth ticket should be provided when
not sent by the allocation owner. Using an auth ticket requires a `lookuphash` to indicate the object on which to list
information. 

| Parameter  | Required | Description                                                              | default | Valid values |
|------------|----------|--------------------------------------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                                                            |         | string       |
| authticket | no       | auth ticked if not owner of the allocation, use share to get auth ticket |         | sting        |
| json       | no       | output the response in json format                                       | false   | boolean      |
| lookuphash | no       | hash of object to list, use with auth ticket.                            |         | string       |
| remotepath | no       | remote path of objects to list, for auth ticket use lookuphash instead   |         | string       |

<details>
  <summary>list</summary>

![image](https://user-images.githubusercontent.com/6240686/124466241-5a005d80-dd8e-11eb-9122-30dbbd98d8e3.png)

</details>

Example

```
./zbox list --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /  
```

Response:

```
auth ticket :eyJjbGllbnRfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwib3duZXJfaWQiOiJiNmRlNTYyYjU3YTBiNTkzZDA0ODA2MjRmNzlhNTVlZDQ2ZGJhNTQ0NDA0NTk1YmVlMDI3MzE0NGUwMTAzNGFlIiwiYWxsb2NhdGlvbl9pZCI6Ijg2OTViOWU3Zjk4NmQ0YTQ0N2I2NGRlMDIwYmE4NmY1M2IzYjVlMmM0NDJhYmNlYjZjZDY1NzQyNzAyMDY3ZGMiLCJmaWxlX3BhdGhfaGFzaCI6IjIwZGM3OThiMDRlYmFiMzAxNTgxN2M4NWQyMmFlYTY0YTUyMzA1YmFkNmY3NDQ5YWNkMzgyOGM4ZDcwYzc2YTMiLCJmaWxlX25hbWUiOiIxLnR4dCIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MjY0MjA1NzQsInRpbWVzdGFtcCI6MTYxODY0NDU3NCwicmVfZW5jcnlwdGlvbl9rZXkiOiJ7XCJyMVwiOlwiOUpnci9aVDh6VnpyME1BcWFidlczdnhoWEZoVkdMSGpzcVZtVUQ1QTJEOD1cIixcInIyXCI6XCIrVEk2Z1pST3JCR3ZURG9BNFlicmNWNXpoSjJ4a0I4VU5SNTlRckwrNUhZPVwiLFwicjNcIjpcInhySjR3bENuMWhqK2Q3RXU5TXNJRzVhNnEzRXVzSlZ4a2N6YXN1K0VqQW89XCJ9Iiwic2lnbmF0dXJlIjoiZTk3NTYyOTAyODU4OTBhY2QwYTcyMzljNTFhZjc0YThmNjU2OTFjOTUwMzRjOWM0ZDJlMTFkMTQ0MTk0NmExYSJ9
```

Response will be a list with information for each file/folder in the given path. The information includes lookuphash which is require for download via authticket.

## Copy

Use `copy` command to copy file to another folder path in dStorage. Only the owner of the allocation can copy an object.

| Parameter  | Required | Description                                       | default | Valid values |
|------------|----------|---------------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                                     |         | string       |
| commit     | no       | save metadata to blockchain                       | false   | boolean      |
| remotepath | yes      | remote path of object to copy                     |         | string       |
| destpath   | yes      | destination, an existing directory to copy object |         | string       |

<details>
  <summary>copy</summary>

![image](https://user-images.githubusercontent.com/6240686/124470632-c0d44580-dd93-11eb-89ba-f22081429616.png)

</details>

Example

```
./zbox copy --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac --remotepath 
```

Response:

```
/file.txt --destpath /existingFolder
/file.txt copied
```

## Move

Use `move` command to move file to another remote folder path on dStorage.

| Parameter  | Required | Description                                       | default | Valid values |
|------------|----------|---------------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                                     |         | string       |
| commit     | no       | save metadata to blockchain                       | false   | boolean      |
| remotepath | yes      | remote path of object to copy                     |         | string       |
| destpath   | yes      | destination, an existing directory to copy object |         | string       |

<details>
  <summary>move</summary>

![image](https://user-images.githubusercontent.com/6240686/124471576-eca3fb00-dd94-11eb-8e52-441489c7cb55.png)

</details>

Example

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

| Parameter  | Required | Description                                                              | default | Valid values |
|------------|----------|--------------------------------------------------------------------------|---------|--------------|
|| json       | no       | output the response in json format                                       | false   | boolean      |

<details>
  <summary>listallocations</summary>

![image](https://user-images.githubusercontent.com/6240686/124474346-51ad2000-dd98-11eb-96f1-348ac926be3c.png)

</details>

Example

```
./zbox listallocations
```

Response:

```
                                 ID                                |    SIZE    |          EXPIRATION           | DATASHARDS | PARITYSHARDS | FINALIZED | CANCELED |   R  PRICE   |   W  PRICE    
+------------------------------------------------------------------+------------+-------------------------------+------------+--------------+-----------+----------+--------------+--------------+
  8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc | 6442450944 | 2021-05-24 00:27:23 +0700 +07 |          4 |            2 | false     | false    | 0.0599999994 | 0.0599999994  
```

## Sync

sync command helps in syncing all files from the local folder recursively to the remote.

| Parameter   | Required | Description                                                                                   | default | Valid values |
|-------------|----------|-----------------------------------------------------------------------------------------------|---------|--------------|
| allocation  | yes      | allocation id                                                                                 |         | string       |
| commit      | no       | commet metadata to blockchain                                                                 | false   | boolean      |
| encryptpath | no       | local directory path to be uploaded as encrypted                                              | false   | boolean      |
| exludepath  | no       | paths to exclude from sync                                                                    |         | string array |
| localchache | no       | local chache of remote snapshot. Used for comparsion with remote. After sync will be updated. |         | string       |
| localpath   | yes      | local directory to which to sync                                                              |         | file path    |
| uploadonly  | no       | only upload and update files                                                                  | false   | boolean      |



Example

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

## Get differences 
`get-diff` returns the differences between local files, and the files stored
with the allocation.

| Parameter   | Required | Description                                   | default | Valid values |
|-------------|----------|-----------------------------------------------|---------|--------------|
| allocation  | yes      | allocation id                                 |         | string       |
| excludepath | no       | remote folder paths to exclude during syncing |         | string array |
| localcache  | no       | local chache of remote snapshot               |         | string       |
| localpath   | yes      | local director to sync                        |         | string       |

## Get wallet

Use `getwallet` command to get additional wallet information including Encryption 
Public Key,Client ID which are required for Private File Sharing.

| Parameter  | Required | Description                  | default | Valid values |
|------------|----------|------------------------------|---------|--------------|
| json       | no       | print response in json format | false   | boolean      |

Example

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

## Get

Use `get` command to get the information about the allocation such as total size , used size, number of challenges 
and challenges passed/failed/open/redeemed.

| Parameter  | Required | Description                  | default | Valid values |
|------------|----------|------------------------------|---------|--------------|
| allocation | yes      | allocation id                |         | string       |
| json       | no       | print response in json format | false   | boolean      |

<details>
  <summary>get</summary>

![image](https://user-images.githubusercontent.com/6240686/124476040-4f4bc580-dd9a-11eb-939c-464ffc6936db.png)

</details>

Example

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

## Get metadata

Use `meta` command to get metadata for a given remote file. Use must either be the
owner of the allocation on have an auth ticket or be a collaborator. 
Use [share](#share) to create an auth ticket for someone or [add-collab](#add-collaborator)
to add a user as a collaborator. To indicate the object use `remotepath` or
`lookuphash` with an auth ticket.

| Parameter  | Required | Description                                                              | default | Valid values |
|------------|----------|--------------------------------------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                                                            |         | string       |
| authticket | no       | auth ticked if not owner of the allocation |         | string       |
| json       | no       | print result in json format                                              | false   | boolean      |
| lookuphash | no       | hash of object, use with auth ticket                                     |         | string       |
| remotepath | no       | remote path of objecte, do not use with authticket                       |         | string       |

<details>
  <summary>meta</summary>

![image](https://user-images.githubusercontent.com/6240686/124484414-4c090780-dda3-11eb-818c-d95477618cfd.png)

</details>

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


Response will be metadata for the given filepath/lookuphash (if using authTicket)

## Rename

`rename` command helps in renaming a file existing already on dStorage.

| Parameter  | Required | Description                                       | default | Valid values |
|------------|----------|---------------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                                     |         | string       |
| commit     | no       | save metadata to blockchain                       | false   | boolean      |
| destname   | yes      | new neame of the object                           |         | string       |
| remotepath | yes      | remote path of object, do not use with authticket |         | string       |

<details>
  <summary>rename</summary>

![image](https://user-images.githubusercontent.com/6240686/124487119-3ea14c80-dda6-11eb-93df-1e084653f212.png)

</details>

Example

```
./zbox rename --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --destname x.txt
```

Response:

```
/1.txt renamed
```

## Stats

`stats` command helps in getting upload, download and challenge information for a file.
Only the owner can get a files stats.

| Parameter  | Required | Description                 | default | Valid values |
|------------|----------|-----------------------------|---------|--------------|
| allocation | yes      | allocation id               |         | string       |
| json       | no       | print result in json format | false   | boolean      |
| remotepath | yes      | file of which to get stats  |         | string       |

<details>
  <summary>stats</summary>

![image](https://user-images.githubusercontent.com/6240686/124490093-9beacd00-dda9-11eb-8673-cf8a53475aec.png)

</details>

Example

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

## Repair

Use `start-repair` command to repair a file on dStorage.
![repair](https://user-images.githubusercontent.com/65766301/120052600-b364c680-c043-11eb-9bf2-038ab244fed6.png)
\

| Parameter  | Required | Description               | default | Valid values |
|------------|----------|---------------------------|---------|--------------|
| allocation | yes      | allocation id             |         | string       |
| repairpath | yes      | remote path to repair     |         | string       |
| rootpath   | yes      | file path for local files |         | string       |

Example

```
./zbox start-repair --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --repairpath / --rootpath /home/dung/Desktop/alloc
```

Response:

```
Repair file completed, Total files repaired:  0
```

## Add collaborator

Use `add-collab` command to add a collaborator for a file on dStorage. 
Collaborators can perform read actions on the collaboration file, with the owner paying.

![collaboration](https://user-images.githubusercontent.com/65766301/120052678-0f2f4f80-c044-11eb-8ca6-1a032659eac3.png)

| Parameter  | Required | Description                  | default | Valid values |
|------------|----------|------------------------------|---------|--------------|
| allocation | yes      | allocation id                |         | string       |
| collabid   | yes      | id of collaberator           |         | string       |
| remotepath | yes      | file on which to collaberate |         | string       |

<details>
  <summary>add-collab</summary>

![image](https://user-images.githubusercontent.com/6240686/124504210-e9be0000-ddbe-11eb-819c-e74c8bf340dd.png)

</details>

Example

```
./zbox add-collab --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --collabid d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329
```

Response will be a confirmation that collaborator is added on all blobbers for the given file .

```
Collaborator d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329 added successfully for the file /1.txt
```

You can check all collaborators for a file in metadata json response.

## Delete collaborator

Use command delete-collab to remove a collaborator for a file

| Parameter  | Required | Description                  | default | Valid values |
|------------|----------|------------------------------|---------|--------------|
| allocation | yes      | allocation id                |         | string       |
| collabid   | yes      | id of collaberator           |         | string       |
| remotepath | yes      | file on which to collaberate |         | string       |

<details>
  <summary>delete-collab</summary>

![image](https://user-images.githubusercontent.com/6240686/124505356-3571a900-ddc1-11eb-9dd8-72927cefa790.png)

</details>

Example

```
./zbox delete-collab --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --remotepath /1.txt --collabid d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329
```

Response will be a confirmation that collaborator is removed on all blobbers for the given file.

```
Collaborator d477d12134c2d7ba5ab71ac8ad37f244224695ef3215be990c3215d531c5a329 removed successfully for the file /1.txt
```

### Challenge pool information

Use `cp-info` command to get the challenge pool brief information.

| Parameter  | Required | Description                 | default | Valid values |
|------------|----------|-----------------------------|---------|--------------|
| allocation | yes      | allocation id               |         | string       |
| json       | no       | print result in json format | false   | boolean      |

<details>
  <summary>cp-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124506637-fe50c700-ddc3-11eb-9e8e-f59f88c89b6c.png)

</details>

Example

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

## Create read pool

Use `rp-create` to create a read pool, `rp-create` has no parameters.

<details>
  <summary>rp-create</summary>

![image](https://user-images.githubusercontent.com/6240686/123973204-77f74800-d9b3-11eb-8165-96741cc0b291.png)

</details>

```
./zbox rp-create
```

## Read pool info

Use `rp-info` to get read pool information.

| Parameter  | Required | Description                 | default | Valid values |
|------------|----------|-----------------------------|---------|--------------|
| allocation | no       | allocation id               |         | string       |
| json       | no       | print result in json format | false   | boolean      |

<details>
  <summary>rp-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124507524-d8c4bd00-ddc5-11eb-853e-513957cf3dbb.png)

</details>

```
./zbox rp-info
```
## Lock tokens into read pool

Lock some tokens in read pool associated with an allocation. 
* Uses two different formats, you can either define a specific blobber
  to lock all tokens, or spread across all the allocations blobbers automatically.
* If the user does not have a pre-existing read pool, then the smart-contract
  creates one.

| Parameter  | Required | Description            | default | Valid values |
|------------|----------|------------------------|---------|--------------|
| allocation | yes      | allocation id          |         | string       |
| blobber    | no       | blobber id to lock for |         | string       |
| duration   | yes      | lock duration          |         | duratation   |
| fee        |          | transaction fee        | 0       | int          |
| tokens     | yes      | tokens to lock         |         | int          |

```
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

<details>
  <summary>rp-lock with a specific blobber</summary>

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1 --blobber f65af5d64000c7cd2883f4910eb69086f9d6e6635c744e62afcfab58b938ee25 
```
![image](https://user-images.githubusercontent.com/6240686/125474085-c57c29a5-127e-4e8e-b560-c235ade869f1.png)

</details>

<details>
  <summary>rp-lock spread across all blobbers</summary>

Tokens are spread between the blobber pools weighted by 
each blobber's Terms.ReadPrice.

```shell
./zbox rp-lock --allocation <allocation_id> --duration 40m --tokens 1
```

![image](https://user-images.githubusercontent.com/6240686/125474486-1c2e1dba-7e61-4e9c-94f4-2a1ebf06d2de.png)

</details>

## Unlock tokens from read pool

Use `rp-unlock` to unlock tokens from an expired read pool by pool id. 
See `rp-info` for the POOL_ID and the expiration.

| Parameter | Required | Description          | default | Valid values |
|-----------|----------|----------------------|---------|--------------|
| fee       | no       | transaction fee      | 0       | float        |
| pool_id   | yes      | id of pool to unlock |         | string       |

<details>
  <summary>rp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/124578670-53352180-de46-11eb-99a5-07debf17e351.png)

</details>

```
./zbox rp-unlock --pool_id <pool_id>
```

## Storage SC configurations

Show storage SC configuration.

| Parameter  | Required | Description                 | default | Valid values |
|------------|----------|-----------------------------|---------|--------------|
| allocation | yes      | allocation id               |         | string       |
| json       | no       | print result in json format | false   | boolean      |

<details>
  <summary>sc-config</summary>

![image](https://user-images.githubusercontent.com/6240686/124578670-53352180-de46-11eb-99a5-07debf17e351.png)

</details>

```
./zbox sc-config
```

## Stake pool info

Use `sp-info` to get your stake pool information and settings.

| Parameter  | Required | Description                 | default        | Valid values |
|------------|----------|-----------------------------|----------------|--------------|
| blobber_id |          | id of blobber               | current client | string       |
| json       | no       | print result in json format | false          | boolean      |

<details>
  <summary>sp-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124581849-63023500-de49-11eb-8927-50d9ff97671b.png)

</details>

```
./zbox sp-info --blobber_id <blobber_id>
```

## Lock tokens into stake pool

Lock creates delegate pool for current client and given blobber. 
The tokens locked for the blobber stake can be unlocked any time, excluding 
where the tokens held by opened offers. The tokens collect interests.
`sp-lock` returns the id of the new stake pool, this will be needed to reference
to stake pool later.

| Parameter  | Required | Description     | default        | Valid values |
|------------|----------|-----------------|----------------|--------------|
| blobber_id |          | id of blobber   | current client | string       |
| fee        | no       | transaction fee | 0              | float        |
| tokens     | yes      | tokens to lock  |                | float        |

<details>
  <summary>sp-lock</summary>

![image](https://user-images.githubusercontent.com/6240686/124585686-73b4aa00-de4d-11eb-83cb-334f7c54543e.png)

</details>

```
./zbox sp-lock --blobber_id <blobber_id> --tokens 1.0
```

### Unlock tokens from stake pool

Unlock a stake pool by pool owner. If the stake pool cannot be unlocked as 
it would leave insufficient funds for opened offers, then `sp-unlock` tags 
the stake pool to be unlocked later. This tag prevents the stake pool affecting 
blobber allocation for any new allocations.

| Parameter  | Required | Description          | default        | Valid values |
|------------|----------|----------------------|----------------|--------------|
| blobber_id |          | id of blobber        | current client | string       |
| fee        | no       | transaction fee      | 0              | float        |
| pool id    | yes      | id of pool to unlock |                | string       |

<details>
  <summary>sp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/124597566-8e8e1b00-de5b-11eb-8926-867687aaa06a.png)

</details>


```
./zbox sp-unlock --blobber_id <blobber_id> --pool_id <pool_id>
```

## Stake pools info of user

Get information about all stake pools of current user.

| Parameter  | Required | Description                 | default        | Valid values |
|------------|----------|-----------------------------|----------------|--------------|
| client_id |          | id of client               | current client | string       |
| json       | no       | print result in json format | false          | boolean      |


<details>
  <summary>sp-user-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124600324-7ff53300-de5e-11eb-9b78-5a4f9c59a536.png)

</details>
```
./zbox sp-user-info
```

## Pay interests

Changes in stake pool pays all pending rewards, But if there are no changes interests will not be paid.
`sp-pay-interests`  command can be used to 
pay interest for all delegates. Use `sp-info` to check interests can be paid or not.

| Parameter  | Required | Description          | default        | Valid values |
|------------|----------|----------------------|----------------|--------------|
| blobber_id |          | id of blobber        | current client | string       |

<details>
  <summary>sp-pay-interests</summary>

![image](https://user-images.githubusercontent.com/6240686/124602256-93a19900-de60-11eb-9ddd-74f6e0570e47.png)

</details>

```
./zbox sp-pay-interests --blobber_id <blobber_id>
```

## Write pool info

Write pool information. Use allocation id to filter results to a singe allocation.

| Parameter     | Required | Description                 | default | Valid values |
|---------------|----------|-----------------------------|---------|--------------|
| allocation id | no       | allocation id               |         | string       |
| json          | no       | print result in json format | false   | boolean      |

<details>
  <summary>sp-pay-interests</summary>

![image](https://user-images.githubusercontent.com/6240686/124603444-d9ab2c80-de61-11eb-82f2-900d540ba63f.png)

</details>


For all write pools.


```
./zbox wp-info
```

Filtering by allocation.

```
./zbox wp-info --allocation <allocation_id>
```

## Lock tokens into write pool

`wp-lock` can be used to lock tokens in a write pool associated with an allocation. 
All tokens will be divided between allocation blobbers depending on their write price.
* Uses two different formats, you can either define a specific blobber
  to lock all tokens, or spread across all the allocations blobbers automatically.
* If the user does not have a pre-existing read pool, then the smart-contract
  creates one.

| Parameter     | Required | Description                       | default | Valid values |
|---------------|----------|-----------------------------------|---------|--------------|
| allocation id | no       | allocation id                     |         | string       |
| blobber       | no       | blobber id                        |         | string       |
| duration      | yes      | duration for which to lock tokens |         | duration     |
| fee           | no       | transaction fee                   | 0       | float        |
| tokens        | yes      | number of tokens to lock          |         | float        |

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
```
./zbox wp-lock --allocation <allocation_id> --duration 40m --tokens 1
```


## Unlock tokens from write pool

`wp-unlock` unlocks an expired write pool by its POOL_ID. See `wp-info` for the pool id and the expiration. 
An expired write pool, associated with an allocation, can be locked until allocation finalization even if it's expired. It possible in cases where related blobber doesn't give their min lock demands. The finalization will pay the demand and unlock the pool.

| Parameter | Required | Description          | default | Valid values |
|-----------|----------|----------------------|---------|--------------|
| fee       | no       | transaction fee      | 0       | float        |
| pool_id   | yes      | id of pool to unlock |         | string       |

<details>
  <summary>rp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/123980742-b09a2000-d9b9-11eb-8987-c18ff90ee705.png)

</details>

```
./zbox wp-unlock --pool_id <pool_id>
```

## Download cost

`get-download-cost` determines the cost for downloading the remote file from dStorage. The clinet
must either be the owner, a collaborator or be using an auth ticket.

| Parameter  | Required | Description                               | default | Valid values |
|------------|----------|-------------------------------------------|---------|--------------|
| allocation | yes      | allocation id                             |         | string       |
| authticket | no       | auth ticket to use if not the owner       |         | string       |
| lookuphash | no       | hash of remote file, use with auth ticket |         | string       |
| remotepath | no       | file of which to get stats, use if owner  |         | string       |

<details>
  <summary>get-download-cost</summary>

![image](https://user-images.githubusercontent.com/6240686/124497750-41ef0500-ddb3-11eb-99ea-115a4e234eda.png)

</details>

```
./zbox get-download-cost --allocation <allocation_id> --remotepath /path/file.ext
```

## Upload cost

`get-upload-cost` determines the cost for uploading a local file on dStorage. 
`--duration` Ignored if `--end` true, in which case the cost of upload calculated until
the allocation expires.

| Parameter  | Required | Description                          | default | Valid values |
|------------|----------|--------------------------------------|---------|--------------|
| allocation | yes      | allocation id                        |         | string       |
| duration   | no       | duration for which to upload file    |         | duration     |
| end        | no       | upload file until allocation expires | false   | boolean      |
| localpath   | yes      | local of path to calculate upload    |         | file path    |

<details>
  <summary>get-upload-cost</summary>

![image](https://user-images.githubusercontent.com/6240686/124501898-51be1780-ddba-11eb-8c1a-d238cfd8f43f.png)

</details>


```
./zbox get-upload-cost --allocation <allocation_id> --localpath ./path/file.ext
```

------

## Streaming

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

