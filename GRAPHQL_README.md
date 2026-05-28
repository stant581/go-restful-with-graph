# GraphQL API Documentation

## Overview
This REST API has been converted to a **GraphQL API**. The key benefit of GraphQL is that only the fields you query will be returned in the response, reducing data transfer and improving efficiency.

## Getting Started

### Installation
```bash
go mod download
go mod tidy
```

### Running the Server
```bash
go run main.go
```

The server starts on `http://localhost:8080` by default.

## Endpoints

### GraphQL Endpoint
- **URL**: `POST /graphql`
- **Playground**: `GET /graphql` (Interactive GraphQL IDE)

Access the interactive GraphQL Playground at `http://localhost:8080/graphql`

## Query Examples

### Query 1: Get User with Only Specific Fields
```graphql
query {
  user(id: 1) {
    userId
    userName
  }
}
```
**Response** (only requested fields):
```json
{
  "data": {
    "user": {
      "userId": 1,
      "userName": "John Doe"
    }
  }
}
```

### Query 2: Get User with Account Information
```graphql
query {
  user(id: 1) {
    userId
    userName
    account {
      accountId
      accountName
      status
    }
  }
}
```

### Query 3: Get Account
```graphql
query {
  account(id: 1) {
    accountId
    accountName
    status
  }
}
```

### Query 4: Get User Profile
```graphql
query {
  userProfile(id: 1) {
    user {
      userId
      userName
    }
    account {
      accountName
      status
    }
  }
}
```

## Mutations

### Create User
```graphql
mutation {
  createUser(userName: "Jane Smith", accountId: 2) {
    userId
    userName
    accountId
  }
}
```

### Create Account
```graphql
mutation {
  createAccount(accountName: "Acme Corp", status: true) {
    accountId
    accountName
    status
  }
}
```

## Key Benefits of GraphQL

1. **Over-fetching Prevention**: Only request the fields you need
2. **Under-fetching Prevention**: Get all required data in a single request
3. **Strongly Typed**: Self-documenting API with type definitions
4. **Flexible Queries**: Clients define their data requirements
5. **Better Performance**: Reduced bandwidth usage by fetching only required fields

## Schema

### Types

**User**
- `userId: Int!` - Unique user identifier
- `userName: String!` - User's name
- `accountId: Int!` - Associated account ID
- `account: Account` - Related account (loaded on demand)

**Account**
- `accountId: Int!` - Unique account identifier
- `accountName: String!` - Account name
- `status: Boolean!` - Account status

**UserProfile**
- `user: User!` - User information
- `account: Account` - Related account information

## Testing in Playground

1. Navigate to `http://localhost:8080/graphql`
2. Enter queries in the left panel
3. Click the play button to execute
4. View results in the right panel
