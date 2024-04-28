Contents : 
[Sync](#sync)
[Get Differences](#get-differences)
[Stream](#stream)
[Get MPT](#get-mpt)
[Rollback Allocation](#rollback-allocation)
[Streaming](#streaming)
    - [How it works:](#how-it-works)
    - [Usage](#usage)
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
| chunknumber | no       | how many chunks should be uploaded in a http request                                          |         | int          |
| remotepath  | no       | Remote dir path from where it sync                                                            |   "/"   | string       |
| verifydownload | no       | how many chunks should be uploaded in a http request                                       |  true   | boolean      |

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

`get-diff` command returns the differences between the local files specified by `localpath` and the files stored
on the root remotepath of the allocation.`localcache` flag can also be specified to use the local cache of remote snapshot created during [Sync](#sync) for file comparison.

| Parameter   | Required | Description                                   | default | Valid values |
| ----------- | -------- | --------------------------------------------- | ------- | ------------ |
| allocation  | yes      | allocation id                                 |         | string       |
| excludepath | no       | remote folder paths to exclude during syncing |         | string array |
| localcache  | no       | local cache of remote snapshot                |         | string       |
| localpath   | yes      | local directory to sync                       |         | string       |

Example

```sh
./zbox get-diff --allocation $ALLOC --localpath $local
```

Response:

```sh
[{"operation":"Upload","path":"/file1.txt","type":"f","attributes":{}},
{"operation":"Upload","path":"/file2.txt","type":"f","attributes":{}},
{"operation":"Upload","path":"/file3.txt","type":"f","attributes":{}},
{"operation":"Download","path":"/myfiles/file1.txt","type":"f","attributes":{}},
{"operation":"Download","path":"/myfiles/file2.txt","type":"f","attributes":{}}]
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
| attr-who-pays-for-reads | no       | Who pays for reads: owner or 3rd_party         | owner   | owner / 3rd_party|

<details>
  <summary>stream</summary>

![image](https://github.com/0chain/blobber/wiki/uml/usecase/live_upload_live.png)

</details>


#### Get MPT

`get-mpt` is used to directly get blockchain data from the MPT key

| Parameter | Required | Description    | default | Valid values |
| --------- | -------- | -------------- | ------- | ------------ |
| key       | yes      | Key in MPT datastore |         | string       |

Sample Command : 
```sh
./zbox get-mpt --key {MPT_KEY}
```

Sample Response : 
```sh

```

#### Rollback Allocation

`rollback` is used to directly get blockchain data from the MPT key

| Parameter | Required | Description    | default | Valid values |
| --------- | -------- | -------------- | ------- | ------------ |
| allocation       | yes      | Allocation ID |         | string       |

Sample Command : 
```sh
./zbox rollback --allocation $ALLOC
```

Sample Response : 
```sh
Rollback successful
```

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
