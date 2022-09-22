# kubels

kubels (kube ls) is a Kubernetes tool that allows you to list Kubernetes resources with a simple command(s).

## Usage

| Commands                          | Descriptions                                       |
|-----------------------------------|----------------------------------------------------|
| kubels or kubels -p               | list of pods in current namespace                  |     
| kubels -p -n {namespace}          | list of pods in specified namespace                |
| kubels -p w                       | list of pods in current namespace with watch       |
| kubels -p o                       | list of pods in current namespace with wide output |
| kubels -p m                       | list of pods in current namespace with metrics     |
| kubels -n                         | list of namespaces                                 |
| kubels -d                         | list of deployments in current namespace           |
| kubels -d -n {namespace}          | list of deployments in specified namespace         |
| kubels -s                         | list of services in current namespace              |
| kubels -s (or svc) -n {namespace} | list of services in specified namespace            |
| kubels -sec                       | list of secrets in current namespace               |
| kubels -sec -n {namespace}        | list of secrets in specified namespace             |

## Installation

### Homebrew

```bash
brew tap semihtok/kubels

brew install kubels
```