## tanzu feature completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	feature completion fish | source

To load completions for every new session, execute once:

	feature completion fish > ~/.config/fish/completions/feature.fish

You will need to start a new shell for this setup to take effect.


```
tanzu feature completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [tanzu feature completion](tanzu_feature_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 14-Sep-2022
