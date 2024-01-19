# Lets-Learn-Git
Learning git

-----Let's learn GIT/GITHUB-------

01)git --version   -->check installed or not on cmd
02)git init  -->to initialize(convert a folder to working directory)
03)git config --global user.name "Afzal Ansari"
04)git config --global user.email "abc@gmail.com"
05)git config --global --edit  -->edit the name and email
06)git congfig --global --list  -->verify the configuration
07)escape :wq press enter  --> come back to gitbash window
08)git add sum.go  -->Now the file is tracked, files goes to staging area(before repository).
09)git commit -m "initial commit"  -->will commit the changes
10git log  -->show how much commit happened
11)git add . -->it insert all the files in staging area
12)git checkout hexCode  -->will take come back to initial state based on hexCode
13)git checkout master   -->will take go forward again
14)git branch dev    -->another branch created(one master branch can have many branches)
15)git branch or git branch -a   ->to see a branch or all branches
16)git branch afzal/multiply  -->do code of this functionality and commit
17)git merge afzal/multiply  -->will insert all the code from afzal/multiply branch to dev branch
18)git merge dev             -->will insert all the code from dev branch to master branch
19)git branch -d branche_name  -->delete the local branch


Git Rebase:-when u want to do create a latest commit from child branch to parent--
   git checkout test
   git rebase master

   git chcekout master
   git merge test  -->from master <- test (merging from test to master)

   git rebase -i HEAD~4 -->you can rearrange the commits, after that press :wq to come out


----Once the repository created---
git remote -v
git remote add origin https://github.com/afzalansari028/Lets-Learn-Git.git
git branch -M master
git push -u origin master


---Want some changes in diff.go file---
do changes in the file or add functionality
git add diff.go
git commit -m "added one function in diff.go"
git push or  git push -u origin master




---Some basic commands---
mkdir -->create folder
cd\   -->back space
clear -->clear screen
