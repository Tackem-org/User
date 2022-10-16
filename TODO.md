# DOING NOW

# TODO
- Normal user systems (edit profile etc)
  - need to allow the user to pick an Icon and look at making the Icon folder a special folder or make it base64 in the DB
- View Other users option with permission
- make the permissions table in user edit more flexable allow "search", pagination, groups
  - a way of splitting permissions up (maybe some kind of permission type groups that are fixed.) (enum help <https://github.com/go-gorm/gorm/issues/1978>)
- Need a way of sending group and permission adding requests
## Future
- look at rules for allowed password too.
- possably need to add in an email field for comunications to a user through email. this will then need some form of list of allowed emails to recieve
- allow choice of DB for this system from SQLite (current) and mysql
