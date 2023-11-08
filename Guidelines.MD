### Branch Naming

When creating a branch from develop, it is adviced to take the name of the task that it will work on from the kanban board

- Create & switch to branch on local
```
git checkout -b <task-name>
```
- Push & register branch to remote repo
```
git push -u origin <task-name>
```

### Merging back to develop branch

To protect the branches from conflicts, and to have more team memebers understand the added features branch protection was enabled for the repo

Merging to master or develop branches is done as follows:
1.  Go to the merging branch on gitea e.g. test branch
2. Click "New Pull Request" button next to the branch name
3. Choose merge into develop
4. All collaboraters will receive a notification on gitea
5. Any collaborate can review & approve

### Frontend Design

To make it easier to envision the elements and overall design, we have agreed that we will follow the basic design patterns from the below website.

For the time being, no effort should be put into the project, as the final step will be to finalize the design in the CSS.

[Sample website: Men's Journal](https://www.mensjournal.com)

