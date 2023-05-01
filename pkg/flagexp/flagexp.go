package flagexp

import "flag"

func Subset(super *flag.FlagSet, prefix string, separator string, setup func(sub *flag.FlagSet)) error {
	sub := flag.NewFlagSet(prefix, 0)
	setup(sub)
	sub.VisitAll(func(f *flag.Flag) {
		name := prefix + ":" + f.Name
		super.Var(f.Value, name, f.Usage)
	})

	return nil
}
