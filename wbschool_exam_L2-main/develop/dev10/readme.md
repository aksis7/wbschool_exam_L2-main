команда
./go-telnet yandex.ru 443
результат
Connected to yandex.ru:443
ввод
GET / HTTP/1.1
Host: yandex.ru
Connection: close
Положительный результат
HTTP/1.1 302 Moved temporarily
Transfer-Encoding: chunked
X-Robots-Tag: unavailable_after: 12 Sep 2022 00:00:00 PST
Portal: Home
Report-To: { "group": "network-errors", "max_age": 100, "endpoints": [{"url": "https://dr.yandex.net/nel", "priority": 1}, {"url": "https://dr2.yandex.net/nel", "priority": 2}]}
X-Content-Type-Options: nosniff
set-cookie: is_gdpr=0; Path=/; Domain=.yandex.ru; Expires=Thu, 24 Dec 2026 19:14:19 GMT
set-cookie: is_gdpr_b=CJfOVRD+pQIoAg==; Path=/; Domain=.yandex.ru; Expires=Thu, 24 Dec 2026 19:14:19 GMT
set-cookie: _yasc=Sjas6Yh8uEJoD3jCNziorTYxhis1SiGRuARaQAnlU06gdQdnk9xEiPh5V53aqHhhEXgo; domain=.yandex.ru; path=/; expires=Fri, 22 Dec 2034set-cookie: receive-cookie-deprecation=1; Path=/; Domain=.yandex.ru; Expires=Wed, 24 Dec 2025 19:14:19 GMT; SameSite=None; Secure; HttpOnly; Partitioned
Date: Tue, 24 Dec 2024 19:14:19 GMT
Accept-CH: Sec-CH-UA-Platform-Version, Sec-CH-UA-Mobile, Sec-CH-UA-Model, Sec-CH-UA, Sec-CH-UA-Full-Version-List, Sec-CH-UA-WoW64, Sec-CH-UA-Arch, Sec-CH-UA-Bitness, Sec-CH-UA-Platform, Sec-CH-UA-Full-Version, Viewport-Width, DPR, Device-Memory, RTT, Downlink, ECT, Width       
X-Yandex-Req-Id: 1735067659099978-6771300649550216118-balancer-l7leveler-kubr-yp-sas-132-BAL
P3P: policyref="/w3c/p3p.xml", CP="NON DSP ADM DEV PSD IVDo OUR IND STP PHY PRE NAV UNI"
Location: https://dzen.ru/?yredirect=true
Cache-Control: max-age=86400,private
NEL: {"report_to": "network-errors", "max_age": 100, "success_fraction": 0.001, "failure_fraction": 0.1}
Connection: Close

0


Connection closed by remote host
Connection closed

