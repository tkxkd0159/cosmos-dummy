# checkers
**checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Requirements
[Checkers Rules](https://www.ducksters.com/games/checkers_rules.php)
* What Ignite CLI commands will get you a long way when it comes to implementation?
* How do you adjust what Ignite CLI created for you?
* How would you unit-test your modest additions?
* How would you use Ignite CLI to locally run a one-node blockchain and interact with it via the CLI to see what you get?

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/checkers@latest! | sudo bash
```
`username/checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

# Settings
```shell
ignite scaffold single systemInfo nextId:uint --module checkers --no-message
ignite scaffold map storedGame board turn black red --index index --module checkers --no-message
ignite scaffold message createGame black red --module checkers --response gameIndex

export alice=$(checkersd keys show alice -a)
export bob=$(checkersd keys show bob -a)
checkersd tx checkers create-game $alice $bob --from $alice --dry-run
checkersd tx checkers create-game $alice $bob --from $alice --broadcast-mode block
checkersd query checkers show-stored-game 1 --output json | jq ".storedGame.board" | sed 's/"//g' | sed 's/|/\n/g'

ignite scaffold message playMove gameIndex fromX:uint fromY:uint toX:uint toY:uint --module checkers --response capturedX:int,capturedY:int,winner
checkersd tx checkers play-move 1 0 5 1 4 --from $bob   # out of turn
checkersd tx checkers play-move 1 1 0 0 1 --from $alice # wrong move
checkersd tx checkers play-move 1 1 2 2 3 --from $alice # success
checkersd query checkers show-stored-game 1 --output json | jq ".storedGame.board" | sed 's/"//g' | sed 's/|/\n/g' # check alice's piece

ignite scaffold message rejectGame gameIndex --module checkers


# etc.
# 1) test specific package
go test -v checkers/x/checkers/keeper

# 2) reset states & start chain with single node. Reload when update but only reset states at first starting (not for reloading)
ignite chain serve --reset-once

# 3) recompile proto
ignite generate proto-go
```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
