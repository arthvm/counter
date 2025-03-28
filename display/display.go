package display

type Options struct {
	args NewOptionsArgs
}

type NewOptionsArgs struct {
	ShowBytes  bool
	ShowWords  bool
	ShowLines  bool
	ShowHeader bool
}

func NewOptions(args NewOptionsArgs) Options {
	return Options{
		args: args,
	}
}

func (d Options) ShouldShowBytes() bool {
	if !d.args.ShowBytes && !d.args.ShowWords && !d.args.ShowLines {
		return true
	}

	return d.args.ShowBytes
}

func (d Options) ShouldShowWords() bool {
	if !d.args.ShowBytes && !d.args.ShowWords && !d.args.ShowLines {
		return true
	}

	return d.args.ShowWords
}

func (d Options) ShouldShowLines() bool {
	if !d.args.ShowBytes && !d.args.ShowWords && !d.args.ShowLines {
		return true
	}

	return d.args.ShowLines
}

func (d Options) ShouldShowHeader() bool {
	return d.args.ShowHeader
}
