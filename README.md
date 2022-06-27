# Lets-Learn-Git
Learning git

-----Let's learn GIT/GITHUB-------

01)git --version   -->check installed or not on cmd
02)git config --global user.name "Afzal Ansari"
03)git config --global user.email "abc@gmail.com"
04)git config --global --edit  -->edit the name and email
05)escape :wq press enter  --> come back to gitbash window
06)git init  -->to initialise
07)git add sum.go  -->Now the file is tracked, files goes to staging area(before repository).
08)git commit -m "initial commit"  -->will commit the changes
09)git log  -->show how much commit happened
10)git add . -->it insert all the files in staging area
11)git checkout hexCode  -->will take come back to initial state based on hexCode
12)git checkout master   -->will take go forward again
13)git branch dev    -->another branch created(one master branch can have many branches)
14)git branch    ->to see how much branches are availble
15)git branch afzal/multiply  -->do code of this functionality and commit
16)git merge afzal/multiply  -->will insert all the code from afzal/multiply branch to dev branch
17)git merge dev             -->will insert all the code from dev branch to master branch


git branch -a  -->show all branches
git branch -d branche_name  -->delete the local branch


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
