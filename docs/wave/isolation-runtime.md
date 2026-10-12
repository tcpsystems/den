# Isolation runtime

## Question

How should den isolate a seat with Podman or smolvm, versus Firecracker, Docker, and Lambda? One volume per instance, no shared writable volume. The path jail binds to the packet root, not a process-global cwd.

## Sources

- Podman `--volume` chown and `:U`: https://docs.podman.io/en/v6.0.2/markdown/podman-run.1.html
- Rootless `podman volume mount` (file volumes, or `podman unshare`): https://docs.podman.io/en/v5.4.0/markdown/podman-volume-mount.1.html
- Rootless UID map, `:U`, named volumes under `~/.local/share/containers/storage/volumes/`: https://oneuptime.com/blog/post/2026-03-17-configure-volume-options-rootless-podman/view
- Firecracker design (KVM microVM, jailer, VirtIO block, no guest GPU): https://github.com/firecracker-microvm/firecracker/blob/main/docs/design.md
- Settled earlier: llama-server stays on the host kernel; a KVM guest does not contain an amdgpu fault; the smolvm virtio GPU window is 4096 MiB. Local seat shape: `sections/workers/dev-env.md`.

## Finding

Podman rootless is a user-namespace container on the host kernel. A named volume's first mount is chowned to the container process user, mapped through that namespace, and only while the volume is still empty and not an external driver. `:U` walks the source and chowns every inode to the container UID/GID; Podman warns that this rewrites the host tree and delays start. `idmap` on a volume is rootful only. Rootless `podman volume mount` mounts file volumes; other types need `podman unshare`. A seat volume should be created for that instance, never bind-mounted from another seat's tree, and never given `:z` (shared label). `:Z` is a private label, not a substitute for a separate volume.

smolvm is the seat today: its own kernel via KVM, argv-only exec, network only when `--net` is set. The writable disk is the guest disk, one machine, not a host directory shared across guests. Firecracker is the same class of boundary with a smaller device model: VirtIO net and VirtIO block, a serial console, and a jailer around the VMM. It does not filter guest egress; the host must. It is what Lambda and Fargate run underneath, not a writable volume API.

Docker adds a daemon and a shared writable graph. Lambda is someone else's microVM with no local amdgpu and no packet-root jail. Neither is the seat.

GPU and the path jail stay outside the guest. llama-server stays on the host kernel. A KVM guest does not contain an amdgpu fault. The smolvm virtio GPU window is 4096 MiB, so the guest is not where the embedder runs. The path jail is a prefix on the packet root for that call, not a process-global cwd shared by the next packet.

## What den should not copy

- Podman `:U` on a host path, or a named volume reused by a second seat.
- `:z` / a shared writable volume, Docker's graph, or Lambda's account filesystem.
- Firecracker's "host filters all egress" left undone, or its missing GPU as a place to put llama-server.
- A process-global cwd as the jail. The jail is the packet root.

## One plan line

One smolvm machine (or one rootless Podman container) per seat, one empty writable volume created for that instance, path jail set to the packet root; llama-server stays on the host.

## Unproven

No measurement here that a rootless Podman volume plus a packet-root jail matches smolvm's KVM boundary on this host, or that the 4096 MiB virtio GPU window can be raised without moving llama-server.
