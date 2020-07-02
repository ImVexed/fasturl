
//line parser.rl:1
// fasturl is a Go URL parser using a [Ragel](http://www.colm.net/open-source/ragel/) state-machine instead of regex, or the built in standard library `url.Parse`.
// 
package fasturl

import "fmt"


//line parser.go:11
const url_parser_start int = 45
const url_parser_first_final int = 45
const url_parser_error int = 0

const url_parser_en_main int = 45


//line parser.rl:96


// URL represents the different parts of a parsed URL
type URL struct {
  Protocol string
  Host     string
  Port     string
  Path     string
  Query    string
  Fragment string
}

// ParseURL parses a given URL and returns a `URL` representing the different parts
func ParseURL(data string) (*URL, error){
  mark, host_mark, port_mark, cs, p, pe, eof := 0, 0, 0, url_parser_en_main, 0, len(data), len(data)

  u := &URL{}
  
  
//line parser.go:39
	{
	cs = url_parser_start
	}

//line parser.rl:115
  
//line parser.go:46
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 45:
		goto st_case_45
	case 0:
		goto st_case_0
	case 46:
		goto st_case_46
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 3:
		goto st_case_3
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 58:
		goto st_case_58
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 85:
		goto st_case_85
	case 23:
		goto st_case_23
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 88:
		goto st_case_88
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 36:
		goto st_case_36
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 110:
		goto st_case_110
	case 111:
		goto st_case_111
	case 112:
		goto st_case_112
	case 113:
		goto st_case_113
	case 114:
		goto st_case_114
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 42:
		goto st_case_42
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 123:
		goto st_case_123
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 131:
		goto st_case_131
	case 132:
		goto st_case_132
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	case 141:
		goto st_case_141
	case 142:
		goto st_case_142
	case 143:
		goto st_case_143
	case 144:
		goto st_case_144
	case 145:
		goto st_case_145
	case 146:
		goto st_case_146
	case 147:
		goto st_case_147
	case 148:
		goto st_case_148
	case 149:
		goto st_case_149
	case 150:
		goto st_case_150
	case 151:
		goto st_case_151
	case 152:
		goto st_case_152
	case 153:
		goto st_case_153
	case 154:
		goto st_case_154
	case 155:
		goto st_case_155
	case 156:
		goto st_case_156
	case 157:
		goto st_case_157
	case 158:
		goto st_case_158
	case 159:
		goto st_case_159
	case 160:
		goto st_case_160
	case 161:
		goto st_case_161
	}
	goto st_out
	st_case_45:
		switch data[p] {
		case 33:
			goto tr14
		case 37:
			goto tr15
		case 47:
			goto tr43
		case 58:
			goto st8
		case 61:
			goto tr14
		case 64:
			goto st11
		case 91:
			goto tr16
		case 92:
			goto tr45
		case 95:
			goto tr14
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr14:
//line parser.rl:23

    host_mark = p;
  
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line parser.go:429
		switch data[p] {
		case 33:
			goto st46
		case 37:
			goto st1
		case 47:
			goto tr47
		case 58:
			goto tr48
		case 61:
			goto st46
		case 64:
			goto st11
		case 92:
			goto tr47
		case 95:
			goto st46
		case 126:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr49
			}
		default:
			goto tr49
		}
		goto st0
tr15:
//line parser.rl:23

    host_mark = p;
  
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line parser.go:474
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st2
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st46
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st46
			}
		default:
			goto st46
		}
		goto st0
tr45:
//line parser.rl:10
 mark = p 
	goto st47
tr47:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:10
 mark = p 
	goto st47
tr61:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st47
tr136:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line parser.go:547
		switch data[p] {
		case 35:
			goto tr51
		case 63:
			goto tr52
		}
		switch {
		case data[p] > 46:
			if 48 <= data[p] && data[p] <= 126 {
				goto st48
			}
		case data[p] >= 32:
			goto st48
		}
		goto st0
tr78:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line parser.go:578
		switch data[p] {
		case 35:
			goto tr51
		case 63:
			goto tr52
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto st48
		}
		goto st0
tr51:
//line parser.rl:37

    if u.Path == "" {
      u.Path = data[mark:p]
    }
  
	goto st49
tr56:
//line parser.rl:10
 mark = p 
//line parser.rl:31

    if u.Query == "" {
      u.Query = data[mark:p]
    }
  
	goto st49
tr58:
//line parser.rl:31

    if u.Query == "" {
      u.Query = data[mark:p]
    }
  
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line parser.go:620
		if 32 <= data[p] && data[p] <= 126 {
			goto tr53
		}
		goto st0
tr53:
//line parser.rl:10
 mark = p 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line parser.go:634
		if 32 <= data[p] && data[p] <= 126 {
			goto st50
		}
		goto st0
tr52:
//line parser.rl:37

    if u.Path == "" {
      u.Path = data[mark:p]
    }
  
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line parser.go:652
		if data[p] == 35 {
			goto tr56
		}
		switch {
		case data[p] < 48:
			if 32 <= data[p] && data[p] <= 46 {
				goto tr55
			}
		case data[p] > 62:
			if 64 <= data[p] && data[p] <= 126 {
				goto tr55
			}
		default:
			goto tr55
		}
		goto st0
tr55:
//line parser.rl:10
 mark = p 
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line parser.go:678
		if data[p] == 35 {
			goto tr58
		}
		switch {
		case data[p] < 48:
			if 32 <= data[p] && data[p] <= 46 {
				goto st52
			}
		case data[p] > 62:
			if 64 <= data[p] && data[p] <= 126 {
				goto st52
			}
		default:
			goto st52
		}
		goto st0
tr48:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line parser.go:706
		switch data[p] {
		case 33:
			goto tr4
		case 37:
			goto tr5
		case 60:
			goto tr3
		case 62:
			goto tr3
		case 64:
			goto tr6
		case 95:
			goto tr4
		case 126:
			goto tr4
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto tr3
				}
			case data[p] > 46:
				if 48 <= data[p] && data[p] <= 61 {
					goto tr4
				}
			default:
				goto tr4
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto tr3
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr3
				}
			default:
				goto tr4
			}
		default:
			goto tr4
		}
		goto st0
tr3:
//line parser.rl:11
 port_mark = p 
	goto st53
tr121:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st53
tr201:
//line parser.rl:11
 port_mark = p 
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line parser.go:785
		switch data[p] {
		case 37:
			goto st74
		case 47:
			goto tr61
		case 64:
			goto st54
		case 92:
			goto tr63
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st54
				}
			case data[p] >= 32:
				goto st54
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st54
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st54
				}
			default:
				goto tr62
			}
		default:
			goto tr62
		}
		goto st0
tr62:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line parser.go:838
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st55
				}
			case data[p] >= 32:
				goto st55
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st55
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st55
				}
			default:
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
tr66:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line parser.go:891
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st56
				}
			case data[p] >= 32:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st56
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st56
				}
			default:
				goto tr70
			}
		default:
			goto tr70
		}
		goto st0
tr70:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st56
tr239:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line parser.go:958
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st57
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st57
				}
			default:
				goto tr73
			}
		default:
			goto tr73
		}
		goto st0
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr75
			}
		case data[p] >= 65:
			goto tr75
		}
		goto st0
tr117:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:10
 mark = p 
	goto st4
tr75:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st4
tr190:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line parser.go:1053
		if data[p] == 58 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 47:
			goto st47
		case 92:
			goto st47
		}
		goto st0
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr77
				}
			case data[p] >= 48:
				goto st6
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr77
			}
		default:
			goto tr75
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st57
			}
		default:
			goto st57
		}
		goto st0
