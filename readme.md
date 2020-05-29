# Golang FizzBuzz API

## Project Purpose

I would like to create own API with Golang after I learned udemy course, so this project have been starting.

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

## How to run ATDD

I would like to drive my project with **Test First Programming**, so I create test first before I start implementation.

In the plan, I'm gonna move them to run into docker container, because I don't need to install golang in CI/CD environment.

```sh
# Run ATDD by go test framework
go test ./atdd
```

## How to run Stub API

I created stub api for make sure my ATDD can run correctly. Used development only

```sh
# When need to build and update image
docker build -t fizzbuzz-stub-api stub_api

# When need to run for test
docker run -d -p 8080:8080 fizzbuzz-stub-api
```

## API Spectification

### POST /setModulationSequence

#### Request

- Send input via **Request Body**
- Format: JSON

```json
{
    "modulationSequences": [
        {
            "key": 3,
            "value": "Fizz"
        }
    ]
}
```

#### Response

- Get successful result
- Return http status 200

```json
{
    "updateCount": 1,
    "modulationSequences": [
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

### GET /saying/{number}

#### Request

- Send number via **QueryString**
- Replace {number} by positive number
- Example: `/saying/3`

#### Response

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
