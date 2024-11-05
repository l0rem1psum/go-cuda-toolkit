package cudadrv

func combineFlags[F ~uint](flags []F) (f F) {
	for _, flag := range flags {
		f |= flag
	}
	return
}
