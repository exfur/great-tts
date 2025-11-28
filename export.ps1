# ==========================================
# Export-Code.ps1
# Combines project source files into one text file for LLM sharing.
# Excludes CSVs, Binaries, and Git data.
# ==========================================

$OutputFile = "Full_Project_Source.txt"
$CurrentLocation = Get-Location

# 1. Extensions to Include (Whitelist approach is cleaner for LLMs)
# We include Go files, Module definitions, Markdown docs, and PowerShell scripts.
$IncludeExtensions = @(".go", ".mod", ".md", ".ps1", ".json")

# 2. Specific folders to ignore completely
$IgnoreFolders = @(".git", ".vs", "bin", "assets", "data") 

# Clear previous export if it exists
if (Test-Path $OutputFile) {
    Remove-Item $OutputFile
}

Write-Host "Scanning for source files..." -ForegroundColor Cyan

# Get all files recursively
$Files = Get-ChildItem -Path $CurrentLocation -Recurse -File | Where-Object {
    $file = $_
    
    # Check if file is inside an ignored folder
    $parentPath = $file.DirectoryName
    $isIgnoredFolder = $false
    foreach ($ignore in $IgnoreFolders) {
        if ($parentPath -match "\\$ignore" -or $parentPath -match "/$ignore") {
            $isIgnoredFolder = $true
            break
        }
    }

    # Filter Logic:
    # 1. Must NOT be in ignored folder
    # 2. Extension must be in whitelist OR NOT be .csv (Double safety)
    # 3. Must not be the output file itself
    if ($isIgnoredFolder) { return $false }
    if ($file.Name -eq $OutputFile) { return $false }
    if ($file.Extension -eq ".csv") { return $false } # Explicitly excluding CSV
    
    # Only allow specific source extensions
    if ($IncludeExtensions -contains $file.Extension) { return $true }
    
    return $false
}

# Create the output file
New-Item -Path $OutputFile -ItemType File -Force | Out-Null

# Loop through files and append content
foreach ($file in $Files) {
    # Get relative path for cleaner context (e.g., "internal/ui/app.go")
    $relativePath = $file.FullName.Replace($CurrentLocation.Path + "\", "").Replace("\", "/")

    Write-Host "Exporting: $relativePath" -ForegroundColor Gray

    # Create Header
    $header = "`n`n--- START OF FILE: $relativePath ---`n"
    Add-Content -Path $OutputFile -Value $header -Encoding UTF8

    # Read and Append Content
    $content = Get-Content -Path $file.FullName -Raw -Encoding UTF8
    Add-Content -Path $OutputFile -Value $content -Encoding UTF8
}

Write-Host "Done!" -ForegroundColor Green
Write-Host "Source code exported to: $OutputFile" -ForegroundColor Yellow