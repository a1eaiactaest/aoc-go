package main

import (
  "fmt"
  "os"
  "ioutil"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}


/* *** FILES *** */

func file_to_string(fn string) string {
  bytes, err := ioutil.ReadFile(fn)
  check(err)
  return string(bytes)
}

func file_to_bytes(fn string) []byte {
  bytes, err := ioutil.ReadFile(fn)
  check(err)
  return bytes
}

/* *** STRINGS *** */

func atoi(s string) int {
  i, err := strnconv.Atoi(s)
  check(err)
  return i
}

