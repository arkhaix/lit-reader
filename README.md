# [lit-reader](http://reader.arkhaix.com)

## A clean interface for reading online literature from multiple sources

### Supported sites
Site | Example URL
--- | ---
`wanderinginn.com` | [wanderinginn.com](http://wanderinginn.com)
`royalroad.com` | www.royalroad.com/fiction/5701/savage-divinity/chapter/58095/chapter-1-new-beginnings
`fictionpress.com` | http://www.fictionpress.com/s/2961893/1/Mother-of-Learning
`archiveofourown.org` | [archiveofourown.org/works/11478249/chapters/25740126](http://archiveofourown.org/works/11478249/chapters/25740126)

## Deployment

### Local (docker-compose)
1. add 127.0.0.1 = reader.local to /etc/hosts
2. docker-compose -f deployments/docker-compose/reader.yaml up
3. visit http://reader.local

### Google Container Engine (kubernetes)
1. modify the host (domain) settings in deployments/kubernetes/traefik/traefik-ingress.yaml
2. reserve a static ip and modify loadBalancerIp in deployments/kubernetes/traefik/traefik.yaml
3. kubectl create clusterrolebinding $USER-cluster-admin-binding --clusterrole=cluster-admin --user=your.google.account@example.com
4. kubectl apply -f deployments/kubernetes/cockroach
5. wait for pods to reach Running state (kubectl get pods)
6. kubectl apply -f deployments/kubernetes/cockroach-init/cluster-init.yaml
7. kubectl apply -f deployments/kubernetes/cockroach-init/cockroach-init.yaml
8. kubectl apply -f deployments/kubernetes/reader
9. kubectl apply -f deployments/kubernetes/traefik

