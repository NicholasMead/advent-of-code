package common

import "flag"

type Flags struct {
	File    *string
	Verbose *bool
	Args    []string
}

var cache *Flags

func GetFlags() Flags {
	if cache == nil {
		cache = &Flags{}
		cache.File = flag.String("f", "./input.txt", "input file")
		cache.Verbose = flag.Bool("v", false, "verbose")

		flag.Parse()
		cache.Args = flag.Args()
	}

	return *cache
}
