# PEMtoPPK
A Go CLI tool to convert PEM files to PPK files for PuTTY

When you create a key pair in AWS and download it to SSH into a server, it will give you a `.pem` file. Now this is perfectly fine for NIX based systems, but for Windows this isn't as friendly. If you are a Windows user and use [Putty](http://www.putty.org/), you know that you need a PPK file. This tool allows you to convert a PEM file to a PPK file quickly and easily instead of using PuttyGen (we all know CLI tools are quicker) :grin:
