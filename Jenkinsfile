pipeline {
    agent any

    stages {
        stage('Deploy') {
            steps {
                sh 'sudo docker-compose -f docker-compose.release.yml stop'
                sh 'yes y | sudo docker-compose rm'
                sh 'sudo docker-compose -f docker-compose.release.yml build'
                sh 'sudo docker-compose -f docker-compose.release.yml -d up'
            }
        }
    }
}
