```puml
title Add collaborator
boundary zbox 
collections blobbers
database store
control 0chain
zbox -> blobbers : add collaborator
    blobbers -> 0chain : get allocation
    0chain -> blobbers : allocation
    alt check owner == sender
        blobbers ->x zbox : allocation must be owner
    end
    blobbers -> store : add colaborator
    blobbers -> zbox
```

```puml
title Upload cost
boundary zbox 
collections blobbers
control 0chain
zbox -> 0chain : get allocation
0chain -> zbox : allocation
zbox -> zbox : calculate upload cost
```

```puml
title Download cost
boundary zbox 
collections blobbers
control 0chain
zbox -> 0chain : get allocation
0chain -> zbox : allocation
zbox -> blobbers : get file metadata
blobbers -> zbox : file metadata
zbox -> zbox : calculate download cost
```

```puml
title File stats
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> blobbers : file stats
note left
    * allocation id
    * remotepath
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender is owner
        blobbers ->x zbox : unauthorised user
    end
    store -> blobbers : object stats
blobbers -> zbox : object stats
```

```puml
title Rename
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> blobbers : rename
note left
    * allocation id
    * remote path   
    * new name
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : needs to be performed\nby the owner
    end
    blobbers -> store : rename object
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```

```puml
title Get metadata
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> blobbers : get metadata
note left
    * allocation id
    * auth ticket
    * remotepath or lookup hash
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    store -> blobbers : collaborators
    alt check sender is owner or\ncollaborator or\nauth ticket validates
        blobbers ->x zbox : unauthorised user
    end
    store -> blobbers : object metadata
blobbers -> zbox : object metadata
alt commit true
zbox -> 0chain : save metadata
        0chain -> blockchain :  save metadata
        0chain -> zbox
end    
```

```puml
title Get allocation info
boundary zbox
control 0chain
entity blockchain
zbox -> 0chain : get allocation
0chain -> blockchain : allocation
blockchain -> 0chain : allocation
0chain -> zbox : allocation
```

```puml
title List allocations
boundary zbox
control 0chain
entity blockchain
zbox -> 0chain : get allocation
0chain -> blockchain : get allocations for user
blockchain -> 0chain : user allocations
0chain -> zbox : user allocations
```

```puml
title List allocations
boundary zbox
control 0chain
entity blockchain
zbox -> 0chain : get allocations for user
0chain -> blockchain : get allocations for user
blockchain -> 0chain : user allocations
0chain -> zbox : user allocations
```

```puml
title Move
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> zbox : encrypt
zbox -> zbox : add thumbnail
zbox -> blobbers : move
note left
    * allocation id
    * remote path   
    * destination directory 
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : needs to be performed\nby the owner
    end
    blobbers -> store : move remote, destination
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```

```puml
title Copy
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> zbox : encrypt
zbox -> zbox : add thumbnail
zbox -> blobbers : copy
note left
    * allocation id
    * remote path   
    * destination directory 
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : needs to be performed\nby the owner
    end
    blobbers -> store : copy remote, destination
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```

```puml
title List
boundary zbox
control blobber
database store 
zbox -> blobber : list allocation objects
note left
    * allocation id
    * auth ticket
    * lookup hash (with auth ticket)
    * remote path (without auth ticket)
end note
    alt sender not owner        
        blobber -> blobber : validate auth ticket
        blobber -> blobber : lookup remote object hash        
    end
    blobber -> store : request object infomation
    store -> blobber : objects information
    blobber -> zbox : object information
```


```puml
title Share
boundary zbox 
collections blobbers
control 0chain
zbox -> 0chain : request allocation info
0chain -> zbox : allocation info
zbox -> zbox : validate
zbox -> blobbers : request file information
blobbers -> zbox : file hash
zbox -> zbox : generate auto ticket
note left
    * allocation id
    * client id
    * encryption public key
    * remote path
end note
```


```puml
title Update
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> zbox : encrypt
zbox -> zbox : add thumbnail
zbox -> blobbers : upload
note left
    * allocation id
    * file data
    * remotepath    
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : needs to be performed\nby the owner
    end
    blobbers -> store : save updated file
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```


```puml
title Delete
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> zbox : encrypt
zbox -> zbox : add thumbnail
zbox -> blobbers : upload
note left
    * allocation id
    * remotepath    
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : delete needs to be\nperformed by the owner
    end
    blobbers -> store : delete file\nremotepath
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```

