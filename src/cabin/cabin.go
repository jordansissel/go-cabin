package cabin

import (
  "time"
  "fmt"
  "os"
)

/* Make a new struct, but for now it has no members.
 * Later I will add channels and such for subscription stuff
 */
type Cabin struct {
  /* Channels receive Event structs */
  channels map[Subscription] EventChannel
  notify_closed chan Subscription
}

/* A channel which emits a sequence of identifiers */
type Subscription int
var idGeneratorChannel = make(chan Subscription, 1)

func init() {
  /* Start the 'id generator' process */
  go generateIds(idGeneratorChannel)
}

func generateIds(channel chan Subscription) {
  var i Subscription = 0
  for ;; i++ {
    channel <- i
  }
} /* generateIds */

func nextId() Subscription {
  return <- idGeneratorChannel
}

type EventChannel chan *Event

/* A cabin event. Simply a timestamp + an object */
type Event struct {
  Timestamp time.Time
  Message string
  Object interface{}
}

type ReceiverFunc func(EventChannel)

type Receiver interface {
  Run(EventChannel) error
}

func (cabin *Cabin) Initialize() {
  if cabin.channels == nil {
    cabin.channels = make(map[Subscription] EventChannel)
  }
  if cabin.notify_closed == nil {
    cabin.notify_closed = make(chan Subscription)
  }
}

func (cabin *Cabin) SubscribeFunc(receiver ReceiverFunc) Subscription {
  cabin.Initialize()
  channel := make(EventChannel)
  id := nextId()
  go cabin.activateSubscription(receiver, id, channel)
  return id
} /* Cabin.SubscribeFunc */

func (cabin *Cabin) Subscribe(receiver Receiver) Subscription {
  return cabin.SubscribeFunc(func(channel EventChannel) {
    err := receiver.Run(channel)
    /* TODO(sissel): We should really have an internal cabin that points at
    * stderr */
    fmt.Fprintf(os.Stderr, "Subscriber(%#v) failed: %s\n", receiver, err)
  })
} /* Cabin.Subscribe */

func (cabin *Cabin) activateSubscription(receiver ReceiverFunc, id Subscription,
                                         channel EventChannel) {
  cabin.channels[id] = channel
  defer func() { 
    cabin.notify_closed <- id 
    delete(cabin.channels, id)
  }()

  receiver(channel)
} /* Cabin.activateSubscription */

/* Log an object */
func (cabin *Cabin) Log(object interface{}) {
  event := &Event{Timestamp: time.Now().UTC(), Object: object}

  for _, channel := range cabin.channels {
    channel <- event
  }
} /* Cabin.Log */

/* Formatted logging */
func (cabin *Cabin) Logf(format string, args...interface{}) {
  message := fmt.Sprintf(format, args...)
  event := &Event{Timestamp: time.Now().UTC(), Object: message}

  for _, channel := range cabin.channels {
    channel <- event
  }
}

/* Close. This will block until all subscribers have completed */
func (cabin *Cabin) Close() {
  count := 0
  for _, channel := range(cabin.channels) {
    close(channel)
    count++
  }

  /* Wait for all the channels to close */
  for ; count > 0 ; count-- {
    <- cabin.notify_closed
  }
}
