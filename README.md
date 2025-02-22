# Shuttle
**Version 1.13**
Shuttle is a utility that facilitates seamless file transfer from a source device to a target location. It automates the movement of files from a specified source directory to a destination server using WebDAV or SFTP.

- This project has been migrated from [ELN_file_watcher](https://github.com/ComPlat/ELN_file_watcher).
- Please note: if you use the [ShuttleBuilder](https://github.com/ComPlat/shuttlebuilder), many things are preconfigured and do not need to be manually added.
   - Use the [ShuttleBuilder](https://github.com/ComPlat/shuttlebuilder) to build an executable instance of Shuttle.
   - Detailed setup instructions are available in the ShuttleBuilder repository.

---

Steffen's TODOs:
- can we obsure the password with go

---

## **Usage**
- Compile Shuttle using Go version `1.23`: see sub-topics below: manually / convenience
- Copy the executable to your source device and follow the **Usage** section for configuration.
- **Windows XP Compatibility**
   - To run this tool on Windows XP, it must be compiled with Go version `1.10.8`.
   - When using the SFTP protocol on Windows XP, save a portable version of [WinSCP](https://winscp.net/download/WinSCP-5.21.5-Portable.zip) in the same folder as `shuttle.exe`.
- Please note: before running Shuttle, configure the following:
   - WebDAV or SFTP server settings.
   - Source and destination directory paths.
   - File modification duration threshold.
   - User credentials for authentication.
   - Transfer type and protocol.

### Manually execute by adding command line arguments
To run the utility, execute the following command:

```shell
go build
shuttle -src < source_directory_path > -duration < time_in_seconds > -dst < destination_server > -user < username > -pass < password > -type < file|folder|zip|tar|flat_tar > -transfer < webdav|sftp > -commonRegex < Regexp >
```

### Convenience: read default parameters from file
1. Edit the default parameters in "default_parameter.go"
1. Execute build which will create an executable for the desired operating system
```shell
python build.py
```

### Parameters:
- < source_directory_path >: Path to the source directory containing files to transfer.
- < time_in_seconds >: Time threshold (in seconds) for determining if a file is ready for transfer (unchanged for the specified duration).
- < destination_server >: URL of the WebDAV/SFTP server.
- < username >: Username for authentication.
- < password >: Password for authentication.
- Transfer Type (-type):
  - file: Transfers individual files, ignoring the source directory structure.
  - folder: Transfers entire folders only when all files in them are ready.
  - zip: Transfers folders as ZIP archives when all files in the folder are ready.
  - tar: Transfers folders as .tar.gz archives when all files in the folder are ready.
  - flat_tar: Transfers folders as .tar.gz archives, grouping files by a common regex.

If using flat_tar, provide a regular expression (< Regexp >) to group related files into datasets for archiving. Files with matching regex groups (or global matches if no groups are used) are archived together.

Transfer Protocol (-transfer):
- webdav: Use WebDAV protocol.
- sftp: Use SFTP protocol.

---

## Preprocess Scripts
Shuttle supports preprocess scripts, which allow you to execute custom scripts before transferring files.

- Windows: Scripts must be .exe files.
- Linux: All executable file types are supported.

#### Preprocess Script Requirements:
The script files must be located in the directory:
```plaintext
~/shuttle/scripts
```

If the shuttle has already been executed, the folder has already been created.

Such a script can manipulate the file before it is sent. The script is called for each file that will be sent. The script gets one argument with the absolute path to the file. The script should either overwrite the file or write a new copy of the file in the same folder

---

## Contributing
We welcome contributions! If you encounter issues or have suggestions for improvement, please:

1) Open an issue.
2) Submit a pull request.


## License
This project is licensed under the MIT License.
