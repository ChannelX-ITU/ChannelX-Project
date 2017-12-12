pipeline {
    agent any

    stages {
        stage('Deploy') {
            steps {
                sh 'docker-compose stop'
                sh 'yes y | docker-compose rm'
                sh 'docker-compose -f docker-compose.release.yml build'
                sh 'docker-compose -f docker-compose.release.yml -d up'
            }
        }
    }
}
