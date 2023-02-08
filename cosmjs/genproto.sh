#!/usr/bin/env bash

proto_dirs=$(find ./proto/checkers -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
dir_test=$(find . \( -path ./gogoproto -prune -o -path ./cosmos -prune -o -path ./google -prune \) -o -name '*.proto' -print)
#echo "$dir_test"

for d in $proto_dirs; do
  while IFS= read -r -d '' file
  do
    ./node_modules/protoc/protoc/bin/protoc \
      --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
      --ts_proto_out="./client/src/types/generated" \
      --proto_path="./proto" \
      --ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
      "$file"
  done < <(find "$d" -maxdepth 1 -name '*.proto' -print0)
done