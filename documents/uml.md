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
control 0chain
zbox -> 0chain : request allocation info
0chain -> zbox : allocation info
zbox -> zbox : validate
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
zbox -> sc : transfer allocatno owner
note left
    * allocatino id
    * new onwer id
    * new owner pubic key
end note
    blockchain -> sc : allocation
    alt check sender is curator
        sc ->x zbox : only curators can transfer allocation
    end
    group allocation
        sc -> sc : chaner owner\nchange owner public key
        blockchain -> sc : new owner write pool
        sc -> sc : new empty alloaction write pool
        sc -> blockchain : save write pool
    end 
    sc -> blockchain : save allocation
sc -> zbox     
``` 

```puml
title Add curator
zbox -> sc : addcurator
note left
    * allocatino id
    * curator id
end note
    blockchain -> sc : allocation
    group allocation
        sc -> sc : append curator
    end 
    sc -> blockchain : save allocation
sc -> zbox     
``` 


```puml
title Finilize allocation
zbox -> sc : alloc-fini
note left
    * allocation id
end note
    alt check allocation expired
        sc ->x zbox : allocation not expired
    end
    blockchain -> sc :blobbers
    sc -> sc : blobber challenge pass rates
    blockchain -> sc : challenge pool
    blockchain -> sc : write pool   
    group challenge pool
        sc -> sc : min lock demand\ncp -> blobbers + stake holders
        sc -> sc : passed challenges\ncp -> blobbers + stake holders
        sc -> sc : pay interest\nsc -> blobbers' stake holders 
        sc -> blockchain : minted interest payments
        sc -> sc : return any reaming funds\ncp -> write pool
    end
    sc -> blockchain : save write pool
    sc -> blockchain : save challenge pool
    blockchain -> sc : all allocations
    sc -> sc : remove allocation id 
    sc -> blockchain : save all allocatinos    
sc -> zbox : allocation id 
```

 ```puml
title Cancel allocation
zbox -> sc : alloc-cancel 
note left
    * allocation id
end note
    alt check allocation not expired
        sc ->x zbox : cancelling expired allocation
    end
    alt check sufficent challenges have failed
        sc ->x zbox : not enough failed challenges
    end
    blockchain -> sc :blobbers
    sc -> sc : blobber challenge pass rates
    blockchain -> sc : challenge pool
    blockchain -> sc : write pool   
    group challenge pool
        sc -> sc : min lock demand\ncp -> blobbers + stake holders
        sc -> sc : passed challenges\ncp -> blobbers + stake holders
        sc -> sc : pay interest\nsc -> blobbers' stake holders 
        sc -> blockchain : minted interest payments
        sc -> sc : return any reaming funds\ncp -> write pool
    end
    sc -> blockchain : save write pool
    sc -> blockchain : save challenge pool
    blockchain -> sc : save all allocations
    sc -> sc : remove allocation id 
    sc -> blockchain : all allocatinos    
sc -> zbox : allocation id 
```

 ```puml
title Free storage new allocation
zbox -> sc : newallocation --free-storage
note left
    Free storage mareker
    * maker issuer name
    * client to give tokens to
    * number of free tokens     
    * timestamp to prevent reuse
    * signed
end note       
    sc <- blockchain : corporations details
    sc -> sc : validate free storage marker
    blockchain -> sc : blobbers
    sc -> sc : select allocation blobbers
    sc -> sc : new allocation
    group new allocation
        sc -> sc : set:\n* in paramters\n* selected blobbers\n* now
        sc -> sc : new write pool
        group new write pool
            sc -> sc : mint tokens for write pool
        end 
        sc -> blockchain : wrie pool
        sc -> sc : new challenge pool    
        sc -> blockchain : challenge pool    
    end
    sc -> blockchain : new allocation
sc -> zbox : allocation id 
```

```puml
title New allocation
zbox -> sc : new allocation
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
    blockchain -> sc : blobbers
    sc -> sc : select allocation blobbers
    sc -> sc : new allocation
    group new allocation
        sc -> sc : set:\n* in paramters\n* selected blobbers\n* now
        sc -> sc : new write pool
        group new write pool
            sc -> sc : transfer owner tokens to write pool
        end 
        sc -> blockchain : wrie pool
        sc -> sc : new challenge pool    
        sc -> blockchain : challenge pool    
    end
    sc -> blockchain : new allocation
sc -> zbox : new allocation id
```

```puml
title update free storage marker
zbox -> sc : updateallocation --free-storage 
note left
    Free storage mareker
    * maker issuer name
    * client to give tokens to
    * number of free tokens     
    * timestamp to prevent reuse
    * signed
