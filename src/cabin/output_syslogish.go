package cabin

import (
  "net"
  "fmt"
  "os"
)

type Syslog struct {
  Host string
  Port uint
}

func (syslog Syslog) Run(channel EventChannel) error {
  conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", syslog.Host, syslog.Port))
  if err != nil {
    return err
  }

  hostname, _ := os.Hostname()

  for event := range(channel) {
    /* TODO(sissel): Check for errors */
    fmt.Fprintf(conn, "%s %s: %v\n",
                formatTimestamp(event.Timestamp),
                hostname, event.Object)
  }

  return nil /* No error, channel closed */
}
