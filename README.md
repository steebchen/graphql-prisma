# GraphQL Server Example

This example shows how to implement a **GraphQL server with Golang** based on Prisma & [gqlgen](https://github.com/99designs/gqlgen).

## How to use

### 1. Download example & install dependencies

Clone the repository:

```
git clone git@github.com:steebchen/graphql.git
```

Ensure dependencies are available and up-to-date:

```
cd graphql
dep ensure -update
```

### 2. Install the Prisma CLI

To run the example, you need the Prisma CLI. Please install it via Homebrew or [using another method](https://www.prisma.io/docs/prisma-cli-and-configuration/using-the-prisma-cli-alx4/#installation):

```
brew install prisma
brew tap
# or
npm i -g prisma
```

### 3. Set up database & deploy Prisma datamodel

Start the server and the database using docker-compose:

```bash
docker-compose up -d
```

Deploy our schema to our database:

```
prisma deploy # this also runs prisma generate and gqlgen
```

### 4. Start the GraphQL server

```
go run .
```

Navigate to [http://localhost:4000](http://localhost:4000) in your browser to explore the API of your GraphQL server in a [GraphQL Playground](https://github.com/prisma/graphql-playground).

### 5. Using the GraphQL API

The schema that specifies the API operations of your GraphQL server is defined in [`./api/schema.graphqls`](./api/schema.graphqls). Below are a number of operations that you can send to the API using the GraphQL Playground.

Feel free to adjust any operation by adding or removing fields. The GraphQL Playground helps you with its auto-completion and query validation features.

#### Log in

```graphql
mutation {
  login(email: "alice@prisma.io", password: "test") {
    id
    email
    name
  }
}
```

#### Get user

```graphql
query {
  user {
    id
    name
    email
  }
}
```

More coming soon.
