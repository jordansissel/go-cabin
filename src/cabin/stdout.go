package cabin

import (
  "os"
)

type Stdout struct {
}

func (stdout Stdout) Run(channel EventChannel) error {
  for event := range(channel) {
    emit(os.Stdout, event)
  }

  return nil
}

