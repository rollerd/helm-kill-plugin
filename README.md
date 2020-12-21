## helm kill plugin

### Description

Simple helm plugin for deleting specified chart versions from an http chart repo where curl would otherwise be used.

This plugin makes deleting easier by keeping the delete/kill command in the helm tool (for the repos I use anyway)

(why not 'helm delete'? Even though 'delete' does not appear in the helm help list, it is set as an alias for uninstall)

### Installing

To install, just run:

```
helm plugin install 
```

### Usage

You must have the variables `USER=<your username>` and `HELM_HTTP_URL` set in your environment (USER is set by default on most systems)

To delete a specific chart version:

```
helm kill <chart name> <chart version>
```

You can see this same command from helm:

```
helm kill --help
```
