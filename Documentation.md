# Documentation
## Usage of shuttle

### Linux
1. Create data directory and open the parameter file. Open firefox and go to  https://sftpcloud.io/tools/free-sftp-server to create a temporary sftp server. Put the data into the parameter file.
    ``` bash
    >~/.../Repositories/shuttle$ mkdir ~/Temporary/shuttleTest
    >~/.../Repositories/shuttle$ code default_parameter.go
    >~/.../Repositories/shuttle$ firefox
    >~/.../Repositories/shuttle$ more default_parameter.go
    ```
    ``` go
    package main

    type Parameter struct {
        src, user, pass, dst, sendType, tType, name, duration, os string
    }

    func DefaultParameter() Parameter {
        return Parameter{
            src: "/home/steffen/Temporary/shuttleTest",
            name: "Linux test",
            os: "linux",
            dst: "eu-central-1.sftpcloud.io",
            user: "606ff7aba10c441e90891ab39ae64cfb",
            pass: "8No1fUsRqLFBwly1DQ45g9OLhbfkRriN",
            duration: "5",
            sendType: "sftp",
            tType: "file"}
    }
    ```

2. Build shuttle program, list the folder and start shuttle
    ```
    >~/.../Repositories/shuttle$ python build.py
    go: downloading go1.23.0 (linux/amd64)
    go: downloading github.com/StarmanMartin/gowebdav v0.0.0-20220901075112-8721ee532c0c
    go: downloading github.com/pkg/sftp v1.13.5
    go: downloading golang.org/x/crypto v0.1.0
    go: downloading github.com/kr/fs v0.1.0
    >~/.../Repositories/shuttle$ ll -rt
    total 18612
    ... content cropped ...
    -rw-rw-r-- 1 steffen steffen     471 Feb 22 12:46 default_parameter.go
    drwxrwxr-x 8 steffen steffen    4096 Feb 22 12:46 .git/
    drwxrwxr-x 3 steffen steffen    4096 Feb 22 12:46 ./
    -rwxrwxr-x 1 steffen steffen 9429853 Feb 22 12:46 shuttle.out*
    >~/.../Repositories/shuttle$ ./shuttle.out
    -> INFO: 2025/02/22 12:47:06 Starting at  22 Feb 25 12:47 CET
    -> INFO: 2025/02/22 12:47:06
    -----------------------------
    Logfile: /home/steffen/shuttle/shuttle_Linux test/shuttle_log.txt
    -----------------------------
    CMD Args:
    name=Linux test,
    dst=ssh://eu-central-1.sftpcloud.io:22,
    src=/home/steffen/Temporary/shuttleTest,
    duration=5 sec.,
    user=606ff7aba10c441e90891ab39ae64cfb,
    type=file,
    transfer=sftp
    -----------------------------
    -> INFO: 2025/02/22 12:47:06 Started transfer process.
    -> INFO: 2025/02/22 12:47:06 Started watch process.
    -> INFO: 2025/02/22 12:47:06 SSH Connected!
    -> INFO: 2025/02/22 12:48:01 Folder/File ready to send:  linux2.txt
    -> INFO: 2025/02/22 12:48:01 SSH Connected!
    -> INFO: 2025/02/22 12:48:01 Sending... linux2.txt
    -> INFO: 2025/02/22 12:48:16 Folder/File ready to send:  linux2.txt
    -> INFO: 2025/02/22 12:48:16 SSH Connected!
    -> INFO: 2025/02/22 12:48:16 Sending... linux2.txt
    ^C
    ```

3. In another terminal, I created a file 'linux2.txt' is the data folder. Show it here for reference.
    ``` bash
    >~/.../Repositories/shuttle$ more /home/steffen/Temporary/shuttleTest/linux2.txt
    This a linux file with numbers 1,2,3
    ```

4. Go to sftp server, list the content, and get the file to the local desktop. Show its content
    ``` bash
    >~/.../Repositories/shuttle$ sftp 606ff7aba10c441e90891ab39ae64cfb@eu-central-1.sftpcloud.io
    606ff7aba10c441e90891ab39ae64cfb@eu-central-1.sftpcloud.io's password:
    Connected to eu-central-1.sftpcloud.io.
    sftp> ls
    linux2.txt
    sftp> get linux2.txt
    Fetching /linux2.txt to linux2.txt
    linux2.txt                                                  100%   37     0.2KB/s   00:00
    sftp> ^D
    >~/.../Repositories/shuttle$ more linux2.txt
    This a linux file with numbers 1,2,3
    ```