tr77:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line parser.go:1137
		if data[p] == 58 {
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st57
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st57
			}
		default:
			goto st57
		}
		goto st0
tr73:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st59
tr200:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line parser.go:1183
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr75
			}
		case data[p] >= 65:
			goto tr75
		}
		goto st0
tr74:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st60
tr146:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line parser.go:1230
		switch data[p] {
		case 35:
			goto tr51
		case 47:
			goto tr61
		case 63:
			goto tr52
		case 91:
			goto st48
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr78
				}
			case data[p] >= 32:
				goto st48
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st48
				}
			case data[p] >= 97:
				goto tr78
			}
		default:
			goto st48
		}
		goto st0
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st62
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr80
			}
		default:
			goto tr80
		}
		goto st0
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr70
				}
			case data[p] >= 48:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr70
			}
		default:
			goto tr75
		}
		goto st0
tr80:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line parser.go:1374
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr70
				}
			case data[p] >= 48:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr70
			}
		default:
			goto tr75
		}
		goto st0
tr71:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st64
tr143:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line parser.go:1435
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st66
		case 47:
			goto tr61
		case 63:
			goto tr52
		case 91:
			goto st65
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr83
				}
			case data[p] >= 32:
				goto st65
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st65
				}
			case data[p] >= 97:
				goto tr83
			}
		default:
			goto st65
		}
		goto st0
tr83:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line parser.go:1486
		switch data[p] {
		case 35:
			goto tr51
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st48
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr78
				}
			case data[p] >= 32:
				goto st48
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st48
				}
			case data[p] >= 97:
				goto tr78
			}
		default:
			goto st48
		}
		goto st0
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
		switch data[p] {
		case 35:
			goto tr51
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st48
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st48
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr85
					}
				case data[p] >= 58:
					goto st48
				}
			default:
				goto st67
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st48
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st48
					}
				case data[p] >= 103:
					goto tr78
				}
			default:
				goto tr85
			}
		default:
			goto tr78
		}
		goto st0
tr85:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line parser.go:1591
		switch data[p] {
		case 35:
			goto tr51
		case 63:
			goto tr52
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 47 {
					goto st48
				}
			case data[p] > 57:
				if 58 <= data[p] && data[p] <= 64 {
					goto st48
				}
			default:
				goto st65
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				if 71 <= data[p] && data[p] <= 96 {
					goto st48
				}
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 126 {
					goto st48
				}
			default:
				goto st65
			}
		default:
			goto st65
		}
		goto st0
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st69
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr87
			}
		default:
			goto tr87
		}
		goto st0
tr87:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line parser.go:1705
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st55
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
tr67:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st70
tr139:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line parser.go:1791
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st72
		case 47:
			goto tr61
		case 63:
			goto tr52
		case 91:
			goto st71
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr90
				}
			case data[p] >= 32:
				goto st71
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st71
				}
			case data[p] >= 97:
				goto tr90
			}
		default:
			goto st71
		}
		goto st0
tr90:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line parser.go:1842
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st66
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st65
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr83
				}
			case data[p] >= 32:
				goto st65
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st65
				}
			case data[p] >= 97:
				goto tr83
			}
		default:
			goto st65
		}
		goto st0
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st66
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st65
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st65
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr92
					}
				case data[p] >= 58:
					goto st65
				}
			default:
				goto st73
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st65
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st65
					}
				case data[p] >= 103:
					goto tr83
				}
			default:
				goto tr92
			}
		default:
			goto tr83
		}
		goto st0
tr92:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line parser.go:1951
		switch data[p] {
		case 35:
			goto tr51
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st48
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st48
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr90
					}
				case data[p] >= 58:
					goto st48
				}
			default:
				goto st71
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st48
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st48
					}
				case data[p] >= 103:
					goto tr78
				}
			default:
				goto tr90
			}
		default:
			goto tr78
		}
		goto st0
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st75
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr94
			}
		default:
			goto tr94
		}
		goto st0
tr94:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line parser.go:2079
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st54
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr62
			}
		default:
			goto tr62
		}
		goto st0
tr63:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st76
tr227:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st76
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
//line parser.go:2165
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st78
		case 47:
			goto tr61
		case 63:
			goto tr52
		case 91:
			goto st77
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr97
				}
			case data[p] >= 32:
				goto st77
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st77
				}
			case data[p] >= 97:
				goto tr97
			}
		default:
			goto st77
		}
		goto st0
tr97:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st77
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
//line parser.go:2216
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st72
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st71
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr90
				}
			case data[p] >= 32:
				goto st71
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st71
				}
			case data[p] >= 97:
				goto tr90
			}
		default:
			goto st71
		}
		goto st0
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st72
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st71
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st71
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr99
					}
				case data[p] >= 58:
					goto st71
				}
			default:
				goto st79
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st71
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st71
					}
				case data[p] >= 103:
					goto tr90
				}
			default:
				goto tr99
			}
		default:
			goto tr90
		}
		goto st0
tr99:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st79
	st79:
		if p++; p == pe {
			goto _test_eof79
		}
	st_case_79:
//line parser.go:2325
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st66
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st65
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st65
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr97
					}
				case data[p] >= 58:
					goto st65
				}
			default:
				goto st77
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st65
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st65
					}
				case data[p] >= 103:
					goto tr83
				}
			default:
				goto tr97
			}
		default:
			goto tr83
		}
		goto st0
tr4:
//line parser.rl:11
 port_mark = p 
	goto st80
tr129:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st80
tr150:
//line parser.rl:11
 port_mark = p 
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st80
tr220:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:11
 port_mark = p 
	goto st80
	st80:
		if p++; p == pe {
			goto _test_eof80
		}
	st_case_80:
//line parser.go:2418
		switch data[p] {
		case 33:
			goto st81
		case 37:
			goto st142
		case 47:
			goto tr61
		case 60:
			goto st54
		case 62:
			goto st54
		case 64:
			goto st119
		case 92:
			goto tr63
		case 95:
			goto st81
		case 126:
			goto st81
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st81
				}
			case data[p] >= 32:
				goto st54
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st54
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st54
				}
			default:
				goto tr103
			}
		default:
			goto tr103
		}
		goto st0
tr103:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st81
	st81:
		if p++; p == pe {
			goto _test_eof81
		}
	st_case_81:
//line parser.go:2481
		switch data[p] {
		case 33:
			goto st82
		case 37:
			goto st140
		case 47:
			goto tr61
		case 60:
			goto st55
		case 62:
			goto st55
		case 64:
			goto st112
		case 92:
			goto tr67
		case 95:
			goto st82
		case 126:
			goto st82
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st82
				}
			case data[p] >= 32:
				goto st55
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st55
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st55
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto st0
tr107:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st82
	st82:
		if p++; p == pe {
			goto _test_eof82
		}
	st_case_82:
//line parser.go:2544
		switch data[p] {
		case 33:
			goto st83
		case 37:
			goto st137
		case 47:
			goto tr61
		case 60:
			goto st56
		case 62:
			goto st56
		case 64:
			goto st108
		case 92:
			goto tr71
		case 95:
			goto st83
		case 126:
			goto st83
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st83
				}
			case data[p] >= 32:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st56
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st56
				}
			default:
				goto tr111
			}
		default:
			goto tr111
		}
		goto st0
tr111:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st83
	st83:
		if p++; p == pe {
			goto _test_eof83
		}
	st_case_83:
