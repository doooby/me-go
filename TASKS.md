# tasks v2
[x] me21 create db
[ ] me22 create a task - only task & message
[ ] me23 list tasks

# tasks v1
[ ] me6 cli me ls
  - pagination
[ ] me11 db backups rake task
  - rake task to create empty DB (see lib/installation/create_database.sh)
[ ] me15 release v3.0
  - finish docs
  - import dev_log from all projects
  - me6
  - me16
[x] me16 end task
  - me end date:time # ends last task if applicable
[ ] me17 summary page
[ ] me18 explore using TTY toolkit (command, config, etc)
[ ] me19 integrate aws ec2 control
[ ] me20 custom control InputSearch

# ideas v1.1
- when create & edit, color the output green / red
- create import script of dev logs (and check current data)
- unify errors `bin/cli.rb e -a` `bin/cli.rb e -at=a`
- rename task#text na #message, actuall make it #goal
- run rails via cli command
- document well cli commands

# ideas v2.0 # TODO v2.0
- move the cli to zig
- move from web server to a native app with direct Sqlite bindings
