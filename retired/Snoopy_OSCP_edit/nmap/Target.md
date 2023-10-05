***nmap -sCV -p22,53,80 10.10.11.212 -oN nmap/Target.md*** 
- scan:

	# Nmap 7.93 scan initiated Tue Oct  3 16:27:30 2023 as: nmap -sCV -p22,53,80 -oN nmap/Target.md 10.10.11.212
	Nmap scan report for snoopy.htb (10.10.11.212)
	Host is up (0.20s latency).
	
	PORT   STATE SERVICE VERSION
	22/tcp open  ssh     OpenSSH 8.9p1 Ubuntu 3ubuntu0.1 (Ubuntu Linux; protocol 2.0)
	| ssh-hostkey: 
	|   256 ee6bcec5b6e3fa1b97c03d5fe3f1a16e (ECDSA)
	|_  256 545941e1719a1a879c1e995059bfe5ba (ED25519)
	53/tcp open  domain  ISC BIND 9.18.12-0ubuntu0.22.04.1 (Ubuntu Linux)
	| dns-nsid: 
	|_  bind.version: 9.18.12-0ubuntu0.22.04.1-Ubuntu
	80/tcp open  http    nginx 1.18.0 (Ubuntu)
	|_http-title: SnoopySec Bootstrap Template - Index
	|_http-server-header: nginx/1.18.0 (Ubuntu)
	Service Info: OS: Linux; CPE: cpe:/o:linux:linux_kernel
	
	Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
	# Nmap done at Tue Oct  3 16:27:47 2023 -- 1 IP address (1 host up) scanned in 17.20 seconds

[nmap]
