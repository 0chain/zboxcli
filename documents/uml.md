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




























