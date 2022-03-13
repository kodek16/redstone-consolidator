# Redstone consolidation chain

This repo explores different configurations of a custom blockchain designed
to handle consolidating data packages from multiple Redstone data providers
for the same price feed.

## Building

This project uses Cosmos SDK and Starport. You first need to have Go installed
in some manner reasonable for your platform. You can then install Starport.
The official docs suggests `curl ... | bash` which is not something I'm a fan
of, so I recommend you instead build from source:

```bash
git clone https://github.com/tendermint/starport --shallow-since=2021-11-01
cd starport
git checkout v0.19.4
make build
alias starport=$(pwd)/dist/starport
```

You may want to copy the resulting `./dist/starport` binary to some location on
your `PATH`, maybe `/usr/local/bin`.
