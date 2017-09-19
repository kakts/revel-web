package services

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

var (
    Session *mgo.Session
    err error
)

// dialURL to mongodb instance
const dialURL = "mongodb://192.168.33.10"

// Create new session.
// if session has already exists, reuse it.
func cloneSession(dialURL string) *mgo.Session {
    if Session == nil {
        Session, err = mgo.Dial(dialURL)
        if err != nil {
            panic(err)
        }
    }
    return Session.Clone()
}

// Insert document to specified db.
func InsertEntity(dbName string,
    collection string,
    model interface{}) {

    // Get DB session
    session := cloneSession(dialURL)
    defer session.Close()

    err := session.DB(dbName).C(collection).Insert(model)
    if err != nil {
        panic(err)
    }
}

// Remove document from collection
func RemoveEntry(dbName string,
    collection string,
    model interface{}) {

    session := cloneSession(dialURL)
    defer session.Close()

    err := session.DB(dbName).C(collection).Remove(model)
    if err != nil {
        panic(err)
    }
}

func queryCollection(dbName string,
    collection string,
    query func(c *mgo.Collection) error) error {

    session := cloneSession(dialURL)
    defer session.Close()

    c := session.DB(dbName).C(collection)
    return query(c)
}

func FindOne(userId string,
    dbName string,
    collection string,
    ) (results []interface{}, errorCode string) {
    query := bson.M{"userid": userId}
    return Search(query, 0, 1, dbName, collection)
}

func Search(q interface{},
    skip int,
    limit int,
    dbName string,
    collection string,
) (results []interface{}, errorCode string) {

    errorCode = ""
    query := func(c *mgo.Collection) error {
        fn := c.Find(q).Skip(skip).Limit(limit).All(&results)
        if limit < 0 {
            fn = c.Find(q).Skip(skip).All(&results)
        }
        return fn
    }

    search := func() error {
        return queryCollection(dbName, collection, query)
    }

    err := search()
    if err != nil {
        errorCode = "error"
    }
    return
}
