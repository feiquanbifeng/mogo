package mogo

import (
    "labix.org/v2/mgo"
    // "labix.org/v2/mgo/bson"
)

const (
    disconnected  state = "disconnected"
    connected     state = "connected"
    connecting    state = "connecting"
    disconnecting state = "disconnecting"
    uninitialized state = "uninitialized"
)

var STATES map[interface{}]interface{}

type state string
type Session *mgo.Session

// type Collection *mgo.Collection

func init() {

    STATES = make(map[interface{}]interface{})

    STATES[0] = disconnected
    STATES[1] = connected
    STATES[2] = connecting
    STATES[3] = disconnecting
    STATES[99] = uninitialized

    STATES[disconnected] = 0
    STATES[connected] = 1
    STATES[connecting] = 2
    STATES[disconnecting] = 3
    STATES[uninitialized] = 99
}

type Mogo struct {
    connections []*mgo.Session
    models      map[string]interface{}
}

func Mogon() *Mogo {
    return &Mogo{}
}

// Connect to database
func (m *Mogo) Connect(url string) Session {

    session, err := mgo.Dial(url)

    if err != nil {
        panic(err)
    }

    // defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    m.connections = append(m.connections, session)

    return session
}

func (m *Mogo) connection() *mgo.Session {
    return m.connections[0]
}

func (m *Mogo) Model(name string, model interface{}) *mgo.Collection {
    if m.models == nil {
        m.models = make(map[string]interface{})
    }

    m.models[name] = model

    c := m.connection().DB("").C(name)

    return c
}
