pipeline {
    agent {
        label 'agent-1'
    }
    tools {
        go 'Go-1.24-local'
    }
    options {
        ansiColor('xterm')
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Checking out Go project from Git...'
                checkout scm
                echo 'Workspace contents after checkout:'
                sh 'ls -la'
            }
        }

        stage('Setup Go environment'){
            steps {
                echo 'Verify Go instlallation on agent'
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                echo 'Building Go application...'
                sh 'go build .'
            }
        }

        stage('Test') {
            steps {
                echo 'Running Go tests...'
                sh 'go test -v ./...'
            }
        }
    }

    post {
        always {
            echo 'Pipeline finished.'
            cleanWs()
        }
        success {
            echo 'Go project pipeline successfull'
        }
        failure {
            echo 'Go project pipeline failed'
        }
    }
}
