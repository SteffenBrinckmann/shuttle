# Shuttle
*Version 1.0*

*This project has been moved from [ELN_file_watcher](https://github.com/ComPlat/ELN_file_watcher)*

This utility facilitates the transfer of files from a source device to a target location, enabling seamless access for Chemotion. It automatically moves files from a specified source directory to a designated destination server via WebDAV or SFTP.

**Important**
1) If you want to run this tool on Windows XP it has to be compiled with go version 1.10.8.
2) If the tool is used with the SFTP protocol under Windows XP, you must also save a portable
   version of WinSCP [WinSCP](https://winscp.net/download/WinSCP-5.21.5-Portable.zip) in the same folder as the efw.exe.

## Setup the efw on a Windows system

Pleaser us the [ShuttleBuilder](https://github.com/ComPlat/shuttlebuilder) to build an executable instance of the shuttle. It contains detailed instruction on how to set it up.

If you do not want to use ShuttleBuilder you can compile it using GO version 1.19.3. Copy the executable into your source device. 
Follow the section Usage for the configurations.

## Usage
**If you have used the [ShuttleBuilder](https://github.com/ComPlat/shuttlebuilder) it is not necessary to add the cli arguments!**

To use the utility, run the following command:
```shell
shuttle -src <source_directory_path> -duration <time_in_seconds> -dst <destination_server> -user <username> -pass <password> -type <file|folder|zip> -transfer <webdav|sftp>
```

Replace <source_directory_path> with the path to the source directory 
containing the files to be transferred. <time_in_seconds> specifies the
period of time after which a file must be unchanged for it to be considered for transfer. <destination_server>
should be the WebDAV/SFTP server URL. You also need to provide a username (-user)
and password (-pass) for authentication.

Additionally, you can specify the type of transfer (-type) as either 'file', 'folder', or 'zip'. 

- The 'file' option treats each file individually, it ignores the directory structure of the source. All files are transferred directly to the root directory of traget
- The 'folder' option transmits entire folders only when all files in them are ready.
- The 'folder' option sends a folder zipped only when all files in the folder are ready.

Choose the transfer protocol (-transfer) as either 'webdav' or 'sftp'.

## Configuration
Before running the utility, ensure that you configure the following:

WebDAV or SFTP server settings.
Source and destination directory paths.
Duration threshold for file modification.
User credentials for authentication.
Transfer type and protocol.

## Contributing
Contributions are welcome! If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.







  