```puml
title Download
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> blobbers : download
note left
    * allocation id
    * auth ticket + rx_pay
    * block to download
    * remotepath    
    * start block
    * number of blocks
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    store -> blobbers : collaborators
    alt check sender is owner or\ncollaborator or\nauth ticket validates
        blobbers ->x zbox : unauthorised user
    end
    blobbers -> blobbers : set payer owner unless\nauth ticket & (rx_pay = true)
    blobbers -> blobbers : check enugh tokens\nin local read pool
    store -> blobbers : file data blocks    
    blobbers -> 0chain : commit read marker
    note left
        * allocation id
        * payer id (owner or collaborator)
        * timestamp
        * read counter
        * signature
    end note          
    blobbers -> zbox : file data
        blockchain -> 0chain : previous reads in block 
        0chain -> 0chain : validate read marker
        blockchain -> 0chain : blobber's stake pool
        blockchain -> 0chain : payer's read pool
        0chain -> 0chain : pay blobber's stake holders
        0chain -> blockchain : save blobber's stake pool
        0chain -> blockchain : save payer's read pool       
zbox -> zbox : decrypt data
zbox -> zbox : save to local file 
alt commit true
zbox -> 0chain : save metadata
        0chain -> blockchain :  save metadata
        0chain -> zbox
end    
```

```puml
title Upload
boundary zbox 
collections blobbers
database store
control 0chain
entity blockchain
zbox -> zbox : encrypt
zbox -> zbox : add thumbnail
zbox -> blobbers : upload
note left
    * allocation id
    * file data
    * remotepath    
end note
    blobbers -> 0chain : request allocaton
        blockchain -> 0chain : allocation
    0chain -> blobbers : allocation
    alt check sender == owner
        blobbers ->x zbox : needs to be performed\nby the owner
    end
    blobbers -> store : save file data
    blobbers -> zbox
alt commit true
zbox -> 0chain : save metadata
    0chain -> blockchain :  save metadata
0chain -> zbox
end    
```

```puml
title Transfer allocation ownership
boundary zbox 
control storagesc
entity blockchain
zbox ->storagesc: curator_transfer_allocation
note left
    * allocatino id
    * new onwer id
    * new owner pubic key
end note
    blockchain ->storagesc: allocation
    alt check sender is curator
       storagesc->x zbox : only curators can transfer allocation
    end
    group allocation
       storagesc->storagesc: chaner owner\nchange owner public key
        blockchain ->storagesc: new owner write pool
       storagesc->storagesc: new empty alloaction write pool
       storagesc-> blockchain : save write pool
    end 
   storagesc-> blockchain : save allocation
storagesc -> zbox     
``` 

```puml
title Add curator
boundary zbox 
control storagesc
entity blockchain
zbox ->storagesc: add_curator
note left
    * allocation id
    * curator id
end note
    blockchain ->storagesc: allocation
    group allocation
       storagesc->storagesc: append curator
    end 
   storagesc-> blockchain : save allocation
storagesc -> zbox     
``` 
```puml
title Remove curator
boundary zbox 
control storagesc
entity blockchain
zbox ->storagesc: remove_curator
note left
    * allocation id
    * curator id
end note
    blockchain ->storagesc: allocation
    group allocation
       storagesc->storagesc: remove curator
    end 
   storagesc-> blockchain : save allocation
storagesc -> zbox 


```puml
title Finilize allocation
boundary zbox 
control storagesc
entity blockchain
zbox ->storagesc: finalize_allocation
note left
    * allocation id
end note
    blockchain -> storagesc : allocation
    alt check allocation expired
       storagesc->x zbox : allocation not expired
    end
    blockchain ->storagesc:blobbers
    storagesc->storagesc: blobber challenge pass rates
    blockchain ->storagesc: challenge pool
    blockchain ->storagesc: allocation's write pools   
    group challenge pool
       storagesc->storagesc: cover min lock demand\ncp -> blobbers + stake holders
       storagesc->storagesc: passed challenges\ncp -> blobbers + stake holders
       storagesc->storagesc: mint interest\nstoragesc -> blobbers' stake holders 
       storagesc-> blockchain : minted interest payments
       storagesc->storagesc: reaming funds\ncp -> owner's write pool
    end
   storagesc-> blockchain : save write pools
   storagesc-> blockchain : save challenge pool
    blockchain ->storagesc: all allocations
   storagesc->storagesc: remove allocation id 
   storagesc-> blockchain : save all allocatinos    
storagesc -> zbox : allocation id 
```

 ```puml
