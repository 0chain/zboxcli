```puml
title Get read pool stats
boundary zbox 
control 0chain
zbox -> 0chain : getReadPoolStat
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