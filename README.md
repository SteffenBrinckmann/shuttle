# Shuttle
Efficient data transfer from instrument computers to a centrally managed server is a critical aspect of device integration. Shuttle is a utility designed to facilitate seamless file transfers from a source device to a designated destination. It automates the movement of files from a specified source directory to a target server using WebDAV or SFTP. Shuttle instances can be generated and deployed to target systems without requiring installation.

Originally developed by Martin Starmann at the Karlsruhe Institute of Technology and NFDI4Chem, Shuttle has since been integrated into a web application that allows users to log in and create customized shuttle instances. This project has evolved from [ELN_file_watcher](https://github.com/ComPlat/ELN_file_watcher) and [ShuttleBuilder](https://github.com/ComPlat/shuttlebuilder).

The objective of this fork is to further develop a version of shuttle-parent that operates on individual computers, enabling the creation of lightweight executable shuttles for automated file transfers. The system aims to streamline the configuration process, allowing users to easily input transfer parameters within shuttle-parent.

## Usage
**Windows XP Compatibility**
   - To run this tool on Windows XP, it must be compiled with Go version `1.10.8`.
   - When using the SFTP protocol on Windows XP, save a portable version of [WinSCP](https://winscp.net/download/WinSCP-5.21.5-Portable.zip) in the same folder as `shuttle.exe`.

### Before running Shuttle-parent, please collect the following information:
   - Source and destination directory paths: this differs for each instrument computer
   - WebDAV or SFTP server settings:
      User credentials for authentication: always the same
   - File modification duration threshold: 20sec is suggested
   - Transfer type and protocol: sftp is suggested

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

### Manually execute by adding command line arguments
To run the utility, execute the following command:

```shell
go build
shuttle -src < source_directory_path > -duration < time_in_seconds > -dst < destination_server > -user < username > -pass < password > -type < file|folder|zip|tar|flat_tar > -transfer < webdav|sftp > -commonRegex < Regexp >
```

### Compile hints
- Compile Shuttle using Go version `1.23`: see sub-topics below: manually / convenience
- Copy the executable to your source device / instrument computer
