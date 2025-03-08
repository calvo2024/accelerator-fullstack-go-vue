Accelerator creating a Go backend binary with a static Vue fronted embedded.

Dockerfile is meant to build each on its own stage and copy the resulting binary on top of an alpine layer.

---
docker run -it --mount type=bind,src=$PWD/src,dst=/opt/src golang:1.22 bash
docker run -it --mount type=bind,src=$PWD/src,dst=/opt/src node:19 bash

---
I still don't know how I'm going to work locally with Vite serving the frontend on one port and the Go API listening on a different port.