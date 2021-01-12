package main

const (
	actionUnknown                   = 0
	actionDel                       = 1
	actionRead                      = 2
	actionWrite                     = 4
	actionOpen                      = 8
	actionMkdir                     = 16
	actionNFSMountSucc              = 32
	actionNFSMountFail              = 64
	actionRename                    = 128
	actionLoginFail                 = 256
	actionLoginSucc                 = 512
	actionLogout                    = 1024
	actionNFSUmount                 = 2048
	actionCopy                      = 4096
	actionMove                      = 8192
	actionAdd                       = 16384
	actionAuthFail                  = 32768
	actionAuthOk                    = 65536
	actionTrashRecovery             = 131072
	actionAddTranscode              = 262144
	actionDelTranscode              = 524288
	actionUpdateTranscode           = 1048576
	actionWatermark                 = 2097152
	actionRotate                    = 4194304
	actionAddThumbnail              = 8388608
	actionAddFiling                 = 16777216
	actionUpdateFiling              = 33554432
	actionDeleteFiling              = 67108864
	actionPauseFiling               = 134217728
	actionStopFiling                = 268435456
	actionEject                     = 536870912
	actionEncrypt                   = 1073741824
	actionDecrypt                   = 2147483648
	actionConnActionCreateShareLink = 4294967296
	actionConnActionCompress        = 8589934592
	actionConnActionExtract         = 17179869184
	actionConnActionCloudconvert    = 34359738368
	actionAny                       = 68719476735
)

var actionMap = map[int]string{
	actionUnknown:                   `Unknown`,
	actionDel:                       `Del`,
	actionRead:                      `Read`,
	actionWrite:                     `Write`,
	actionOpen:                      `Open`,
	actionMkdir:                     `Mkdir`,
	actionNFSMountSucc:              `NFSMountSucc`,
	actionNFSMountFail:              `NFSMountFail`,
	actionRename:                    `Rename`,
	actionLoginFail:                 `LoginFail`,
	actionLoginSucc:                 `LoginSucc`,
	actionLogout:                    `Logout`,
	actionNFSUmount:                 `NFSUmount`,
	actionCopy:                      `Copy`,
	actionMove:                      `Move`,
	actionAdd:                       `Add`,
	actionAuthFail:                  `AuthFail`,
	actionAuthOk:                    `AuthOk`,
	actionTrashRecovery:             `TrashRecovery`,
	actionAddTranscode:              `AddTranscode`,
	actionDelTranscode:              `DelTranscode`,
	actionUpdateTranscode:           `UpdateTranscode`,
	actionWatermark:                 `Watermark`,
	actionRotate:                    `Rotate`,
	actionAddThumbnail:              `AddThumbnail`,
	actionAddFiling:                 `AddFiling`,
	actionUpdateFiling:              `UpdateFiling`,
	actionDeleteFiling:              `DeleteFiling`,
	actionPauseFiling:               `PauseFiling`,
	actionStopFiling:                `StopFiling`,
	actionEject:                     `Eject`,
	actionEncrypt:                   `Encrypt`,
	actionDecrypt:                   `Decrypt`,
	actionConnActionCreateShareLink: `ConnActionCreateShareLink`,
	actionConnActionCompress:        `ConnActionCompress`,
	actionConnActionExtract:         `ConnActionExtract`,
	actionConnActionCloudconvert:    `ConnActionCloudconvert`,
	actionAny:                       `Any`,
}

const (
	serviceUnknown = 0
	serviceSamba   = 1
	serviceFtp     = 2
	serviceHttp    = 4
	serviceNfs     = 8
	serviceAfp     = 16
	serviceTelnet  = 32
	serviceSsh     = 64
	serviceIscsi   = 128
	serviceRadius  = 256
	serviceVpn     = 512
	serviceHttps   = 1024
	serviceAny     = 2047
)

var serviceMap = map[int]string{
	serviceUnknown: `Unknown`,
	serviceSamba:   `Samba`,
	serviceFtp:     `Ftp`,
	serviceHttp:    `Http`,
	serviceNfs:     `Nfs`,
	serviceAfp:     `Afp`,
	serviceTelnet:  `Telnet`,
	serviceSsh:     `Ssh`,
	serviceIscsi:   `Iscsi`,
	serviceRadius:  `Radius`,
	serviceVpn:     `Vpn`,
	serviceHttps:   `Https`,
	serviceAny:     `Any`,
}
