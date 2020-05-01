package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	sep := flag.String("sep", "-", "seperator used to join branch name parts")
	ns := flag.String("ns", "", "prefixed to the branch name to namespace it")
	nsSep := flag.String("nsSep", "@", "seperator used to join the namespace and the branch name")

	flag.Parse()

	fmt.Println(branchify(flag.Args(), *sep, *ns, *nsSep))
}

func branchify(args []string, sep, ns, nsSep string) string {
	slugified := strings.ToLower(strings.ReplaceAll(strings.Join(args, sep), " ", sep))
	if ns == "" {
		return slugified
	}
	return ns + nsSep + slugified
}
