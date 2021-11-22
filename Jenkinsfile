pipeline {
    agent any

    environment {
        SERVICE_NAME = "healthcheck"
    }

    stages {
        stage('Prepare') {
            steps {
                git credentialsId: 'github', url: "https://github.com/${ORGANIZATION_NAME}/${SERVICE_NAME}"
            }
        }
        stage('Build') {
            steps {
                echo 'Building..!'
                sh 'ls -lah'
            }
        }
        // stage('Test') {
        //     steps {
        //         echo 'Testing..!'
        //     }
        // }
        stage('Deploy') {
            steps {
                echo 'Deploying....!'
            }
        }
    }
}