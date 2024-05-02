# Faros.sh CLI

## Description

This is a CLI tool to interact with the faros.sh API.

## Prerequisites

- [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)


## Installation

```bash
kubectl krew index add faros https://github.com/faroshq/krew-index.git

kubectl krew install faros/faros faros/kcp-ws faros/kcp-kcp
```

## Usage

```bash
# Login to the faros.sh API
kubectl faros login

# Create a workspace
kubectl faros ws create my-workspace

# List workspaces
kubectl get ws
```
