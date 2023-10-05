9th May 2023 / Document No D23.100.237
Prepared By: TheCyberGeek
Machine Author: ctrlzero
Difficulty: Hard
Classification: Official

# Synopsis

Snoopy is a Hard Difficulty Linux machine that involves the exploitation of an LFI vulnerability to extract the configuration secret of ***Bind9***. The obtained secret allows the redirection of the ***mail*** subdomain to the attacker's IP address, facilitating the interception of password reset requests within the Mattermost chat client. Within that service, a custom plugin designed for web admins to log into remote servers is manipulated to direct them to a server set up as an SSH honeypot , leading to the interception of cbrown 's credentials. Exploiting the privileges of cbrown , the attacker utilizes the ability to execute git apply as sbrown , resulting in a unique symlinking attack for privilege escalation. The final stage involves the abuse of ***CVE-2023-20052*** to include the root user's SSH key into a file via XXE , with the payload scanned by clamscan to trigger the XXE output in the debug response.

## Skills Required

- Reconnaissance
- DNS knowledge
- Understanding of the git command line interface usage
- Source code analysis
## Skills Learned
- Abusing leaked Bind9 secret keys to control DNS entries
- Intercepting SSH credentials via SSH honeypots
- Abusing git symlinks for privilege escalation via sudo
- Injecting XXE payloads into DMG files