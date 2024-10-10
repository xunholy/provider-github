# REQUIRED: Only used to check available resources in .github/workflows/update.yml
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "6.3.1"
    }
  }
}

provider "github" {
  # Configuration options
}
