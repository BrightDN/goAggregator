# UpdateLog

## Current versions
v 0.1 -- Initial BETA release
v 0.2 -- Agg now scrapes concurrently so that all feeds are being checked at the same time at each interval

## Future updates
- Adding a help function to help figure out how to use commands and their purpose
- Adding a REPL mode so you can work continiously without the program self-closing after every command
- Cleaning up some DB functionality (using CURRENT_TIMESTAMP on db instead of time.Now() in code)
- Cleaning up and adding more clarification to error handling
- Adding some QoL changes e.g. Register a user automatically when attempting to log in, if user wishes to follow a feed that doesn't exist, prompt to create it, then continue to follow
- Adding pagination to the browse function (ONLY IN REPL mode)