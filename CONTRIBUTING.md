# Project setup

You can read more about the setup of the project in the docs (on README). But, in summary, we installed the Golang, created a new Go project (aka, created a new module), and added the dependencies related to the Protobuf.

# Extensions

That are some extension on VS Code that can ease your learning curve. They're optional, but very recommended.

- Protocol Buffers (https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
- Go (https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - This extension depends on a set of extra command-line tools. If they are missing, the extension will show the "⚠️ Analysis Tools Missing" warning. Click the notification to complete the installation.

# Directories

- `proto`: This directory persists all Protobuf schemas. Since we have the same use-case in all 4 approaches, we're using the same schema in all of them.
- `pb`: This directory persists all the compiled Protobuf schemas. The schemas from the `proto` directory must be compiled, in order for us to use them in our source code.
- `request-response`: Here we have the classic request-response pattern, using gRPC and Golang
- `bidirectional`: Here we have the bidirectional pattern, using gRPC and Golang
- `client-streaming`: Here we have the unidirectional pattern, streaming from the client to the server, using gRPC and Golang
- `server-streaming`: Here we have the unidirectional pattern, streaming from the server to the client, using gRPC and Golang

# Schemas

Our schemas are the entities declared in Protobuf. They are sent through the gRPC tunnel. As said before, our use-case is very simplistic, that's why we only have a few schemas, as follows:

- `user.proto`: Schema file with the definition of our entity/model and services (with its contracts)
