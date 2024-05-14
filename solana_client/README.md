## Solana client to be called by Relayer

Build the program client under ` ./src/index.ts` into "boxednode\_client" binary using [boxednode](https://github.com/mongodb-js/boxednode?tab=readme-ov-file)

Relayer process at `../relayer/publish/skateapp.go` interact with Solana by invoking this binary.

_NOTE: to scale up, we need to change the client communication with relayer process via Unix-socket_
