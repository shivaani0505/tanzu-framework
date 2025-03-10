## tanzu cluster machinehealthcheck node delete

Delete a MachineHealthCheck object

### Synopsis

Delete a MachineHealthCheck object of the nodes of a cluster

```
tanzu cluster machinehealthcheck node delete CLUSTER_NAME [flags]
```

### Options

```
  -h, --help               help for delete
  -m, --mhc-name string    Name of the MachineHealthCheck object
  -n, --namespace string   The namespace where the MachineHealthCheck object was created, default to the cluster's namespace
  -y, --yes                Delete the MachineHealthCheck object without asking for confirmation
```

### Options inherited from parent commands

```
      --log-file string   Log file path
  -v, --verbose int32     Number for the log level verbosity(0-9)
```

### SEE ALSO

* [tanzu cluster machinehealthcheck node](tanzu_cluster_machinehealthcheck_node.md)	 - MachineHealthCheck operations for the nodes of a cluster

###### Auto generated by spf13/cobra on 14-Sep-2022
