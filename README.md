# apollo-challenge 1

Thats a proxy that accepts incoming GraphQL request and sends it to Github API to create/update/read gists

<pre>
                                  ┌─────────────────────┐                      ┌─────────────────────┐
                                  │                     │                      │                     │
                                  │                     │◀──Create gist───────▶│                     │
                                  │                     │                      │                     │
                                  │                     │◀──Edit gist─────────▶│                     │
                                  │                     │                      │                     │
┌────────────┐                    │                     │                      │                     │
│            │     GraphQL        │       graphQL       │                      │  Github Gists API   │
│   client   │◀────request───────▶│    proxy server     │◀──Get gist by id────▶│       (REST)        │
│            │                    │                     │                      │                     │
└────────────┘                    │                     │◀──Get public gists──▶│                     │
                                  │                     │                      │                     │
                                  │                     │◀──Get gist commits──▶│                     │
                                  │                     │                      │                     │
                                  │                     │◀──Get gist comments─▶│                     │
                                  └─────────────────────┘                      └─────────────────────┘
</pre>

### Prerequisites
1. Go (1.21+)
2. Docker
3. Any tool to send GraphQL requests (Postman for example)
4. Github personal access token (PAT) with access to gists

### Local run
Server serving at 8080 by default
1.
    ```shell
    go run ./server.go
    ```
2. run via docker-compose (will build server and then run it.)
    ```shell
    docker-compose up
    ```
3. That's all! Server ready to accept incoming requests

### Example requests

1. Create gist
    ```graphql
    mutation CreateGist {
        createGist(
            input: {
                description: "test graphQL proxy by @AmirDevPostHack"
                public: false
                files: { fileName: "README.md", content: "## Apollo challenge 1" }
            }
        ) {
            id
        }
    }
    ```

2. Update gist:
    ```graphql
    mutation UpdateGist {
        updateGist(
            gist_id: "<gist_id_from_previous_step>"
            input: {
                description: "<new description>"
                public: false
                files: { fileName: "file.<extension>", content: "<file contents>" }
            }
        ) {
            id
            description
            files {
                fileName
            }
        }
    }
    ```