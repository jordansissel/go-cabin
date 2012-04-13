package main

import (
  "cabin"
)

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

  example := Example{Code: 42, Message: "The answer."}
  logger.Log(example)
  logger.Log(example.Code)
}
