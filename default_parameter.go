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
