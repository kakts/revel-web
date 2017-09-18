package main
import (
    "fmt"
    mgo "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
    ID bson.ObjectId `bson:"_id"`
    Name string `bson:"name"`
    Age int `bson:"age"`
}

func main() {
    session, err := mgo.Dial("mongodb://192.168.33.10/:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // test DB session取得
    db := session.DB("local")

    /**
     * みつけるところ
    **/
    p := new(Person)
    query := db.C("startup_log").Find(nil)
    query.One(&p)
    fmt.Printf("%+v\n", p)
}
