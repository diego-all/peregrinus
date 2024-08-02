# peregrinus



    curl -X POST http://localhost:8080/execute -d 'cat ../../../../etc/passwd' -H 'Content-Type: text/plain' (v1)

Contexts (OS and file server)

    curl -X POST http://localhost:8080/execute-system -d 'cat /etc/passwd' -H 'Content-Type: text/plain'  (v2)

    curl -X POST http://localhost:8080/execute-files -d 'cat etc/passwd' -H 'Content-Type: text/plain'  (v2)


    curl -X POST http://localhost:8080/execute-python -d 'system_info.py' -H 'Content-Type: text/plain'  (v3)
