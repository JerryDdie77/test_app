pipeline {
    agent any

    parameters {
        choice(name: 'ENV', choices: ['dev', 'prod'], description: 'Среда деплоя')
    }

    stages {
        stage('Print Message') {
            steps {
                echo "Deploying to ${params.ENV}"
            }
        }

        stage('Deploy via SSH') {
            steps {
                sshPublisher(
                    publishers: [
                        sshPublisherDesc(
                            configName: 'worker',
                            transfers: [
                                sshTransfer(
                                    sourceFiles: '**/*'
                                )
                            ]
                        )
                    ]
                )
            }
        }

        stage('Clean Workspace') {
            steps {
                cleanWs()
            }
        }
    }
}