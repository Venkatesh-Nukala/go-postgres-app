pipeline {
    agent {
        label 'Docker-Agent'
    }

    stages {

            stage('Checkout Code') {
                steps {
                    git branch: 'main', url: 'https://github.com/Venkatesh-Nukala/go-postgres-app.git'
                }
            }

            stage('Build Image') {
                steps {
                    sh 'docker compose build web'
                }
            }

            stage('Run Containers') {
                steps {
                    sh '''
                    docker compose down || true
                    docker compose up -d
                    '''
                }
            }

            stage('Verify Deployment') {
                steps {
                    sh 'docker ps'
                }
            }

            stage('Push to DockerHub') {
                steps {
                    withCredentials([usernamePassword(credentialsId: 'dockerhub', usernameVariable: 'DOCKERHUB_USERNAME', passwordVariable: 'DOCKERHUB_PASSWORD')]) {
                        sh '''
                        echo $DOCKERHUB_PASSWORD | docker login -u "$DOCKERHUB_USERNAME" --password-stdin
                        docker tag go-postgres-app_web $DOCKERHUB_USERNAME/go-postgres-app:latest
                        docker push $DOCKERHUB_USERNAME/go-postgres-app:latest
                        '''
                    }
                }
            }
            
    }
}
