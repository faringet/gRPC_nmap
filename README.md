# gRPC_nmap
gRPC service wrapper over nmap
___
- Sets gRPC between the client and the server
- Checks the correct number of address arguments
- Informs the user about errors (such as failed server start/request) via `<fmt.Errorf>`.

___
### Quick start
- `<make build>` for server/client build
- `<./bin/server :3000>` to start the server
- `<./bin/client :3000>` to run the client



