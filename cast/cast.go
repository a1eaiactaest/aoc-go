package cast

import (
	"fmt"
	"strconv"
)

// supported types: string
func ToInt(arg interface{}) int {
  var val int
  switch arg.(type) {
    case string:
			var err error
      val, err = strconv.Atoi(arg.(string))
			if err != nil {
				panic("there was an error converting string to int " + err.Error())
			}

    default:
      panic(fmt.Sprintf("unhandled type for int casting %T", arg))
  }
  return val
}

// supported types: int, byte, rune
func ToString(arg interface{}) (str string) {
  switch arg.(type) {
    case int:
      str = strconv.Itoa(arg.(int))

    case byte:
      bytes := arg.(byte)
      str = string(rune(bytes))

    case rune:
      str = string(arg.(rune))

    default:
      panic(fmt.Sprintf("unhandled type for int casting %T", arg))
  }
  return str
}

func ascii_int_to_char(acode int) string {
  return string(rune(acode))
}