//line parser.go:2607
		switch data[p] {
		case 33:
			goto st84
		case 37:
			goto st88
		case 47:
			goto tr61
		case 60:
			goto st57
		case 62:
			goto st57
		case 64:
			goto st89
		case 92:
			goto tr74
		case 95:
			goto st84
		case 126:
			goto st84
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st84
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st57
				}
			default:
				goto tr114
			}
		default:
			goto tr114
		}
		goto st0
	st84:
		if p++; p == pe {
			goto _test_eof84
		}
	st_case_84:
		switch data[p] {
		case 33:
			goto st8
		case 37:
			goto st9
		case 47:
			goto tr61
		case 61:
			goto st8
		case 64:
			goto st11
		case 92:
			goto tr61
		case 95:
			goto st8
		case 126:
			goto st8
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 33:
			goto st8
		case 37:
			goto st9
		case 61:
			goto st8
		case 64:
			goto st11
		case 95:
			goto st8
		case 126:
			goto st8
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st8
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st8
				}
			case data[p] >= 65:
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st10
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st10
			}
		default:
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st8
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 33:
			goto tr14
		case 37:
			goto tr15
		case 58:
			goto st8
		case 61:
			goto tr14
		case 64:
			goto st11
		case 91:
			goto tr16
		case 95:
			goto tr14
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto tr14
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr14
				}
			case data[p] >= 65:
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
tr16:
//line parser.rl:23

    host_mark = p;
  
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line parser.go:2816
		if data[p] == 118 {
			goto st28
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto st13
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 58:
			goto st13
		case 93:
			goto st85
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 46:
			goto st15
		case 58:
			goto st13
		case 93:
			goto st85
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if 48 <= data[p] && data[p] <= 57 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 46 {
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st26
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if 48 <= data[p] && data[p] <= 57 {
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 46 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if 48 <= data[p] && data[p] <= 57 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 93 {
			goto st85
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 93 {
			goto st85
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 93 {
			goto st85
		}
		goto st0
	st85:
		if p++; p == pe {
			goto _test_eof85
		}
	st_case_85:
		switch data[p] {
		case 47:
			goto tr47
		case 58:
			goto tr116
		case 92:
			goto tr47
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr117
			}
		case data[p] >= 65:
			goto tr117
		}
		goto st0
tr116:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line parser.go:3000
		if data[p] == 37 {
			goto tr31
		}
		switch {
		case data[p] < 36:
			if 32 <= data[p] && data[p] <= 34 {
				goto tr3
			}
		case data[p] > 46:
			switch {
			case data[p] > 62:
				if 64 <= data[p] && data[p] <= 126 {
					goto tr3
				}
			case data[p] >= 48:
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
tr31:
//line parser.rl:11
 port_mark = p 
	goto st86
	st86:
		if p++; p == pe {
			goto _test_eof86
		}
	st_case_86:
//line parser.go:3031
		switch data[p] {
		case 37:
			goto st74
		case 47:
			goto tr61
		case 64:
			goto st54
		case 92:
			goto tr63
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st54
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st54
					}
				case data[p] >= 48:
					goto st87
				}
			default:
				goto st54
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st54
					}
				case data[p] >= 71:
					goto tr62
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st54
					}
				case data[p] >= 103:
					goto tr62
				}
			default:
				goto tr119
			}
		default:
			goto tr119
		}
		goto st0
tr119:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st87
	st87:
		if p++; p == pe {
			goto _test_eof87
		}
	st_case_87:
//line parser.go:3103
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st53
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr121
			}
		default:
			goto tr121
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if data[p] == 46 {
			goto st19
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 46 {
			goto st19
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 46 {
			goto st17
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 46 {
			goto st17
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 46 {
			goto st30
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 61:
			goto st31
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 48:
			if 38 <= data[p] && data[p] <= 46 {
				goto st31
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st31
				}
			case data[p] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 61:
			goto st31
		case 93:
			goto st85
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 48:
			if 38 <= data[p] && data[p] <= 46 {
				goto st31
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st31
				}
			case data[p] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
tr115:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line parser.go:3328
		switch data[p] {
		case 33:
			goto st8
		case 37:
			goto st9
		case 58:
			goto st33
		case 61:
			goto st8
		case 64:
			goto st11
		case 95:
			goto st8
		case 126:
			goto st8
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto st8
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st8
				}
			case data[p] >= 65:
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 33:
			goto st8
		case 37:
			goto st9
		case 47:
			goto st47
		case 61:
			goto st8
		case 64:
			goto st11
		case 92:
			goto st47
		case 95:
			goto st8
		case 126:
			goto st8
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st8
			}
		default:
			goto st8
		}
		goto st0
	st88:
		if p++; p == pe {
			goto _test_eof88
		}
	st_case_88:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr123
				}
			case data[p] >= 48:
				goto st34
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr123
			}
		default:
			goto tr75
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st84
			}
		default:
			goto st84
		}
		goto st0
tr123:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line parser.go:3466
		if data[p] == 58 {
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st84
			}
		default:
			goto st84
		}
		goto st0
	st89:
		if p++; p == pe {
			goto _test_eof89
		}
	st_case_89:
		switch data[p] {
		case 33:
			goto tr14
		case 37:
			goto tr15
		case 47:
			goto tr61
		case 58:
			goto st8
		case 61:
			goto tr14
		case 64:
			goto st11
		case 91:
			goto tr16
		case 92:
			goto tr61
		case 95:
			goto tr14
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr124
			}
		default:
			goto tr124
		}
		goto st0
tr49:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:10
 mark = p 
	goto st90
tr124:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st90
tr147:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st90
tr213:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:10
 mark = p 
	goto st90
	st90:
		if p++; p == pe {
			goto _test_eof90
		}
	st_case_90:
//line parser.go:3572
		switch data[p] {
		case 33:
			goto st46
		case 37:
			goto st1
		case 47:
			goto tr47
		case 58:
			goto tr125
		case 61:
			goto st46
		case 64:
			goto st11
		case 92:
			goto tr47
		case 95:
			goto st46
		case 126:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr49
			}
		default:
			goto tr49
		}
		goto st0
tr125:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line parser.go:3617
		switch data[p] {
		case 33:
			goto tr4
		case 37:
			goto tr5
		case 47:
			goto st47
		case 60:
			goto tr3
		case 62:
			goto tr3
		case 64:
			goto tr6
		case 92:
			goto tr39
		case 95:
			goto tr4
		case 126:
			goto tr4
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr4
				}
			case data[p] >= 32:
				goto tr3
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto tr3
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr3
				}
			default:
				goto tr4
			}
		default:
			goto tr4
		}
		goto st0
tr5:
//line parser.rl:11
 port_mark = p 
	goto st91
	st91:
		if p++; p == pe {
			goto _test_eof91
		}
	st_case_91:
//line parser.go:3674
		switch data[p] {
		case 37:
			goto st74
		case 47:
			goto tr61
		case 64:
			goto st54
		case 92:
			goto tr63
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st54
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st54
					}
				case data[p] >= 48:
					goto st92
				}
			default:
				goto st54
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st54
					}
				case data[p] >= 71:
					goto tr62
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st54
					}
				case data[p] >= 103:
					goto tr62
				}
			default:
				goto tr127
			}
		default:
			goto tr127
		}
		goto st0
tr127:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st92
	st92:
		if p++; p == pe {
			goto _test_eof92
		}
	st_case_92:
//line parser.go:3746
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st80
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr129
			}
		default:
			goto tr129
		}
		goto st0
tr6:
//line parser.rl:11
 port_mark = p 
	goto st93
tr221:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:11
 port_mark = p 
	goto st93
	st93:
		if p++; p == pe {
			goto _test_eof93
		}
	st_case_93:
