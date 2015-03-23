package http

import (
  "net/http"
)

type Entity interface {
  UnmarshalHTTP(*http.Request) error
}
