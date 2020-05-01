package main

import (
  "flag"
  "fmt"
  "strings"
)

func main() {
  namespace := flag.String("namespace", "", "prefixed to the branch name to namespace it")
  sep := flag.String("sep", "-", "seperator used to join branch name parts")
  namespaceSep := flag.String("nsSep", ":", "seperator used to join the namespace and the branch name")

  flag.Parse()

  slugified := strings.ToLower(strings.ReplaceAll(strings.Join(flag.Args(), *sep), " ", *sep))
  if *namespace != "" {
    fmt.Println(*namespace + *namespaceSep + slugified)
  } else {
    fmt.Println(slugified)
  }
}