//line parser.go:3820
		switch data[p] {
		case 33:
			goto tr130
		case 37:
			goto tr131
		case 47:
			goto tr61
		case 58:
			goto st81
		case 60:
			goto st54
		case 62:
			goto st54
		case 64:
			goto st119
		case 91:
			goto tr133
		case 92:
			goto tr63
		case 95:
			goto tr130
		case 126:
			goto tr130
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr130
				}
			case data[p] >= 32:
				goto st54
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st54
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st54
				}
			default:
				goto tr132
			}
		default:
			goto tr132
		}
		goto st0
tr130:
//line parser.rl:23

    host_mark = p;
  
	goto st94
tr180:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st94
tr132:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st94
tr226:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st94
	st94:
		if p++; p == pe {
			goto _test_eof94
		}
	st_case_94:
//line parser.go:3921
		switch data[p] {
		case 33:
			goto st95
		case 37:
			goto st110
		case 47:
			goto tr136
		case 58:
			goto tr137
		case 60:
			goto st55
		case 62:
			goto st55
		case 64:
			goto st112
		case 92:
			goto tr139
		case 95:
			goto st95
		case 126:
			goto st95
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st95
				}
			case data[p] >= 32:
				goto st55
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st55
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st55
				}
			default:
				goto tr138
			}
		default:
			goto tr138
		}
		goto st0
tr181:
//line parser.rl:23

    host_mark = p;
  
	goto st95
tr168:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st95
tr183:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st95
tr138:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st95
	st95:
		if p++; p == pe {
			goto _test_eof95
		}
	st_case_95:
//line parser.go:4020
		switch data[p] {
		case 33:
			goto st96
		case 37:
			goto st105
		case 47:
			goto tr136
		case 58:
			goto tr137
		case 60:
			goto st56
		case 62:
			goto st56
		case 64:
			goto st108
		case 92:
			goto tr143
		case 95:
			goto st96
		case 126:
			goto st96
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st96
				}
			case data[p] >= 32:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st56
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st56
				}
			default:
				goto tr142
			}
		default:
			goto tr142
		}
		goto st0
tr169:
//line parser.rl:23

    host_mark = p;
  
	goto st96
tr159:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st96
tr171:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st96
tr142:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st96
	st96:
		if p++; p == pe {
			goto _test_eof96
		}
	st_case_96:
//line parser.go:4119
		switch data[p] {
		case 33:
			goto st97
		case 37:
			goto st98
		case 47:
			goto tr136
		case 58:
			goto tr137
		case 60:
			goto st57
		case 62:
			goto st57
		case 64:
			goto st89
		case 92:
			goto tr146
		case 95:
			goto st97
		case 126:
			goto st97
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st97
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st57
				}
			default:
				goto tr145
			}
		default:
			goto tr145
		}
		goto st0
tr160:
//line parser.rl:23

    host_mark = p;
  
	goto st97
	st97:
		if p++; p == pe {
			goto _test_eof97
		}
	st_case_97:
//line parser.go:4180
		switch data[p] {
		case 33:
			goto st46
		case 37:
			goto st1
		case 47:
			goto tr136
		case 58:
			goto tr48
		case 61:
			goto st46
		case 64:
			goto st11
		case 92:
			goto tr136
		case 95:
			goto st46
		case 126:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr147
			}
		default:
			goto tr147
		}
		goto st0
tr161:
//line parser.rl:23

    host_mark = p;
  
	goto st98
	st98:
		if p++; p == pe {
			goto _test_eof98
		}
	st_case_98:
//line parser.go:4225
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr149
				}
			case data[p] >= 48:
				goto st37
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr149
			}
		default:
			goto tr75
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st97
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto st0
tr149:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line parser.go:4288
		if data[p] == 58 {
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st97
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st97
			}
		default:
			goto st97
		}
		goto st0
tr137:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st99
	st99:
		if p++; p == pe {
			goto _test_eof99
		}
	st_case_99:
//line parser.go:4316
		switch data[p] {
		case 33:
			goto tr4
		case 37:
			goto tr5
		case 47:
			goto tr61
		case 60:
			goto tr3
		case 62:
			goto tr3
		case 64:
			goto tr6
		case 92:
			goto tr151
		case 95:
			goto tr4
		case 126:
			goto tr4
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr4
				}
			case data[p] >= 32:
				goto tr3
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto tr3
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr3
				}
			default:
				goto tr150
			}
		default:
			goto tr150
		}
		goto st0
tr39:
//line parser.rl:11
 port_mark = p 
	goto st100
tr151:
//line parser.rl:11
 port_mark = p 
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st100
tr224:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:11
 port_mark = p 
//line parser.rl:10
 mark = p 
	goto st100
	st100:
		if p++; p == pe {
			goto _test_eof100
		}
	st_case_100:
//line parser.go:4395
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st102
		case 47:
			goto tr61
		case 63:
			goto tr52
		case 91:
			goto st101
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr154
				}
			case data[p] >= 32:
				goto st101
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st101
				}
			case data[p] >= 97:
				goto tr154
			}
		default:
			goto st101
		}
		goto st0
tr154:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st101
	st101:
		if p++; p == pe {
			goto _test_eof101
		}
	st_case_101:
//line parser.go:4446
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st78
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st77
		}
		switch {
		case data[p] < 93:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 92 {
					goto tr97
				}
			case data[p] >= 32:
				goto st77
			}
		case data[p] > 96:
			switch {
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st77
				}
			case data[p] >= 97:
				goto tr97
			}
		default:
			goto st77
		}
		goto st0
	st102:
		if p++; p == pe {
			goto _test_eof102
		}
	st_case_102:
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st78
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st77
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st77
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr156
					}
				case data[p] >= 58:
					goto st77
				}
			default:
				goto st103
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st77
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st77
					}
				case data[p] >= 103:
					goto tr97
				}
			default:
				goto tr156
			}
		default:
			goto tr97
		}
		goto st0
tr156:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st103
	st103:
		if p++; p == pe {
			goto _test_eof103
		}
	st_case_103:
//line parser.go:4555
		switch data[p] {
		case 35:
			goto tr51
		case 37:
			goto st72
		case 47:
			goto tr78
		case 63:
			goto tr52
		case 91:
			goto st71
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] < 48:
				if 32 <= data[p] && data[p] <= 46 {
					goto st71
				}
			case data[p] > 57:
				switch {
				case data[p] > 64:
					if 65 <= data[p] && data[p] <= 70 {
						goto tr154
					}
				case data[p] >= 58:
					goto st71
				}
			default:
				goto st101
			}
		case data[p] > 92:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st71
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st71
					}
				case data[p] >= 103:
					goto tr90
				}
			default:
				goto tr154
			}
		default:
			goto tr90
		}
		goto st0
tr162:
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st104
tr145:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st104
	st104:
		if p++; p == pe {
			goto _test_eof104
		}
	st_case_104:
//line parser.go:4642
		switch data[p] {
		case 33:
			goto st46
		case 37:
			goto st1
		case 47:
			goto tr136
		case 58:
			goto tr125
		case 61:
			goto st46
		case 64:
			goto st11
		case 92:
			goto tr136
		case 95:
			goto st46
		case 126:
			goto st46
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr147
			}
		default:
			goto tr147
		}
		goto st0
tr170:
//line parser.rl:23

    host_mark = p;
  
	goto st105
	st105:
		if p++; p == pe {
			goto _test_eof105
		}
	st_case_105:
//line parser.go:4687
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st106
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr158
			}
		default:
			goto tr158
		}
		goto st0
	st106:
		if p++; p == pe {
			goto _test_eof106
		}
	st_case_106:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr159
				}
			case data[p] >= 48:
				goto st96
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr159
			}
		default:
			goto tr75
		}
		goto st0
