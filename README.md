## helm kill plugin

### Description

Simple helm plugin for deleting specified chart versions from an http chart repo where curl would otherwise be used.

This plugin makes deleting easier by keeping the delete/kill command in the helm tool.

(why not 'helm delete'? Even though 'delete' does not appear in the helm help list, it is set as an alias for uninstall)

### Installing

To install, just run:

```
helm plugin install 
```

### Usage

You must have the variable `USER=<your username>` set in your environment (this is set by default on most systems) for basic auth

To delete a specific chart version:

```
helm kill <chart name> <chart version>
```

You can see this same command from helm:

```
helm kill --help
```

#### TODO

Right now, the repo URL is hard-coded into the go binary. This could easily be updated to use an environment variable or to read from the helm config file

