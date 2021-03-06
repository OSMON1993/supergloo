---
title: "supergloo set mesh rootcert"
weight: 5
---
## supergloo set mesh rootcert

set the root certificate used to provision client and server certificates for a mesh

### Synopsis

Updates the target mesh to use the provided root certificate. Root certificate must be stored 
as a TLS secret created with `supergloo create secret tls`. 
used to provision client and server certificates for a mesh

```
supergloo set mesh rootcert [flags]
```

### Options

```
  -h, --help                           help for rootcert
      --target-mesh ResourceRefValue   resource reference the mesh for which you wish to set the root cert. format must be <NAMESPACE>.<NAME> (default { })
      --tls-secret ResourceRefValue    resource reference the TLS Secret (created with supergloo CLI) which you wish to use as the custom root cert for the mesh. if empty, the any existing custom root cert will be removed from this mesh. format must be <NAMESPACE>.<NAME> (default { })
```

### SEE ALSO

* [supergloo set mesh](../supergloo_set_mesh)	 - subcommands set a configuration parameters for one or more meshes