tr158:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st107
	st107:
		if p++; p == pe {
			goto _test_eof107
		}
	st_case_107:
//line parser.go:4793
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr159
				}
			case data[p] >= 48:
				goto st96
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr159
			}
		default:
			goto tr75
		}
		goto st0
	st108:
		if p++; p == pe {
			goto _test_eof108
		}
	st_case_108:
		switch data[p] {
		case 33:
			goto tr160
		case 37:
			goto tr161
		case 47:
			goto tr61
		case 58:
			goto st84
		case 60:
			goto st57
		case 62:
			goto st57
		case 64:
			goto st89
		case 91:
			goto tr163
		case 92:
			goto tr74
		case 95:
			goto tr160
		case 126:
			goto tr160
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr160
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st57
				}
			default:
				goto tr162
			}
		default:
			goto tr162
		}
		goto st0
tr163:
//line parser.rl:23

    host_mark = p;
  
	goto st109
	st109:
		if p++; p == pe {
			goto _test_eof109
		}
	st_case_109:
//line parser.go:4892
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		case 118:
			goto tr165
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 58:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr164
				}
			case data[p] >= 48:
				goto st13
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr164
			}
		default:
			goto tr75
		}
		goto st0
tr164:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line parser.go:4939
		switch data[p] {
		case 58:
			goto st40
		case 93:
			goto st85
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 47:
			goto st47
		case 58:
			goto st13
		case 92:
			goto st47
		case 93:
			goto st85
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st13
			}
		default:
			goto st13
		}
		goto st0
tr165:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line parser.go:5002
		if data[p] == 58 {
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
tr182:
//line parser.rl:23

    host_mark = p;
  
	goto st110
	st110:
		if p++; p == pe {
			goto _test_eof110
		}
	st_case_110:
//line parser.go:5030
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st111
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr167
			}
		default:
			goto tr167
		}
		goto st0
tr167:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st111
	st111:
		if p++; p == pe {
			goto _test_eof111
		}
	st_case_111:
//line parser.go:5102
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st95
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr168
			}
		default:
			goto tr168
		}
		goto st0
	st112:
		if p++; p == pe {
			goto _test_eof112
		}
	st_case_112:
		switch data[p] {
		case 33:
			goto tr169
		case 37:
			goto tr170
		case 47:
			goto tr61
		case 58:
			goto st83
		case 60:
			goto st56
		case 62:
			goto st56
		case 64:
			goto st108
		case 91:
			goto tr172
		case 92:
			goto tr71
		case 95:
			goto tr169
		case 126:
			goto tr169
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr169
				}
			case data[p] >= 32:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st56
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st56
				}
			default:
				goto tr171
			}
		default:
			goto tr171
		}
		goto st0
tr172:
//line parser.rl:23

    host_mark = p;
  
	goto st113
	st113:
		if p++; p == pe {
			goto _test_eof113
		}
	st_case_113:
//line parser.go:5226
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		case 118:
			goto tr175
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 58:
					if 59 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st114
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
	st114:
		if p++; p == pe {
			goto _test_eof114
		}
	st_case_114:
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st13
		case 92:
			goto tr61
		case 93:
			goto st85
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr164
				}
			case data[p] >= 48:
				goto st14
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr164
			}
		default:
			goto tr75
		}
		goto st0
tr174:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st115
	st115:
		if p++; p == pe {
			goto _test_eof115
		}
	st_case_115:
//line parser.go:5338
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st40
		case 92:
			goto tr61
		case 93:
			goto st85
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr164
				}
			case data[p] >= 48:
				goto st14
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr164
			}
		default:
			goto tr75
		}
		goto st0
tr175:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st116
	st116:
		if p++; p == pe {
			goto _test_eof116
		}
	st_case_116:
//line parser.go:5387
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr176
				}
			case data[p] >= 48:
				goto st29
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr176
			}
		default:
			goto tr75
		}
		goto st0
tr176:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line parser.go:5434
		switch data[p] {
		case 46:
			goto st30
		case 58:
			goto st5
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st29
			}
		default:
			goto st29
		}
		goto st0
tr131:
//line parser.rl:23

    host_mark = p;
  
	goto st117
	st117:
		if p++; p == pe {
			goto _test_eof117
		}
	st_case_117:
//line parser.go:5465
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st118
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr178
			}
		default:
			goto tr178
		}
		goto st0
tr178:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st118
	st118:
		if p++; p == pe {
			goto _test_eof118
		}
	st_case_118:
//line parser.go:5537
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st94
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr180
			}
		default:
			goto tr180
		}
		goto st0
	st119:
		if p++; p == pe {
			goto _test_eof119
		}
	st_case_119:
		switch data[p] {
		case 33:
			goto tr181
		case 37:
			goto tr182
		case 47:
			goto tr61
		case 58:
			goto st82
		case 60:
			goto st55
		case 62:
			goto st55
		case 64:
			goto st112
		case 91:
			goto tr184
		case 92:
			goto tr67
		case 95:
			goto tr181
		case 126:
			goto tr181
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr181
				}
			case data[p] >= 32:
				goto st55
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto st55
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st55
				}
			default:
				goto tr183
			}
		default:
			goto tr183
		}
		goto st0
tr184:
//line parser.rl:23

    host_mark = p;
  
	goto st120
	st120:
		if p++; p == pe {
			goto _test_eof120
		}
	st_case_120:
//line parser.go:5661
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		case 118:
			goto tr187
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 58:
					if 59 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st121
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr186
			}
		default:
			goto tr186
		}
		goto st0
tr186:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st121
	st121:
		if p++; p == pe {
			goto _test_eof121
		}
	st_case_121:
//line parser.go:5735
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 58:
			goto st114
		case 64:
			goto st57
		case 92:
			goto tr74
		case 93:
			goto st123
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 59 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st122
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
	st122:
		if p++; p == pe {
			goto _test_eof122
		}
	st_case_122:
		switch data[p] {
		case 46:
			goto st15
		case 47:
			goto tr61
		case 58:
			goto st13
		case 92:
			goto tr61
		case 93:
			goto st85
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr164
				}
			case data[p] >= 48:
				goto st14
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr164
			}
		default:
			goto tr75
		}
		goto st0
	st123:
		if p++; p == pe {
			goto _test_eof123
		}
	st_case_123:
		switch data[p] {
		case 47:
			goto tr136
		case 58:
			goto tr116
		case 92:
			goto tr136
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr190
			}
		case data[p] >= 65:
			goto tr190
		}
		goto st0
tr187:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st124
	st124:
		if p++; p == pe {
			goto _test_eof124
		}
	st_case_124:
//line parser.go:5873
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st125
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st0
	st125:
		if p++; p == pe {
			goto _test_eof125
		}
	st_case_125:
		switch data[p] {
		case 46:
			goto st30
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr176
				}
			case data[p] >= 48:
				goto st29
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr176
			}
		default:
			goto tr75
		}
		goto st0
tr192:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st126
	st126:
		if p++; p == pe {
			goto _test_eof126
		}
	st_case_126:
//line parser.go:5981
		switch data[p] {
		case 46:
			goto st30
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr176
				}
			case data[p] >= 48:
				goto st29
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr176
			}
		default:
			goto tr75
		}
		goto st0
tr133:
//line parser.rl:23

    host_mark = p;
  
	goto st127
	st127:
		if p++; p == pe {
			goto _test_eof127
		}
	st_case_127:
//line parser.go:6026
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		case 118:
			goto tr195
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 58:
					if 59 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st128
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr194
			}
		default:
			goto tr194
		}
		goto st0
tr194:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st128
	st128:
		if p++; p == pe {
			goto _test_eof128
		}
	st_case_128:
