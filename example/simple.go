package main

import (
  "cabin"
)

/* A simple struct. Below I'll show how to use it in your logs */
type Example struct {
  Code int
  Message string
}

func main() {
  logger := cabin.Cabin{}
  c := make(chan *cabin.Event)
  go cabin.StdoutLogger(c)
  logger.Subscribe(c)

  logger.Log("Hello world")
  logger.Log(42)

  example := Example{Code: 42, Message: "The answer."}
  logger.Log(example)
}
