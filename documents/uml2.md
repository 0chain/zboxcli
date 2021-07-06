```puml
title Get write pool information
boundary zbox 
control 0chain
zbox -> 0chain : REST API //getWritePoolStat//
0chain -> zbox : write pool info
```

```puml
title Lock tokens in stake pool
boundary zbox 
control storagesc
entity blockchain
zbox -> storagesc : txn: //stake_pool_pay_interests//
note left
    * blobber id
end note
    blockchain -> storagesc : blobber stake pool
    group blobber stake pools
        storagesc -> storagesc : mint interest
    end
    storagesc -> blockchain : save blobber stake pool
storagesc -> zbox : 
```

```puml
title Get stake pool information for user
boundary zbox 
control 0chain
zbox -> 0chain : REST API //getUserStakePoolStat//
0chain -> zbox : stake pool info
```

```puml
title Unlock stake pool
boundary zbox 
control storagesc
entity blockchain
zbox -> storagesc : txn: //stake_pool_unlock//
note left
    * blobber id
    * pool id
end note
    blockchain -> storagesc : blobber stake pools
    group blobber stake pool
        storagesc -> storagesc : pay interest
        alt stake pool locked for offer
            storagesc -> storagesc : marke to be unlocked
            storagesc -> blockchain : blobber stake pool
            storagesc -> zbox : max time to unlock
        end
        storagesc -> blockchain : transfer funds to user
        storagesc -> storagesc : remove stake pool
    end
    blockchain -> storagesc : user stake pools
    storagesc -> storagesc : remove pool from user staek pools
    storagesc -> blockchain : blobber stake pools
    storagesc -> blockchain : user stake pools   
storagesc -> zbox :  amount transfer back
```

```puml
title Lock tokens in stake pool
boundary zbox 
control storagesc
entity blockchain
zbox -> storagesc : txn: //stake_pool_lock//
note left
    * blobber id
end note
    alt lock > min lock\nmax delegates not exceeded
        0chain ->x zbox    
    end
    blockchain -> storagesc : blobber stake pool
    group blobber stake pools
        storagesc -> storagesc : mint interest
        storagesc -> storagesc : new pool
    end
    storagesc -> blockchain : blobber stake pool
    blockchain -> storagesc : user stake pools
    storagesc -> storagesc : add new stake pool
    storagesc -> blockchain : user stake pools  
storagesc -> zbox : new pool id
```




```puml
title Get stake pool information
boundary zbox 
control 0chain
zbox -> 0chain : REST API //getStakePoolStat//
0chain -> zbox : stake pool info
```

```puml
title Get smart contract configuration
boundary zbox 
control 0chain
zbox -> 0chain : REST API //getConfig//
0chain -> zbox : smart contract configuration
```

```puml
title Get read pool stats
boundary zbox 
control 0chain
zbox -> 0chain : REST API //getReadPoolStat//
0chain -> zbox : read pool stats
zbox -> zboc : filter stats by allocation id
```

```puml
title Delete collaborator
boundary zbox 
collections blobbers
database store
control 0chain
zbox -> blobbers : delete collaborator
    blobbers -> 0chain : get allocation
    0chain -> blobbers : allocation
    alt check owner == sender
        blobbers ->x zbox : allocation must be owner
    end
    blobbers -> store : delete colaborator
    blobbers -> zbox
```

```puml
title Get challenge pool stats
boundary zbox 
control 0chain
zbox -> 0chain : getChallengePoolStat
0chain -> zbox : challenge pool stats
```