#!/usr/bin/python3
import requests
import os
file = input("file: ")

while file != "q":
    url = f"http://icinga.cerberus.local:8080/icingaweb2/lib/icinga/icinga-php-thirdparty{file}"
    r = requests.get(url)
    print(r.text)
    file = input("file: ")

