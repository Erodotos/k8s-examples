# Exploring Horizontal Pod Autoscaler v2

## Motivation
Kubernetes Horizontal Pod Autoscaler (HPA) is a powerful K8s feature that automatically revises the number of pods in a Deployment based on resource utilization, such as CPU, memory usage, or other custom metrics.

The HPA continuously monitors resource utilization and compares it to a user-defined target. When resource usage rises above or falls below this target, the HPA automatically scales the number of pods up or down to match the demand. This dynamic scaling ensures that resources are efficiently allocated during periods of high demand and conserved when usage is lower.

In this tutorial, we will focus on the more recent HPA v2, which offers a few more features compared to v1. It allows scaling based on multiple metrics and more sophisticated scaling policies, making it a handy resource management tool.

## Demo

#### 1. Deploy a sample web server with four endpoints.

| Method | Endpoint     | Description                                               |
|--------|--------------|-----------------------------------------------------------|
| GET    | /cpuLoad     | Executes CPU intensive task                                |
| GET    | /memoryLoad  | Captures memory                                            |
| GET    | /healthz     | Implements endpoint which k8s uses to check the app's availability |
| GET    | /readyz      | Implements endpoint which k8s uses to check the app's readiness   |

Execute the following command to deploy our dummy web server.

```shell
# Deployment
$ kubectl apply -f https://raw.githubusercontent.com/Erodotos/k8s-examples/master/horizontal-pod-autoscaler/manifests/deployment.yaml

# Service
$ kubectl apply -f https://raw.githubusercontent.com/Erodotos/k8s-examples/master/horizontal-pod-autoscaler/manifests/service.yaml
```

#### 2. CPU Based Scaling

This HPA is designed to automatically scale the number of pods between 1 and 10 based on CPU utilization. The target average utilization is set to 50%, meaning that if the average CPU usage exceeds this threshold, additional pods will be deployed until the utilization returns to the desired level.

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: web-app
spec:
  scaleTargetRef:
    apiVersion: apps/vl
    kind: Deployment
    name: web-app
  minReplicas: 1
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          # the following is a percentage (%)
          type: Utilization
          averageUtilization: 50
```

#### 3. Memory Based Scaling


Let's add a second metric to our previous configuration to consider both CPU and Memory utilization. We'll apply the same scaling logic to memory, setting the target average utilization at 80%.

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: web-app
spec:
  scaleTargetRef:
    apiVersion: apps/vl
    kind: Deployment
    name: web-app
  minReplicas: 1
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          # the following is a percentage (%)
          type: Utilization
          averageUtilization: 50
    - type: Resource
      resource:
        name: memory
        target:
          # the following is a percentage (%)
          type: Utilization
          averageUtilization: 80
```

#### 4. Scaling Behaviour

Now, someone might want to fine-tune the scaling behavior. The Scale Up behavior controls how many pods are added and the waiting period before triggering another Scale Up event.

In more detail, we can control both Upscaling and Downscaling. In the example below, during Upscaling, 2 pods are added within 5 seconds, followed by a 10-second wait before the next Scale Up. Downscaling is handled similarly, but only 1 pod is removed at a time, with a stabilization window set to 20 seconds.

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: web-app
spec:
  scaleTargetRef:
    apiVersion: apps/vl
    kind: Deployment
    name: web-app
  minReplicas: 1
  maxReplicas: 10
  metrics:
    - type: Resource
      resource:
        name: memory
        target:
          # the following is a percentage (%)
          type: Utilization
          averageUtilization: 90
    - type: Resource
      resource:
        name: cpu
        target:
          # the following is a percentage (%)
          type: Utilization
          averageUtilization: 50
  behavior:
    scaleUp:
      stabilizationWindowSeconds: 10
      policies:
        - type: Pod
          value: 2
          periodSeconds: 5
    scaleDown:
      stabilizationWindowSeconds: 20
      policies:
        - type: Pod
          value: 1
          periodSeconds: 5

```

#### 4. Action Time (Demo)

Deploy the HPA we created in the previous steps.

```shell
# HPA
$ kubectl apply -f https://raw.githubusercontent.com/Erodotos/k8s-examples/master/horizontal-pod-autoscaler/manifests/hpa.yaml
```

To test our HPA, we need a load generator, which can be set up using a helper pod running BusyBox. Execute the follopwing command:

```shell
$ kubectl run -it --rm load-generator --image=busybox /bin/sh
```

In the Pod's terminal run the following command to test CPU:

```shell
$ while true; do wget -q -O- http://web-app:8080/cpuLoad; done
```

On an other terminal whatch the average utilisation and the Pods scaling up.

```shell
$ kubectl get hpa

NAME         REFERENCE                     TARGET    MINPODS   MAXPODS   REPLICAS
web-app      Deployment/web-app/scale      0% / 50%  1         10        1
```

Stop the load generator using CTRL+C and observe the Pods gradually terminating.

```shell
$ kubectl get pods
```

Repeat the same for testing memory metric. Just replace `http://web-app:8080/cpuLoad` with `http://web-app:8080/memoryLoad` in the workload generating command.
