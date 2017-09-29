https://hackernoon.com/learn-blockchains-by-building-one-117428612f46
https://learncryptography.com/hash-functions/what-are-hash-functions

` pip install Flask==0.12.2 requests==2.18.4 `

Blockchain is an immutable, sequential chain of records called Blocks.
They can contain transactions, files or any data. But the important
thing is that they're chained together using hashes.

Each block has an index, a timestamp (unix time), a list of
transactions, a proof, and the hash of the previous block.

```
block = {
    'index': 1,
    'timestamp': 1506057125.900785,
    'transactions': [
        {
            'sender': "8527147fe1f5426f9dd545de4b27ee00",
            'recipient': "a77f5cdfa2934df3954a5c7c7da5df1f",
            'amount': 5,
        }
    ],
    'proof': 324984774000,
    'previous_hash': "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
}
```

each new block contains within itself the hash of the previous block.
This is crucial because it's what gives blockchains immutability. If an
attacker corrupted an earlier Block in the chain then all subsequent
blocks will contain incorrect hashes.

After `new_transaction()` adds a transaction to the list, it returns
the index of the block which the transaction will be added to - the next
one to be mined.

When Blockchain is instantitated we'll need to seed it with a genesis
block - a block with no predecessors. We'll also need to add a proof to
our genesis block which is the result of mining (or proof of work).


proof of work
A proof of work algorithm (PoW) is how new Blocks are created or mined
on the block chain. The goal of PoW is to discover a number which solves
a problem. The number must be difficult to find but easy to verify by
anyone on the network. This is the core idea behind the proof of work.

eg:
let's decide that the hash of some integer x multiplied by another y
must end in 0. So, hash(x*y) = abcde...0. Let's fix x = 5

```
from hashlib import sha256

x = 5
y = 0 # we don't know what y should be yet

while sha256(f'{x*y}'.encode()).hexdigest()[-1] != "0":
    y += 1
```

the solution is hash(5*21) =  1253e9373e...5e3600155e860

y = 21

In bitcoin the proof of work algorithm is called HashCash. It's the
algorithm miners race to solve in order to create a new block. In
general, the difficulty is determined by the number of characters
searched for in a string. The miners are then rewarded for their
solution by receiving a coin - in a transaction.

The network is able to easily verify the solution.

implementing basic proof of work
find a number p that when hashed with the previous block's solution in a
hash with 4 leading 0s is produces.

To adjust the difficulty of the algorithm, we could modify the number of
leading zeroes. But 4 is sufficient. Addition of a single leading zero
makes a mammoth difference to the time required to find a solution.


block chain as an api
/transactions/new - to create a new transaction
/mine - to tell the server to mine a new block
/chain - to return the full blockchain


mining endpoing is where the magic happens. it has to do 3 things.
1) calculate the proof of work
2) reward the miner by adding a transaction granting 1 coin
3) forge the new block by adding it to the chain

