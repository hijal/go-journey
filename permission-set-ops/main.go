package main

import (
	"fmt"
	"maps"
	"slices"
)

func has(perms map[string]struct{}, p string) bool {
	_, ok := perms[p]
	return ok
}

func main() {
	adminPerms := map[string]struct{}{
		"read": {}, "write": {}, "delete": {}, "invite": {},
	}
	editorPerms := map[string]struct{}{
		"read": {}, "write": {},
	}

	fmt.Println("admin can delete?", has(adminPerms, "delete"))
	fmt.Println("editor can delete?", has(editorPerms, "delete"))

	audit := maps.Clone(adminPerms)

	maps.Copy(audit, map[string]struct{}{"billing": {}})
	fmt.Println("audit can billing ", has(audit, "billing"))

	baseline := map[string]struct{}{}

	for p := range editorPerms {
		if _, ok := adminPerms[p]; ok {
			baseline[p] = struct{}{}
		}
	}

	for _, p := range slices.Sorted(maps.Keys(baseline)) {
		fmt.Println("shared permission:", p)
	}
}
