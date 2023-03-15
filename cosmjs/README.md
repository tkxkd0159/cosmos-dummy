# keplr setup
```shell
npx create-next-app@latest --typescript # cosmjs-keplr
npm run dev # reload when change code
```

# proto setup
```shell
npm install ts-proto protoc --save-dev
./node_modules/protoc/protoc/bin/protoc --version

mkdir -p ./proto/cosmos/base/query/v1beta1/
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.46.6/proto/cosmos/base/query/v1beta1/pagination.proto -o ./proto/cosmos/base/query/v1beta1/pagination.proto
mkdir -p ./proto/google/api
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto -o ./proto/google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto -o ./proto/google/api/http.proto
mkdir -p ./proto/gogoproto
curl https://raw.githubusercontent.com/cosmos/gogoproto/v1.4.4/gogoproto/gogo.proto -o ./proto/gogoproto/gogo.proto

# compile
mkdir -p client/src/types/generated
find ./proto -name "*.proto" | xargs -I {} ./node_modules/protoc/protoc/bin/protoc \
  --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
  --ts_proto_out="./client/src/types/generated" \
  --proto_path="./proto" \
  --ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" {}
```
