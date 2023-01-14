# Setup
```shell
docker build -f Dockerfile-ubuntu --build-arg BUILDARCH=amd64 -t iignite .
docker run --rm -it iignite ignite version # check version after install
docker run --rm -it -v $(pwd):/app -w /app iignite ignite scaffold chain <project_gomod_name>

# 1317 : REST API, 26657 : Tendermint API
# 3000 : GUI for chain, 4500 : token faucet

# Throwaway version
cd checkers
docker run --rm -it -v $(pwd):/app -w /app -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name jschain iignite ignite chain serve
# Persistent version
docker create --name jschain -i -v $(pwd):/app -w /app -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 iignite
docker start jschain
docker exec -it jschain ignite chain serve

# Set GUI
docker exec -it jschain bash -c "cd vue && npm install"
docker exec -it jschain bash -c "cd vue && npm run dev -- --host"

# Interaction
docker exec -it jschain bash -c "checkersd status 2>&1 | jq"
docker exec -it jschain <name>d --help
docker exec -it jschain <name>d status --help
docker exec -it jschain <name>d query --help
```
## Generate message
```shell
docker run --rm -it -v $(pwd):/app -w /app iignite ignite scaffold message createPost title body

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