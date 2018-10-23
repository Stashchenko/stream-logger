### Create image from docker file
    docker build -t "simplelog:final" .
#### Add tag 
    docker tag "simplelog:final" gcr.io/<project_name>/simplelog:final
#### Push 
    docker push gcr.io/<project_name>/simplelog:final
#### Kubernetis apply
    kubectl apply -f deployment.yaml -n services
