package main

type Parameter struct {
	src, user, pass, dst, sendType, tType, name, duration, os string
}

func DefaultParameter() Parameter {
	return Parameter{
		src: "C:\\Test",
		name: "Windows test",
		os: "win10",
		dst: "server123",
		user: "user123",
		pass: "password123",
		duration: "5",
		sendType: "sftp",
		tType: "folder"}
}
