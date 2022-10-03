# Setup
```shell
docker build -f Dockerfile-ubuntu . -t checkers_i
docker run --rm -it checkers_i ignite version # check version after install
docker run --rm -it -v $(pwd):/checkers -w /checkers checkers_i ignite scaffold chain <project_gomod_name>

# 1317 : REST API, 26657 : Tendermint API
# 3000 : GUI for chain, 4500 : token faucet

# Throwaway version
docker run --rm -it -v $(pwd):/checkers -w /checkers -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name checkers checkers_i ignite chain serve

# Persistent version
docker create --name checkers -i -v $(pwd):/checkers -w /checkers -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 checkers_i
docker start checkers
docker exec -it checkers ignite chain serve

# Set GUI
docker exec -it checkers bash -c "cd vue && npm install"
docker exec -it checkers bash -c "cd vue && npm run dev -- --host"

# Interaction
docker exec -it checkers bash -c "checkersd status 2>&1 | jq"
docker exec -it checkers checkersd --help
docker exec -it checkers checkersd status --help
docker exec -it checkers checkersd query --help
```
## Generate message
```shell
docker run --rm -it -v $(pwd):/checkers -w /checkers checkers_i ignite scaffold message createPost title body

```
```protobuf
message MsgCreatePost {
  string creator = 1;
  string title = 2;
  string body = 3;
}
```
```go
func CmdCreatePost() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "create-post [title] [body]",
    Short: "Broadcast message createPost",
    Args:  cobra.ExactArgs(2),
  }
}
```