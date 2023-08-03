# zbox - a CLI for Züs dStorage

zbox is a command line interface (CLI) tool to understand the capabilities of Züs dStorage and prototype your app. The utility is built using Züs [GoSDK](https://github.com/0chain/gosdk) . 
![Storage](https://user-images.githubusercontent.com/65766301/120052450-0ab66700-c043-11eb-91ab-1f7aa69e133a.png)

- [zbox - a CLI for Züs dStorage](#zbox---a-cli-for-züs-dstorage)
  - [Züs Overview](#züs-overview)
  - [Installation Guides](#installation-instructions)
    - [Install on Linux Windows Mac](#build-instructions-for-linux-windows-mac)
    - [Other Platform Builds](#other-platform-builds)
    - [Use custom miner/sharder](#use-custom-minersharder)
    - [Running zbox](#running-zbox)
  - [Global Flags](#global-flags)
  - [Commands](#commands)
    - [Creating and Managing Allocations](#creating-and-managing-allocations)
      - [Create new allocation](#create-new-allocation)
        - [Free storage allocation](#free-storage-allocation)
      - [Update allocation](#update-allocation)
        - [Forbid allocation](#forbid-allocation)
        - [Add Blobber](#add-blobber)
        - [Replace Blobber](#replace-blobber)
      - [Cancel allocation](#cancel-allocation)
      - [Finalise allocation](#finalise-allocation)
      - [Get Allocation Info](#get)
      - [List blobbers](#list-blobbers)
      - [Detailed blobber information](#detailed-blobber-information)
      - [List all files](#list-all-files)
      - [List owner's allocations](#list-owners-allocations)
      - [Update blobber settings](#update-blobber-settings)
      - [List all Validators](#list-all-validators)
      - [Get Validator Configuration](#get-validator-configuration)
      - [Kill Blobber](#kill-blobber)
      - [Kill Validator](#kill-validator)

    - [Uploading and Managing files](#uploading-and-managing-files)
      - [Upload](#upload)
        - [Upload file with no encryption](#upload-file-with-no-encryption)
        - [Upload file with encryption](#upload-file-with-encryption)
        - [Upload file with web streaming](#upload-file-with-web-streaming)
        - [Multi Upload](#multi-upload)
      - [Stream](#stream)
      - [Feed](#feed)
      - [Download](#download)
        - [Multi Download](#multi-download)
      - [Update](#update)
      - [Delete](#delete)
      - [Share](#share)
        - [share-encrypted revoke](#share-encrypted-revoke)
      - [List](#list)
      - [Copy](#copy)
      - [Move](#move)
      - [Sync](#sync)
      - [Get differences](#get-differences)
      - [Get wallet](#get-wallet)
      - [Get metadata](#get-metadata)
      - [Rename](#rename)
      - [Stats](#stats)
      - [Repair](#repair)
      - [Sign data](#sign-data)
      - [Download cost](#download-cost)
      - [Upload cost](#upload-cost)
      - [Streaming](#streaming)
        - [How it works:](#how-it-works)
        - [Usage](#usage)
    - [Lock and Unlock Tokens](#lock-and-unlock-tokens)
      - [Challenge pool information](#challenge-pool-information)
      - [Create read pool](#create-read-pool)
      - [Collect rewards](#collect-rewards)
      - [Read pool info](#read-pool-info)
      - [Lock tokens into read pool](#lock-tokens-into-read-pool)
      - [Unlock tokens from read pool](#unlock-tokens-from-read-pool)
      - [Storage SC configurations](#storage-sc-configurations)
      - [Stake pool info](#stake-pool-info)
      - [Lock tokens into stake pool](#lock-tokens-into-stake-pool)
      - [Unlock tokens from stake pool](#unlock-tokens-from-stake-pool)
      - [Stake pools info of user](#stake-pools-info-of-user)
      - [Write pool info](#write-pool-info)
      - [Lock tokens into write pool](#lock-tokens-into-write-pool)
      - [Unlock tokens from write pool](#unlock-tokens-from-write-pool)
      
  - [Troubleshooting](#troubleshooting)

## Züs Overview
[Züs](https://zus.network/) is a high-performance cloud on a fast blockchain offering privacy and configurable uptime. It is an alternative to traditional cloud S3 and has shown better performance on a test network due to its parallel data architecture. The technology uses erasure code to distribute the data between data and parity servers. Züs storage is configurable to provide flexibility for IT managers to design for desired security and uptime, and can design a hybrid or a multi-cloud architecture with a few clicks using [Blimp's](https://blimp.software/) workflow, and can change redundancy and providers on the fly.

For instance, the user can start with 10 data and 5 parity providers and select where they are located globally, and later decide to add a provider on-the-fly to increase resilience, performance, or switch to a lower cost provider.

Users can also add their own servers to the network to operate in a hybrid cloud architecture. Such flexibility allows the user to improve their regulatory, content distribution, and security requirements with a true multi-cloud architecture. Users can also construct a private cloud with all of their own servers rented across the globe to have a better content distribution, highly available network, higher performance, and lower cost.

[The QoS protocol](https://medium.com/0chain/qos-protocol-weekly-debrief-april-12-2023-44524924381f) is time-based where the blockchain challenges a provider on a file that the provider must respond within a certain time based on its size to pass. This forces the provider to have a good server and data center performance to earn rewards and income.

The [privacy protocol](https://zus.network/build) from Züs is unique where a user can easily share their encrypted data with their business partners, friends, and family through a proxy key sharing protocol, where the key is given to the providers, and they re-encrypt the data using the proxy key so that only the recipient can decrypt it with their private key.

Züs has ecosystem apps to encourage traditional storage consumption such as [Blimp](https://blimp.software/), a S3 server and cloud migration platform, and [Vult](https://vult.network/), a personal cloud app to store encrypted data and share privately with friends and family, and [Chalk](https://chalk.software/), a zero upfront cost permanent storage solution for NFT artists.

Other apps are [Bolt](https://bolt.holdings/), a wallet that is very secure with air-gapped 2FA split-key protocol to prevent hacks from compromising your digital assets, and it enables you to stake and earn from the storage providers; [Atlus](https://atlus.cloud/), a blockchain explorer and [Chimney](https://demo.chimney.software/), which allows anyone to join the network and earn using their server or by just renting one, with no prior knowledge required.

## Installation Instructions

### [Build Instructions for Linux Windows Mac](https://github.com/0chain/zboxcli/wiki/Build-Instructions)

### [Other Platform Builds](https://github.com/0chain/zboxcli/wiki/Alternative-Platform-Builds)

### Use custom miner/sharder

As mentioned in build guides, a ./zcn folder is created to store configuration files for zboxcli. Here is a sample network config file

```
  ---
  block_worker: https://demo.zus.network/dns
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

When you run the `./zbox` command in terminal with no arguments, it will list all the available commands and the global flags.For working of specific command check [commands](#commands) section. 

```
Usage:
  zbox [command]

Available Commands:
  add                Adds free storage assigner
  alloc-cancel       Cancel an allocation
  alloc-fini         Finalize an expired allocation
  bl-info            Get blobber info
  bl-update          Update blobber settings by its delegate_wallet owner
  collect-reward     Collect accrued rewards for a stake pool.
  completion         Generate the autocompletion script for the specified shell
  copy               copy an object(file/folder) to another folder on blobbers
  cp-info            Challenge pool information.
  createdir          Create directory
  decrypt            Decrypt text with passphrase
  delete             delete file from blobbers
  download           download file from blobbers
  feed               download segment files from remote live feed, and upload
  get-diff           Get difference of local and allocation root
  get-download-cost  Get downloading cost
  get-mpt            Directly view blockchain data
  get-upload-cost    Get uploading cost
  getallocation      Gets the allocation info
  getwallet          Get wallet information
  help               Help about any command
  kill-blobber       punitively deactivate a blobber
  kill-validator     punitively deactivate a validator
  list               list files from blobbers
  list-all           list all files from blobbers
  listallocations    List allocations for the client
  ls-blobbers        Show active blobbers in storage SC.
  ls-validators      Show active Validators.
  meta               get meta data of files from blobbers
  move               move an object(file/folder) to another folder on blobbers
  newallocation      Creates a new allocation
  recent-refs        get list of recently added refs
  rename             rename an object(file/folder) on blobbers
  rp-create          Create read pool if missing
  rp-info            Read pool information.
  rp-lock            Lock some tokens in read pool.
  rp-unlock          Unlock some expired tokens in a read pool.
  sc-config          Show storage SC configuration.
  share              share files from blobbers
  shutdown-blobber   deactivate a blobber
  shutdown-validator deactivate a validator
  sign-data          Sign given data
  sp-info            Stake pool information.
  sp-lock            Lock tokens lacking in stake pool.
  sp-unlock          Unlock tokens in stake pool.
  sp-user-info       Stake pool information for a user.
  start-repair       start repair file to blobbers
  stats              stats for file from blobbers
  stream             capture video and audio streaming form local devices, and upload
  sync               Sync files to/from blobbers
  transferallocation Transfer an allocation from one account to another
  update             update file to blobbers
  updateallocation   Updates allocation's expiry and size
  upload             upload file to blobbers
  validator-info     Get validator info
  validator-update   Update validator settings by its delegate_wallet owner
  version            Prints version information
  wp-lock            Lock some tokens in write pool.
  wp-unlock          Unlock some expired tokens in a write pool.

Flags:
      --config string              config file (default is config.yaml)
      --configDir string           configuration directory (default is $HOME/.zcn)
      --fee float                  transaction fee for the given transaction (if unset, it will be set to blockchain min fee)
  -h, --help                       help for zbox
      --network string             network file to overwrite the network details (if required, default is network.yaml)
      --silent                     Do not show interactive sdk logs (shown by default)
      --wallet string              wallet file (default is wallet.json)
      --wallet_client_id string    wallet client_id
      --wallet_client_key string   wallet client_key
      --withNonce int              nonce that will be used in transaction (default is 0)
```

### Global Flags

Global Flags are parameters in zbox that can be used with any command to override the default configuration.zbox supports the following global parameters.

| Flags                      | Description                                                                                                    | Usage                                            |
| -------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------ |
| --config string            | Specify a zbox configuration file (default is [$HOME/.zcn/config.yaml](#zcnconfigyaml))                        | zbox [command] --config config1.yaml             |
| --configDir string         | Specify a zbox configuration directory (default is $HOME/.zcn)                                                 | zbox [command] --configDir /$HOME/.zcn2          |
| -h, --help                 | Gives more information about a particular command.                                                             | zbox [command] --help                            |
| --network string           | Specify a network file to overwrite the network details(default is [$HOME/.zcn/network.yaml](#zcnnetworkyaml)) | zbox [command] --network network1.yaml           |
| --verbose                  | Provides additional details as to what the particular command is doing.                                        | zbox [command] --silent                         |
| --wallet string            | Specify a wallet file or 2nd wallet (default is $HOME/.zcn/wallet.json)                                        | zbox [command] --wallet wallet2.json             |
| --wallet_client_id string  | Specify a wallet client id (By default client_id specified in $HOME/.zcn/wallet.json is used)                  | zbox [command] --wallet_client_id <client_id>    |
| --wallet_client_key string | Specify a wallet client_key (By default client_key specified in $HOME/.zcn/wallet.json is used)                | zbox [command] --wallet_client_key < client_key> |

## Commands

Note in this document, we will only show the commands for particular functionalities,
the response will vary depending on your usage and may not be provided in all places.
To get a more descriptive view of all the zbox functionalities check zbox cli
documentation at docs.0chain.net.

### Creating and Managing Allocations

#### Create new allocation

Command `newallocation` reserves hard disk space on the blobbers. Later `upload`
can be used to save files to the blobber. `newallocation` has three modes triggered
by the presence or absence of the `cost`
and `free_storage` parameters.

- `cost` Converts `newallocation` into a query that returns the cost of the allocation
  determined by the remaining parameters.
- [`free_storage`](#free-storage-allocation) Creates an allocation using a free storage marker. All other
  parameters except `cost` will be ignored. The allocation settings will be set
  automatically by `0chain`, from preconfigured values.
- `otherwise` Creates an allocation applying the settings indicated by the
  remaining parameters.
- Use `owner` if you want to create and fund an allocation for someone else. If
  using this option then you need to provide the new owner's public key. Otherwise,
  the owner defaults to the client.

| Parameter              | Description                                                             | Default        | Valid Values |
| ---------------------- | ----------------------------------------------------------------------- | -------------- | ------------ |
| allocationFileName     | local file to store allocation information                              | allocation.txt | file path    |
| cost                   | returns the cost of the allocation, no allocation created               |                | flag         |
| data                   | number of data shards, effects upload and download speeds               | 2              | int          |
| expire                 | duration to allocation expiration                                       | 720h           | duration     |
| free_storage           | free storage marker file.                                               |                | file path    |
| owner                  | owner's id, use for funding an allocation for another                   |                | string       |
| owner_public_key       | public key, use for funding an allocation for another                   |                | string       |
| lock                   | lock write pool with given number of tokens                             |                | float        |
| parity                 | number of parity shards, effects availability                           | 2              | int          |
| read_price             | filter blobbers by read price range                                     | 0-inf          | range        |
| size                   | size of space reserved on blobbers                                      | 2147483648     | bytes        |
| usd                    | give token value in USD                                                 |                | flag         |
| write_price            | filter blobbers by write price range                                    | 0-inf          | range        |
| false                  | bool                                                                    |
| third_party_extendable | specify if the allocation can be extended by users other than the owner | false          | bool         |
| forbid_upload          | specify if users cannot upload to this allocation                       | false          | bool         |
| forbid_delete          | specify if the users cannot delete objects from this allocation         | false          | bool         |
| forbid_update          | specify if the users cannot update objects in this allocation           | false          | bool         |
| forbid_move            | specify if the users cannot move objects from this allocation           | false          | bool         |
| forbid_copy            | specify if the users cannot copy object from this allocation            | false          | bool         |
| forbid_rename          | specify if the users cannot rename objects in this allocation           | false          | bool         |

<details>
  <summary>newallocation </summary>

![allocation](https://user-images.githubusercontent.com/65766301/120052477-27529f00-c043-11eb-91bb-573558325b20.png)

![image](https://user-images.githubusercontent.com/6240686/125315476-15953480-e32f-11eb-8a11-b069079911d3.png)

</details>

<details>
  <summary>Free storage newallocation </summary>

![image](https://user-images.githubusercontent.com/6240686/127857969-1aa1a56c-4a65-4ba5-b724-943cf12594b4.png)

</details>

##### Free storage allocation

Entities can give free `0chain` storage in the form of markers. A marker takes the
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

- `assigner` A label for the entity providing the free storage.
- `recipient` The marker has to be run by the recipient to be valid.
- `free_tokens` The amount of free tokens. When creating a new allocation the
  free tokens will be split between the allocation's write pool,
  and a new read pool; the ratio of this split configured on the blockchain.
- `timestamp` A unique timestamp. Used to prevent multiple applications of the same marker.
- `signature` Signed by the assigner, validated using the stored public key on the blockchain.
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

To use a free storage marker, you only need to provide the path to the marker file.

```shell
./zbox newallocation --lock 0.5 --free_storage markers/my_marker.json
```

Response:

```
Allocation created : d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

#### Update allocation

`updateallocation` updates allocation settings. It has two modes depending on
the presence of the `free_storage` field.

- `free_storage` Uses a free storage marker to fund this allocation update; settings
  predefined by `0chain`. See [newallocation](#free-storage-allocation) for further details.
- `otherwise` Update an allocation applying the settings indicated by the
  remaining parameters.

If not a `free_storage` update, then tokens will come from those locked.
Further we can add a blobber to the allocation,
adding a blobber will allow a blobber to be removed.
An increase in blobber count will increment the parity shards.

| Parameter      | Required | Description                                                          | Valid Values |
| -------------- | -------- | -------------------------------------------------------------------- | ------------ |
| allocation     | yes      | allocation id                                                        | string       |
| expiry         |          | adjust storage expiration time                                       | duration     |
| free_storage   |          | free storage marker file                                             | string       |
| lock           | yes\*    | lock additional tokens in write pool                                 | int          |
| update_terms   |          | will update the allocation with the latest blobber terms             | boolean      |
| size           |          | adjust allocation size                                               | bytes        |
| add_blobber    |          | add a new blobber to the allocation, required for remove_blobber     | string       |
| remove_blobber |          | remove a blobber from the allocation, requires an add_blobber option | string2      |

`*` only required if free_storage not set.
| third_party_extendable | specify if the allocation can be extended by users other than the owner | false | bool
| forbid_upload | specify if users cannot upload to this allocation |false | bool
| forbid_delete | specify if the users cannot delete objects from this allocation | false | bool
| forbid_update | specify if the users cannot update objects in this allocation |false | bool
| forbid_move | specify if the users cannot move objects from this allocation |false | bool
| forbid_copy | specify if the users cannot copy object from this allocation |false | bool
| forbid_rename | specify if the users cannot rename objects in this allocation |false | bool

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

##### Forbid Allocation

There are various operations which you can forbid on an allocation. Forbid flag works with [update allocation](#update-allocation) command. Check its working first.
Here are the operations:

| Parameter       | Description                                                  |
| --------------- | ------------------------------------------------------------ |
| --forbid_copy   | specify if the users cannot copy object from this allocation |
| --forbid_update | specify if the users cannot update objects in this allocation |
| --forbid_delete | specify if the users cannot delete objects from this allocation |
| --forbid_move   | specify if the users cannot move objects from this allocation |
| --forbid_rename | specify if the users cannot rename objects in this allocation |
| --forbid_upload | specify if users cannot upload to this allocation            |

Here is a sample command for --forbid_upload .Other parameters can be done the same way.

```
./zbox updateallocation --allocation $ALLOC --forbid_upload
```
Sample Response :
```
Allocation Updated with txID : b84185dae620bbba8386286726f1efcd20d2516bcf1a448215434d87be3b30d
```
To test functionality try uploading file to allocation. You should get the following response :
```
Upload failed. this options for this file is not permitted for this allocation: 
file_option_not_permitted.
```
##### Add Blobber

Use `add_blobber` flag with [update allocation](#update-allocation) command to add blobber to allocation.The new blobber will be added as a parity blobber.For more details [check how a file is stored on blobbers](https://docs.zus.network/concepts/store).

Here are the necessary parameters for adding blobber.

| Parameter     | Description                                                 | Valid Values |
| ------------- | ----------------------------------------------------------- | ------------ |
| --allocation  | Provide Allocation ID for adding blobber to allocation      | string       |
| --add_blobber | Provide Blobber ID to add. Can be fetched using [List blobbers](#list-blobbers). | string       |

Sample Command:

```
./zbox updateallocation --allocation $ALLOC --add_blobber 98f14362f075caf467653044cf046eb9e8a5dfee88dc8b78cad1891748245003
```

Sample Response:

```
Allocation updated with txId : d853a82907453d37ed978b9fc1a55663be99bb351d18cca31068c0dc95566edd
```

**Note:** Files will automatically be uploaded,splitted and stored on added blobber. 
**Note:** An allocation is already hosted on a set of blobbers. To find a blobber that is available to add you should exclude the current set of blobbers hosting your allocation by checking them via [Get Allocation Info](#get)command.

##### Replace Blobber

Sometimes, a blobber might be mallfunctioning or faulty or the blobber might be slow becuase its far from your geolocation, in such cases you might have to replace the blobber with a new one.

Use `add_blobber` and `remove_blobber` flag with [update allocation](#update-allocation) command to replace a blobber hosting an allocation .

| Parameter        | Description                                                  | Valid Values |
| ---------------- | ------------------------------------------------------------ | ------------ |
| --allocation     | Provide Allocation ID for replacing blobber                  | string       |
| --add_blobber    | Provide new Blobber ID . Can be fetched using [List blobbers](#list-blobbers). | string       |
| --remove_blobber | Provide ID for the blobber which has to remove               | string       |

Sample Command:

```
./zbox updateallocation --allocation $ALLOC --add_blobber 8d19a8fd7147279d1dfdadd7e3ceecaf91c63ad940dae78731e7a64b104441a6 --remove_blobber 06166f3dfd72a90cd0b51f4bd7520d4434552fc72880039b1ee1e8fe4b3cd7ea
```

Sample Response:

```
allocation updated successfully
```
**Note:**To find a blobber that can be replaced you should check the current set of blobbers hosting your allocation with [Get Allocation Info](#get)command.


#### Cancel allocation

`alloc-cancel` immediately return all remaining tokens from challenge pool back to the
allocation's owner and cancels the allocation. If blobbers already got some tokens,
the tokens will not be returned. Remaining min lock payment to the blobber will be
funded from the allocation's write pools.

Cancelling an allocation can only occur if the amount of failed challenges exceed a preset threshold.

| Parameter  | Required | Description   | Valid Values |
| ---------- | -------- | ------------- | ------------ |
| allocation | yes      | allocation id | string       |

<details>
  <summary>alloc-cancel</summary>

![image](https://user-images.githubusercontent.com/6240686/127854863-eff675aa-17ec-4251-a084-ffa400e8daa1.png)

</details>

Example

```
./zbox alloc-cancel --allocation <allocation_id>
```

#### Finalise allocation

`alloc-fini` finalises an expired allocation. An allocation becomes expired when
the expiry time has passed followed by a period equal to the challenge completion
period.

Any remaining min lock payment to the blobber will be funded from the
allocation's write pools. Any available money in the challenge pool returns to
the allocation's owner.

An allocation can be finalised by the owner or one of the allocation blobbers.

| Parameter  | Required | Description   | Valid Values |
| ---------- | -------- | ------------- | ------------ |
| allocation | yes      | allocation id | string       |

<details>
  <summary>alloc-fini</summary>

![image](https://user-images.githubusercontent.com/6240686/127855881-e7578512-d972-49c3-8920-843a04d740f1.png)

</details>

Example

```
./zbox alloc-fini --allocation <allocation_id>
```
#### List blobbers

Use `ls-blobbers` command to show active blobbers.

| Parameter | Required | Description                          | Valid Values |
| --------- | -------- | ------------------------------------ | ------------ |
| all       | no       | shows active and non active blobbers | flag         |
| json      | no       | display result in .json format       | flag         |

<details>
  <summary>ls-blobbers</summary>

![image](https://user-images.githubusercontent.com/6240686/127860789-94d8118c-93d4-4afe-ae09-05d1e4f053d7.png)

</details>

Example

```
./zbox ls-blobbers
- id:                    0ece681f6b00221c5567865b56040eaab23795a843ed629ce71fb340a5566ba3
  url:                   http://demo.zus.network:31302
  used / total capacity: 101.7 GiB / 1000.0 GiB
  terms:
    read_price:          0.01 tok / GB
    write_price:         0.01 tok / GB / time_unit
    min_lock_demand:     0.1
    max_offer_duration:  744h0m0s
- id:                    788b1deced159f12d3810c61b4b8d381e80188c470e9798939f2e5036d964ffc
  url:                   http://demo.zus.network:31301
  used / total capacity: 102.7 GiB / 1000.0 GiB
  terms:
    read_price:          0.01 tok / GB
    write_price:         0.01 tok / GB / time_unit
    min_lock_demand:     0.1
    max_offer_duration:  744h0m0s
```

#### Detailed blobber information

Use `bl-info` command to get detailed blobber information.

| Parameter  | Required | Description                         | default | Valid values |
| ---------- | -------- | ----------------------------------- | ------- | ------------ |
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
url:               http://localhost:5051
capacity:          1.0 GiB
last_health_check: 2021-04-08 22:54:50 +0700 +07
capacity_used:     0 B
terms:
  read_price:         0.01 tok / GB
  write_price:        0.1 tok / GB
  min_lock_demand:    10 %
  max_offer_duration: 744h0m0s
settings:
  delegate_wallet: 8b87739cd6c966c150a8a6e7b327435d4a581d9d9cc1d86a88c8a13ae1ad7a96
  min_stake:       1 tok
  max_stake:       100 tok
  num_delegates:   50
  service_charge:  30 %
```

#### List all files

`list-all` lists all the files stored with an allocation

| Parameter  | Required | Description                                    | Default | Valid values |
| ---------- | -------- | ---------------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id, sender must be allocation owner |         | string       |

```shell
./zbox list-all --allocation 4ebeb69feeaeb3cd308570321981d61beea55db65cbeba4ba3b75c173c0f141b
```

#### List owner's allocations

`listallocations` provides a list of all allocations owned by the user.

| Parameter | Required | Description                 | default | Valid values |
| --------- | -------- | --------------------------- | ------- | ------------ |
| json      | no       | print output in json format |         | boolean      |

<details>
  <summary>listallocations</summary>

![image](https://user-images.githubusercontent.com/6240686/127861831-ba36f343-0210-442e-8d12-d580e46415a3.png)

</details>

```shell
./zbox listallocations
ZED | CANCELED | R  PRICE |   W  PRICE
+------------------------------------------------------------------+-----------+-------------------------------+------------+--------------+-----------+----------+----------+--------------+
  4ebeb69feeaeb3cd308570321981d61beea55db65cbeba4ba3b75c173c0f141b | 104857600 | 2021-07-16 13:34:29 +0100 BST |          1 |            1 | false     | false    |     0.02 | 0.1999999998
```

#### Update blobber settings

Use `./zbox bl-update ` to update a blobber's configuration settings. This updates the settings
on the blockchain not the blobber.

| Parameter          | Required | Description                               | default | Valid values |
| ------------------ | -------- | ----------------------------------------- | ------- | ------------ |
| blobber_id         | yes      | id of blobber of which to update settings |         | string       |
| capacity           | no       | update blobber capacity                   |         | int          |
| max_offer_duration | no       | update max offer duration                 |         | duration     |
| max_stake          | no       | update maximum stake                      |         | float        |
| min_lock_demand    | no       | update minimum lock demand                |         | float        |
| min_stake          | no       | update minimum stake                      |         | float        |
| num_delegates      | no       | update maximum number of delegates        |         | int          |
| read_price         | no       | update read price                         |         | float        |
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

Get Version

The version of Zbox and Gosdk can be fetched using the ./zbox version command.

```
./zbox version
```
Sample Response:

```
zbox....:  v1.4.3
gosdk...:  v1.8.14
```

#### List All Validators

List all active validators on the network

Command:
```
./zbox ls-validators
```
Response :
```
id:                b9f4f244e2e483548795e42dad0c5b5bb8f5c25d70cadeafc202ce6011b7ff8c
url:               https://demo.zus.network/validator03/
settings:
  delegate_wallet: 9c693cb14f29917968d6e8c909ebbea3425b4c1bc64b6732cadc2a1869f49be9
  min_stake:       1.000 ZCN
  max_stake:       100.000 ZCN
  num_delegates:   50
  service_charge:  30 %
id:                c025fad27d3daa6fbe6a10ef38f1075dc5a6386760951816ece953391ff9804b
url:               https://demo.zus.network/validator02/
settings:
  delegate_wallet: 9c693cb14f29917968d6e8c909ebbea3425b4c1bc64b6732cadc2a1869f49be9
  min_stake:       1.000 ZCN
  max_stake:       100.000 ZCN
  num_delegates:   50
  service_charge:  30 %
```

#### Get Validator Configuration

`./zbox validator-info` command is used to get a particular validator configuration . Here are the parameters for the command . 

| Parameter          | Required | Description                               
| ------------------ | -------- | ----------------------------------------- 
| --validator_id     | yes      | id of validator whose configuration has to be fetched          
| --json             | optional | Print Response as json data                            
| --help             | no       | Provide information about the command                          

Sample Command :
```
./zbox validator-info --validator_id f82ab34a98406b8757f11513361752bab9cb679a5cb130b81
```
Sample Response :  
```
id:                f82ab34a98406b8757f11513361752bab9cb679a5cb130b81a4e86cec50eefc3
url:               https://demo2.zus.network/validator01
last_health_check:  2023-05-12 20:09:15 +0530 IST
is killed:         false
is shut down:      false
settings:
  delegate_wallet: 9c693cb14f29917968d6e8c909ebbea3425b4c1bc64b6732cadc2a1869f49be9
  min_stake:       0 SAS
  max_stake:       0 SAS
  total_stake:     200000000000
  total_unstake:   0
  num_delegates:   50
  service_charge:  10 %
```
#### Kill Blobber
`./zbox kill-blobber` command deactivates a blobber to avoid storage of data . Required parameters are :

| Parameter          | Required | Description                               
| ------------------ | -------- | ----------------------------------------- 
| --blobber_id       | yes      | Blobber Id to kill a specific blobber.Can be retrieved using [List blobbers](#list-blobbers).    
| --json             | optional | Print Response as json data                            
| --help             | no       | Provide information about the command 
 
 Sample Command :
```
./zbox kill-blobber --id $BLOBBER_ID --wallet $CHAIN_OWNER_WALLET
```
Note : Kill Blobber command should be evoked from chain owner wallet only

Sample Response :  
```
killed blobber $BLOBBER_ID
```
 
#### Kill Validator

`./zbox kill-validator` command deactivates a specific validator available on the network. Required parameters are : 

| Parameter          | Required | Description                               
| ------------------ | -------- | ----------------------------------------- 
| --validator_id     | yes      | Validator Id to kill a specific blobber.Can be retrieved using [List all Validators](#list-all-validators).    
| --json             | optional | Print Response as json data                            
| --help             | no       | Provide information about the command 


Sample Command :
```
./zbox kill-validator --id $VALIDATOR_ID --wallet $CHAIN_OWNER_WALLET
```
Sample Response :
```
killed validator, id: $VALIDATOR_ID
```
### Uploading and Managing files

#### Upload

Use `upload` command to upload file(s).

- upload a local file
- download segment files from remote live feed, and upload them
- start live streaming from local devices, encode it into segment files with `ffmpeg`, and upload them.

The user must be the owner of the allocation.You can request the file be encrypted before upload, and can send thumbnails with the file.

| Parameter     | Required | Description                                             | Default | Valid values |
| ------------- | -------- | ------------------------------------------------------- | ------- | ------------ |
| allocation    | yes      | allocation id, sender must be allocation owner          |         | string       |
| encrypt       | no       | encrypt file before upload                              | false   | boolean      |
| web-streaming | no       | transcode file before upload to fragmented mp4          | false   | boolean      |
| localpath     | yes      | local path of the file to upload                        |         | file path    |
| remotepath    | yes      | remote path to upload file to, use to access file later |         | string       |
| thumbnailpath | no       | local path of thumbnaSil                                |         | file path    |
| chunknumber   | no       | how many chunks should be uploaded in a http request    | 1       | int          |

<details>
  <summary>upload</summary>

![image](https://user-images.githubusercontent.com/6240686/124287350-cf2e2180-db47-11eb-8079-40f069a5e0c2.png)

</details>



#####Upload file with no encryption

```
./zbox upload --localpath /absolute-path-to-local-file/hello.txt --remotepath /myfiles/hello.txt --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

Response:

```
12390 / 12390 [================================================================================] 100.00% 3s
Status completed callback. Type = application/octet-stream. Name = hello.txt
```

#####Upload file with encryption

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

#####Upload file with web-streaming

Use upload command with optional web-streaming parameter to upload a video file in fragmented
mp4 format to support streaming from browser.

```
./zbox upload --web-streaming --localpath <absolute path to file>/samplevideo.mov --remotepath /myfile/ --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac
```

Response:

```
15691733 / 15691733 [=====================================================================================] 100.00% 32s
Status completed callback. Type = video/fmp4. Name = raw.samplevideo.mp4
```
##### Multi Upload 

Use `./zbox upload ` to upload multiple files to allocation at once via json file.

Here are the parameters for multi-upload:

| Parameter         | Description                                                | Valid Values |
| ----------------- | ---------------------------------------------------------- | ------------ |
| --allocation      | Provide Allocation ID for uploading multiple files         | string       |
| --multiuploadjson | Path to  JSON file containing details for files to upload. | string       |

Here is a sample json file for Multi Upload:

```
[
  {
    "remotePath": "/raw.file_example_MP4_1920_18MG.mp4",
    "localPath": "./a.mp4",
    "downloadOp": 1
  },
  {
    "remotePath": "/a.mp4",
    "localPath": "./b.mp4",
    "downloadOp": 1
  },
  {
    "remotePath": "/raw.file_example_MP4_1920_18MG.mp4",
    "localPath": "./c.mp4",
    "downloadOp": 1
  }
]
```

`remotepath`: Remote path to upload file to on Allocation.

`localpath`: Local path of the file to upload

`downloadOp`: Can pass two values 1 or 2. 1 is for downloading the actual file, 2 is for downloading the thumbnail only.

Sample Command:

```
./zbox upload --allocation $alloc --multiuploadjson ./multi-upload.json
```
#### Stream

Use `stream` to capture video and audio streaming form local devices, and upload

The user must be the owner of the allocation.You can request the file be encrypted before upload, and can send thumbnails with the file.

| Parameter     | Required | Description                                                  | Default | Valid values |
| ------------- | -------- | ------------------------------------------------------------ | ------- | ------------ |
| allocation    | yes      | allocation id, sender must be allocation owner               |         | string       |
| encrypt       | no       | encrypt file before upload                                   | false   | boolean      |
| localpath     | yes      | local path of segment files to download, generate and upload |         | file path    |
| remotepath    | yes      | remote path to upload file to, use to access file later      |         | string       |
| thumbnailpath | no       | local path of thumbnaSil                                     |         | file path    |
| chunknumber   | no       | how many chunks should be uploaded in a http request         | 1       | int          |
| delay         | no       | set segment duration to seconds.                             | 5       | int          |

<details>
  <summary>stream</summary>

![image](https://github.com/0chain/blobber/wiki/uml/usecase/live_upload_live.png)

</details>

#### Feed

Use `feed` command to automatically download segment files from remote live feed with `--downloader-args "-q -f best"`

- encode them into new segment files with `--delay` and `--ffmpeg-args`, and upload.
- please use `youtube-dl -F https://www.youtube.com/watch?v=pC5mGB5enkw` to list formats of video (see below).

```
[youtube] pC5mGB5enkw: Downloading webpage
[info] Available formats for pC5mGB5enkw:
format code  extension  resolution note
249          webm       audio only tiny   44k , webm_dash container, opus @ 44k (48000Hz), 95.21MiB
250          webm       audio only tiny   59k , webm_dash container, opus @ 59k (48000Hz), 127.05MiB
251          webm       audio only tiny  123k , webm_dash container, opus @123k (48000Hz), 264.98MiB
140          m4a        audio only tiny  129k , m4a_dash container, mp4a.40.2@129k (44100Hz), 277.82MiB
278          webm       256x136    144p   87k , webm_dash container, vp9@  87k, 30fps, video only, 188.78MiB
160          mp4        256x136    144p  118k , mp4_dash container, avc1.4d400c@ 118k, 30fps, video only, 253.62MiB
242          webm       426x224    240p  190k , webm_dash container, vp9@ 190k, 30fps, video only, 409.20MiB
133          mp4        426x224    240p  252k , mp4_dash container, avc1.4d400d@ 252k, 30fps, video only, 541.15MiB
243          webm       640x338    360p  326k , webm_dash container, vp9@ 326k, 30fps, video only, 701.53MiB
134          mp4        640x338    360p  576k , mp4_dash container, avc1.4d401e@ 576k, 30fps, video only, 1.21GiB
244          webm       854x450    480p  649k , webm_dash container, vp9@ 649k, 30fps, video only, 1.36GiB
135          mp4        854x450    480p 1028k , mp4_dash container, avc1.4d401f@1028k, 30fps, video only, 2.16GiB
247          webm       1280x676   720p 1320k , webm_dash container, vp9@1320k, 30fps, video only, 2.77GiB
136          mp4        1280x676   720p 1988k , mp4_dash container, avc1.64001f@1988k, 30fps, video only, 4.17GiB
248          webm       1920x1012  1080p 2527k , webm_dash container, vp9@2527k, 30fps, video only, 5.30GiB
137          mp4        1920x1012  1080p 4125k , mp4_dash container, avc1.640028@4125k, 30fps, video only, 8.64GiB
271          webm       2560x1350  1440p 7083k , webm_dash container, vp9@7083k, 30fps, video only, 14.84GiB
313          webm       3840x2026  2160p 13670k , webm_dash container, vp9@13670k, 30fps, video only, 28.65GiB
18           mp4        640x338    360p  738k , avc1.42001E, 30fps, mp4a.40.2 (44100Hz), 1.55GiB
22           mp4        1280x676   720p 2117k , avc1.64001F, 30fps, mp4a.40.2 (44100Hz) (best)
```

`--downloader-args "-f 22"` dowloads video with `22           mp4        1280x676   720p 2117k , avc1.64001F, 30fps, mp4a.40.2 (44100Hz) (best)`

The user must be the owner of the allocation.You can request the file to be encrypted before upload, and can send thumbnails with the file.

| Parameter       | Required | Description                                                           | Default           | Valid values                                                                       |
| --------------- | -------- | --------------------------------------------------------------------- | ----------------- | ---------------------------------------------------------------------------------- |
| allocation      | yes      | allocation id, sender must be allocation owner                        |                   | string                                                                             |
| encrypt         | no       | encrypt file before upload                                            | false             | boolean                                                                            |
| localpath       | yes      | local path of segment files to download, generate and upload          |                   | file path                                                                          |
| remotepath      | yes      | remote path to upload file to, use to access file later               |                   | string                                                                             |
| thumbnailpath   | no       | local path of thumbnaSil                                              |                   | file path                                                                          |
| chunknumber     | no       | how many chunks should be uploaded in a http request                  | 1                 | int                                                                                |
| delay           | no       | set segment duration to seconds.                                      | 5                 | int                                                                                |
| feed            | no       | set remote live feed to url.                                          | false             | url                                                                                |
| downloader-args | no       | pass args to youtube-dl to download video. default is \"-q -f best\". | -q -f best        | [youtube-dl](https://github.com/ytdl-org/youtube-dl/blob/master/README.md#options) |
| ffmpeg-args     | no       | pass args to ffmpeg to build segments.                                | -loglevel warning | [ffmpeg](https://www.ffmpeg.org/ffmpeg.html)                                       |

<details>
  <summary>feed</summary>

![image](https://github.com/0chain/blobber/wiki/uml/usecase/live_upload_sync.png)

</details>

Example

```
./zbox feed --localpath <absolute path to file>/tvshow.m3u8 --remotepath /videos/tvsho --allocation d0939e912851959637257573b08c748474f0dd0ebbc8e191e4f6ad69e4fdc7ac  --delay 10 --downloader-args "-f 22" --feed https://www.youtube.com/watch?v=pC5mGB5enkw
```

#### Download

Use `download` command to download your own or a shared file.

- `owner` The owner of the allocation can always download files, in this case the owner pays for the download.
- `collaborator` A collaborator can download files, the owner pays. To add collaborators to an allocation, use
  [add-collab](#add-collaborator).
- `authticket` To download a file using `authticket`, you must have previous be given an auth
  ticket using the [share](#share) command. Use rx_pay to indicate who pays, `rx_pay = true` you pay,
  `rx_pay = false` the allocation owner pays.
  Use `startblock` and `endblock` to only download part of the file.

| Parameter       | Required | Description                                                                                             | Default | Valid values |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------- | ------- | ------------ |
| allocation      | yes      | allocation id                                                                                           |         | string       |
| authticket      | no       | auth ticked if not owner of the allocation, use share to get auth ticket                                |         | string       |
| blockspermarker | no       | download multiple blocks per marker                                                                     | 10      | int          |
| endblock        | no       | download until specified block number                                                                   |         | int          |
| localpath       | yes      | local path to which to download the file to                                                             |         | file path    |
| remotepath      | yes      | remote path to which the file was uploaded                                                              |         | string       |
| rx_pay          | no       | `authticket` must be valid, true = sender pays, false = allocation owner pays                           | false   | boolean      |
| startblock      | no       | start download from specified block                                                                     |         | int          |
| thumbail        | no       | only download the thumbnail                                                                             | false   | boolean      |
| live            | no       | start m3u8 downloader,and automatically generate media playlist(m3u8) on --localpath                    | false   | boolean      |
| delay           | no       | pass segment duration to generate media playlist(m3u8). only works with --live. default duration is 5s. | 5       | int          |

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

Downloaded file will be in the location specified by the `localpath` argument.

##### Multi Download

Use `./zbox download ` to download multiple files to local system at once via json file.

Here are the parameters for multi-download:

| Parameter           | Description                                                  | Valid Values |
| ------------------- | ------------------------------------------------------------ | ------------ |
| --allocation        | Provide Allocation ID for downloading multiple files from allocation | string       |
| --multidownloadjson | Path to  JSON file containing details for files to download. | string       |

Here is a sample json file for multi download:

```
[
  {
    "remotePath": "/raw.file_example_MP4_1920_18MG.mp4",
    "localPath": "./a.mp4",
    "downloadOp": 1
  },
  {
    "remotePath": "/a.mp4",
    "localPath": "./b.mp4",
    "downloadOp": 1
  },
  {
    "remotePath": "/raw.file_example_MP4_1920_18MG.mp4",
    "localPath": "./c.mp4",
    "downloadOp": 1
  }
]
```
`remotepath`: Remote path of downloaded file.

`localpath`: Local path where file has to be downloaded.

`downloadOp`: Can pass two values 1 or 2. 1 is for downloading the actual file, 2 is for downloading the thumbnail only.

Sample Command:

```
zbox download --multidownloadjson ./multi-download.json --allocation $ALLOC
```

#### Update

Use `update` command to update content of an existing file in the remote path.
Like [upload](#upload) command. Only the owner of the allocation or a collaborator
can update a file. To add collaborators to an allocation, use
[add-collab](#add-collaborator).

| Parameter     | Required | Description                                          | Default | Valid values |
| ------------- | -------- | ---------------------------------------------------- | ------- | ------------ |
| allocation    | yes      | allocation id                                        |         | string       |
| encrypt       | no       | encrypt file before upload                           | false   | boolean      |
| localpath     | yes      | local file to upload                                 |         | file path    |
| remotepath    | yes      | remote file to upload                                |         | string       |
| thumbnailpath | no       | local fumbnail file to upload                        |         | file path    |
| chunknumber   | no       | how many chunks should be uploaded in a http request | 1       | int          |

<details>
  <summary>update</summary>

![image](https://user-images.githubusercontent.com/6240686/124354473-14b02480-dc04-11eb-9463-5a91d4f6f02d.png)

</details>

#### Delete

Use `delete` command to delete your file on the allocation. Only the owner
of the application can delete a file.

| Parameter  | Required | Description                   | Default | Valid values |
| ---------- | -------- | ----------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                 |         | string       |
| remotepath | yes      | remote path of file to delete |         | string       |

<details>
  <summary>delete</summary>

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

#### Share

![Alt text](documents/share_cli.png?raw=true 'Share')

Use share command to generate an authtoken that provides authorization to the holder to the specified file on the remotepath.

- --allocation string Allocation ID
- --clientid string ClientID of the user to share with. Leave blank for public share
- --encryptionpublickey string Encryption public key of the client you want to share with (from [getwallet](#Get-wallet) command )
- --remotepath string Remote path to share
- --expiration-seconds number The seconds after which the ticket will expire(defaults to number of seconds in 90 days
  if option not provided)
- --available-after Timelock for private file that makes the file available for download at certain time. 4 input formats are supported: +1h30m, +30, 1647858200 and 2022-03-21 10:21:38

`auth ticket` can be used with [download](#download), and [list](#list),
[meta](#get-metadata) and [get_download_cost](#download-cost), but only for files in
the pre-defined remote path.

| Parameter           | Required | Description                                                                                                                                                                                               | Valid values |
| ------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ |
| allocation          | yes      | allocation id                                                                                                                                                                                             | string       |
| clientid            | no       | id of user to share file with, leave blank for public share                                                                                                                                               | string       |
| encryptionpublickey | no       | public key of the client to share file with, required if clientId                                                                                                                                         | string       |
| expiration-seconds  | no       | seconds before `auth ticket` expires                                                                                                                                                                      | int          |
| remotepath          | yes      | remote path of file to share                                                                                                                                                                              | string       |
| revoke              | no       | revoke share for remote path                                                                                                                                                                              | flag         |
| available-after     | no       | timelock for private file that makes the file available for download at certain time. 4 input formats are supported: +1h30m, +30, 1647858200 and 2022-03-21 10:21:38. default value is current local time | string       |

<details>
  <summary>share</summary>

![image](https://user-images.githubusercontent.com/6240686/127869637-323e5eae-7306-40a2-a552-86726f19a4a4.png)

</details>

Example

**Public share**

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt
```

Response:

```
Auth token eyJjbGllbnRfaWQiOiIiLCJvd25lcl9pZCI6IjE3ZTExOTQwNmQ4ODg3ZDAyOGIxNDE0YWNmZTQ3ZTg4MDhmNWIzZjk4Njk2OTk4Nzg3YTIwNTVhN2VkYjk3YWYiLCJhbGxvY2F0aW9uX2lkIjoiODlkYjBjZDI5NjE4NWRkOTg2YmEzY2I0ZDBlODE0OTE3NmUxNmIyZGIyMWEwZTVjMDZlMTBmZjBiM2YxNGE3NyIsImZpbGVfcGF0aF9oYXNoIjoiM2NhNzIyNTQwZTY1M2Y3NTQ1NjI5ZjBkYzE5ZGY2ODk5ZTI0MDRjNDI4ZDRiMWZlMmM0NjI3ZGQ3MWY3ZmQ2NCIsImFjdHVhbF9maWxlX2hhc2giOiIyYmM5NWE5Zjg0NDlkZDEyNjFmNmJkNTg3ZjY3ZTA2OWUxMWFhMGJiIiwiZmlsZV9uYW1lIjoidGVzdC5wZGYiLCJyZWZlcmVuY2VfdHlwZSI6ImYiLCJleHBpcmF0aW9uIjoxNjM1ODQ5MzczLCJ0aW1lc3RhbXAiOjE2MjgwNzMzNzMsInJlX2VuY3J5cHRpb25fa2V5IjoiIiwiZW5jcnlwdGVkIjpmYWxzZSwic2lnbmF0dXJlIjoiZDRiOTM4ZTE0MDk0ZmZkOGFiMDcwOWFmN2QyMDAyZTdlMGFmNmU3MWJlNGFmMmRjNmUxMGYxZWJmZTUwOTMxOSJ9
```

```
Auth token decoded

{"client_id":"","owner_id":"17e119406d8887d028b1414acfe47e8808f5b3f98696998787a2055a7edb97af","allocation_id":"89db0cd296185dd986ba3cb4d0e8149176e16b2db21a0e5c06e10ff0b3f14a77","file_path_hash":"3ca722540e653f7545629f0dc19df6899e2404c428d4b1fe2c4627dd71f7fd64","actual_file_hash":"2bc95a9f8449dd1261f6bd587f67e069e11aa0bb","file_name":"test.pdf","reference_type":"f","expiration":1635849373,"timestamp":1628073373,"re_encryption_key":"","encrypted":false,"signature":"d4b938e14094ffd8ab0709af7d2002e7e0af6e71be4af2dc6e10f1ebfe509319"}
```

**Encrypted share**

Upload file with _--encrypted_ tag.

Get encryptionpublickey first, by calling from user you sharing with:

```
./zbox getwallet
```

Response:

```
PUBLIC KEY | CLIENTID | ENCRYPTION PUBLIC KEY
-----------------------------------------------------------------------------------------------------------------------------------+------------------------------------------------------------------+-----------------------------------------------
  19cd2396df9b8b77358a1110492ff65cbb5b55cae06b8bd204e0969b2454851ca620ae74aebe9ed641166be3bca056a1855610f6154f4f4435a29565a2111282 | b734ef935e2a02892b2fa31e3488b360ef300d3b0b32c03834cea3a83e2453f0 | 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=
```

You have to pickup _ENCRYPTION PUBLIC KEY_

Use _clientid_ of the user to share with. _encryptionpublickey_ - key from command above.

![Private File Sharing](https://user-images.githubusercontent.com/65766301/120052575-962ff800-c043-11eb-9cf7-433383d532a3.png)

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt --clientid b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae --encryptionpublickey 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=
```

Response:

```
Auth token eyJjbGllbnRfaWQiOiIwMGZmODhkY2IxNjQ2Y2RlZjA2OWE4MGE0MGQwMWNlOTYyMmQ3ZmUzYmQ0ZWNjMzIzYTcwZTdkNmVkMWE2YjY3Iiwib3duZXJfaWQiOiIxN2UxMTk0MDZkODg4N2QwMjhiMTQxNGFjZmU0N2U4ODA4ZjViM2Y5ODY5Njk5ODc4N2EyMDU1YTdlZGI5N2FmIiwiYWxsb2NhdGlvbl9pZCI6Ijg5ZGIwY2QyOTYxODVkZDk4NmJhM2NiNGQwZTgxNDkxNzZlMTZiMmRiMjFhMGU1YzA2ZTEwZmYwYjNmMTRhNzciLCJmaWxlX3BhdGhfaGFzaCI6IjM2Mjk0MGMwMTZlOWZlZTQ4ZmI5MTA0OGI4MzJjOGFlNWQ2MGUyYzUzMmQ1OGNlYzdmNGM0YjBmZTRkZjM2MzYiLCJhY3R1YWxfZmlsZV9oYXNoIjoiMmJjOTVhOWY4NDQ5ZGQxMjYxZjZiZDU4N2Y2N2UwNjllMTFhYTBiYiIsImZpbGVfbmFtZSI6InRlc3QyLnBkZiIsInJlZmVyZW5jZV90eXBlIjoiZiIsImV4cGlyYXRpb24iOjE2MzU4NDk4NDMsInRpbWVzdGFtcCI6MTYyODA3Mzg0MywicmVfZW5jcnlwdGlvbl9rZXkiOiIiLCJlbmNyeXB0ZWQiOnRydWUsInNpZ25hdHVyZSI6IjNlNGMwOTAwMzAwN2M5NzUzZjFiNGIwODExMWM4OGRlY2JmZjU2MDRmNTIwZDZjMmYyMTdhMzUyZTFkMmE0MTEifQ==
```

```
Auth token decoded

{"client_id":"00ff88dcb1646cdef069a80a40d01ce9622d7fe3bd4ecc323a70e7d6ed1a6b67","owner_id":"17e119406d8887d028b1414acfe47e8808f5b3f98696998787a2055a7edb97af","allocation_id":"89db0cd296185dd986ba3cb4d0e8149176e16b2db21a0e5c06e10ff0b3f14a77","file_path_hash":"362940c016e9fee48fb91048b832c8ae5d60e2c532d58cec7f4c4b0fe4df3636","actual_file_hash":"2bc95a9f8449dd1261f6bd587f67e069e11aa0bb","file_name":"test2.pdf","reference_type":"f","expiration":1635849843,"timestamp":1628073843,"re_encryption_key":"","encrypted":true,"signature":"3e4c09003007c9753f1b4b08111c88decbff5604f520d6c2f217a352e1d2a411"}
```

Response contains an auth ticket- an encrypted string that can be shared.

**Directory share**

Follow up steps above to get _encryptionpublickey_

Upload multiple files to directory with _zbox upload /folder1/file1.z ..._

```
./zbox share --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /folder1 --clientid b6de562b57a0b593d0480624f79a55ed46dba544404595bee0273144e01034ae --encryptionpublickey 1JuT4AbQnmIaOMTuWn07t98xQRsSqXAxZYfwCI1yQLM=
```

Response:

Encoded

```
eyJjbGllbnRfaWQiOiJiNzM0ZWY5MzVlMmEwMjg5MmIyZmEzMWUzNDg4YjM2MGVmMzAwZDNiMGIzMmMwMzgzNGNlYTNhODNlMjQ1M2YwIiwib3duZXJfaWQiOiI2MzlmMjcxZmU1MTFjZDE4ODBjMmE0ZDhlYTRhNGYyNDBmYWYzMzY1YzYxYjY1YjQyNWZhYjVlMDIzMTcxM2MzIiwiYWxsb2NhdGlvbl9pZCI6IjkzN2FkNjlmYjIwZGMxMTFiY2ZkMDFkZTQyYzc5MmEwYzJiNDQxZGUzZDNjZjRjZGIzZjI1YzIxYzFhYjRiN2IiLCJmaWxlX3BhdGhfaGFzaCI6ImFkMThmMzg1Y2I2MWM4MTNjMzE0NDU2OTM0NWYxYzQ2ODE1ODljNzM0N2JkNzI4NjkyZTg1ZjFiNzM4NmI2OWQiLCJhY3R1YWxfZmlsZV9oYXNoIjoiIiwiZmlsZV9uYW1lIjoiZm9sZGVyMSIsInJlZmVyZW5jZV90eXBlIjoiZCIsImV4cGlyYXRpb24iOjE2MzYyODgwNDMsInRpbWVzdGFtcCI6MTYyODUxMjA0MywicmVfZW5jcnlwdGlvbl9rZXkiOiIiLCJlbmNyeXB0ZWQiOnRydWUsInNpZ25hdHVyZSI6ImNiYTZlMjA2OTBjOGZjZTk5YmFjZTMzYjFjMGY3ODQ5ZDE4YmJlMTdhODkyNjczODg1MjI2MDc3MGQzNzgzMGQifQ==
```

Decoded

```
{"client_id":"b734ef935e2a02892b2fa31e3488b360ef300d3b0b32c03834cea3a83e2453f0","owner_id":"639f271fe511cd1880c2a4d8ea4a4f240faf3365c61b65b425fab5e0231713c3","allocation_id":"937ad69fb20dc111bcfd01de42c792a0c2b441de3d3cf4cdb3f25c21c1ab4b7b","file_path_hash":"ad18f385cb61c813c3144569345f1c4681589c7347bd728692e85f1b7386b69d","actual_file_hash":"","file_name":"folder1","reference_type":"d","expiration":1636288043,"timestamp":1628512043,"re_encryption_key":"","encrypted":true,"signature":"cba6e20690c8fce99bace33b1c0f7849d18bbe17a8926738852260770d37830d"}
```

Make sure _"reference_type":"d"_ is "d" (directory)

Now you able to download files inside this directory with same auth*ticket. To download it just point to excact file location in *--remotepath\_ param:

```
zbox download --allocation 76ad9fa86f9b6685880553588a250586806ba5d7d20fc229d6905998be55d64a --localpath ~/file1.z --authticket $auth --remotepath /folder1/file1.z
```

This method works for both: encrypted and non-encrypted files.

##### share-encrypted revoke

This will cancel the share for particular buyer that was performed by the seller using zbox share. _Works only for files with --encrypted tag._

Use clientid of the user that was share the remotepath.
Required parameters are allocation, remotepath and clientid.

Command

    ./zbox share --revoke --allocation 3c0d32560ea18d9d0d76808216a9c634f661979d29ba59cc8dafccb3e5b95341 --remotepath /myfiles/hello.txt --clientid d52d82133177ec18505145e784bc87a0fb811d7ac82aa84ae6b013f96b93cfaa

Response

Returns status message showing whether the operation was successful or not.

#### List

Use `list` command to list files in given remote path of the dStorage. An auth ticket should be provided when
not sent by the allocation's owner. Using an auth ticket requires a `lookuphash` to indicate the path for which to list
contents.

| Parameter  | Required | Description                                                              | default | Valid values |
| ---------- | -------- | ------------------------------------------------------------------------ | ------- | ------------ |
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

#### Copy

Use `copy` command to copy file to another folder path in dStorage.
Only the owner of the allocation can copy an object.

| Parameter  | Required | Description                                       | default | Valid values |
| ---------- | -------- | ------------------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                                     |         | string       |
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

#### Move

Use `move` command to move file to another remote folder path on dStorage.
Only the owner of the allocation can copy an object.

| Parameter  | Required | Description                                       | default | Valid values |
| ---------- | -------- | ------------------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                                     |         | string       |
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

#### Sync

`sync` command syncs all files from the local folder recursively to the remote.
Only the allocation's owner can successfully run `sync`.

| Parameter   | Required | Description                                                                                   | default | Valid values |
| ----------- | -------- | --------------------------------------------------------------------------------------------- | ------- | ------------ |
| allocation  | yes      | allocation id                                                                                 |         | string       |
| encryptpath | no       | local directory path to be uploaded as encrypted                                              | false   | boolean      |
| excludepath | no       | paths to exclude from sync                                                                    |         | string array |
| localchache | no       | local chache of remote snapshot. Used for comparsion with remote. After sync will be updated. |         | string       |
| localpath   | yes      | local directory to which to sync                                                              |         | file path    |
| uploadonly  | no       | only upload and update files                                                                  | false   | boolean      |
| chunknumber | no       | how many chunks should be uploaded in a http request                                          | 1       | int          |

<details>
  <summary>sync</summary>

![image](https://user-images.githubusercontent.com/6240686/127884376-a95c4f27-4b2a-4d9b-91b6-c7e7919f88bc.png)

</details>

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

#### Get differences

`./zbox get-diff` command returns the differences between the local files specified by `localpath` and the files stored
on the root remotepath of the allocation.`localcache` flag can also be specified to use the local cache of remote snapshot created during [Sync](#sync) for file comparison.

| Parameter   | Required | Description                                   | default | Valid values |
| ----------- | -------- | --------------------------------------------- | ------- | ------------ |
| allocation  | yes      | allocation id                                 |         | string       |
| excludepath | no       | remote folder paths to exclude during syncing |         | string array |
| localcache  | no       | local cache of remote snapshot                |         | string       |
| localpath   | yes      | local directory to sync                       |         | string       |

Example

```
./zbox get-diff --allocation $ALLOC --localpath $local
```

Response:

```
[{"operation":"Upload","path":"/file1.txt","type":"f","attributes":{}},
{"operation":"Upload","path":"/file2.txt","type":"f","attributes":{}},
{"operation":"Upload","path":"/file3.txt","type":"f","attributes":{}},
{"operation":"Download","path":"/myfiles/file1.txt","type":"f","attributes":{}},
{"operation":"Download","path":"/myfiles/file2.txt","type":"f","attributes":{}}]
```

#### Get wallet

Use `getwallet` command to get additional wallet information including Encryption
Public Key,Client ID which are required for Private File Sharing.

| Parameter | Required | Description                   | default | Valid values |
| --------- | -------- | ----------------------------- | ------- | ------------ |
| json      | no       | print response in json format | false   | boolean      |

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

#### Get

Use `getallocation` command to get the information about the allocation such as total size , used size, number of challenges
and challenges passed/failed/open/redeemed.

| Parameter  | Required | Description                   | default | Valid values |
| ---------- | -------- | ----------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                 |         | string       |
| json       | no       | print response in json format | false   | boolean      |

<details>
  <summary>get</summary>

![image](https://user-images.githubusercontent.com/6240686/124476040-4f4bc580-dd9a-11eb-939c-464ffc6936db.png)

</details>

Example

```
./zbox getallocation --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc
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
      base URL:         http://demo.zus.network:31304
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

#### Get metadata

Use `meta` command to get metadata for a given remote file. Use must either be the
owner of the allocation on have an auth ticket or be a collaborator.
Use [share](#share) to create an auth ticket for someone or [add-collab](#add-collaborator)
to add a user as a collaborator. To indicate the object use `remotepath` or
`lookuphash` with an auth ticket.

| Parameter  | Required | Description                                        | default | Valid values |
| ---------- | -------- | -------------------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                                      |         | string       |
| authticket | no       | auth ticked if not owner of the allocation         |         | string       |
| json       | no       | print result in json format                        | false   | boolean      |
| lookuphash | no       | hash of object, use with auth ticket               |         | string       |
| remotepath | no       | remote path of objecte, do not use with authticket |         | string       |

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

#### Rename

`rename` command renames a file existing already on dStorage. Only the
allocation's owner can rename a file.

| Parameter  | Required | Description                                       | default | Valid values |
| ---------- | -------- | ------------------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                                     |         | string       |
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

#### Stats

`stats` command gets upload, download and challenge statistics for a file.
Only the owner can get a files stats.

| Parameter  | Required | Description                 | default | Valid values |
| ---------- | -------- | --------------------------- | ------- | ------------ |
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

#### Repair

Use `start-repair` command to repair a file on dStorage.
![repair](https://user-images.githubusercontent.com/65766301/120052600-b364c680-c043-11eb-9bf2-038ab244fed6.png)
\

| Parameter  | Required | Description               | default | Valid values |
| ---------- | -------- | ------------------------- | ------- | ------------ |
| allocation | yes      | allocation id             |         | string       |
| repairpath | yes      | remote path to repair     |         | string       |
| rootpath   | yes      | file path for local files |         | string       |

<details>
  <summary>start-repair</summary>

![image](https://user-images.githubusercontent.com/6240686/127882066-83d7e641-56a7-4ae1-bbf1-547ca389481c.png)

</details>

Example

```
./zbox start-repair --allocation 8695b9e7f986d4a447b64de020ba86f53b3b5e2c442abceb6cd65742702067dc --repairpath / --rootpath /home/dung/Desktop/alloc
```

Response:

```
Repair file completed, Total files repaired:  0
```

#### Sign data

`sign-data` uses the information from your wallet to sign the input data string

| Parameter | Required | Description    | default | Valid values |
| --------- | -------- | -------------- | ------- | ------------ |
| data      | yes      | string to sign |         | string       |

```shell
./zbox sign-data "data to sign"
Signature : 9432ab2ee602062afaf48c4016b373a65db48a8546a81c09dead40e54966399e
```

---

#### Streaming

Video streaming with Zbox CLI can be implemented with players for different operating platforms(iOS, Android Mac).Zbox CLI does not have a player itself and use the the downloadFileByBlocks helper function to properly returns file-chunks with correct byte range.

![streaming-android](https://user-images.githubusercontent.com/65766301/120052635-ce373b00-c043-11eb-94a5-a9711078ee54.png)

##### How it works:

When the user starts the video player (ExoPlayer for Android or AVPlayer for iOS), A ZChainDataSource starts chunked download and requests chunks of video from the buffer(a Middleman between streaming player and Zbox).

After the arrival of the first chunk, the player starts requesting more chunks from the buffer, which requests the Zbox SDK. Zbox SDK, which is built using GO, makes use of the downloadFileByBlocks method to reliably download large files by chunking them into a sequence of parts that can be downloaded individually. Once the blocks are downloaded, they are read into input streams and added to the media source of the streaming player.

The task of downloading files and writing them to buffer using Zbox SDK happens constantly, and If players request random bits of video, they are delivered instantly by a buffer.

In a case, if the player didn't receive chunks (for example, it's still not downloaded), then the player switches to STALE state, and the video stream will pause. During the STALE state, a player tries to make multiple requests for chunks; if didn't receive a response, the video stream stops.

##### Usage

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

### Lock and Unlock Tokens

#### Challenge pool information

Use `cp-info` command to get the challenge pool brief information.

| Parameter  | Required | Description                 | default | Valid values |
| ---------- | -------- | --------------------------- | ------- | ------------ |
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

#### Create read pool

Use `rp-create` to create a read pool, `rp-create` has no parameters.

<details>
  <summary>rp-create</summary>

![image](https://user-images.githubusercontent.com/6240686/127875827-f0301162-5c62-4964-989a-d56d4b2292af.png)

</details>

```
./zbox rp-create
```

#### Collect rewards

Use `collect-reward` to transfer reward tokens from a stake pool in which you have
invested to your wallet.

You earn rewards for:
Blobbers

- File space used by allocation owners and associates.
- A min lock demand for each allocation.
- Block rewards. Each block a reward gets paid out to blobber stakeholders in the form of a random lottery.
  Validators
- Payment for validating blobber challenge responses.

The stake pool keeps an account for all stakeholders to maintain accrued rewards.
These rewards can be accessed using this `collect-reward` command.

| Parameter     | Required | Description          | default | Valid values |
| ------------- | -------- | -------------------- | ------- | ------------ |
| provider_type | no       | blobber or validator | blobber | string       |

```bash
./zbox colect-reward --provider_type blobber
```

#### Read pool info

Use `rp-info` to get read pool information.

| Parameter | Required | Description                 | default | Valid values |
| --------- | -------- | --------------------------- | ------- | ------------ |
| json      | no       | print result in json format | false   | boolean      |

<details>
  <summary>rp-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124507524-d8c4bd00-ddc5-11eb-853e-513957cf3dbb.png)

</details>

```
./zbox rp-info
```

#### Lock tokens into read pool

Lock some tokens in read pool. ReadPool is not linked to specific allocations anymore.
Each wallet has a singular, non-expiring, untethered ReadPool. Locked tokens in ReadPool can be unlocked at any time and returned to the original wallet balance.

- If the user does not have a pre-existing read pool, then the smart-contract
  creates one.

Locked tokens can be used to pay for read access to file(s) stored with different allocations.
To use these tokens the user must be the allocation owner, collaborator or have an auth ticket.

| Parameter | Required | Description     | default | Valid values |
| --------- | -------- | --------------- | ------- | ------------ |
| fee       |          | transaction fee | 0       | int          |
| tokens    | yes      | tokens to lock  |         | int          |

```
./zbox rp-lock --tokens 1
```

#### Unlock tokens from read pool

Use `rp-unlock` to unlock tokens from `read pool by ownership.

| Parameter | Required | Description     | default | Valid values |
| --------- | -------- | --------------- | ------- | ------------ |
| fee       | no       | transaction fee | 0       | float        |

Unlocked tokens get returned to the original wallet balance.

#### Storage SC configurations

Show storage SC configuration.

| Parameter  | Required | Description                 | default | Valid values |
| ---------- | -------- | --------------------------- | ------- | ------------ |
| allocation | yes      | allocation id               |         | string       |
| json       | no       | print result in json format | false   | boolean      |

<details>
  <summary>sc-config</summary>

![image](https://user-images.githubusercontent.com/6240686/124578670-53352180-de46-11eb-99a5-07debf17e351.png)

</details>

```
./zbox sc-config
```

#### Stake pool info

Use `sp-info` to get your stake pool information and settings.

| Parameter  | Required | Description                 | default        | Valid values |
| ---------- | -------- | --------------------------- | -------------- | ------------ |
| blobber_id |          | id of blobber               | current client | string       |
| json       | no       | print result in json format | false          | boolean      |

<details>
  <summary>sp-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124581849-63023500-de49-11eb-8927-50d9ff97671b.png)

</details>

```
./zbox sp-info --blobber_id <blobber_id>
```

#### Lock tokens into stake pool

Lock creates delegate pool for current client and a given provider (blobber or validator).
The tokens locked for the provider stake can be unlocked any time, excluding times
when the tokens held by opened offers. These tokens will earn rewards depending on the
actions of the linked provider.

`sp-lock` returns the id of the new stake pool, this will be needed to reference
to stake pool later.

| Parameter    | Required | Description     | default | Valid values |
| ------------ | -------- | --------------- | ------- | ------------ |
| blobber_id   |          | id of blobber   | n/a     | string       |
| validator_id |          | id of validator | n/a     | string       |
| fee          | no       | transaction fee | 0       | float        |
| tokens       | yes      | tokens to lock  |         | float        |

<details>
  <summary>sp-lock</summary>

![image](https://user-images.githubusercontent.com/6240686/124585686-73b4aa00-de4d-11eb-83cb-334f7c54543e.png)

</details>

To stake tokens for blobbers:

```
./zbox sp-lock --blobber_id <blobber_id> --tokens 1.0
```

To stake tokens for validators:

```
./zbox sp-lock --validator_id <validator_id> --tokens 1.0
```

#### Unlock tokens from stake pool

Unlock a stake pool by pool owner. If the stake pool cannot be unlocked as
it would leave insufficient funds for opened offers, then `sp-unlock` tags
the stake pool to be unlocked later. This tag prevents the stake pool affecting
blobber allocation for any new allocations.

| Parameter    | Required | Description     | default | Valid values |
| ------------ | -------- | --------------- | ------- | ------------ |
| blobber_id   |          | id of blobber   | n/a     | string       |
| validator_id |          | id of validator | n/a     | string       |
| fee          | no       | transaction fee | 0       | float        |

<details>
  <summary>sp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/124597566-8e8e1b00-de5b-11eb-8926-867687aaa06a.png)

</details>

To unstake blobber tokens:

```
./zbox sp-unlock --blobber_id <blobber_id>
```

To unstake validator tokens:

```
./zbox sp-unlock --validator_id <validator_id> --pool_id <pool_id>
```

#### Stake pools info of user

Get information about all stake pools of current user.

| Parameter | Required | Description                 | default | Valid values |
| --------- | -------- | --------------------------- | ------- | ------------ |
| json      | no       | print result in json format | false   | boolean      |

<details>
  <summary>sp-user-info</summary>

![image](https://user-images.githubusercontent.com/6240686/124600324-7ff53300-de5e-11eb-9b78-5a4f9c59a536.png)

</details>

```
./zbox sp-user-info
```

#### Write pool info

Write pool information. Use allocation id to filter results to a singe allocation.

| Parameter     | Required | Description                 | default | Valid values |
| ------------- | -------- | --------------------------- | ------- | ------------ |
| allocation id | no       | allocation id               |         | string       |
| json          | no       | print result in json format | false   | boolean      |

<details>
  <summary>wp-info</summary>

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

#### Lock tokens into write pool

`wp-lock` can be used to lock tokens in a write pool associated with an allocation.
All tokens will be divided between allocation blobbers depending on their write price.

- Uses two different formats, you can either define a specific blobber
  to lock all tokens, or spread across all the allocations blobbers automatically.
- If the user does not have a pre-existing read pool, then the smart-contract
  creates one.

Anyone can lock tokens with a write pool attached an allocation. These tokens can
be used to pay for the allocation updates and min lock demand as needed. Any tokens
moved into the challenge pool to underwrite blobbers' min lock demands return to the
allocation's owner on closing the allocation.

| Parameter     | Required | Description                       | default | Valid values |
| ------------- | -------- | --------------------------------- | ------- | ------------ |
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

#### Unlock tokens from write pool

`wp-unlock` unlocks an expired write pool.
An expired write pool, associated with an allocation, can be locked until allocation finalization even if it's expired. It possible in cases where related blobber doesn't give their min lock demands. The finalization will pay the demand and unlock the pool.

| Parameter | Required | Description     | default | Valid values |
| --------- | -------- | --------------- | ------- | ------------ |
| fee       | no       | transaction fee | 0       | float        |

<details>
  <summary>rp-unlock</summary>

![image](https://user-images.githubusercontent.com/6240686/123980742-b09a2000-d9b9-11eb-8987-c18ff90ee705.png)

</details>

```
./zbox wp-unlock
```

#### Download cost

`get-download-cost` determines the cost for downloading the remote file from dStorage. The client must be an
owner, collaborator, or using an auth ticket to determine the download cost of the file.

| Parameter  | Required | Description                               | default | Valid values |
| ---------- | -------- | ----------------------------------------- | ------- | ------------ |
| allocation | yes      | allocation id                             |         | string       |
| authticket | no       | auth ticket to use if not the owner       |         | string       |
| lookuphash | no       | hash of remote file, use with auth ticket |         | string       |
| remotepath | no       | file of which to get stats, use if owner  |         | string       |

<details>
  <summary>get-download-cost</summary>

![image](https://user-images.githubusercontent.com/6240686/124497750-41ef0500-ddb3-11eb-99ea-115a4e234eda.png)

</details>
Command:
```
./zbox get-download-cost --allocation <allocation_id> --remotepath /path/file.ext
```
Response:
```
0.0000107434 tokens for 10 64KB blocks (24 B) of <remote_path_of_file> .
```

#### Upload cost

`get-upload-cost` determines the cost for uploading a local file on dStorage.
`--duration` Ignored if `--end` true, in which case the cost of upload calculated until
the allocation expires.

| Parameter  | Required | Description                          | default | Valid values |
| ---------- | -------- | ------------------------------------ | ------- | ------------ |
| allocation | yes      | allocation id                        |         | string       |
| duration   | no       | duration for which to upload file    |         | duration     |
| end        | no       | upload file until allocation expires | false   | boolean      |
| localpath  | yes      | local of path to calculate upload    |         | file path    |

<details>
  <summary>get-upload-cost</summary>

![image](https://user-images.githubusercontent.com/6240686/124501898-51be1780-ddba-11eb-8c1a-d238cfd8f43f.png)

</details>

Command:

```
./zbox get-upload-cost --allocation <allocation_id> --localpath ./path/file.ext
```

Response:

```
 0.0000000028 tokens / 720h0m0s for 24 B of <remote_path_of_file>
```

## Troubleshooting

1. Both `rp-info` and `rp-lock` are not working.

```
./zbox rp-info
```

Response:

```
Failed to get read pool info: error requesting read pool info: consensus_failed: consensus failed on sharders
```

This can happen if read pool is not yet created for wallet. Read pool is usually created when new wallet is created by `zbox` or `zwallet`. However, if wallet is recovered through `zwallet recoverwallet`, read pool may not have been created. Simply run `zbox rp-create` to create a read pool.
