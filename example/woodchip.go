package main

import (
  "io"
  "bufio"
  "os"
  "flag"
  "strconv"
  "net" /* for SplitHostPort */
  "cabin" // "github.com/jordansissel/go-cabin/src/cabin"
  //"fmt"
)

var syslog_host = flag.String("syslog-host", "", "the host:port to send events to")

/**
 * Take stdin as a cabin input 
 */
func main() {
  logger := &cabin.Cabin{}
  stdin := bufio.NewReaderSize(os.Stdin, 16834)

  /* TODO(sissel): Make this tunable */
  logger.Subscribe(cabin.Stdout{})

  flag.Parse()

  /* TODO(sissel): Improve the API for this */
  if len(*syslog_host) > 0 {
    logger.Logf("Subscribing to syslog %s", *syslog_host)
    host, port_str, _ := net.SplitHostPort(*syslog_host)
    port, _ := strconv.ParseUint(port_str, 10, 16)
    syslog := cabin.Syslog{Host: host, Port: uint(port)}
    logger.Subscribe(syslog)
  }

  for {
    /* Read lines from stdin */
    line, err := stdin.ReadString('\n')

    if err != nil {
      exitCode := 0
      switch err {
        case io.EOF: /* Expected if stdin closes, don't complain. */
        default:
          exitCode = 1
          logger.Logf("Unexpected error: %v", err)
      }

      /* On any error, close up the logger and shutdown */
      logger.Close()

      /* This will terminate us. */
      os.Exit(exitCode)
    } /* end error checking */

    /* log the line received */
    logger.Log(line[0:len(line) - 1])
  } /* loop forever */

  panic("should not get here")
} /* main */
