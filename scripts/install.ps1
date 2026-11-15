param([Parameter(Mandatory=$true)][string]$Version)
$ErrorActionPreference = "Stop"
if ($Version -notmatch '^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$') { throw "Use version v0.1.0-rc.1" }
$arch = [System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture.ToString()
switch ($arch) { "X64" { $target="amd64" } "Arm64" { $target="arm64" } default { throw "Unsupported CPU: $arch" } }
$asset = "factory_${Version}_windows_${target}.zip"
$base = "https://github.com/atstaeff/ai-factory-cli/releases/download/$Version"
$tmp = Join-Path ([IO.Path]::GetTempPath()) ([guid]::NewGuid().ToString())
New-Item -ItemType Directory -Path $tmp | Out-Null
try {
  $archive = Join-Path $tmp $asset
  $sums = Join-Path $tmp "SHA256SUMS"
  Invoke-WebRequest -Uri "$base/$asset" -OutFile $archive
  Invoke-WebRequest -Uri "$base/SHA256SUMS" -OutFile $sums
  $match = Get-Content $sums | Where-Object { $_ -match ("^[0-9a-fA-F]{64}\s+\*?" + [regex]::Escape($asset) + "$") } | Select-Object -First 1
  if (-not $match) { throw "Missing checksum" }
  $expected = ($match -split "\s+")[0].ToLowerInvariant()
  $actual = (Get-FileHash -Path $archive -Algorithm SHA256).Hash.ToLowerInvariant()
  if ($expected -ne $actual) { throw "Checksum mismatch" }
  $unpack = Join-Path $tmp "unpacked"
  Expand-Archive -Path $archive -DestinationPath $unpack
  $dest = Join-Path $HOME ".local\bin"
  New-Item -ItemType Directory -Path $dest -Force | Out-Null
  Copy-Item (Join-Path $unpack "factory.exe") (Join-Path $dest "factory.exe") -Force
  $userPath = [Environment]::GetEnvironmentVariable("Path","User")
  if (($userPath -split ";") -notcontains $dest) { [Environment]::SetEnvironmentVariable("Path", "$userPath;$dest", "User") }
  if (($env:Path -split ";") -notcontains $dest) { $env:Path += ";$dest" }
  & (Join-Path $dest "factory.exe") version
  Write-Host "Installed: factory scan ."
} finally { Remove-Item -LiteralPath $tmp -Recurse -Force -ErrorAction SilentlyContinue }
