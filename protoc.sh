#! /bin/bash
PROTOC_BIN='protoc'
PROTOC_HEAD_PARAMS='--proto_path=.'
PROTOC_TAIL_PARAMS='--go_out=plugins=grpc,paths=source_relative:.'
PROTO_FILES=(
  api/common/status.proto
  api/scraper/scraper.proto
  api/story/story.proto
)
for proto in "${PROTO_FILES[@]}"
do
  #echo "$PROTOC_BIN $PROTOC_HEAD_PARAMS ${proto} $PROTOC_TAIL_PARAMS"
  $PROTOC_BIN $PROTOC_HEAD_PARAMS ${proto} $PROTOC_TAIL_PARAMS
done