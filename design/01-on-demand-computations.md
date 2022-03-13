# On-demand computations

This chain is meant to be an experiment for two ideas at the same time:

* Consolidating actual price feed data from Redstone data provider nodes.
* Generalizable on-demand computation architecture.

Consolidating price data is a much more limited problem that could probably
be solved more efficiently with something similar to Chainlink OCR. But we
want to use a Cosmos blockchain in order to have an example chain that
performs a useful distributed computation and enforces consensus. This
computation happens to be a very simple one: median of data from different
providers, but we can still use it to test more generic ideas.

Now let's discuss different modes of computation that we can potentially
support. To formalize our thoughts, let's think of computation in terms
of "tasks" that clients may ask the chain to run. Each task has some number of
parameters (think: asset, time interval), and running it requires
some input data (which we generally assume is stored on-chain) and produces
an output (think: median price). We avoid the term "query" because it has a
special meaning in most blockchains and implies reading data from one node
without relying on any consensus mechanism.

## Eager computation

For domains where the set of tasks is limited and predictable, we can imagine
eagerly running all of them. This may be the case for price consolidation:
it is reasonable to assume that at least initially, if a consolidated feed
was set up for some assets, there will be clients using its outputs very
often. So it may make sense to just compute all of them.

The downside here is, of course, that the amount of required computation and
storage (for results) scales linearly with domain space (number of assets).
This makes it harder to add less popular assets, which will only be consumed
occasionally.

## Same-block on-demand computation

Now let's consider a different model: where the chain stores the input data
required for computations, but doesn't run any computation tasks unless
explicitly requested. A request comes from an address in a form of transaction.
It includes the task parameters, and likely some fee (which at this point
can be implemented as gas price).

The transaction reaches the mempool, and is eventually selected for processing
in some block. All nodes in the network will consider it in the `RunTx`
Tendermint callback, and produce a result, which will be stored for future
queries and emitted as an event.

There are two obvious limitations here:

* This can only work for tasks that we know are cheap. Doing any potentially
  expensive computation in `RunTx` is a recipe for failure.

* We need to produce a result immediately. If there are not enough input data
  (prices from individual Redstone data providers), then we have to respond
  with a `null`, conceptually.

First is fine for price consolidation, second is a bit sad.

## Delayed synchronous computation

Let's first try to solve the input data availability issue. A natural thing
to do seems to be producing the results whenever enough input data become
available, not necessarily in the same block where the task was received.

Since input data can only become available as a result of transactions, we
can design the chain so that it tracks the pending tasks and which data they
require, and as soon as new data arrives and some tasks can be completed, they
must be completed in the same block. We can do this by storing results and
emitting events in `EndBlock`. `EndBlock` must process results for all
previously pending tasks that just became completed, so we still need to be
very careful to ensure that our computation is fast.

This modification is useful for tasks similar to price consolidation, where
input data availability is an issue, but the computation itself is cheap.

## Delayed asynchronous computation

(This is approximately the part of this document where we move from
ideas we could actually use for price consolidation to ideas mostly interesting
from a generic computation perspective)

If we want to support any kinds of computation that may be too expensive for
a single block (so maybe more than 1 second of run time), we need to somehow
perform them outside of the main Tendermint consensus loop. That said, we still
want some redundancy and consensus about the results of these computations.

A natural step in this direction seems to be adding a "sidecar" process to every
node in the chain that listens to new tasks, performs the computation, and
sends the results to the node. It is assumed that the node and its sidecar
process completely trust each other, so they should most likely run on the
same machine to prevent any network attack vectors.

The interesting question now is: how exactly are the results published to the
chain and agreed upon? Since our computation is asynchronous (it happens in
a background process), the time when result becomes available on different
nodes may be vastly different. In particular, nothing guarantees that all
nodes will get the results of their computation in the same block.

There are two general ways in which we can handle this. One is fully decoupling the
result consensus algorithm from the Tendermint block consensus algorithm.
This means that every node publishes its result as a transaction, and when
enough results have been provided, the chain agrees on the consensus result.
This seems easy until we start thinking about what happens in the case of
different results from different nodes. This situation is interesting: it
implies that the chain has consensus on the Tendermint level (transactions,
blocks, state), but not on the "application" level. So if we wanted to handle
this somehow (exclude faulty nodes, slash their stake, etc.), we would need
to duplicate whatever Tendermint and Cosmos SDK already do for nodes that
provide faulty blocks.

Another way could be reusing the Tendermint consensus for computation results
as well. To handle the issue of asynchronisity, each node sends a "task is ready"
transaction first as soon as it gets the result from its sidecar. After enough
"task is ready" transactions were processed, in the block where the threshold
was reached, the nodes "reveal" their results via sending an event in
`EndBlock`. This way if a minority of nodes send results different from the
majority, Tendermint will treat them the same as any other kind of faulty
nodes.

This latter idea seems nice, but unfortunately it has a major drawback. If the
nodes are required to publish their results in `EndBlock`, then absolutely
all valid nodes must compute results. Triggerring the reveal before a "ready"
transaction is collected from every node would be unfair, so we have to wait
until absolutely every node is finished, which makes things slower. We also
need to know on the application level how many nodes are currently in the
network, which seems very iffy if we consider different kinds of nodes (full
vs. lite), validator set changes, etc. Besides that, we also lose the possibility
to treat different failure modes differently: we would have to hand over control
to Tendermint and rely on the same punishment logic for:

* Actually creating invalid blocks.
* Computing wrong results.
* Not computing results in time. (timeouts are not explicitly discussed, but
  we obviously need to handle them in this model)

Considering the downsides of the latter idea, it seems that the former
(app-level result consensus) is much more viable. We would likely track
the consensus state per-task in storage, and provide queries to read results
that have reaches consensus.

## Decoupling computation coordination from input data

An astute reader may notice that by this point we keep input data on the same
chain as the computation tasks and their results mostly by inertia from a much
earlier design. We don't really need to do this: we can use different chains
with different properties for storing data and coordinating computations.

Now, if we look at this "chain for coordinating computations", and zoom out
a little, we will realize that we just arrived at the original pre-OCR
Chainlink design. The only difference is that Chainlink performed this
"coordination of computation", or (less abstractly) collecting data from
oracles on the same chain where end users needed the "results" (Oracle data),
while our design uses a separate custom Cosmos chain.

So how do we reliably transfer results from our coordination chain to client
chains that need them? Well... there is this thing called "oracles"... :)

## Practical considerations

So besides being an interesting thought experiment, is any of this actually
useful for our three-step price consolidation roadmap?

Well, not completely. As we mentioned in the beginning, price consolidation does
not need a blockchain per se, though it could be useful for governance and
financial incentives. But if we squint a little and think of Cosmos as a
consensus mechanism with some extra features (some of which may actually be
useful some time down the road), it doesn't seem crazy to use it.

Concretely, this could fit into the second step: one package, multiple
signatures. We need to decide who the signatories are, and how we keep
client chains aware of the signatories (hardcoding? upgradable contracts?
reading from a central registry?). Once this is done (notice that this is
orthogonal to producing consolidated prices), the signatories, whoever they
are, have to just listen for events coming out of the consolidation chain,
sign them, and publish their signatures to some PubSub where client code
can collect them from and send to client chains.
