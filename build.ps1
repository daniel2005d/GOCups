param (
    [ValidateSet("windows", "linux", "darwin", "all", "clean")]
    [string]$Target = "windows"
)

$BinaryName = "gocups"
$BuildDir = "bin"

if ($Target -eq "clean") {
    if (Test-Path $BuildDir) {
        Remove-Item -Recurse -Force $BuildDir
        Write-Host "Build directory cleaned." -ForegroundColor Yellow
    }
    exit 0
}

if (-not (Test-Path $BuildDir)) {
    New-Item -ItemType Directory -Path $BuildDir | Out-Null
}

function Build-Target([string]$Os, [string]$Arch, [string]$OutputName) {
    Write-Host "Building for $Os/$Arch -> $OutputName" -ForegroundColor Cyan

    $env:CGO_ENABLED = "0"
    $env:GOOS = $Os
    $env:GOARCH = $Arch

    $outputPath = Join-Path $BuildDir $OutputName
    $buildArgs = @(
        "build",
        "-ldflags=-s -w",
        "-trimpath",
        "-o", $outputPath,
        "."
    )

    & go @buildArgs

    if ($LASTEXITCODE -ne 0) {
        Write-Error "Build failed for $Os/$Arch with exit code $LASTEXITCODE"
        exit $LASTEXITCODE
    }
}

switch ($Target) {
    "windows" {
        Build-Target "windows" "amd64" "$BinaryName.exe"
    }
    "linux" {
        Build-Target "linux" "amd64" "$BinaryName-linux-amd64"
    }
    "darwin" {
        Build-Target "darwin" "arm64" "$BinaryName-darwin-arm64"
    }
    "all" {
        Build-Target "windows" "amd64" "$BinaryName.exe"
        Build-Target "linux" "amd64" "$BinaryName-linux-amd64"
        Build-Target "darwin" "arm64" "$BinaryName-darwin-arm64"
    }
}

Write-Host "Build completed successfully." -ForegroundColor Green