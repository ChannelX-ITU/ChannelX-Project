pipeline {
    agent any

    stages {
        stage('Deploy') {
            steps {
                bash 'docker-compose stop'
                bash 'yes y | docker-compose rm'
                bash 'docker-compose -f docker-compose.release.yml build'
                bash 'docker-compose -f docker-compose.release.yml -d up'
            }
        }
    }
}
