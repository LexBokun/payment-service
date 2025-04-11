# Загрузим .env, если он есть
if (Test-Path ".env") {
    Get-Content .env | ForEach-Object {
        if ($_ -match "^\s*([^#][^=]+)=(.+)$") {
            $key = $matches[1].Trim()
            $value = $matches[2].Trim()
            [System.Environment]::SetEnvironmentVariable($key, $value)
        }
    }
}

# Путь до бинарников
$env:GOBIN = "C:/Users/user/go/bin"
$env:GOPATH = "C:/Users/user/go"


# Запуск buf с токеном
docker run --rm `
  -v "${PWD}:/workspace" `
  -w /workspace `
  -e BUF_TOKEN=$Env:BUF_TOKEN `
  -e GOBIN=$Env:GOBIN `
  -e GOPATH=$Env:GOPATH `
  bufbuild/buf:latest @Args