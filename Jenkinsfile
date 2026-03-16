pipeline {
    agent any

    stages {

        stage('Checkout Code') {
            steps {
                git branch: 'main', url: 'https://github.com/Venkatesh-Nukala/go-postgres-app.git'
            }
        }

        stage('Build Containers') {
            steps {
                sh 'docker compose build'
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

    }
}
