# apollo-challenge 1

This project is a GraphQL server built with Go. It serves as a proxy to the GitHub REST API,
specifically for operations related to GitHub Gists. 
The server accepts incoming GraphQL requests and forwards them to the GitHub API to create, update, and read gists.

### Table of contents
- [Features](#features)
- [Architecture overview](#system-architecture-overview)
- [Available operations](#available-operations)
- [Prerequisites](#prerequisites)
- [Local run](#local-run)
- [Example requests](#example-requests)
---

### Features
* GraphQL API: This project exposes a GraphQL API for creating, updating, and reading GitHub Gists. GraphQL offers more efficient data querying and manipulation than traditional REST APIs. It allows clients to specify exactly what data they need, which can significantly reduce the amount of data that needs to be transferred over the network and improve performance.
* GitHub Integration: The server uses the GitHub REST API under the hood. GitHub Gists are a great way to share snippets of code or other notes, but the REST API can be verbose and require multiple round trips to the server to gather all the necessary data. By integrating with the GitHub API and exposing the functionality through a GraphQL API, this project provides a more efficient interface to work with Gists.

This project is particularly useful for developers who frequently work with GitHub Gists and want a more efficient way to interact with them. By providing a GraphQL interface to the GitHub Gist API, it allows developers to get exactly the data they need in a single request, rather than having to make multiple requests to the REST API. This can lead to performance improvements and a better developer experience.

---

### System Architecture Overview
The given schema illustrates a three-tier architecture involving a client, a GraphQL proxy server, and the GitHub Gists API. This setup enables the client to communicate with the GitHub Gists API using GraphQL, even though the GitHub Gists API natively uses REST. The GraphQL proxy server acts as an intermediary layer, converting GraphQL queries and mutations from the client into REST API calls to the GitHub Gists API, and vice versa.
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

---

### Available operations
* Create gist
* Edit gist
* Retrieve gist by id
* Retrieve public gists
* Retrieve gist commits
* Retrieve gist comments

---

### Prerequisites
1. Go (1.21+)
2. Docker
3. Any tool to send GraphQL requests (Postman for example)
4. Github personal access token (PAT) with access to gists

---

### Local run
Server serving at 8080 by default
1. run via go tools
    ```shell
    go run ./server.go
    ```
2. run via docker-compose (will build server and then run it.)
    ```shell
    docker-compose up
    ```
3. That's all! Server ready to accept incoming requests

---

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