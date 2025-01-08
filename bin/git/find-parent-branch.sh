#!/bin/bash
# 根据git commit记录查找和当前分支具有最近距离的feature分支或dev分支
current_branch=$(git rev-parse --abbrev-ref HEAD)
branch_type=$(echo $current_branch | cut -d'/' -f1)
version=$(echo $current_branch | cut -d'/' -f2)
jira_id=$(echo $current_branch | cut -d'/' -f3)
commit_id=$(git rev-parse HEAD)
parent_branch=""

for i in {1..50}
do
    if [ -n "$parent_branch" ]; then
        echo $parent_branch
        exit 0
    else
        parent_commit_id=$(git rev-parse $commit_id^)
        parent_branch=$(git branch --contains $parent_commit_id | grep -v '\*' | grep -E "dev/$version|feature/$version|feat/$version" | head -n 1)
        commit_id=$parent_commit_id
    fi
done

echo "dev/$version"