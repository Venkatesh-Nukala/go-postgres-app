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

        stage('Run Containers') {
            steps {
                sh '''
                docker compose down || true
                docker compose up --build -d
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
                            docker login -u "$DOCKERHUB_USERNAME" -p "$DOCKERHUB_PASSWORD"
                            docker tag go-postgres-app_web $DOCKERHUB_USERNAME/go-postgres-app:latest
                            docker push $DOCKERHUB_USERNAME/go-postgres-app:latest
                            '''
                        }
                    }
            }

    }
}
