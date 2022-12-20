# EmailSender  
一个批量邮件发送器 半自用 有建议Issues中提出 很可能下个版本就会修正/添加功能  
___________              .__.__    _________                  .___              
\_   _____/ _____ _____  |__|  |  /   _____/ ____   ____    __| _/___________  
 |    __)_ /     \\__  \ |  |  |  \_____  \_/ __ \ /    \  / __ |/ __ \_  __ \  
 |        \  Y Y  \/ __ \|  |  |__/        \  ___/|   |  \/ /_/ \  ___/|  | \/  
/_______  /__|_|  (____  /__|____/_______  /\___  >___|  /\____ |\___  >__|  
GLOBAL OPTIONS:  
   --smtpServer value, -s value       smtp服务器host  
   --smtpPort value, -p value         smtp服务器端口 (default: 465)  
   --username value, -u value         发送者用户名  
   --password value, --pa value       发送者密码  
   --recverList value, -l value       接收者列表  
   --mailbodyfile value, --mbf value  邮件BODY 的文件  
   --mailheader value, --mh value     邮件标题  
   --attachfile value, --af value     附件文件(会大幅降低邮件发送速度)  
   --DBGMOD, --DBG                    DBG MOD (default: false)  
   --help, -h                         show help (default: false)  
   --version, -v                      print the version (default: false)  

body文件中可以使用的环境变量:  
%RECV_EMAIL% 对应接收者邮箱  
%SEND_EMAIL% 对应发送者邮箱  
EG:  
![image](https://user-images.githubusercontent.com/31125137/208306366-3387088c-49b7-4822-87cd-5f8685f60aab.png)
![image](https://user-images.githubusercontent.com/31125137/208306377-2f8b6b92-d74b-4d2d-a371-281ea7d64d7d.png)
