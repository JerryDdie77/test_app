pipeline {
    agent any
    
    parameters {
        choice(name: 'ENV', choices: ['dev', 'prod'], description: 'Select deployment environment')
    }
    
    options {
        cleanWs()
    }
    
    stages {
        stage('Print Environment') {
            steps {
                echo "Deploying to ${params.ENV}"
            }
        }
        
        stage('Copy Files via SSH') {
            steps {
                sshPublisher(
                    publishers: [
                        sshPublisherDesc(
                            configName: 'worker',
                            transfers: [
                                sshTransfer(
                                    sourceFiles: '**/*',
                                    remoteDirectory: "${params.ENV}",
                                    execCommand: 'echo "Files copied to ${params.ENV}"'
                                )
                            ],
                            dryRun: false
                        )
                    ]
                )
            }
        }
    }
}