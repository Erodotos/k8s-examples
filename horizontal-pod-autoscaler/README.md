# Exploring Horizontal Pod Autoscaler v2

## Motivation
Kubernetes Horizontal Pod Autoscaler (HPA) is a powerful K8s feature that automatically revises the number of pods in a Deployment based on resource utilization, such as CPU, memory usage, or other custom metrics.

The HPA continuously monitors resource utilization and compares it to a user-defined target. When resource usage rises above or falls below this target, the HPA automatically scales the number of pods up or down to match the demand. This dynamic scaling ensures that resources are efficiently allocated during periods of high demand and conserved when usage is lower.

In this tutorial, we will focus on the more recent HPA v2, which offers a few more features compared to v1. It allows scaling based on multiple metrics and more sophisticated scaling policies, making it a handy resource management tool.

## Demo

1. Deploy a sample web server with four endpoints.

`GET | /cpuLoad | Executes CPU intensive task`

`GET | /memoryLoad | Cuptures memory`

`GET | /healthz | Implements endpoint which k8s uses to check the app's availability `

`GET | /readyz | Implements endpoint which k8s uses to check the app's readiness`

Execute the following command to deploy sample web server.

```shell
$ kubectl apply -f https://raw.githubusercontent.com/Erodotos/k8s-examples/master/horizontal-pod-autoscaler/manifests/deployment.yaml
```

2. CPU Based Scaling

2. Memory Based Scaling

3. ScaleUp Behaviour

4. ScaleDown Behaviour

## Key Takeways

## References
- Kubernetes Documentation
- Kubernetes API Reference
- GoFiber Documentation
- https://github.com/bnkamalesh/golang-dockerfile Docker file
