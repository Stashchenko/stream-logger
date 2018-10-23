### Create image from docker file
    docker build -t "simplelog:7" .
#### Add tag 
    docker tag "simplelog:7" gcr.io/nimbussheridan/simplelog:7
#### Push 
    docker push gcr.io/nimbussheridan/simplelog:7
#### Kubernetis apply
    kubectl apply -f deployment.yaml -n services
