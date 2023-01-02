main.main STEXT size=118 args=0x0 locals=0x50 funcid=0x0 align=0x0
	0x0000 00000 (main.go:5)	TEXT	main.main(SB), ABIInternal, $80-0
	0x0000 00000 (main.go:5)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:5)	PCDATA	$0, $-2
	0x0004 00004 (main.go:5)	JLS	111
	0x0006 00006 (main.go:5)	PCDATA	$0, $-1
	0x0006 00006 (main.go:5)	SUBQ	$80, SP
	0x000a 00010 (main.go:5)	MOVQ	BP, 72(SP)
	0x000f 00015 (main.go:5)	LEAQ	72(SP), BP
	0x0014 00020 (main.go:5)	FUNCDATA	$0, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0014 00020 (main.go:5)	FUNCDATA	$1, gclocals·5aa34RaZcmo0NkRpBHp2fg==(SB)
	0x0014 00020 (main.go:5)	FUNCDATA	$2, main.main.stkobj(SB)
	0x0014 00020 (<unknown line number>)	NOP
	0x0014 00020 (main.go:8)	MOVUPS	X15, main..autotmp_15+56(SP)
	0x001a 00026 (main.go:9)	MOVL	$17, AX
	0x001f 00031 (main.go:9)	PCDATA	$1, $1
	0x001f 00031 (main.go:9)	NOP
	0x0020 00032 (main.go:9)	CALL	runtime.convT64(SB)
	0x0025 00037 (main.go:9)	LEAQ	type.int(SB), CX
	0x002c 00044 (main.go:9)	MOVQ	CX, main..autotmp_15+56(SP)
	0x0031 00049 (main.go:9)	MOVQ	AX, main..autotmp_15+64(SP)
	0x0036 00054 (<unknown line number>)	NOP
	0x0036 00054 ($GOROOT/src/fmt/print.go:213)	MOVQ	os.Stdout(SB), BX
	0x003d 00061 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0044 00068 ($GOROOT/src/fmt/print.go:213)	LEAQ	go.string."Assembly!: %d\n"(SB), CX
	0x004b 00075 ($GOROOT/src/fmt/print.go:213)	MOVL	$14, DI
	0x0050 00080 ($GOROOT/src/fmt/print.go:213)	LEAQ	main..autotmp_15+56(SP), SI
	0x0055 00085 ($GOROOT/src/fmt/print.go:213)	MOVL	$1, R8
	0x005b 00091 ($GOROOT/src/fmt/print.go:213)	MOVQ	R8, R9
	0x005e 00094 ($GOROOT/src/fmt/print.go:213)	PCDATA	$1, $0
	0x005e 00094 ($GOROOT/src/fmt/print.go:213)	NOP
	0x0060 00096 ($GOROOT/src/fmt/print.go:213)	CALL	fmt.Fprintf(SB)
	0x0065 00101 (main.go:10)	MOVQ	72(SP), BP
	0x006a 00106 (main.go:10)	ADDQ	$80, SP
	0x006e 00110 (main.go:10)	RET
	0x006f 00111 (main.go:10)	NOP
	0x006f 00111 (main.go:5)	PCDATA	$1, $-1
	0x006f 00111 (main.go:5)	PCDATA	$0, $-2
	0x006f 00111 (main.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x0074 00116 (main.go:5)	PCDATA	$0, $-1
	0x0074 00116 (main.go:5)	JMP	0
	0x0000 49 3b 66 10 76 69 48 83 ec 50 48 89 6c 24 48 48  I;f.viH..PH.l$HH
	0x0010 8d 6c 24 48 44 0f 11 7c 24 38 b8 11 00 00 00 90  .l$HD..|$8......
	0x0020 e8 00 00 00 00 48 8d 0d 00 00 00 00 48 89 4c 24  .....H......H.L$
	0x0030 38 48 89 44 24 40 48 8b 1d 00 00 00 00 48 8d 05  8H.D$@H......H..
	0x0040 00 00 00 00 48 8d 0d 00 00 00 00 bf 0e 00 00 00  ....H...........
	0x0050 48 8d 74 24 38 41 b8 01 00 00 00 4d 89 c1 66 90  H.t$8A.....M..f.
	0x0060 e8 00 00 00 00 48 8b 6c 24 48 48 83 c4 50 c3 e8  .....H.l$HH..P..
	0x0070 00 00 00 00 eb 8a                                ......
	rel 2+0 t=23 type.int+0
	rel 2+0 t=23 type.*os.File+0
	rel 33+4 t=7 runtime.convT64+0
	rel 40+4 t=14 type.int+0
	rel 57+4 t=14 os.Stdout+0
	rel 64+4 t=14 go.itab.*os.File,io.Writer+0
	rel 71+4 t=14 go.string."Assembly!: %d\n"+0
	rel 97+4 t=7 fmt.Fprintf+0
	rel 112+4 t=7 runtime.morestack_noctxt+0
main.add STEXT nosplit size=4 args=0x10 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (main.go:12)	TEXT	main.add(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (main.go:12)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:12)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:12)	FUNCDATA	$5, main.add.arginfo1(SB)
	0x0000 00000 (main.go:12)	FUNCDATA	$6, main.add.argliveinfo(SB)
	0x0000 00000 (main.go:12)	PCDATA	$3, $1
	0x0000 00000 (main.go:13)	ADDQ	BX, AX
	0x0003 00003 (main.go:13)	RET
	0x0000 48 01 d8 c3                                      H...
go.cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.main.add$abstract SDWARFABSFCN dupok size=29
	0x0000 05 6d 61 69 6e 2e 61 64 64 00 01 01 13 61 00 00  .main.add....a..
	0x0010 00 00 00 00 13 62 00 00 00 00 00 00 00           .....b.......
	rel 16+4 t=31 go.info.int+0
	rel 24+4 t=31 go.info.int+0
go.info.fmt.Printf$abstract SDWARFABSFCN dupok size=54
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 66 00 01 01 13 66  .fmt.Printf....f
	0x0010 6f 72 6d 61 74 00 00 00 00 00 00 13 61 00 00 00  ormat.......a...
	0x0020 00 00 00 13 6e 00 01 00 00 00 00 13 65 72 72 00  ....n.......err.
	0x0030 01 00 00 00 00 00                                ......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 0+0 t=22 type.string+0
	rel 23+4 t=31 go.info.string+0
	rel 31+4 t=31 go.info.[]interface {}+0
	rel 39+4 t=31 go.info.int+0
	rel 49+4 t=31 go.info.error+0
main..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.string."Assembly!: %d\n" SRODATA dupok size=14
	0x0000 41 73 73 65 6d 62 6c 79 21 3a 20 25 64 0a        Assembly!: %d.
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 5a 22 ee 60 00 00 00 00 00 00 00 00 00 00 00 00  Z".`............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3b fc f8 8f 08 08 08 36 00 00 00 00 00 00 00 00  ;......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 39 7a 09 0f 02 08 08 14 00 00 00 00 00 00 00 00  9z..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9d 9c 0e 59 08 08 08 36 00 00 00 00 00 00 00 00  ...Y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 76 de 99 0d 02 08 08 17 00 00 00 00 00 00 00 00  v...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=18
	0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface 
	0x0010 7b 7d                                            {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=-32763 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·5aa34RaZcmo0NkRpBHp2fg== SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 00 02                    ..........
main.main.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
main.add.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
main.add.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
