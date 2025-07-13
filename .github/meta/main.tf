resource "github_repository" "bookworm_repo" {
  name          = "bookworm"
  description   = "Bookworm is a CLI application to manage bookmarks"
  auto_init     = false
  has_downloads = true
  has_issues    = true
  has_projects  = true
}

resource "github_branch_protection" "main_branch_protection" {
  count            = github_repository.bookworm_repo.visibility == "private" ? 0 : 1
  repository_id    = github_repository.bookworm_repo.node_id
  pattern          = "main"
  enforce_admins   = true
  allows_deletions = false
}
