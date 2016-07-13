const express = require(`express`),
    bodyParser = require(`body-parser`),
    port = 3000;

let app = express();
app.use(bodyParser.json());

app.get('*', (req, res) => {
    let guid = Math.random().toString().replace(`.`, ``).slice(0, 4);
    for (let key in req.query) {
        if (req.query.hasOwnProperty(key)) {
            console.log(`Query GUID ${guid} IS ${key} : ${req.query[key]}`);
        }
    }
    res.status(200).send("All good");
});

app.post('*', (req, res) => {
    let guid = Math.random().toString().replace(`.`, ``).slice(0, 4);
    for (let key in req.body) {
        if (req.body.hasOwnProperty(key)) {
            console.log(`Data GUID ${guid} IS ${key} : ${req.body[key]}`);
        }
    }
    res.status(200).send("All good");
});

app.listen(port, () => console.log(`Listening on port ${port}`));