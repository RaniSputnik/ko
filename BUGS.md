# Bugs

### Neelance GraphQL

Schema;

```
createMatch(boardSize:Int! = 19): Match!
```

Response;

```
{
  "errors": [
    {
      "message": "Field \"createMatch\" argument \"boardSize\" of type \"Int!\" is required but not provided.",
      "locations": [
        {
          "line": 2,
          "column": 3
        }
      ]
    }
  ]
}
```