package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	sep := flag.String("sep", "-", "seperator used to join branch name parts")
	namespace := flag.String("namespace", "", "prefixed to the branch name to namespace it")
	namespaceSep := flag.String("nsSep", ":", "seperator used to join the namespace and the branch name")

	flag.Parse()

	fmt.Println(branchify(flag.Args(), *sep, *namespace, *namespaceSep))
}

func branchify(args []string, sep, namespace, namespaceSep string) string {
	slugified := strings.ToLower(strings.ReplaceAll(strings.Join(flag.Args(), sep), " ", sep))
	if namespace == "" {
		return slugified
	}
	return namespace + namespaceSep + slugified
}
