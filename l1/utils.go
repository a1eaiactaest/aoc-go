package main

import (
  "os"
  "bufio"
  "io/ioutil"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}


/* *** FILES *** */

func file_to_lines(path string) (lines []string) {
  const MAX_CAPACITY = 1000000000 // max line length

  f, err := os.Open(path)
  check(err)
  defer f.Close() // defer task

  scanner := bufio.NewScanner(f) 
  buf := make([]byte, MAX_CAPACITY)
  scanner.Buffer(buf, MAX_CAPACITY)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  err = scanner.Err()
  check(err)

  return lines
}

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
  i, err := strconv.Atoi(s)
  check(err)
  return i
}

