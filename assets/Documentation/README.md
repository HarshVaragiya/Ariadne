# Example Output
```log
Ariadne Starting On Target [127.0.0.1] with ProjectIndex [test] and httpExtensions [php,html,txt]
[NMAP] Found OPEN Port 80 running service http
[NMAP] Found OPEN Port 80 running service http
[NMAP] Found OPEN Port 5601 running service esmagent
[NMAP] Found OPEN Port 6942 running service unknown
[NMAP] Found OPEN Port 9200 running service wap-wsp
[NMAP] Found OPEN Port 9300 running service vrace
[NMAP] Found OPEN Port 63342 running service 
============================================================
NMAP Port Scan Report for target: 127.0.0.1 
------------------------------------------------------------
 http            Running on [80] 
 esmagent        Running on [5601] 
 unknown         Running on [6942] 
 wap-wsp         Running on [9200] 
 vrace           Running on [9300] 
 unknown         Running on [63342] 
------------------------------------------------------------
End of report.
============================================================

Following services appear to be running :  map[:[63342] esmagent:[5601] ftp:[] http:[80] unknown:[6942] vrace:[9300] wap-wsp:[9200]]
Starting HTTP enumeration on ports  [80]
[GOBUSTER-DIR] Endpoint found : http://127.0.0.1:80/index.html       [200] 
[GOBUSTER-DIR] Endpoint found : http://127.0.0.1:80/javascript       [301] 
[GOBUSTER-DIR] Endpoint found : http://127.0.0.1:80/secret.txt       [200] 
[GOBUSTER-DIR] Endpoint found : http://127.0.0.1:80/server-status    [200] 

Http Report/s :
============================================================
Gobuster dir search report for URL: http://127.0.0.1:80/ 
------------------------------------------------------------
 http://127.0.0.1:80/index.html           - STATUS [200] 
 http://127.0.0.1:80/javascript           - STATUS [301] 
 http://127.0.0.1:80/secret.txt           - STATUS [200] 
 http://127.0.0.1:80/server-status        - STATUS [200] 
------------------------------------------------------------
End of report.
============================================================


Process finished with exit code 0

```