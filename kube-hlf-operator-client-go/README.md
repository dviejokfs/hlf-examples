# HLF Kubernetes Operator Helper

## Overview
This utility tool is designed to interact with the Hyperledger Fabric (HLF) network orchestrated on Kubernetes. It leverages the `hlf-operator` client to fetch details about the Fabric Peers currently deployed within the Kubernetes cluster.

## Features
- **Kubernetes Configuration**: Uses a Kubernetes configuration file to establish a connection with the cluster.
- **Hyperledger Fabric Client**: Creates a client instance for Hyperledger Fabric to interact with the network components.
- **List Fabric Peers**: Retrieves and displays a list of all Fabric peers in the network along with their respective MSP IDs.

## Requirements
- Kubernetes cluster with HLF operator installed.
- `kubectl` and context set for the Kubernetes cluster.
- Access to the Kubernetes configuration file (default location is `~/.kube/config`).

## Usage

Set up your Kubernetes configuration file path in `kubeConfig` constant. Then run the program to list all the Fabric peers:

```shell
go run main.go
```

You will receive an output of all the Fabric peers in the format:

```text
Peer: <peer-name>/<namespace> mspID: <msp-id>
```

## Installation

To use this tool, clone the repository and build the binary, or directly run the script using `go run`:

```shell
git clone <repository-url>
cd kube-hlf-operator-client-go
export KUBECONFIG=<path-to-your-kubeconfig-file>
go run main.go

```

## Configuration

The tool uses the following environment variables:

- `KUBECONFIG`: Path to your kubeconfig file. Default is set to `/Users/davidviejo/.kube/k3d-k8s-hlf` but can be changed according to your setup.

## Contributions

To contribute to this tool, please fork the repository, make your changes, and submit a pull request.

## License

This tool is open-sourced under the MIT License.

---

Note: Replace `<repository-url>` and `<repository-directory>` with the actual URL and directory name of your repository.