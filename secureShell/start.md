
# go get golang.org/x/crypto/ssh

# Executing a command over SSH

Once ssh.Client is created, you can begin creating sessions. A client supports
multiple sessions at once. A session has its own standard input, output, and error.
They are standard reader and writer interfaces.

To execute a command there are a few options: Run(), Start(), Output(), and
CombinedOutput(). They are all very similar, but behave a little differently:

* session.Output(cmd): The Output() function will execute the command, and
    return session.Stdout as a byte slice.
* session.CombinedOutput(cmd) : This does the same as Output() , but it returns both
    standard output and standard error combined.
* session.Run(cmd) : The Run() function will execute the command and wait for it
    to finish. It will fill the standard output and error buffers, but it won't do
    anything with them. You have to manually read the buffers or set the
    session output to go to the Terminal output before calling Run() (for
    example, session.Stdout = os.Stdout ). It will only return without an error if the
    program exited with an error code of 0 and there were no issues copying the
    standard output buffers.
* session.Start(cmd) : The Start() function is similar to Run() , except that it will
    not wait for the command to finish. You must explicitly call session.Wait() if
    you want to block execution until the command is complete. This is useful
    for starting long running commands or if you want more control over the
    application flow.

A session can only perform one action. Once you call Run() , Output() ,
CombinedOutput() , Start() , or Shell() , you can't use the session for executing any
other commands. If you need to run multiple commands, you can string them
together separated with a semicolon. For example, you can pass multiple
commands in a single command string like this:
## df -h; ps aux; pwd; whoami;