//line parser.go:6100
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 58:
			goto st121
		case 64:
			goto st56
		case 92:
			goto tr71
		case 93:
			goto st131
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 59 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st129
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr186
			}
		default:
			goto tr186
		}
		goto st0
	st129:
		if p++; p == pe {
			goto _test_eof129
		}
	st_case_129:
		switch data[p] {
		case 37:
			goto st58
		case 46:
			goto st130
		case 47:
			goto tr61
		case 58:
			goto st114
		case 64:
			goto st57
		case 92:
			goto tr74
		case 93:
			goto st123
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 45:
				switch {
				case data[p] > 57:
					if 59 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st122
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr174
			}
		default:
			goto tr174
		}
		goto st0
	st130:
		if p++; p == pe {
			goto _test_eof130
		}
	st_case_130:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr75
			}
		default:
			goto tr75
		}
		goto st0
	st131:
		if p++; p == pe {
			goto _test_eof131
		}
	st_case_131:
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr136
		case 58:
			goto tr199
		case 64:
			goto st57
		case 92:
			goto tr146
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st57
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st57
				}
			default:
				goto tr200
			}
		default:
			goto tr200
		}
		goto st0
tr199:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st132
	st132:
		if p++; p == pe {
			goto _test_eof132
		}
	st_case_132:
//line parser.go:6307
		switch data[p] {
		case 37:
			goto tr31
		case 47:
			goto tr61
		case 64:
			goto tr3
		case 92:
			goto tr151
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto tr3
				}
			case data[p] >= 32:
				goto tr3
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto tr3
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto tr3
				}
			default:
				goto tr201
			}
		default:
			goto tr201
		}
		goto st0
tr195:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st133
	st133:
		if p++; p == pe {
			goto _test_eof133
		}
	st_case_133:
//line parser.go:6360
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st134
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr203
			}
		default:
			goto tr203
		}
		goto st0
tr203:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st134
	st134:
		if p++; p == pe {
			goto _test_eof134
		}
	st_case_134:
//line parser.go:6432
		switch data[p] {
		case 37:
			goto st58
		case 46:
			goto st135
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 45:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st125
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr192
			}
		default:
			goto tr192
		}
		goto st0
	st135:
		if p++; p == pe {
			goto _test_eof135
		}
	st_case_135:
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 47:
			goto tr61
		case 61:
			goto st31
		case 92:
			goto tr61
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 38 <= data[p] && data[p] <= 59 {
				goto st31
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr205
			}
		default:
			goto tr205
		}
		goto st0
tr205:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line parser.go:6540
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 58:
			goto st44
		case 61:
			goto st31
		case 93:
			goto st85
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 48:
			if 38 <= data[p] && data[p] <= 46 {
				goto st31
			}
		case data[p] > 59:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st31
				}
			case data[p] >= 65:
				goto st31
			}
		default:
			goto st31
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 47:
			goto st47
		case 61:
			goto st31
		case 92:
			goto st47
		case 93:
			goto st85
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 38 <= data[p] && data[p] <= 59 {
				goto st31
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st31
			}
		default:
			goto st31
		}
		goto st0
tr114:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st136
	st136:
		if p++; p == pe {
			goto _test_eof136
		}
	st_case_136:
//line parser.go:6626
		switch data[p] {
		case 33:
			goto st8
		case 37:
			goto st9
		case 47:
			goto tr61
		case 58:
			goto st33
		case 61:
			goto st8
		case 64:
			goto st11
		case 92:
			goto tr61
		case 95:
			goto st8
		case 126:
			goto st8
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto st8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr115
			}
		default:
			goto tr115
		}
		goto st0
	st137:
		if p++; p == pe {
			goto _test_eof137
		}
	st_case_137:
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st138
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr207
			}
		default:
			goto tr207
		}
		goto st0
	st138:
		if p++; p == pe {
			goto _test_eof138
		}
	st_case_138:
		switch data[p] {
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr111
				}
			case data[p] >= 48:
				goto st83
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr111
			}
		default:
			goto tr75
		}
		goto st0
tr207:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st139
	st139:
		if p++; p == pe {
			goto _test_eof139
		}
	st_case_139:
//line parser.go:6770
		switch data[p] {
		case 47:
			goto tr61
		case 58:
			goto st5
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 71:
			switch {
			case data[p] > 57:
				if 65 <= data[p] && data[p] <= 70 {
					goto tr111
				}
			case data[p] >= 48:
				goto st83
			}
		case data[p] > 90:
			switch {
			case data[p] > 102:
				if 103 <= data[p] && data[p] <= 122 {
					goto tr75
				}
			case data[p] >= 97:
				goto tr111
			}
		default:
			goto tr75
		}
		goto st0
	st140:
		if p++; p == pe {
			goto _test_eof140
		}
	st_case_140:
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st141
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr209
			}
		default:
			goto tr209
		}
		goto st0
tr209:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st141
	st141:
		if p++; p == pe {
			goto _test_eof141
		}
	st_case_141:
//line parser.go:6878
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st57
					}
				case data[p] >= 48:
					goto st82
				}
			default:
				goto st57
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st57
					}
				case data[p] >= 71:
					goto tr73
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 103:
					goto tr73
				}
			default:
				goto tr107
			}
		default:
			goto tr107
		}
		goto st0
	st142:
		if p++; p == pe {
			goto _test_eof142
		}
	st_case_142:
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st143
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr211
			}
		default:
			goto tr211
		}
		goto st0
tr211:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st143
	st143:
		if p++; p == pe {
			goto _test_eof143
		}
	st_case_143:
//line parser.go:7011
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st81
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr103
			}
		default:
			goto tr103
		}
		goto st0
tr43:
//line parser.rl:10
 mark = p 
	goto st144
tr219:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:10
 mark = p 
	goto st144
	st144:
		if p++; p == pe {
			goto _test_eof144
		}
	st_case_144:
//line parser.go:7085
		switch data[p] {
		case 35:
			goto tr51
		case 47:
			goto st145
		case 63:
			goto tr52
		}
		if 32 <= data[p] && data[p] <= 126 {
			goto st48
		}
		goto st0
	st145:
		if p++; p == pe {
			goto _test_eof145
		}
	st_case_145:
		switch data[p] {
		case 33:
			goto tr14
		case 37:
			goto tr15
		case 47:
			goto tr45
		case 58:
			goto st8
		case 61:
			goto tr14
		case 64:
			goto st11
		case 91:
			goto tr16
		case 92:
			goto tr45
		case 95:
			goto tr14
		case 126:
			goto tr14
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 59 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr213
			}
		default:
			goto tr213
		}
		goto st0
tr44:
//line parser.rl:10
 mark = p 
//line parser.rl:23

    host_mark = p;
  
	goto st146
tr216:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:10
 mark = p 
	goto st146
	st146:
		if p++; p == pe {
			goto _test_eof146
		}
	st_case_146:
//line parser.go:7159
		switch data[p] {
		case 33:
			goto st46
		case 37:
			goto st1
		case 43:
			goto st146
		case 47:
			goto tr47
		case 58:
			goto tr215
		case 59:
			goto st46
		case 61:
			goto st46
		case 64:
			goto st11
		case 92:
			goto tr47
		case 95:
			goto st46
		case 126:
			goto st46
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 44 {
				goto st46
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr216
				}
			case data[p] >= 65:
				goto tr216
			}
		default:
			goto st146
		}
		goto st0
tr215:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
	goto st147
	st147:
		if p++; p == pe {
			goto _test_eof147
		}
	st_case_147:
