resource "github_repository" "bookworm_repo" {
  name          = "bookworm"
  description   = "Bookworm is a CLI application to manage bookmarks"
  auto_init     = false
  has_downloads = true
  has_issues    = true
  has_projects  = true
}

import {
  to = github_repository.bookworm_repo
  id = "bookworm"
}