title Cancel allocation
boundary zbox 
control storagesc
entity blockchain
zbox ->storagesc: alloc-cancel 
note left
    * allocation id
end note
    blockchain -> storagesc : allocation
    alt check allocation not expired
       storagesc->x zbox : cancelling expired allocation
    end
    blockchain ->storagesc: challenge pool
    alt check sufficent challenges have failed
       storagesc->x zbox : not enough failed challenges
    end
    blockchain ->storagesc:blobbers
    storagesc->storagesc: blobber challenge pass rates
    blockchain ->storagesc: allocation's write pools   
    group challenge pool
       storagesc->storagesc: cover min lock demand\ncp -> blobbers + stake holders
       storagesc->storagesc: passed challenges\ncp -> blobbers + stake holders
       storagesc->storagesc: mint interest\nstoragesc -> blobbers' stake holders 
       storagesc-> blockchain : minted interest payments
       storagesc->storagesc: return any reaming funds\ncp -> owner's write pool
    end
   storagesc-> blockchain : save write pools
   blockchain ->storagesc: list of all allocations
   storagesc->storagesc: remove allocation id 
   storagesc-> blockchain : save all allocatinos    
storagesc -> zbox : allocation id 
```

 ```puml
title Free storage new allocation
boundary zbox
control storagesc
entity blockchain
zbox ->storagesc: free_allocation_request
note left
    Free storage mareker
    * maker issuer name
    * client to give tokens to
    * number of free tokens     
    * timestamp to prevent reuse
    * signed
end note       
   storagesc<- blockchain : corporations details
   storagesc->storagesc: validate free storage marker
   blockchain ->storagesc: blobbers
   storagesc->storagesc: select allocation blobbers
   storagesc->storagesc: new allocation
   group new allocation
       storagesc->storagesc: set:\n* in paramters\n* selected blobbers\n* now
       storagesc->storagesc: new write pool; mint tokens
       storagesc->storagesc: new read pool; mint tokens
       storagesc-> blockchain : write pool
       storagesc-> blockchain : read pool
       storagesc->storagesc: new challenge pool    
       storagesc-> blockchain : challenge pool    
   end
   storagesc-> blockchain : new allocation
storagesc -> zbox : allocation id 
```

```puml
title New allocation
boundary zbox
control storagesc
entity blockchain
zbox -> storagesc : new_allocation_request
note left
    * data shards
    * paraty shards
    * size
    * expiration
    * preferred blobbers
    * read price range
    * write price range
    * max challenge time
end note
    blockchain -> storagesc : blobbers
    storagesc -> storagesc : select allocation blobbers
    storagesc -> storagesc : new allocation
    group new allocation
        storagesc -> storagesc : set:\n* in paramters\n* selected blobbers\n* now
        storagesc -> storagesc : new write pool
        group new write pool
            storagesc -> storagesc : transfer owner tokens to write pool
        end 
        storagesc -> blockchain : write pool
        storagesc -> storagesc : new challenge pool    
        storagesc -> blockchain : challenge pool    
    end
    storagesc -> blockchain : new allocation
storagesc -> zbox : new allocation id
```

```puml
title update free storage marker
boundary zbox
control storagesc
entity blockchain
zbox ->storagesc: free_update_allocation
note left
    Free storage mareker
    * maker issuer name
    * client to give tokens to
    * number of free tokens     
    * timestamp to prevent reuse
    * signed
end note   
   storagesc<- blockchain : maker issuer's details
   storagesc->storagesc: validate free storage marker
    blockchain ->storagesc: allocation
    alt confirm all allocation blobbers have enough capacity
       storagesc->x zbox : blobber doesn't have enough free space
    end
    alt check expiration agaisnt allocation blobbers' max offer duration
       storagesc->x zbox : blobber doesn't allow so long offers
    end
    group allocation
       storagesc->storagesc: update alllocation as required\nsize expireation and immutable
       blockchain ->storagesc: owner wrtie pool
       storagesc->storagesc: ammend write pool lock duration
       blockchain ->storagesc: challenge pool
       storagesc->storagesc: mint new tokens for challenge pool
       storagesc-> blockchain : challenge pool
       storagesc-> blockchain : write pool
    end
   storagesc-> blockchain : allocation
storagesc -> zbox : transaction id
```

```puml
title Update allocation
boundary zbox
control storagesc
entity blockchain
zbox ->storagesc: update_allocation_request
note left
    owner = sender
    allocation id
    size
    expiration
    set immutable?
