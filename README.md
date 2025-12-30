# Lets-Learn-Git
Learning git

-----Let's learn GIT/GITHUB-------

```bash
# Setup:
git --version                      # Check if Git is installed
git config --global user.name "Afzal Ansari"
git config --global user.email "abc@gmail.com"
git config --global --edit        # Edit name and email
git config --global --list        # Verify configuration

#T o save and exit the config editor:
:wq                                # (Press `Esc`, then type `:wq` and press Enter)

# Git Basics:
git init                           # Initialize a Git repository in current folder
git add sum.go                     # Add a file to staging area
git add .                          # Add all files to staging area
git commit -m "initial commit"     # Commit the staged changes
git log                            # View commit history

# Branching and Merging:
git branch dev                     # Create a new branch named 'dev'
git branch                         # List local branches
git branch -a                      # List all branches (local and remote)
git checkout dev                   # Switch to the 'dev' branch
git checkout <commit_hash>        # Checkout a specific commit by its hash
git checkout master                # Switch back to the 'master' branch
git branch afzal/multiply          # Create a new feature branch from current branch
git checkout afzal/multiply        # Switch to the 'afzal/multiply' branch
# (make changes and commit them on afzal/multiply)

git checkout dev                   # Switch to 'dev' branch
git merge afzal/multiply           # Merge 'afzal/multiply' into 'dev'

git checkout master                # Switch to 'master' branch
git merge dev                      # Merge 'dev' into 'master'
git branch -d branch_name          # Delete a local branch
-----------------------------------------------------------
#Git Rebase:
git checkout test
git rebase master                  # Apply commits from master onto test

git checkout master
git merge test                     # Merge test → master

git rebase -i HEAD~4               # Interactive rebase last 4 commits

Selectively apply commits to another branch:
git cherry-pick commit_id1 commit_id2

# Hide your local uncommitted changes temporarily.
git stash                          # Stash tracked files
git stash -u                       # Stash tracked + untracked files
git stash list                     # Show all stashes
git stash show -p stash@{0}        # Show stashed code changes too
git pull                           # pull the latest changes from repo.
git stash pop                      # get back the stashed files and changes.

# Git Amend:
Modify the most recent commit (e.g., for a small fix).
git add .
git commit --amend -m "Updated commit message"

# Git Reset:
Go back to a specific commit (destructive operation).
git reset --hard commit_id

# Pushing to GitHub:
git remote -v                                          # Verify remote repo
git remote add origin https://github.com/afzalansari028/Lets-Learn-Git.git
git branch -M master                                   # Rename main to master (if needed)
git push -u origin master                              # Push code to GitHub

# Make some changes to diff.go
git add diff.go
git commit -m "Added one function in diff.go"
git push

# Some Basic Shell Commands:
mkdir folder_name         # Create a folder
cd ..                     # Go back one directory
clear                     # Clear the terminal screen
