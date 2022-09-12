// go run mksysnum.go /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk/usr/include/sys/syscall.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build amd64,darwin

package unix

const (
	SYS_SYSCALL                        = 0
	SYS_EXIT                           = 1
	SYS_FORK                           = 2
	SYS_READ                           = 3
	SYS_WRITE                          = 4
	SYS_OPEN                           = 5
	SYS_CLOSE                          = 6
	SYS_WAIT4                          = 7
	SYS_LINK                           = 9
	SYS_UNLINK                         = 10
	SYS_CHDIR                          = 12
	SYS_FCHDIR                         = 13
	SYS_MKNOD                          = 14
	SYS_CHMOD                          = 15
	SYS_CHOWN                          = 16
	SYS_GETFSSTAT                      = 18
	SYS_GETPID                         = 20
	SYS_SETUID                         = 23
	SYS_GETUID                         = 24
	SYS_GETEUID                        = 25
	SYS_PTRACE                         = 26
	SYS_RECVMSG                        = 27
	SYS_SENDMSG                        = 28
	SYS_RECVFROM                       = 29
	SYS_ACCEPT                         = 30
	SYS_GETPEERNAME                    = 31
	SYS_GETSOCKNAME                    = 32
	SYS_ACCESS                         = 33
	SYS_CHFLAGS                        = 34
	SYS_FCHFLAGS                       = 35
	SYS_SYNC                           = 36
	SYS_KILL                           = 37
	SYS_GETPPID                        = 39
	SYS_DUP                            = 41
	SYS_PIPE                           = 42
	SYS_GETEGID                        = 43
	SYS_SIGACTION                      = 46
	SYS_GETGID                         = 47
	SYS_SIGPROCMASK                    = 48
	SYS_GETLOGIN                       = 49
	SYS_SETLOGIN                       = 50
	SYS_ACCT                           = 51
	SYS_SIGPENDING                     = 52
	SYS_SIGALTSTACK                    = 53
	SYS_IOCTL                          = 54
	SYS_REBOOT                         = 55
	SYS_REVOKE                         = 56
	SYS_SYMLINK                        = 57
	SYS_READLINK                       = 58
	SYS_EXECVE                         = 59
	SYS_UMASK                          = 60
	SYS_CHROOT                         = 61
	SYS_MSYNC                          = 65
	SYS_VFORK                          = 66
	SYS_MUNMAP                         = 73
	SYS_MPROTECT                       = 74
	SYS_MADVISE                        = 75
	SYS_MINCORE                        = 78
	SYS_GETGROUPS                      = 79
	SYS_SETGROUPS                      = 80
	SYS_GETPGRP                        = 81
	SYS_SETPGID                        = 82
	SYS_SETITIMER                      = 83
	SYS_SWAPON                         = 85
	SYS_GETITIMER                      = 86
	SYS_GETDTABLESIZE                  = 89
	SYS_DUP2                           = 90
	SYS_FCNTL                          = 92
	SYS_SELECT                         = 93
	SYS_FSYNC                          = 95
	SYS_SETPRIORITY                    = 96
	SYS_SOCKET                         = 97
	SYS_CONNECT                        = 98
	SYS_GETPRIORITY                    = 100
	SYS_BIND                           = 104
	SYS_SETSOCKOPT                     = 105
	SYS_LISTEN                         = 106
	SYS_SIGSUSPEND                     = 111
	SYS_GETTIMEOFDAY                   = 116
	SYS_GETRUSAGE                      = 117
	SYS_GETSOCKOPT                     = 118
	SYS_READV                          = 120
	SYS_WRITEV                         = 121
	SYS_SETTIMEOFDAY                   = 122
	SYS_FCHOWN                         = 123
	SYS_FCHMOD                         = 124
	SYS_SETREUID                       = 126
	SYS_SETREGID                       = 127
	SYS_RENAME                         = 128
	SYS_FLOCK                          = 131
	SYS_MKFIFO                         = 132
	SYS_SENDTO                         = 133
	SYS_SHUTDOWN                       = 134
	SYS_SOCKETPAIR                     = 135
	SYS_MKDIR                          = 136
	SYS_RMDIR                          = 137
	SYS_UTIMES                         = 138
	SYS_FUTIMES                        = 139
	SYS_ADJTIME                        = 140
	SYS_GETHOSTUUID                    = 142
	SYS_SETSID                         = 147
	SYS_GETPGID                        = 151
	SYS_SETPRIVEXEC                    = 152
	SYS_PREAD                          = 153
	SYS_PWRITE                         = 154
	SYS_NFSSVC                         = 155
	SYS_STATFS                         = 157
	SYS_FSTATFS                        = 158
	SYS_UNMOUNT                        = 159
	SYS_GETFH                          = 161
	SYS_QUOTACTL                       = 165
	SYS_MOUNT                          = 167
	SYS_CSOPS                          = 169
	SYS_CSOPS_AUDITTOKEN               = 170
	SYS_WAITID                         = 173
	SYS_KDEBUG_TYPEFILTER              = 177
	SYS_KDEBUG_TRACE_STRING            = 178
	SYS_KDEBUG_TRACE64                 = 179
	SYS_KDEBUG_TRACE                   = 180
	SYS_SETGID                         = 181
	SYS_SETEGID                        = 182
	SYS_SETEUID                        = 183
	SYS_SIGRETURN                      = 184
	SYS_THREAD_SELFCOUNTS              = 186
	SYS_FDATASYNC                      = 187
	SYS_STAT                           = 188
	SYS_FSTAT                          = 189
	SYS_LSTAT                          = 190
	SYS_PATHCONF                       = 191
	SYS_FPATHCONF                      = 192
	SYS_GETRLIMIT                      = 194
	SYS_SETRLIMIT                      = 195
	SYS_GETDIRENTRIES                  = 196
	SYS_MMAP                           = 197
	SYS_LSEEK                          = 199
	SYS_TRUNCATE                       = 200
	SYS_FTRUNCATE                      = 201
	SYS_SYSCTL                         = 202
	SYS_MLOCK                          = 203
	SYS_MUNLOCK                        = 204
	SYS_UNDELETE                       = 205
	SYS_OPEN_DPROTECTED_NP             = 216
	SYS_GETATTRLIST                    = 220
	SYS_SETATTRLIST                    = 221
	SYS_GETDIRENTRIESATTR              = 222
	SYS_EXCHANGEDATA                   = 223
	SYS_SEARCHFS                       = 225
	SYS_DELETE                         = 226
	SYS_COPYFILE                       = 227
	SYS_FGETATTRLIST                   = 228
	SYS_FSETATTRLIST                   = 229
	SYS_POLL                           = 230
	SYS_WATCHEVENT                     = 231
	SYS_WAITEVENT                      = 232
	SYS_MODWATCH                       = 233
	SYS_GETXATTR                       = 234
	SYS_FGETXATTR                      = 235
	SYS_SETXATTR                       = 236
	SYS_FSETXATTR                      = 237
	SYS_REMOVEXATTR                    = 238
	SYS_FREMOVEXATTR                   = 239
	SYS_LISTXATTR                      = 240
	SYS_FLISTXATTR                     = 241
	SYS_FSCTL                          = 242
	SYS_INITGROUPS                     = 243
	SYS_POSIX_SPAWN                    = 244
	SYS_FFSCTL                         = 245
	SYS_NFSCLNT                        = 247
	SYS_FHOPEN                         = 248
	SYS_MINHERIT                       = 250
	SYS_SEMSYS                         = 251
	SYS_MSGSYS                         = 252
	SYS_SHMSYS                         = 253
	SYS_SEMCTL                         = 254
	SYS_SEMGET                         = 255
	SYS_SEMOP                          = 256
	SYS_MSGCTL                         = 258
	SYS_MSGGET                         = 259
	SYS_MSGSND                         = 260
	SYS_MSGRCV                         = 261
	SYS_SHMAT                          = 262
	SYS_SHMCTL                         = 263
	SYS_SHMDT                          = 264
	SYS_SHMGET                         = 265
	SYS_SHM_OPEN                       = 266
	SYS_SHM_UNLINK                     = 267
	SYS_SEM_OPEN                       = 268
	SYS_SEM_CLOSE                      = 269
	SYS_SEM_UNLINK                     = 270
	SYS_SEM_WAIT                       = 271
	SYS_SEM_TRYWAIT                    = 272
	SYS_SEM_POST                       = 273
	SYS_SYSCTLBYNAME                   = 274
	SYS_OPEN_EXTENDED                  = 277
	SYS_UMASK_EXTENDED                 = 278
	SYS_STAT_EXTENDED                  = 279
	SYS_LSTAT_EXTENDED                 = 280
	SYS_FSTAT_EXTENDED                 = 281
	SYS_CHMOD_EXTENDED                 = 282
	SYS_FCHMOD_EXTENDED                = 283
	SYS_ACCESS_EXTENDED                = 284
	SYS_SETTID                         = 285
	SYS_GETTID                         = 286
	SYS_SETSGROUPS                     = 287
	SYS_GETSGROUPS                     = 288
	SYS_SETWGROUPS                     = 289
	SYS_GETWGROUPS                     = 290
	SYS_MKFIFO_EXTENDED                = 291
	SYS_MKDIR_EXTENDED                 = 292
	SYS_IDENTITYSVC                    = 293
	SYS_SHARED_REGION_CHECK_NP         = 294
	SYS_VM_PRESSURE_MONITOR            = 296
	SYS_PSYNCH_RW_LONGRDLOCK           = 297
	SYS_PSYNCH_RW_YIELDWRLOCK          = 298
	SYS_PSYNCH_RW_DOWNGRADE            = 299
	SYS_PSYNCH_RW_UPGRADE              = 300
	SYS_PSYNCH_MUTEXWAIT               = 301
	SYS_PSYNCH_MUTEXDROP               = 302
	SYS_PSYNCH_CVBROAD                 = 303
	SYS_PSYNCH_CVSIGNAL                = 304
	SYS_PSYNCH_CVWAIT                  = 305
	SYS_PSYNCH_RW_RDLOCK               = 306
	SYS_PSYNCH_RW_WRLOCK               = 307
	SYS_PSYNCH_RW_UNLOCK               = 308
	SYS_PSYNCH_RW_UNLOCK2              = 309
	SYS_GETSID                         = 310
	SYS_SETTID_WITH_PID                = 311
	SYS_PSYNCH_CVCLRPREPOST            = 312
	SYS_AIO_FSYNC                      = 313
	SYS_AIO_RETURN                     = 314
	SYS_AIO_SUSPEND                    = 315
	SYS_AIO_CANCEL                     = 316
	SYS_AIO_ERROR                      = 317
	SYS_AIO_READ                       = 318
	SYS_AIO_WRITE                      = 319
	SYS_LIO_LISTIO                     = 320
	SYS_IOPOLICYSYS                    = 322
	SYS_PROCESS_POLICY                 = 323
	SYS_MLOCKALL                       = 324
	SYS_MUNLOCKALL                     = 325
	SYS_ISSETUGID                      = 327
	SYS___PTHREAD_KILL                 = 328
	SYS___PTHREAD_SIGMASK              = 329
	SYS___SIGWAIT                      = 330
	SYS___DISABLE_THREADSIGNAL         = 331
	SYS___PTHREAD_MARKCANCEL           = 332
	SYS___PTHREAD_CANCELED             = 333
	SYS___SEMWAIT_SIGNAL               = 334
	SYS_PROC_INFO                      = 336
	SYS_SENDFILE                       = 337
	SYS_STAT64                         = 338
	SYS_FSTAT64                        = 339
	SYS_LSTAT64                        = 340
	SYS_STAT64_EXTENDED                = 341
	SYS_LSTAT64_EXTENDED               = 342
	SYS_FSTAT64_EXTENDED               = 343
	SYS_GETDIRENTRIES64                = 344
	SYS_STATFS64                       = 345
	SYS_FSTATFS64                      = 346
	SYS_GETFSSTAT64                    = 347
	SYS___PTHREAD_CHDIR                = 348
	SYS___PTHREAD_FCHDIR               = 349
	SYS_AUDIT                          = 350
	SYS_AUDITON                        = 351
	SYS_GETAUID                        = 353
	SYS_SETAUID                        = 354
	SYS_GETAUDIT_ADDR                  = 357
	SYS_SETAUDIT_ADDR                  = 358
	SYS_AUDITCTL                       = 359
	SYS_BSDTHREAD_CREATE               = 360
	SYS_BSDTHREAD_TERMINATE            = 361
	SYS_KQUEUE                         = 362
	SYS_KEVENT                         = 363
	SYS_LCHOWN                         = 364
	SYS_BSDTHREAD_REGISTER             = 366
	SYS_WORKQ_OPEN                     = 367
	SYS_WORKQ_KERNRETURN               = 368
	SYS_KEVENT64                       = 369
	SYS___OLD_SEMWAIT_SIGNAL           = 370
	SYS___OLD_SEMWAIT_SIGNAL_NOCANCEL  = 371
	SYS_THREAD_SELFID                  = 372
	SYS_LEDGER                         = 373
	SYS_KEVENT_QOS                     = 374
	SYS_KEVENT_ID                      = 375
	SYS___MAC_EXECVE                   = 380
	SYS___MAC_SYSCALL                  = 381
	SYS___MAC_GET_FILE                 = 382
	SYS___MAC_SET_FILE                 = 383
	SYS___MAC_GET_LINK                 = 384
	SYS___MAC_SET_LINK                 = 385
	SYS___MAC_GET_PROC                 = 386
	SYS___MAC_SET_PROC                 = 387
	SYS___MAC_GET_FD                   = 388
	SYS___MAC_SET_FD                   = 389
	SYS___MAC_GET_PID                  = 390
	SYS_PSELECT                        = 394
	SYS_PSELECT_NOCANCEL               = 395
	SYS_READ_NOCANCEL                  = 396
	SYS_WRITE_NOCANCEL                 = 397
	SYS_OPEN_NOCANCEL                  = 398
	SYS_CLOSE_NOCANCEL                 = 399
	SYS_WAIT4_NOCANCEL                 = 400
	SYS_RECVMSG_NOCANCEL               = 401
	SYS_SENDMSG_NOCANCEL               = 402
	SYS_RECVFROM_NOCANCEL              = 403
	SYS_ACCEPT_NOCANCEL                = 404
	SYS_MSYNC_NOCANCEL                 = 405
	SYS_FCNTL_NOCANCEL                 = 406
	SYS_SELECT_NOCANCEL                = 407
	SYS_FSYNC_NOCANCEL                 = 408
	SYS_CONNECT_NOCANCEL               = 409
	SYS_SIGSUSPEND_NOCANCEL            = 410
	SYS_READV_NOCANCEL                 = 411
	SYS_WRITEV_NOCANCEL                = 412
	SYS_SENDTO_NOCANCEL                = 413
	SYS_PREAD_NOCANCEL                 = 414
	SYS_PWRITE_NOCANCEL                = 415
	SYS_WAITID_NOCANCEL                = 416
	SYS_POLL_NOCANCEL                  = 417
	SYS_MSGSND_NOCANCEL                = 418
	SYS_MSGRCV_NOCANCEL                = 419
	SYS_SEM_WAIT_NOCANCEL              = 420
	SYS_AIO_SUSPEND_NOCANCEL           = 421
	SYS___SIGWAIT_NOCANCEL             = 422
	SYS___SEMWAIT_SIGNAL_NOCANCEL      = 423
	SYS___MAC_MOUNT                    = 424
	SYS___MAC_GET_MOUNT                = 425
	SYS___MAC_GETFSSTAT                = 426
	SYS_FSGETPATH                      = 427
	SYS_AUDIT_SESSION_SELF             = 428
	SYS_AUDIT_SESSION_JOIN             = 429
	SYS_FILEPORT_MAKEPORT              = 430
	SYS_FILEPORT_MAKEFD                = 431
	SYS_AUDIT_SESSION_PORT             = 432
	SYS_PID_SUSPEND                    = 433
	SYS_PID_RESUME                     = 434
	SYS_PID_HIBERNATE                  = 435
	SYS_PID_SHUTDOWN_SOCKETS           = 436
	SYS_SHARED_REGION_MAP_AND_SLIDE_NP = 438
	SYS_KAS_INFO                       = 439
	SYS_MEMORYSTATUS_CONTROL           = 440
	SYS_GUARDED_OPEN_NP                = 441
	SYS_GUARDED_CLOSE_NP               = 442
	SYS_GUARDED_KQUEUE_NP              = 443
	SYS_CHANGE_FDGUARD_NP              = 444
	SYS_USRCTL                         = 445
	SYS_PROC_RLIMIT_CONTROL            = 446
	SYS_CONNECTX                       = 447
	SYS_DISCONNECTX                    = 448
	SYS_PEELOFF                        = 449
	SYS_SOCKET_DELEGATE                = 450
	SYS_TELEMETRY                      = 451
	SYS_PROC_UUID_POLICY               = 452
	SYS_MEMORYSTATUS_GET_LEVEL         = 453
	SYS_SYSTEM_OVERRIDE                = 454
	SYS_VFS_PURGE                      = 455
	SYS_SFI_CTL                        = 456
	SYS_SFI_PIDCTL                     = 457
	SYS_COALITION                      = 458
	SYS_COALITION_INFO                 = 459
	SYS_NECP_MATCH_POLICY              = 460
	SYS_GETATTRLISTBULK                = 461
	SYS_CLONEFILEAT                    = 462
	SYS_OPENAT                         = 463
	SYS_OPENAT_NOCANCEL                = 464
	SYS_RENAMEAT                       = 465
	SYS_FACCESSAT                      = 466
	SYS_FCHMODAT                       = 467
	SYS_FCHOWNAT                       = 468
	SYS_FSTATAT                        = 469
	SYS_FSTATAT64                      = 470
	SYS_LINKAT                         = 471
	SYS_UNLINKAT                       = 472
	SYS_READLINKAT                     = 473
	SYS_SYMLINKAT                      = 474
	SYS_MKDIRAT                        = 475
	SYS_GETATTRLISTAT                  = 476
	SYS_PROC_TRACE_LOG                 = 477
	SYS_BSDTHREAD_CTL                  = 478
	SYS_OPENBYID_NP                    = 479
	SYS_RECVMSG_X                      = 480
	SYS_SENDMSG_X                      = 481
	SYS_THREAD_SELFUSAGE               = 482
	SYS_CSRCTL                         = 483
	SYS_GUARDED_OPEN_DPROTECTED_NP     = 484
	SYS_GUARDED_WRITE_NP               = 485
	SYS_GUARDED_PWRITE_NP              = 486
	SYS_GUARDED_WRITEV_NP              = 487
	SYS_RENAMEATX_NP                   = 488
	SYS_MREMAP_ENCRYPTED               = 489
	SYS_NETAGENT_TRIGGER               = 490
	SYS_STACK_SNAPSHOT_WITH_CONFIG     = 491
	SYS_MICROSTACKSHOT                 = 492
	SYS_GRAB_PGO_DATA                  = 493
	SYS_PERSONA                        = 494
	SYS_WORK_INTERVAL_CTL              = 499
	SYS_GETENTROPY                     = 500
	SYS_NECP_OPEN                      = 501
	SYS_NECP_CLIENT_ACTION             = 502
	SYS___NEXUS_OPEN                   = 503
	SYS___NEXUS_REGISTER               = 504
	SYS___NEXUS_DEREGISTER             = 505
	SYS___NEXUS_CREATE                 = 506
	SYS___NEXUS_DESTROY                = 507
	SYS___NEXUS_GET_OPT                = 508
	SYS___NEXUS_SET_OPT                = 509
	SYS___CHANNEL_OPEN                 = 510
	SYS___CHANNEL_GET_INFO             = 511
	SYS___CHANNEL_SYNC                 = 512
	SYS___CHANNEL_GET_OPT              = 513
	SYS___CHANNEL_SET_OPT              = 514
	SYS_ULOCK_WAIT                     = 515
	SYS_ULOCK_WAKE                     = 516
	SYS_FCLONEFILEAT                   = 517
	SYS_FS_SNAPSHOT                    = 518
	SYS_TERMINATE_WITH_PAYLOAD         = 520
	SYS_ABORT_WITH_PAYLOAD             = 521
	SYS_NECP_SESSION_OPEN              = 522
	SYS_NECP_SESSION_ACTION            = 523
	SYS_SETATTRLISTAT                  = 524
	SYS_NET_QOS_GUIDELINE              = 525
	SYS_FMOUNT                         = 526
	SYS_NTP_ADJTIME                    = 527
	SYS_NTP_GETTIME                    = 528
	SYS_OS_FAULT_WITH_PAYLOAD          = 529
	SYS_KQUEUE_WORKLOOP_CTL            = 530
	SYS___MACH_BRIDGE_REMOTE_TIME      = 531
	SYS_MAXSYSCALL                     = 532
	SYS_INVALID                        = 63
)
