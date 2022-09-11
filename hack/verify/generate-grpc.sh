#!/bin/bash

# Copyright 2021 The Kubernetes Authors.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Install the protobuf and protoc-gen for GO
function install_grpc {

    sudo apt install -y protobuf-compiler

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

    export PATH="$PATH:$(go env GOPATH)/bin"
}

# To generate the gRPC services
function generate_stubs {
    echo "---Printingggggggggggggg------"
    ls -al api/localnetv1/verify-protobuf/
    protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $(find . -name '*.proto')
    echo "--------+++++++++++++++--------------"
    ls -al api/localnetv1/verify-protobuf/
    echo "End of generate"

}

# Comment out if the install_gprc if already present
install_grpc
generate_stubs
