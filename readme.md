# Golang FizzBuzz API

## Project Purpose

I'd like to create own API with Golang after I learned udemy course, so this project have been starting.

## Git commit message pattern

Use only **Present Tense** with commit message

Mostly used below words as message prefix

- Create
- Add
- Update
- Fix
- Delete

## Git branch pattern

- Main branch is **master**
- Feature branch is **feature/what-do-need-to-do-into-this-branch**

## API Spectification

`POST /setModulationSequence`

### Request

- Send input via **Request Body**
- Format: JSON

```json
{
    "modulationSequence": [
        {
            "key": 3,
            "value": "Fizz"
        }
    ]
}
```

### Response

- Get successful result
- Return http status 200

```json
{
    "updateCount": 1,
    "modulationSequence": [
        {
            "key": 3,
            "value": "Fizz"
        }
    ]
}
```

- Get failed result
- Return http status 400, 500

```json
{
    "error": "Error Message"
}
```

`GET /saying/{number}`

### Request

- Send number via **QueryString**
- Replace {number} by positive number
- Example: `/saying/3`

### Response

- Get successful result
- Return http status 200

```json
{
    "saidCount": 1,
    "say": "Fizz"
}
```

- Get failed result
- Return http status 400, 500

```json
{
    "error": "Error Message"
}
