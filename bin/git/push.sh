#!/bin/bash
current_branch=$(git rev-parse --abbrev-ref HEAD)
branch_type=$(echo $current_branch | awk -F'/' '{print $1}')
version=$(echo $current_branch | awk -F'/' '{print $2}')
jira_id=$(echo $current_branch | awk -F'/' '{print $3}')
commit_message=$(git log -1 --pretty=%B)
title=$commit_message
target_branch="dev/$version"

if [[ $commit_message != *$jira_id* ]]; then
    title="[$jira_id]$commit_message"
fi

if [[ ! "$branch_type" =~ ^(feature|feat|bugfix|hotfix|task|subtask|subTask)$ ]]; then
    git push --set-upstream origin $current_branch
    exit 0
fi

if [[ "$branch_type" =~ ^(task|subtask|subTask)$ ]]; then
    project_path=$(cd `dirname $0`; pwd)
    target_branch=$($project_path/find-parent-branch.sh)
fi

git push --set-upstream origin $current_branch -o merge_request.create -o merge_request.title="$title" -o merge_request.target=$target_branch