//line parser.go:7213
		switch data[p] {
		case 33:
			goto tr217
		case 37:
			goto tr218
		case 47:
			goto tr219
		case 58:
			goto tr220
		case 60:
			goto tr3
		case 62:
			goto tr3
		case 64:
			goto tr221
		case 91:
			goto tr223
		case 92:
			goto tr224
		case 95:
			goto tr217
		case 126:
			goto tr217
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto tr217
				}
			case data[p] >= 32:
				goto tr3
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 93 <= data[p] && data[p] <= 96 {
					goto tr3
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto tr3
				}
			default:
				goto tr222
			}
		default:
			goto tr222
		}
		goto st0
tr231:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st148
tr217:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:11
 port_mark = p 
	goto st148
tr222:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:11
 port_mark = p 
//line parser.rl:10
 mark = p 
	goto st148
	st148:
		if p++; p == pe {
			goto _test_eof148
		}
	st_case_148:
//line parser.go:7306
		switch data[p] {
		case 33:
			goto st94
		case 37:
			goto st117
		case 47:
			goto tr136
		case 58:
			goto tr137
		case 60:
			goto st54
		case 62:
			goto st54
		case 64:
			goto st119
		case 92:
			goto tr227
		case 95:
			goto st94
		case 126:
			goto st94
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st94
				}
			case data[p] >= 32:
				goto st54
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st54
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st54
				}
			default:
				goto tr226
			}
		default:
			goto tr226
		}
		goto st0
tr218:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:11
 port_mark = p 
	goto st149
	st149:
		if p++; p == pe {
			goto _test_eof149
		}
	st_case_149:
//line parser.go:7373
		switch data[p] {
		case 37:
			goto st74
		case 47:
			goto tr61
		case 64:
			goto st54
		case 92:
			goto tr63
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st54
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st54
					}
				case data[p] >= 48:
					goto st150
				}
			default:
				goto st54
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st54
					}
				case data[p] >= 71:
					goto tr62
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st54
					}
				case data[p] >= 103:
					goto tr62
				}
			default:
				goto tr229
			}
		default:
			goto tr229
		}
		goto st0
tr229:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st150
	st150:
		if p++; p == pe {
			goto _test_eof150
		}
	st_case_150:
//line parser.go:7445
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st148
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr231
			}
		default:
			goto tr231
		}
		goto st0
tr223:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
//line parser.rl:23

    host_mark = p;
  
//line parser.rl:11
 port_mark = p 
	goto st151
	st151:
		if p++; p == pe {
			goto _test_eof151
		}
	st_case_151:
//line parser.go:7519
		switch data[p] {
		case 37:
			goto st74
		case 47:
			goto tr61
		case 64:
			goto st54
		case 92:
			goto tr63
		case 118:
			goto tr234
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st54
				}
			case data[p] > 46:
				switch {
				case data[p] > 58:
					if 59 <= data[p] && data[p] <= 62 {
						goto st54
					}
				case data[p] >= 48:
					goto st152
				}
			default:
				goto st54
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st54
					}
				case data[p] >= 71:
					goto tr62
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st54
					}
				case data[p] >= 103:
					goto tr62
				}
			default:
				goto tr233
			}
		default:
			goto tr233
		}
		goto st0
tr233:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st152
	st152:
		if p++; p == pe {
			goto _test_eof152
		}
	st_case_152:
//line parser.go:7593
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 58:
			goto st128
		case 64:
			goto st55
		case 92:
			goto tr67
		case 93:
			goto st156
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 59 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st153
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr194
			}
		default:
			goto tr194
		}
		goto st0
	st153:
		if p++; p == pe {
			goto _test_eof153
		}
	st_case_153:
		switch data[p] {
		case 37:
			goto st61
		case 46:
			goto st154
		case 47:
			goto tr61
		case 58:
			goto st121
		case 64:
			goto st56
		case 92:
			goto tr71
		case 93:
			goto st131
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 45:
				switch {
				case data[p] > 57:
					if 59 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st129
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr186
			}
		default:
			goto tr186
		}
		goto st0
	st154:
		if p++; p == pe {
			goto _test_eof154
		}
	st_case_154:
		switch data[p] {
		case 37:
			goto st58
		case 47:
			goto tr61
		case 64:
			goto st57
		case 92:
			goto tr74
		}
		switch {
		case data[p] < 58:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st57
				}
			case data[p] > 46:
				if 48 <= data[p] && data[p] <= 57 {
					goto st155
				}
			default:
				goto st57
			}
		case data[p] > 62:
			switch {
			case data[p] < 91:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr73
				}
			case data[p] > 96:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st57
					}
				case data[p] >= 97:
					goto tr73
				}
			default:
				goto st57
			}
		default:
			goto st57
		}
		goto st0
	st155:
		if p++; p == pe {
			goto _test_eof155
		}
	st_case_155:
		switch data[p] {
		case 46:
			goto st17
		case 47:
			goto tr61
		case 92:
			goto tr61
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr75
			}
		default:
			goto tr75
		}
		goto st0
	st156:
		if p++; p == pe {
			goto _test_eof156
		}
	st_case_156:
		switch data[p] {
		case 37:
			goto st61
		case 47:
			goto tr136
		case 58:
			goto tr199
		case 64:
			goto st56
		case 92:
			goto tr143
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 62 {
					goto st56
				}
			case data[p] >= 32:
				goto st56
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st56
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 126 {
					goto st56
				}
			default:
				goto tr239
			}
		default:
			goto tr239
		}
		goto st0
tr234:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st157
	st157:
		if p++; p == pe {
			goto _test_eof157
		}
	st_case_157:
//line parser.go:7857
		switch data[p] {
		case 37:
			goto st68
		case 47:
			goto tr61
		case 64:
			goto st55
		case 92:
			goto tr67
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st55
				}
			case data[p] > 46:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st55
					}
				case data[p] >= 48:
					goto st158
				}
			default:
				goto st55
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st55
					}
				case data[p] >= 71:
					goto tr66
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st55
					}
				case data[p] >= 103:
					goto tr66
				}
			default:
				goto tr241
			}
		default:
			goto tr241
		}
		goto st0
tr241:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st158
	st158:
		if p++; p == pe {
			goto _test_eof158
		}
	st_case_158:
//line parser.go:7929
		switch data[p] {
		case 37:
			goto st61
		case 46:
			goto st159
		case 47:
			goto tr61
		case 64:
			goto st56
		case 92:
			goto tr71
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] < 36:
				if 32 <= data[p] && data[p] <= 34 {
					goto st56
				}
			case data[p] > 45:
				switch {
				case data[p] > 57:
					if 58 <= data[p] && data[p] <= 62 {
						goto st56
					}
				case data[p] >= 48:
					goto st134
				}
			default:
				goto st56
			}
		case data[p] > 70:
			switch {
			case data[p] < 97:
				switch {
				case data[p] > 90:
					if 91 <= data[p] && data[p] <= 96 {
						goto st56
					}
				case data[p] >= 71:
					goto tr70
				}
			case data[p] > 102:
				switch {
				case data[p] > 122:
					if 123 <= data[p] && data[p] <= 126 {
						goto st56
					}
				case data[p] >= 103:
					goto tr70
				}
			default:
				goto tr203
			}
		default:
			goto tr203
		}
		goto st0
	st159:
		if p++; p == pe {
			goto _test_eof159
		}
	st_case_159:
		switch data[p] {
		case 33:
			goto st160
		case 37:
			goto st58
		case 47:
			goto tr61
		case 60:
			goto st57
		case 62:
			goto st57
		case 64:
			goto st57
		case 92:
			goto tr74
		case 95:
			goto st160
		case 126:
			goto st160
		}
		switch {
		case data[p] < 65:
			switch {
			case data[p] > 34:
				if 36 <= data[p] && data[p] <= 61 {
					goto st160
				}
			case data[p] >= 32:
				goto st57
			}
		case data[p] > 90:
			switch {
			case data[p] < 97:
				if 91 <= data[p] && data[p] <= 96 {
					goto st57
				}
			case data[p] > 122:
				if 123 <= data[p] && data[p] <= 125 {
					goto st57
				}
			default:
				goto tr244
			}
		default:
			goto tr244
		}
		goto st0
	st160:
		if p++; p == pe {
			goto _test_eof160
		}
	st_case_160:
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 47:
			goto tr61
		case 61:
			goto st31
		case 92:
			goto tr61
		case 93:
			goto st85
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 38 <= data[p] && data[p] <= 59 {
				goto st31
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr205
			}
		default:
			goto tr205
		}
		goto st0
