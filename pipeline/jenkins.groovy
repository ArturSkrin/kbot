pipeline {
    agent any

    parameters {
        choice(
            name: 'OS',
            choices: ['linux', 'darwin', 'windows'],
            description: 'Target operating system'
        )
        choice(
            name: 'ARCH',
            choices: ['amd64', 'arm64'],
            description: 'Target architecture'
        )
        booleanParam(
            name: 'SKIP_TESTS',
            defaultValue: false,
            description: 'Skip running tests'
        )
        booleanParam(
            name: 'SKIP_LINT',
            defaultValue: false,
            description: 'Skip running linter'
        )
    }

    environment {
        REGISTRY   = 'ghcr.io'
        REPOSITORY = 'arturskrin/kbot'
    }

    stages {
        stage('clone') {
            steps {
                git branch: 'develop', url: 'https://github.com/ArturSkrin/kbot.git'
            }
        }

        stage('lint') {
            when { expression { return !params.SKIP_LINT } }
            steps {
                sh 'make lint'
            }
        }

        stage('test') {
            when { expression { return !params.SKIP_TESTS } }
            steps {
                sh 'make test'
            }
        }

        stage('build') {
            steps {
                sh "make build TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
            }
        }

        stage('image') {
            steps {
                sh "make image TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
            }
        }

        stage('push') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'ghcr-credentials',
                    usernameVariable: 'USER',
                    passwordVariable: 'PASSWORD'
                )]) {
                    sh 'echo $PASSWORD | docker login ghcr.io -u $USER --password-stdin'
                    sh "make push TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
                }
            }
        }
    }
}