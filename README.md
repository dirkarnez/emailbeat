EmailBeat
===========
Heartbeat sender using email 

### Usage
```
emailbeat 
--email={sender and receiver email} \
--password={password for the email} \
--smtp={smtp host} \
[--port={smtp port | default 587}] \
[--interval={heartbeat interval in seconds | default 1800}]
```
### Supported email services
- [x] Outlook
- [ ] Gmail