end note
    blockchain ->storagesc: allocation
    alt confirm all allocation blobbers have enough capacity
       storagesc->x zbox : blobber doesn't have enough free space
    end
    alt check expiration agaisnt allocation blobbers' max offer duration
       storagesc->x zbox : blobber doesn't allow so long offers
    end
    group allocation
       storagesc->storagesc: update alllocation as required\nsize expireation and immutable
        blockchain ->storagesc: owner wrtie pool
       storagesc->storagesc: ammend write pool lock duration
        blockchain ->storagesc: challenge pool
       storagesc->storagesc: transfer tokens between\nwrite pool amd challenge pool
       storagesc-> blockchain : challenge pool
       storagesc-> blockchain : write pool
    end
   storagesc-> blockchain : allocation
storagesc -> zbox : transaction id
```
    
```puml
title Create read pool
zbox ->storagesc: //new_read_pool//
storagesc ->storagesc: new read pool
storagesc -> blockchain : read pool
storagesc -> zbox
```

```puml
title Lock tokens in read pool for given bobber
boundary zbox
control storagesc
entity blockchain
zbox ->storagesc: rp-lock, token value
note left
    * lock duration
    * allocation id
    * blobber id to lock for
end note 
storagesc ->storagesc: new allocation pool
group new allocation pool
   storagesc->storagesc: add new blobber pool\nbalance value
   storagesc->storagesc: transfer tokens
   storagesc->storagesc: add set expiration
end
blockchain ->storagesc: readpool
group read pool
storagesc ->storagesc: add new allocation pool
end 
storagesc -> blockchain : save read pool
storagesc -> zbox
```

```puml
title Lock tokens in read pool no blobber specified
boundary zbox
control storagesc
entity blockchain
zbox ->storagesc: rp-lock, token value
note left
    * lock duration
    * allocation id
end note 
storagesc ->storagesc: new allocation pool
group new allocation pool
    blockchain ->storagesc: get allocation blobbers
    loop each blobber
       storagesc->storagesc: new blobber pool\nbalance value \n split by read price
    end
   storagesc->storagesc: add new blobber pools
   storagesc->storagesc: transfer tokens
   storagesc->storagesc: add set expiration
end
blockchain ->storagesc: readpool
group read pool
storagesc ->storagesc: add new allocation pool
end 
storagesc -> blockchain : save read pool
storagesc -> zbox
```

```puml
boundary zbox
control sc
entity blockchain
title Unlock read pool
zbox ->storagesc: rp-unlock
note left
    * read pool id
end note 
blockchain ->storagesc: read pool
alt read pool expired
    group read pool
       storagesc->storagesc: remove pool
    end
   storagesc-> blockchain : transfer tokens from\nread pool to user
   storagesc-> blockchain : save read pool
   storagesc-> zbox
else read pool not expired
   storagesc->x zbox : read pool not expired
end
```

```puml
title Lock tokens in write pool for given bobber
zbox ->storagesc: rp-lock, token value
note left
    * lock duration
    * allocation id
    * blobber id to lock for
end note 
storagesc ->storagesc: new allocation pool
group new allocation pool
   storagesc->storagesc: add new blobber pool\nbalance value
   storagesc->storagesc: transfer tokens
   storagesc->storagesc: add set expiration
end
blockchain ->storagesc: writepool
group write pool
storagesc ->storagesc: add new allocation pool
end 
storagesc -> blockchain : save write pool
storagesc -> zbox
```

```puml
title Lock tokens in write pool no blobber specified
zbox ->storagesc: rp-lock, token value
note left
    * lock duration
    * allocation id
end note 
storagesc ->storagesc: new allocation pool
group new allocation pool
    blockchain ->storagesc: get allocation blobbers
    loop each blobber
       storagesc->storagesc: new blobber pool\nbalance value \n split by write price
    end
   storagesc->storagesc: add new blobber pools
   storagesc->storagesc: transfer tokens
   storagesc->storagesc: add set expiration
end
blockchain ->storagesc: readpool
group write pool
storagesc ->storagesc: add new allocation pool
end 
storagesc -> blockchain : save write pool
storagesc -> zbox
```

```puml
title Unlock write pool
zbox ->storagesc: rp-unlock
note left
    * write pool id
end note 
blockchain ->storagesc: write pool
alt write pool with id expired
    group write pool
       storagesc->storagesc: remove pool id
    end
   storagesc-> blockchain : transfer tokens from\nwrite pool id to user
   storagesc-> blockchain : save write pool
   storagesc-> zbox
else write pool id not expired
   storagesc->x zbox : write pool not expired
end
```
