Jenkins là một công cụ tự động hóa

Jenkins hoạt động dựa trên một kiến trúc phân tán với 2 thành phần chính:
1. Jenkins master (controller)
  - là phần trung tâm điều khiển
  - chịu trách nhiệm nhận và quản lí các tác vụ từ người dùng
  - quản lí giao diện người dùng, xem trạng thái các công việc, cấu hình hệ thống, plugin
  - tích hợp với các hệ thống khác như git, docker,...
  - điều phối các agent thực thi công việc và giám sát quá trình đó

2. Jenkins agent
  - là các phiên bản jenkins độc lập thực thi các tác vụ được giao từ jenkins master
  - các agent giúp chia nhỏ công việc, giúp mở rộng quy mô và xử lí nhiều tác vụ đồng thời

=> Tóm lại
  - Controller (jenkins core + ui + plugin, ...)
  - Agent (jenkins core)

# quy trình cài đặt
- cài jenkins controller
- cài jenkins agent
- khai báo một agent trong controller, kết nối jenkins agent đến controller

# Quy trình hoạt động
- Cài đặt việc cần làm
- jenkins controller điều phối công việc: jenkins master nhận công việc và kiểm tra xem cần chạy trên agent nào
- jenkins agent thực thi công việc
- jenkins controller hiển thị kết quả

# Các cách kích hoạt build job
1. Thủ công (Manual build)
  - click 'Build now'
2. Scheduled Build (Cron job)
3. Poll SCM
  - jenkins sẽ từ poll source code, nếu có commit mới thì build
4. Webhook
  - source repo sẽ gửi thông báo đến jenkins
5. Trigger từ job khác
6. Trigger từ Pipeline Script (Jenkinsfile)
7. Trigger từ bên ngoài (API call)
  - gửi http request đến jenkins api để trigger build từ bên ngoài
8. kết hợp các loại v.v.

# Các cách cấu hình job
1. Freestyle Project (cấu hình qua UI)
2. Pipeline job (cấu hình bằng code groovy)
- Declarative pipeline (DSL groovy)
```groovy
  pipeline {
    agent any
    stages {
      stage('Build') {
        steps {
          sh 'go build .'
        }
      }
    }
  }
```
- scripted pipeline (thuần groovy)
```groovy
node {
  stage('Build') {
    sh 'go build .'
  }
}
```
3. Muiltibranch Pipeline
- dùng khi có nhiều branch và muốn jenkins tự tạo jobs cho từng branch có Jenkinsfile
- Jenkins sẽ scan repo và tự tạo các job con
- phù hợp cho miucroservices hoặc git flow với nhiều nhánh
4. Folder + Job DSL (viết code groovy DSL để sinh job)
- Dùng plugin Job DSL, viết Groovy script để generate job (chứ không phải chạy build).
- Rất mạnh cho quản lý hàng trăm job tự động.
```groovy
job('example-job') {
  steps {
    shell('echo Hello Jenkins!')
  }
}
```
5. External Configuration (JCasC, Jenkins Configuration as Code)
- dùng file YAML để cấu hình Jenkins toàn bộ
```yaml
jenkins:
  systemMessage: "Welcome to Jenkins"
  jobs:
    - script: >
        job('my-job') {
          steps {
            shell('echo Hello')
          }
        }
```

### Ví dụ về Declarative Pipeline và Scripted code
ví dụ với một job gồm các bước sau:
- Checkout code từ GitHub
- Build project Go
- Chạy test
- Lưu kết quả test
- Gửi notify nếu build thất bại
1. Declarative Pipeline
```groovy
pipeline {
  agent any

  environment {
    GO_ENV = "production"
    GITHUB_REPO = "https://github.com/example/my-go-app.git"
  }

  options {
    timestamps()
    skipDefaultCheckout()
  }

  triggers {
    pollSCM('H/5 * * * *') // check code mỗi 5 phút (nếu không dùng webhook)
  }

  stages {
    stage('Checkout') {
      steps {
        git url: "${GITHUB_REPO}", branch: 'main'
      }
    }

    stage('Build') {
      steps {
        sh 'go mod tidy'
        sh 'go build -o bin/app ./...'
      }
    }

    stage('Test') {
      steps {
        sh 'go test -v ./... | tee test-report.txt'
      }
      post {
        always {
          junit 'test-report.txt' // nếu có plugin hỗ trợ format JUnit
        }
      }
    }
  }

  post {
    failure {
      echo "Build failed! Sending notification..."
      // ví dụ: gửi Slack, email...
    }
  }
}
```

2. Scripted Pipeline
```groovy
node {
  def GITHUB_REPO = "https://github.com/example/my-go-app.git"
  def GO_ENV = "production"

  try {
    stage('Checkout') {
      git url: GITHUB_REPO, branch: 'main'
    }

    stage('Build') {
      env.GO_ENV = GO_ENV
      sh 'go mod tidy'
      sh 'go build -o bin/app ./...'
    }

    stage('Test') {
      sh 'go test -v ./... | tee test-report.txt'
      // Ghi chú: phải dùng plugin hoặc tool convert test-report.txt sang JUnit XML nếu cần publish
    }

  } catch (err) {
    echo "Build failed: ${err}"
    // Gửi notification tại đây nếu cần
    currentBuild.result = 'FAILURE'
    throw err
  }
}
```
