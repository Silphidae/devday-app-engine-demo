const express = require('express');
const router = express.Router();
const bodyParser = require('body-parser');
const model = require('../model/candy');

router.use(bodyParser.json());

router.get('/', function(req, res, next) {
    model.list(100, req.query.pageToken, (err, entities, cursor) => {
        if (err) {
            next(err);
            return;
        }
        res.json({ items: entities });
    });
});

router.post('/', function(req, res, next) {
    model.create(req.body, (err, entity) => {
        if (err) {
            next(err);
            return;
        }
        res.json(entity);
    })
})

module.exports = router;
