package piscine

func Rot14(s string) string {
  var k string
  
  for _, ch := range s {
    if ch >= 'A' && ch <= 'Z' {
      k += k + string('A' + (ch-'A'+14)%26)
    } else if ch >= 'a' && ch <= 'z' {
      k += string('a' + (ch-'a'+14)%26)
    } else {
      k += string(ch)
    }
  }
  return k
}

