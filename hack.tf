# REQUIRED: Only used to check available resources in .github/workflows/update.yml
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "6.5.0"
    }
  }
}

provider "github" {
  # Configuration options
}