### Windows
1. Edit the parameter file for windows, and show it here in linux. Build the windows file and list the files.
    ``` bash
    >~/.../Repositories/shuttle$ more default_parameter.go
    package main

    type Parameter struct {
        src, user, pass, dst, sendType, tType, name, duration, os string
    }

    func DefaultParameter() Parameter {
        return Parameter{
            src: "C:\\shuttleTest",
            name: "Windows test",
            os: "win10",
            dst: "eu-central-1.sftpcloud.io",
            user: "606ff7aba10c441e90891ab39ae64cfb",
            pass: "8No1fUsRqLFBwly1DQ45g9OLhbfkRriN",
            duration: "5",
            sendType: "sftp",
            tType: "file"}
    }
    >~/.../Repositories/shuttle$ python build.py
    >~/.../Repositories/shuttle$ ll -rt
    total 18616
    ... content cropped ...
    -rwxrwxr-x 1 steffen steffen 9429853 Feb 22 12:46 shuttle.out*
    -rw-r--r-- 1 steffen steffen      37 Feb 22 12:51 linux2.txt
    -rw-rw-r-- 1 steffen steffen     453 Feb 22 12:52 default_parameter.go
    drwxrwxr-x 8 steffen steffen    4096 Feb 22 12:53 .git/
    drwxrwxr-x 3 steffen steffen    4096 Feb 22 12:53 ./
    -rwxrwxr-x 1 steffen steffen 9520640 Feb 22 12:53 shuttle.exe*
    ```
2.  Use nautilus to create the data folder on the windows partition (the user rights will not be ok, but no problem). Copy the shuttle.exe to another location.
   ``` bash
   >~/.../Repositories/shuttle$ nautilus .
   ```

3. Reboot to windows partition, and use cmd.exe there. Start shuttle.exe. Using the windows explorer, create a folder in the data directory, and create a next text file 'windows.txt' there. Edit it with the editor.
    ``` cmd
    C:\Test>shuttle.exe
    -> INFO: 2025/02/22 11:57:14 Starting at  22 Feb 25 11:57 CET
    -> INFO: 2025/02/22 11:57:14
    -----------------------------
    Logfile: C:\Users\Steffen/shuttle/shuttle_Windows test/shuttle_log.txt
    -----------------------------
    CMD Args:
    name=Windows test,
    dst=ssh://eu-central-1.sftpcloud.io:22,
    src=C:\shuttleTest,
    duration=5 sec.,
    user=606ff7aba10c441e90891ab39ae64cfb,
    type=file,
    transfer=sftp
    -----------------------------
    -> INFO: 2025/02/22 11:57:14 Started transfer process.
    -> INFO: 2025/02/22 11:57:14 Started watch process.
    -> INFO: 2025/02/22 11:57:15 SSH Connected!
    -> INFO: 2025/02/22 11:59:59 Folder/File ready to send:  Test\windows.txt
    -> INFO: 2025/02/22 12:00:00 SSH Connected!
    -> INFO: 2025/02/22 12:00:00 Sending... windows.txt
    -> INFO: 2025/02/22 12:00:14 Folder/File ready to send:  Test\windows.txt
    -> INFO: 2025/02/22 12:00:15 SSH Connected!
    -> INFO: 2025/02/22 12:00:15 Sending... windows.txt
    ^C
    ```
4. Reboot to linux to verify and finish. Change to the folder, use sftp to list and get the windows.txt file.
    ``` bash
    cd FZJ/DataScience/Repositories/shuttle/
    >~/.../Repositories/shuttle$ sftp 606ff7aba10c441e90891ab39ae64cfb@eu-central-1.sftpcloud.io
    606ff7aba10c441e90891ab39ae64cfb@eu-central-1.sftpcloud.io's password:
    Connected to eu-central-1.sftpcloud.io.
    sftp> ls
    linux2.txt   windows.txt
    sftp> get windows.txt
    Fetching /windows.txt to windows.txt
    windows.txt                                                 100%   30     0.3KB/s   00:00
    sftp> ^D
    ```
5. Show its content
    ``` bash
    >~/.../Repositories/shuttle$ more windows.txt
    This is the windows test 9,8,7
    ```
