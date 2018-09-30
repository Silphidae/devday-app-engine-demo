'use strict';

const Datastore = require('@google-cloud/datastore');

const datastore = new Datastore({
    projectId: 'kristas-demo-2018'
});
const kind = 'Candy';

function fromDatastore (obj) {
  obj.id = obj[Datastore.KEY].id;
  return obj;
}

function list(limit, token, cb) {
    const q = datastore.createQuery(kind).order('Created');

    datastore.runQuery(q, (err, entities, nextQuery) => {
        cb(err, err ? null : entities.map(fromDatastore))
    })
}

function create(data, cb) {
    const key = datastore.key(kind);
    const entity = {
        key: key,
        data: [
            { name: 'Created', value: new Date().toJSON() },
            { name: 'Text', value: data.text, excludeFromIndexes: true }
        ]
    }

    datastore.save(entity, (err) => {
        cb(err, err ? null : entity.data);
    });
}

module.exports = {
    create,
    list
};
