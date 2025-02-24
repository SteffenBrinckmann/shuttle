# Preprocess Scripts
Shuttle supports preprocess scripts, which allow you to execute custom scripts before transferring files.

Such a script can manipulate the file before it is sent. The script is called for each file that will be sent. The script gets one argument with the absolute path to the file. The script should either overwrite the file or write a new copy of the file in the same folder

- Windows: Scripts must be .exe files.
- Linux: All executable file types are supported.

## Preprocess Script Requirements:
The script files must be located in the directory:
```plaintext
~/shuttle/scripts
```

If the shuttle has already been executed, the folder has already been created.
