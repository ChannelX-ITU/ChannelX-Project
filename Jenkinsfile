pipeline {
    agent any

    stages {
        stage('Clean') {
            steps {
                sh 'sudo docker-compose -f docker-compose.release.yml stop'
                sh 'yes y | sudo docker-compose rm'
            }
        }
        stage('Build') {
            steps {
                sh 'sudo docker-compose -f docker-compose.release.yml build'
            }
        }
        stage('Run') {
            steps {
                sh 'sudo docker-compose -f docker-compose.release.yml up -d'
            }
        }
    }
}
