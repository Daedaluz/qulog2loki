# QuLog2Loki

QNAP QuLog has a function to send logs to a rfc5424 receiver.
Since the log messages can contain QuLog specific IDs, I found it difficult to export proper, readable log rows to loki
with promtail.

I found some constants in `/mnt/ext/opt/QuLog/opt/www/app.js` on the nas, which i used for translation in the messages.


