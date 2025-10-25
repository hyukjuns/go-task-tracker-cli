# Task Tracker CLI
- Task Management CLI Tool
- All data is in json file

### Usage
```
# Build
go build -o task-cli .

# Adding, Updating, Deleting tasks
task-cli add "Study everyone"
task-cli update 1 "Play Soccer"
task-cli delete 1

# Change task ttatus
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing tasks
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

### Ref
https://roadmap.sh/projects/task-tracker?fl=1
