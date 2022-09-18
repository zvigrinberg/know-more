# know-more
Application that stores knowledge from command line, and knows how to track it, intercept it, keep context data and make it more accessible

## Planned Features
1. Application will intercept all commands invoke by a certain user from bash terminal , including PID of shell, command invoked, working directory, real return code of command, date and time of invocation.
2. It will parse all data and insert it into data structure , including all data from previous section, including whether the command was successfull or not(if not, will keep the return code).
3. It will give option to list all data , all data according to free search ,all data according to dates , according to tags, and by default will show only successfull invocations.
4. It will enable tagging intercepted commands with certain labels' keys and values, to enhance their fetching in the future.


## Prototype Technical details

- Prototype will be based on bash terminal script, which located per user in ~/.bashrc, and there we'll leverage bash' environment variable PROMPT_COMMAND   to aggregate all invoked commands from all bash terminals into a single file.

```shell
## append these two lines to the end of the script ~/.bashrc(per user)
PROMPT_COMMAND=' RC=$(echo $?) ; echo "$(date +"%Y/%m/%d (%H:%M)") $(history 1 |cut -c 7-) $(pwdx $(echo $$)) ,RC= $RC " >> /tmp/trace'
export PROMPT_COMMAND
```
- The chosen programming language for implementing this application is Golang.
