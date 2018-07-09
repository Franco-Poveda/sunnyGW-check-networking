pipeline {
  agent any
  stages {
    stage('install') {
      steps {
        sh 'pwd'
        script {
          def root = tool name: 'go10.3', type: 'go'
          withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"


            stage 'preTest'
            sh 'go version'
            sh 'go get github.com/fatih/color'
            sh 'go get github.com/reiver/go-telnet'


          }
        }

      }
    }
    stage('build') {
      steps {
        sh './xcompile.sh'
      }
    }
    stage('deliver') {
      steps {
        build 'deploy1'
      }
    }
  }
}