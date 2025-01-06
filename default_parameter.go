package main

type Parameter struct {
	src, user, pass, dst, sendType, tType, name, duration, os string
}

func DefaultParameter() Parameter {
	return Parameter{
		src: "",
		name: "Instrument 1",
		os: "winXP",
		dst: "",
		user: "",
		pass: "",
		duration: "5",
		sendType: "sftp",
		tType: "file"}
}