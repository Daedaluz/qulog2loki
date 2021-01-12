
let enums = {
     ACTION: {
      UNKNOWN: 0,
      DEL: Math.pow(2, 0),
      READ: Math.pow(2, 1),
      WRITE: Math.pow(2, 2),
      OPEN: Math.pow(2, 3),
      MKDIR: Math.pow(2, 4),
      N_F_S_MOUNT_SUCC: Math.pow(2, 5),
      N_F_S_MOUNT_FAIL: Math.pow(2, 6),
      RENAME: Math.pow(2, 7),
      LOGIN_FAIL: Math.pow(2, 8),
      LOGIN_SUCC: Math.pow(2, 9),
      LOGOUT: Math.pow(2, 10),
      N_F_S_UMOUNT: Math.pow(2, 11),
      COPY: Math.pow(2, 12),
      MOVE: Math.pow(2, 13),
      ADD: Math.pow(2, 14),
      AUTH_FAIL: Math.pow(2, 15),
      AUTH_OK: Math.pow(2, 16),
      TRASH_RECOVERY: Math.pow(2, 17),
      ADD_TRANSCODE: Math.pow(2, 18),
      DEL_TRANSCODE: Math.pow(2, 19),
      UPDATE_TRANSCODE: Math.pow(2, 20),
      WATERMARK: Math.pow(2, 21),
      ROTATE: Math.pow(2, 22),
      ADD_THUMBNAIL: Math.pow(2, 23),
      ADD_FILING: Math.pow(2, 24),
      UPDATE_FILING: Math.pow(2, 25),
      DELETE_FILING: Math.pow(2, 26),
      PAUSE_FILING: Math.pow(2, 27),
      STOP_FILING: Math.pow(2, 28),
      EJECT: Math.pow(2, 29),
      ENCRYPT: Math.pow(2, 30),
      DECRYPT: Math.pow(2, 31),
      CONN_ACTION_CREATE_SHARE_LINK: Math.pow(2, 32),
      CONN_ACTION_COMPRESS: Math.pow(2, 33),
      CONN_ACTION_EXTRACT: Math.pow(2, 34),
      CONN_ACTION_CLOUDCONVERT: Math.pow(2, 35),
      ANY: Math.pow(2, 36) - 1
    },
    SERVICE: {
      UNKNOWN: 0,
      SAMBA: 1 << 0,
      FTP: 1 << 1,
      HTTP: 1 << 2,
      NFS: 1 << 3,
      AFP: 1 << 4,
      TELNET: 1 << 5,
      SSH: 1 << 6,
      ISCSI: 1 << 7,
      RADIUS: 1 << 8,
      VPN: 1 << 9,
      HTTPS: 1 << 10,
      ANY: (1 << 11) - 1
      // ANY:     ((1 << 9) - 1) ^ (1 << 3)
    },
}

String.prototype.toPascalCase = function() {
    const words = this.match(/[a-z]+/gi);
    if (!words) return "";
    return words
        .map(function(word) {
            return word.charAt(0).toUpperCase() + word.substr(1).toLowerCase();
        })
        .join("");
};

console.log("package main");
console.log("");

console.log("const(");
for(let x in enums.ACTION) {
    console.log("    ", "action" + x.toPascalCase(), "=", enums.ACTION[x]);
}
console.log(")");

console.log("");

console.log("var actionMap = map[int]string{");
for(let x in enums.ACTION) {
    console.log("    ", "action"+x.toPascalCase(), ":", "`" + x.toPascalCase() + "`,");
}
console.log("}");

console.log("");

console.log("const(");
for(let x in enums.SERVICE) {
    console.log("    ", "service" + x.toPascalCase(), "=", enums.SERVICE[x]);
}
console.log(")");

console.log("");

console.log("var serviceMap = map[int]string{");
for(let x in enums.SERVICE) {
    console.log("    ", "service"+x.toPascalCase(), ":", "`" + x.toPascalCase() + "`,");
}
console.log("}");