tr244:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:10
 mark = p 
	goto st161
	st161:
		if p++; p == pe {
			goto _test_eof161
		}
	st_case_161:
//line parser.go:8091
		switch data[p] {
		case 33:
			goto st31
		case 36:
			goto st31
		case 47:
			goto tr61
		case 58:
			goto st44
		case 61:
			goto st31
		case 92:
			goto tr61
		case 93:
			goto st85
		case 95:
			goto st31
		case 126:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 38 <= data[p] && data[p] <= 59 {
				goto st31
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr205
			}
		default:
			goto tr205
		}
		goto st0
	st_out:
	_test_eof46: cs = 46; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof79: cs = 79; goto _test_eof
	_test_eof80: cs = 80; goto _test_eof
	_test_eof81: cs = 81; goto _test_eof
	_test_eof82: cs = 82; goto _test_eof
	_test_eof83: cs = 83; goto _test_eof
	_test_eof84: cs = 84; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof85: cs = 85; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof86: cs = 86; goto _test_eof
	_test_eof87: cs = 87; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof88: cs = 88; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof89: cs = 89; goto _test_eof
	_test_eof90: cs = 90; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof91: cs = 91; goto _test_eof
	_test_eof92: cs = 92; goto _test_eof
	_test_eof93: cs = 93; goto _test_eof
	_test_eof94: cs = 94; goto _test_eof
	_test_eof95: cs = 95; goto _test_eof
	_test_eof96: cs = 96; goto _test_eof
	_test_eof97: cs = 97; goto _test_eof
	_test_eof98: cs = 98; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof99: cs = 99; goto _test_eof
	_test_eof100: cs = 100; goto _test_eof
	_test_eof101: cs = 101; goto _test_eof
	_test_eof102: cs = 102; goto _test_eof
	_test_eof103: cs = 103; goto _test_eof
	_test_eof104: cs = 104; goto _test_eof
	_test_eof105: cs = 105; goto _test_eof
	_test_eof106: cs = 106; goto _test_eof
	_test_eof107: cs = 107; goto _test_eof
	_test_eof108: cs = 108; goto _test_eof
	_test_eof109: cs = 109; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof110: cs = 110; goto _test_eof
	_test_eof111: cs = 111; goto _test_eof
	_test_eof112: cs = 112; goto _test_eof
	_test_eof113: cs = 113; goto _test_eof
	_test_eof114: cs = 114; goto _test_eof
	_test_eof115: cs = 115; goto _test_eof
	_test_eof116: cs = 116; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof117: cs = 117; goto _test_eof
	_test_eof118: cs = 118; goto _test_eof
	_test_eof119: cs = 119; goto _test_eof
	_test_eof120: cs = 120; goto _test_eof
	_test_eof121: cs = 121; goto _test_eof
	_test_eof122: cs = 122; goto _test_eof
	_test_eof123: cs = 123; goto _test_eof
	_test_eof124: cs = 124; goto _test_eof
	_test_eof125: cs = 125; goto _test_eof
	_test_eof126: cs = 126; goto _test_eof
	_test_eof127: cs = 127; goto _test_eof
	_test_eof128: cs = 128; goto _test_eof
	_test_eof129: cs = 129; goto _test_eof
	_test_eof130: cs = 130; goto _test_eof
	_test_eof131: cs = 131; goto _test_eof
	_test_eof132: cs = 132; goto _test_eof
	_test_eof133: cs = 133; goto _test_eof
	_test_eof134: cs = 134; goto _test_eof
	_test_eof135: cs = 135; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof136: cs = 136; goto _test_eof
	_test_eof137: cs = 137; goto _test_eof
	_test_eof138: cs = 138; goto _test_eof
	_test_eof139: cs = 139; goto _test_eof
	_test_eof140: cs = 140; goto _test_eof
	_test_eof141: cs = 141; goto _test_eof
	_test_eof142: cs = 142; goto _test_eof
	_test_eof143: cs = 143; goto _test_eof
	_test_eof144: cs = 144; goto _test_eof
	_test_eof145: cs = 145; goto _test_eof
	_test_eof146: cs = 146; goto _test_eof
	_test_eof147: cs = 147; goto _test_eof
	_test_eof148: cs = 148; goto _test_eof
	_test_eof149: cs = 149; goto _test_eof
	_test_eof150: cs = 150; goto _test_eof
	_test_eof151: cs = 151; goto _test_eof
	_test_eof152: cs = 152; goto _test_eof
	_test_eof153: cs = 153; goto _test_eof
	_test_eof154: cs = 154; goto _test_eof
	_test_eof155: cs = 155; goto _test_eof
	_test_eof156: cs = 156; goto _test_eof
	_test_eof157: cs = 157; goto _test_eof
	_test_eof158: cs = 158; goto _test_eof
	_test_eof159: cs = 159; goto _test_eof
	_test_eof160: cs = 160; goto _test_eof
	_test_eof161: cs = 161; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 53, 54, 55, 56, 57, 58, 59, 61, 62, 63, 68, 69, 74, 75, 80, 81, 82, 83, 84, 86, 87, 88, 89, 91, 92, 93, 98, 99, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 124, 125, 126, 127, 128, 129, 130, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 149, 150, 151, 152, 153, 154, 155, 157, 158, 159, 160, 161:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
		case 147:
//line parser.rl:19

    u.Protocol = data[0:p-1]
  
		case 46, 85, 90, 146:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
		case 52:
//line parser.rl:31

    if u.Query == "" {
      u.Query = data[mark:p]
    }
  
		case 47, 48, 67, 144:
//line parser.rl:37

    if u.Path == "" {
      u.Path = data[mark:p]
    }
  
		case 50:
//line parser.rl:43

    u.Fragment = data[mark:p]
  
		case 51:
//line parser.rl:10
 mark = p 
//line parser.rl:31

    if u.Query == "" {
      u.Query = data[mark:p]
    }
  
		case 49:
//line parser.rl:10
 mark = p 
//line parser.rl:43

    u.Fragment = data[mark:p]
  
		case 60, 64, 65, 66, 70, 71, 72, 73, 76, 77, 78, 79, 100, 101, 102, 103:
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.rl:37

    if u.Path == "" {
      u.Path = data[mark:p]
    }
  
		case 94, 95, 96, 97, 104, 123, 131, 148, 156:
//line parser.rl:27

    u.Host = data[host_mark:p]
  
//line parser.rl:13

    if port_mark > host_mark{
      u.Port = data[port_mark:p]
    }
  
//line parser.go:8366
		}
	}

	_out: {}
	}

//line parser.rl:116
  if cs < url_parser_first_final {
    return nil, fmt.Errorf("Failed to match URL")
  } else {
    return u, nil
  }
}

