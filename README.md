# Boot Blog Aggregator in Go

Blog Aggregator project build in GO (For use on Linux)

1. You will need the following software and files installed on your machine in order for the aggregator to work:
    - GO Runtime environment
    - Postgres: SQL database software that can be installed via `sudo apt install postgresql postgresql-contrib`
        - set system pass with `sudo passwd postgres`
        - enter Postgres shell with `sudo -u postgres psql`
            - Create a database with `CREATE DATABASE [database name];`
            - access with `\c [database name]`
            - Set databasae user and pass with `ALTER USER postgres PASSWORD '[password]';`
            - use `exit` or `\q` to exit the interface
    - Goose database migration tool installed with `go install github.com/pressly/goose/v3/cmd/goose@latest`
        - This will allow for migrations to be made up and down for the database
    - A config file named ".gatorconfig.json" in the home of your account file directory with {"db_url":"servername://username:password@[Server IP or localhost if on local machnine]:[port number: default 5432]/[database name]?sslmode=disable"} inserted into it
    - Install the program with `go install [program name]`
2. How the program works
- The blog aggregator allows for users and feeds to be added, users follow feeds and after aggregator has been run will be able to view posts from said feeds.
- all commands are preceded with `[program name]` and some have additional arguments
- to add users: `register [username]`
- to change users: `login [username]`
- to get list of all users: `users`
- to add feeds: `addfeed [name of feed] [url to feed]`
- to get list of feeds: `feeds`
- to follow feed: `follow [feed url]`
- to unfollowfeed: `unfollow [feed url]`
- to get all feeds user is following: `following`
- to aggregate feeds into posts: `agg [time interval in string format, such as "1s"]`
- to browse posts: `browse [limit of posts: default 2]`
- to reset DB: `reset`