Contents : 
[Sync](#sync)
[Get Differences](#get-differences)
[Stream](#stream)
[Get MPT](#get-mpt)
[Rollback Allocation](#rollback-allocation)

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
