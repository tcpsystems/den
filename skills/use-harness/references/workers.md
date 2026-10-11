# Workers

Read this only when the task runs inside a worker.

The dev environment is a smolvm microVM on the network named `harness`. Guests can reach each other and the mesh. Packages and files stay in the guest.

```
smolvm machine exec --name <seat> -- <argv>
```

Argv only. No shell. Create a seat with `smolvm machine create --net --network harness --name <seat> --image alpine`.

A fresh Debian host runs `deploy/debian-fresh.sh`, then logs in again so the kvm group applies. Nix is the same package set when Nix is already on the machine. It is not a second installer.
