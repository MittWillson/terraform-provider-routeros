# routeros_scheduler (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `on_event` (String) Name of the script to execute. It must be presented at /system script.

### Optional

- `___id___` (Number) <em>Resource ID type (.id / name). This is an internal service field, setting a value is not required.</em>
- `___path___` (String) <em>Resource path for CRUD operations. This is an internal service field, setting a value is not required.</em>
- `comment` (String)
- `disabled` (Boolean)
- `interval` (String) Interval between two script executions, if time interval is set to zero, the script is only executed at its start time, otherwise it is executed repeatedly at the time interval is specified.
- `policy` (List of String) List of applicable policies:
    * dude - Policy that grants rights to log in to dude server.  
    * ftp - Policy that grants full rights to log in remotely via FTP, to read/write/erase files and to transfer files from/to the router. Should be used together with read/write policies.  
    * password - Policy that grants rights to change the password.  
    * policy - Policy that grants user management rights. Should be used together with the write policy. Allows also to see global variables created by other users (requires also 'test' policy).  
    * read - Policy that grants read access to the router's configuration. All console commands that do not alter router's configuration are allowed. Doesn't affect FTP.  
    * reboot - Policy that allows rebooting the router.  
    * romon - Policy that grants rights to connect to RoMon server.  
    * sensitive - Policy that grants rights to change "hide sensitive" option, if this policy is disabled sensitive information is not displayed.  
    * sniff - Policy that grants rights to use packet sniffer tool.  
    * test - Policy that grants rights to run ping, traceroute, bandwidth-test, wireless scan, snooper, and other test commands.  
    * write - Policy that grants write access to the router's configuration, except for user management. This policy does not allow to read the configuration, so make sure to enable read policy as well.  
policy = ["ftp", "read", "write"]
- `start_date` (String) Date of the first script execution.
- `start_time` (String) Time of the first script execution. If scheduler item has start-time set to startup, it behaves as if start-time and start-date were set to time 3 seconds after console starts up. It means that all scripts having start-time is startup and interval is 0 will be executed once each time router boots. If the interval is set to value other than 0 scheduler will not run at startup.

### Read-Only

- `id` (String) The ID of this resource.
- `next_run` (String)
- `owner` (String)
- `run_count` (String) This counter is incremented each time the script is executed.


