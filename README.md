# know-more
Application that stores knowledge from command line, and knows how to track it, intercept it, keep context data and make it more accessible

## Background

Today, if you want to store knowledge you're dependant of, in order to run command lines utilities, the common intuitive way is to write it in notes, but this is very tedious and cumbersome . \ 
Another alternative is to use the linux history feature, which is good, but it has some limitations, for example:

1. The size of the history file is limited(although it can be configured and extended to some limit).
2. Although you have there the whole command, you're missing some context, for example,  in which directory the command was executed.
3. Another problem is that from history , you can't tell whether the command was success or not(RC is not logged there).
4. When working from several terminals, the logs are written to history only after you exit the terminal, hence commands are not intercepted into log in real time.

Because all of the above, we need a new application that will tackle all of these issues

## Planned Features
1. Application will intercept all commands invoke by a certain user from bash terminal In Real time , including PID of shell, command invoked, working directory, real return code of command, date and time of invocation.
2. It will parse all data and insert it into data structure , including all data from previous section, including whether the command was successfull or not(if not, will keep the return code).
3. It will give option to list all data , all data according to free search ,all data according to dates , according to tags, and by default will show only successfull invocations.
4. It will enable tagging intercepted commands with certain labels' keys and values, to enhance their fetching in the future.


## Prototype Technical details

- Prototype will be based on bash terminal script, which located per user in ~/.bashrc, and there we'll leverage bash' environment variable `PROMPT_COMMAND`   to aggregate all invoked commands from all bash terminals into a single file.
- `PROMPT_COMMAND` environment variable , contains the command that will be invoked just after executing each command, and just before writing the command prompt for the next command, on each bash, so it's like a hook that invoked in between each command and the next command execution.

```shell
## append these two lines to the end of the script ~/.bashrc(per user)
PROMPT_COMMAND=' RC=$(echo $?) ; echo "$(date +"%Y/%m/%d (%H:%M)") $(history 1 |cut -c 7-) $(pwdx $(echo $$)) ,RC= $RC " >> /tmp/trace'
export PROMPT_COMMAND
```
- The chosen programming language for implementing this application is Golang.
- There will be two components:
  1. a linux service background utility.
  2. a command line utility to improve knowledge and to get/fetch knowledge.
  3. a lightweight database(TBD)/files with smart datastructures that will store all knowledge data.
