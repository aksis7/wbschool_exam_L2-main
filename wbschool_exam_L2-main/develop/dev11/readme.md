C:\WINDOWS\system32>curl -X POST -d "id=1&user_id=123&title=Meeting&date=2024-06-10" http://localhost:8080/create_event
{"result":"event created"}

C:\WINDOWS\system32>curl -X POST -d "id=1&user_id=123&title=UpdatedMeeting&date=2024-06-11" http://localhost:8080/update_event
{"result":"event updated"}

C:\WINDOWS\system32>http POST http://localhost:8080/delete_event id=1
"http" не является внутренней или внешней
командой, исполняемой программой или пакетным файлом.

C:\WINDOWS\system32>curl -X POST -d "id=1" http://localhost:8080/delete_event
{"result":"event deleted"}

PS C:\Users\марат\Desktop\4 курс\WB_learn\wbschool_exam_L2-main\develop\dev11> go run main.go
2024/12/25 01:16:56 Server started on :8080
2024/12/25 01:18:35 [POST] /create_event [::1]:49286 0s
2024/12/25 01:18:40 [POST] /update_event [::1]:49298 0s
2024/12/25 01:19:15 [POST] /delete_event [::1]:49338 0s