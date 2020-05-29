const express = require('express')
const app = express()

app.post('/setModulationSequence', (req, res) => {
    res.json({
        "updateCount": 1,
        "modulationSequences": [
            {
                "key": 3,
                "value": "Fizz"
            },
            {
                "key": 5,
                "value": "Buzz"
            }
        ]
    })
})

app.get('/say/:number', (req, res) => {
    const lookupTable = {
        1: "1",
        3: "Fizz",
        5: "Buzz",
        6: "Fizz",
        10: "Buzz",
        15: "FizzBuzz",
    }
    const number = Number(req.params.number)
    res.json({
        "saidCount": 1,
        "say": lookupTable[number]
    })
})

app.listen(8080)