end note   
    sc <- blockchain : maker issuer's details
    sc -> sc : validate free storage marker
    blockchain -> sc : allocation
    alt confirm all allocation blobbers have enough capacity
        sc ->x zbox : blobber doesn't have enough free space
    end
    alt check expiration agaisnt allocation blobbers' max offer duration
        sc ->x zbox : blobber doesn't allow so long offers
    end
    group allocation
        sc -> sc : update alllocation as required\nsize expireation and immutable
        blockchain -> sc : owner wrtie pool
        sc -> sc : ammend write pool lock duration
        blockchain -> sc : challenge pool
        sc -> sc : mint new tokens for challenge pool
        sc -> blockchain : challenge pool
        sc -> blockchain : write pool
    end
    sc -> blockchain : allocation
sc -> zbox : transaction id
```

```puml
title Update allocation
zbox -> sc : update allocation
note left
    owner = sender
    allocation id
    size
    expiration
    set immutable?
end note
    blockchain -> sc : allocation
    alt confirm all allocation blobbers have enough capacity
        sc ->x zbox : blobber doesn't have enough free space
    end
    alt check expiration agaisnt allocation blobbers' max offer duration
        sc ->x zbox : blobber doesn't allow so long offers
    end
    group allocation
        sc -> sc : update alllocation as required\nsize expireation and immutable
        blockchain -> sc : owner wrtie pool
        sc -> sc : ammend write pool lock duration
        blockchain -> sc : challenge pool
        sc -> sc : transfer tokens between\nwrite pool amd challenge pool
        sc -> blockchain : challenge pool
        sc -> blockchain : write pool
    end
    sc -> blockchain : allocation
sc -> zbox : transaction id
```
    
```puml
title Create read pool
zbox -> sc : rp-create
sc -> sc : new read pool
sc -> blockchain : read pool
sc -> zbox
```

```puml
title Lock tokens in read pool for given bobber
zbox -> sc : rp-lock, token value
note left
    * lock duration
    * allocation id
    * blobber id to lock for
end note 
sc -> sc : new allocation pool
group new allocation pool
    sc -> sc : add new blobber pool\nbalance value
    sc -> sc : transfer tokens
    sc -> sc : add set expiration
end
blockchain -> sc : readpool
group read pool
sc -> sc : add new allocation pool
end 
sc -> blockchain : save read pool
sc -> zbox
```

```puml
title Lock tokens in read pool no blobber specified
zbox -> sc : rp-lock, token value
note left
    * lock duration
    * allocation id
end note 
sc -> sc : new allocation pool
group new allocation pool
    blockchain -> sc : get allocation blobbers
    loop each blobber
        sc -> sc : new blobber pool\nbalance value \n split by read price
    end
    sc -> sc : add new blobber pools
    sc -> sc : transfer tokens
    sc -> sc : add set expiration
end
blockchain -> sc : readpool
group read pool
sc -> sc : add new allocation pool
end 
sc -> blockchain : save read pool
sc -> zbox
```

```puml
title Unlock read pool
zbox -> sc : rp-unlock
note left
    * read pool id
end note 
blockchain -> sc : read pool
alt read pool with id expired
    group read pool
        sc -> sc : remove pool id
    end
    sc -> blockchain : transfer tokens from\nread pool id to user
    sc -> blockchain : save read pool
    sc -> zbox
else read pool id not expired
    sc ->x zbox : read pool not expired
end
```

```puml
title Lock tokens in write pool for given bobber
zbox -> sc : rp-lock, token value
note left
    * lock duration
    * allocation id
    * blobber id to lock for
end note 
sc -> sc : new allocation pool
group new allocation pool
    sc -> sc : add new blobber pool\nbalance value
    sc -> sc : transfer tokens
    sc -> sc : add set expiration
end
blockchain -> sc : writepool
group write pool
sc -> sc : add new allocation pool
end 
sc -> blockchain : save write pool
sc -> zbox
```

```puml
title Lock tokens in write pool no blobber specified
zbox -> sc : rp-lock, token value
note left
    * lock duration
    * allocation id
end note 
sc -> sc : new allocation pool
group new allocation pool
    blockchain -> sc : get allocation blobbers
    loop each blobber
        sc -> sc : new blobber pool\nbalance value \n split by write price
    end
    sc -> sc : add new blobber pools
    sc -> sc : transfer tokens
    sc -> sc : add set expiration
end
blockchain -> sc : readpool
group write pool
sc -> sc : add new allocation pool
end 
sc -> blockchain : save write pool
sc -> zbox
```

```puml
title Unlock write pool
zbox -> sc : rp-unlock
note left
    * write pool id
end note 
blockchain -> sc : write pool
alt write pool with id expired
    group write pool
        sc -> sc : remove pool id
    end
    sc -> blockchain : transfer tokens from\nwrite pool id to user
    sc -> blockchain : save write pool
    sc -> zbox
else write pool id not expired
    sc ->x zbox : write pool not expired
end
```














































































