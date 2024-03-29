package erajp

type EraItem struct {
	Name        string
	Ruby        string
	RubyInitial string
	Year        int
	Month       int
	Day         int
}

var eras = []EraItem {
	{Name: "白雉", Ruby: "はくち", RubyInitial: "H", Year: 650, Month: 2, Day: 15},
	{Name: "朱鳥", Ruby: "しゅちょう（すちょう）", RubyInitial: "S", Year: 686, Month: 7, Day: 20},
	{Name: "大宝", Ruby: "たいほう（だいほう）", RubyInitial: "T", Year: 701, Month: 3, Day: 21},
	{Name: "慶雲", Ruby: "けいうん（きょううん）", RubyInitial: "K", Year: 704, Month: 5, Day: 10},
	{Name: "和銅", Ruby: "わどう", RubyInitial: "W", Year: 708, Month: 1, Day: 11},
	{Name: "養老", Ruby: "ようろう", RubyInitial: "Y", Year: 717, Month: 11, Day: 17},
	{Name: "神亀", Ruby: "じんき", RubyInitial: "J", Year: 724, Month: 2, Day: 4},
	{Name: "天平", Ruby: "てんぴょう（てんびょう）", RubyInitial: "T", Year: 729, Month: 8, Day: 5},
	{Name: "天平感宝", Ruby: "てんぴょうかんぽう", RubyInitial: "T", Year: 749, Month: 4, Day: 14},
	{Name: "天平勝宝", Ruby: "てんぴょうしょうほう", RubyInitial: "T", Year: 749, Month: 7, Day: 2},
	{Name: "天平宝字", Ruby: "てんぴょうほうじ", RubyInitial: "T", Year: 757, Month: 8, Day: 18},
	{Name: "天平神護", Ruby: "てんぴょうしんご", RubyInitial: "T", Year: 765, Month: 1, Day: 7},
	{Name: "神護景雲", Ruby: "しんごけいうん", RubyInitial: "S", Year: 767, Month: 8, Day: 16},
	{Name: "宝亀", Ruby: "ほうき", RubyInitial: "H", Year: 770, Month: 10, Day: 1},
	{Name: "天応", Ruby: "てんおう", RubyInitial: "T", Year: 781, Month: 1, Day: 1},
	{Name: "延暦", Ruby: "えんりゃく", RubyInitial: "E", Year: 782, Month: 8, Day: 19},
	{Name: "弘仁", Ruby: "こうにん", RubyInitial: "K", Year: 810, Month: 9, Day: 19},
	{Name: "天長", Ruby: "てんちょう", RubyInitial: "T", Year: 824, Month: 1, Day: 5},
	{Name: "承和", Ruby: "じょうわ（しょうわ）", RubyInitial: "J", Year: 834, Month: 1, Day: 3},
	{Name: "嘉祥", Ruby: "かしょう（かじょう）", RubyInitial: "K", Year: 848, Month: 6, Day: 13},
	{Name: "仁寿", Ruby: "にんじゅ", RubyInitial: "N", Year: 851, Month: 4, Day: 28},
	{Name: "斎衡", Ruby: "さいこう", RubyInitial: "S", Year: 854, Month: 11, Day: 30},
	{Name: "天安", Ruby: "てんあん（てんなん）", RubyInitial: "T", Year: 857, Month: 2, Day: 21},
	{Name: "貞観", Ruby: "じょうがん", RubyInitial: "J", Year: 859, Month: 4, Day: 15},
	{Name: "元慶", Ruby: "がんぎょう（げんけい）", RubyInitial: "G", Year: 877, Month: 4, Day: 16},
	{Name: "仁和", Ruby: "にんな（じんな）", RubyInitial: "N", Year: 885, Month: 2, Day: 21},
	{Name: "寛平", Ruby: "かんぴょう（かんぺい）", RubyInitial: "K", Year: 889, Month: 4, Day: 27},
	{Name: "昌泰", Ruby: "しょうたい", RubyInitial: "S", Year: 898, Month: 4, Day: 26},
	{Name: "延喜", Ruby: "えんぎ", RubyInitial: "E", Year: 901, Month: 7, Day: 15},
	{Name: "延長", Ruby: "えんちょう", RubyInitial: "E", Year: 923, Month: 4, Day: 11},
	{Name: "承平", Ruby: "じょうへい（しょうへい）", RubyInitial: "J", Year: 931, Month: 4, Day: 26},
	{Name: "天慶", Ruby: "てんぎょう（てんきょう）", RubyInitial: "T", Year: 938, Month: 5, Day: 22},
	{Name: "天暦", Ruby: "てんりゃく（てんれき）", RubyInitial: "T", Year: 947, Month: 4, Day: 22},
	{Name: "天徳", Ruby: "てんとく", RubyInitial: "T", Year: 957, Month: 10, Day: 27},
	{Name: "応和", Ruby: "おうわ", RubyInitial: "O", Year: 961, Month: 2, Day: 16},
	{Name: "康保", Ruby: "こうほう", RubyInitial: "K", Year: 964, Month: 7, Day: 10},
	{Name: "安和", Ruby: "あんな（あんわ）", RubyInitial: "A", Year: 968, Month: 8, Day: 13},
	{Name: "天禄", Ruby: "てんろく", RubyInitial: "T", Year: 970, Month: 3, Day: 25},
	{Name: "天延", Ruby: "てんえん", RubyInitial: "T", Year: 973, Month: 12, Day: 20},
	{Name: "貞元", Ruby: "じょうげん（ていげん）", RubyInitial: "J", Year: 976, Month: 7, Day: 13},
	{Name: "天元", Ruby: "てんげん", RubyInitial: "T", Year: 978, Month: 11, Day: 29},
	{Name: "永観", Ruby: "えいかん", RubyInitial: "E", Year: 983, Month: 4, Day: 15},
	{Name: "寛和", Ruby: "かんな（かんわ）", RubyInitial: "K", Year: 985, Month: 4, Day: 27},
	{Name: "永延", Ruby: "えいえん（ようえん）", RubyInitial: "E", Year: 987, Month: 4, Day: 5},
	{Name: "永祚", Ruby: "えいそ", RubyInitial: "E", Year: 989, Month: 8, Day: 8},
	{Name: "正暦", Ruby: "しょうりゃく（じょうりゃく）", RubyInitial: "S", Year: 990, Month: 11, Day: 7},
	{Name: "長徳", Ruby: "ちょうとく", RubyInitial: "T", Year: 995, Month: 2, Day: 22},
	{Name: "長保", Ruby: "ちょうほう", RubyInitial: "T", Year: 999, Month: 1, Day: 13},
	{Name: "寛弘", Ruby: "かんこう", RubyInitial: "K", Year: 1004, Month: 7, Day: 20},
	{Name: "長和", Ruby: "ちょうわ", RubyInitial: "T", Year: 1012, Month: 12, Day: 25},
	{Name: "寛仁", Ruby: "かんにん", RubyInitial: "K", Year: 1017, Month: 4, Day: 23},
	{Name: "治安", Ruby: "じあん（ちあん）", RubyInitial: "J", Year: 1021, Month: 2, Day: 2},
	{Name: "万寿", Ruby: "まんじゅ", RubyInitial: "M", Year: 1024, Month: 7, Day: 13},
	{Name: "長元", Ruby: "ちょうげん", RubyInitial: "T", Year: 1028, Month: 7, Day: 25},
	{Name: "長暦", Ruby: "ちょうりゃく（ちょうれき）", RubyInitial: "T", Year: 1037, Month: 4, Day: 21},
	{Name: "長久", Ruby: "ちょうきゅう", RubyInitial: "T", Year: 1040, Month: 11, Day: 10},
	{Name: "寛徳", Ruby: "かんとく", RubyInitial: "K", Year: 1044, Month: 11, Day: 24},
	{Name: "永承", Ruby: "えいしょう（えいじょう）", RubyInitial: "E", Year: 1046, Month: 4, Day: 14},
	{Name: "天喜", Ruby: "てんぎ（てんき）", RubyInitial: "T", Year: 1053, Month: 1, Day: 11},
	{Name: "康平", Ruby: "こうへい", RubyInitial: "K", Year: 1058, Month: 8, Day: 29},
	{Name: "治暦", Ruby: "じりゃく（ちりゃく）", RubyInitial: "J", Year: 1065, Month: 8, Day: 2},
	{Name: "延久", Ruby: "えんきゅう", RubyInitial: "E", Year: 1069, Month: 4, Day: 13},
	{Name: "承保", Ruby: "じょうほう（しょうほう）", RubyInitial: "J", Year: 1074, Month: 8, Day: 23},
	{Name: "承暦", Ruby: "じょうりゃく（しょうりゃく）", RubyInitial: "J", Year: 1077, Month: 11, Day: 17},
	{Name: "永保", Ruby: "えいほう", RubyInitial: "E", Year: 1081, Month: 2, Day: 10},
	{Name: "応徳", Ruby: "おうとく", RubyInitial: "O", Year: 1084, Month: 2, Day: 7},
	{Name: "寛治", Ruby: "かんじ", RubyInitial: "K", Year: 1087, Month: 4, Day: 7},
	{Name: "嘉保", Ruby: "かほう", RubyInitial: "K", Year: 1094, Month: 12, Day: 15},
	{Name: "永長", Ruby: "えいちょう（ようちょう）", RubyInitial: "E", Year: 1096, Month: 12, Day: 17},
	{Name: "承徳", Ruby: "じょうとく（しょうとく）", RubyInitial: "J", Year: 1097, Month: 11, Day: 21},
	{Name: "康和", Ruby: "こうわ", RubyInitial: "K", Year: 1099, Month: 8, Day: 28},
	{Name: "長治", Ruby: "ちょうじ", RubyInitial: "T", Year: 1104, Month: 2, Day: 10},
	{Name: "嘉承", Ruby: "かじょう（かしょう）", RubyInitial: "K", Year: 1106, Month: 4, Day: 9},
	{Name: "天仁", Ruby: "てんにん", RubyInitial: "T", Year: 1108, Month: 8, Day: 3},
	{Name: "天永", Ruby: "てんえい", RubyInitial: "T", Year: 1110, Month: 7, Day: 13},
	{Name: "永久", Ruby: "えいきゅう", RubyInitial: "E", Year: 1113, Month: 7, Day: 13},
	{Name: "元永", Ruby: "げんえい", RubyInitial: "G", Year: 1118, Month: 4, Day: 3},
	{Name: "保安", Ruby: "ほうあん", RubyInitial: "H", Year: 1120, Month: 4, Day: 10},
	{Name: "天治", Ruby: "てんじ", RubyInitial: "T", Year: 1124, Month: 4, Day: 3},
	{Name: "大治", Ruby: "だいじ（たいじ）", RubyInitial: "D", Year: 1126, Month: 1, Day: 22},
	{Name: "天承", Ruby: "てんしょう（てんじょう）", RubyInitial: "T", Year: 1131, Month: 1, Day: 29},
	{Name: "長承", Ruby: "ちょうしょう（ちょうじょう）", RubyInitial: "T", Year: 1132, Month: 8, Day: 11},
	{Name: "保延", Ruby: "ほうえん", RubyInitial: "H", Year: 1135, Month: 4, Day: 27},
	{Name: "永治", Ruby: "えいじ", RubyInitial: "E", Year: 1141, Month: 7, Day: 10},
	{Name: "康治", Ruby: "こうじ", RubyInitial: "K", Year: 1142, Month: 4, Day: 28},
	{Name: "天養", Ruby: "てんよう", RubyInitial: "T", Year: 1144, Month: 2, Day: 23},
	{Name: "久安", Ruby: "きゅうあん", RubyInitial: "K", Year: 1145, Month: 7, Day: 22},
	{Name: "仁平", Ruby: "にんぺい（にんびょう）", RubyInitial: "N", Year: 1151, Month: 1, Day: 26},
	{Name: "久寿", Ruby: "きゅうじゅ", RubyInitial: "K", Year: 1154, Month: 10, Day: 28},
	{Name: "保元", Ruby: "ほうげん", RubyInitial: "H", Year: 1156, Month: 4, Day: 27},
	{Name: "平治", Ruby: "へいじ（びょうじ）", RubyInitial: "H", Year: 1159, Month: 4, Day: 20},
	{Name: "永暦", Ruby: "えいりゃく（ようりゃく）", RubyInitial: "E", Year: 1160, Month: 1, Day: 10},
	{Name: "応保", Ruby: "おうほう", RubyInitial: "O", Year: 1161, Month: 9, Day: 4},
	{Name: "長寛", Ruby: "ちょうかん", RubyInitial: "T", Year: 1163, Month: 3, Day: 29},
	{Name: "永万", Ruby: "えいまん（ようまん）", RubyInitial: "E", Year: 1165, Month: 6, Day: 5},
	{Name: "仁安", Ruby: "にんあん（にんなん）", RubyInitial: "N", Year: 1166, Month: 8, Day: 27},
	{Name: "嘉応", Ruby: "かおう", RubyInitial: "K", Year: 1169, Month: 4, Day: 8},
	{Name: "承安", Ruby: "じょうあん（しょうあん）", RubyInitial: "J", Year: 1171, Month: 4, Day: 21},
	{Name: "安元", Ruby: "あんげん", RubyInitial: "A", Year: 1175, Month: 7, Day: 28},
	{Name: "治承", Ruby: "じしょう（じじょう）", RubyInitial: "J", Year: 1177, Month: 8, Day: 4},
	{Name: "養和", Ruby: "ようわ", RubyInitial: "Y", Year: 1181, Month: 7, Day: 14},
	{Name: "寿永", Ruby: "じゅえい", RubyInitial: "J", Year: 1182, Month: 5, Day: 27},
	{Name: "元暦", Ruby: "げんりゃく", RubyInitial: "G", Year: 1184, Month: 4, Day: 16},
	{Name: "文治", Ruby: "ぶんじ（もんじ）", RubyInitial: "B", Year: 1185, Month: 8, Day: 14},
	{Name: "正治", Ruby: "しょうじ", RubyInitial: "S", Year: 1199, Month: 4, Day: 27},
	{Name: "建仁", Ruby: "けんにん", RubyInitial: "K", Year: 1201, Month: 2, Day: 13},
	{Name: "元久", Ruby: "げんきゅう", RubyInitial: "G", Year: 1204, Month: 2, Day: 20},
	{Name: "建永", Ruby: "けんえい", RubyInitial: "K", Year: 1206, Month: 4, Day: 27},
	{Name: "承元", Ruby: "じょうげん（しょうげん）", RubyInitial: "J", Year: 1207, Month: 10, Day: 25},
	{Name: "建暦", Ruby: "けんりゃく", RubyInitial: "K", Year: 1211, Month: 3, Day: 9},
	{Name: "建保", Ruby: "けんぽう（けんほう）", RubyInitial: "K", Year: 1213, Month: 12, Day: 6},
	{Name: "承久", Ruby: "じょうきゅう（しょうきゅう）", RubyInitial: "J", Year: 1219, Month: 4, Day: 12},
	{Name: "貞応", Ruby: "じょうおう（ていおう）", RubyInitial: "J", Year: 1222, Month: 4, Day: 13},
	{Name: "元仁", Ruby: "げんにん", RubyInitial: "G", Year: 1224, Month: 11, Day: 20},
	{Name: "嘉禄", Ruby: "かろく", RubyInitial: "K", Year: 1225, Month: 4, Day: 20},
	{Name: "安貞", Ruby: "あんてい", RubyInitial: "A", Year: 1227, Month: 12, Day: 10},
	{Name: "寛喜", Ruby: "かんき", RubyInitial: "K", Year: 1229, Month: 3, Day: 5},
	{Name: "貞永", Ruby: "じょうえい（ていえい）", RubyInitial: "J", Year: 1232, Month: 4, Day: 2},
	{Name: "天福", Ruby: "てんぷく（てんふく）", RubyInitial: "T", Year: 1233, Month: 4, Day: 15},
	{Name: "文暦", Ruby: "ぶんりゃく（もんりゃく）", RubyInitial: "B", Year: 1234, Month: 11, Day: 5},
	{Name: "嘉禎", Ruby: "かてい", RubyInitial: "K", Year: 1235, Month: 9, Day: 19},
	{Name: "暦仁", Ruby: "りゃくにん（れきにん）", RubyInitial: "R", Year: 1238, Month: 11, Day: 23},
	{Name: "延応", Ruby: "えんおう（えんのう）", RubyInitial: "E", Year: 1239, Month: 2, Day: 7},
	{Name: "仁治", Ruby: "にんじ（にんち）", RubyInitial: "N", Year: 1240, Month: 7, Day: 16},
	{Name: "寛元", Ruby: "かんげん", RubyInitial: "K", Year: 1243, Month: 2, Day: 26},
	{Name: "宝治", Ruby: "ほうじ", RubyInitial: "H", Year: 1247, Month: 2, Day: 28},
	{Name: "建長", Ruby: "けんちょう", RubyInitial: "K", Year: 1249, Month: 3, Day: 18},
	{Name: "康元", Ruby: "こうげん", RubyInitial: "K", Year: 1256, Month: 10, Day: 5},
	{Name: "正嘉", Ruby: "しょうか", RubyInitial: "S", Year: 1257, Month: 3, Day: 14},
	{Name: "正元", Ruby: "しょうげん", RubyInitial: "S", Year: 1259, Month: 3, Day: 26},
	{Name: "文応", Ruby: "ぶんおう", RubyInitial: "B", Year: 1260, Month: 4, Day: 13},
	{Name: "弘長", Ruby: "こうちょう", RubyInitial: "K", Year: 1261, Month: 2, Day: 20},
	{Name: "文永", Ruby: "ぶんえい", RubyInitial: "B", Year: 1264, Month: 2, Day: 28},
	{Name: "建治", Ruby: "けんじ", RubyInitial: "K", Year: 1275, Month: 4, Day: 25},
	{Name: "弘安", Ruby: "こうあん", RubyInitial: "K", Year: 1278, Month: 2, Day: 29},
	{Name: "正応", Ruby: "しょうおう", RubyInitial: "S", Year: 1288, Month: 4, Day: 28},
	{Name: "永仁", Ruby: "えいにん", RubyInitial: "E", Year: 1293, Month: 8, Day: 5},
	{Name: "正安", Ruby: "しょうあん", RubyInitial: "S", Year: 1299, Month: 4, Day: 25},
	{Name: "乾元", Ruby: "けんげん", RubyInitial: "K", Year: 1302, Month: 11, Day: 21},
	{Name: "嘉元", Ruby: "かげん", RubyInitial: "K", Year: 1303, Month: 8, Day: 5},
	{Name: "徳治", Ruby: "とくじ", RubyInitial: "T", Year: 1306, Month: 12, Day: 14},
	{Name: "延慶", Ruby: "えんきょう（えんぎょう）", RubyInitial: "E", Year: 1308, Month: 10, Day: 9},
	{Name: "応長", Ruby: "おうちょう", RubyInitial: "O", Year: 1311, Month: 4, Day: 28},
	{Name: "正和", Ruby: "しょうわ", RubyInitial: "S", Year: 1312, Month: 3, Day: 20},
	{Name: "文保", Ruby: "ぶんぽう（ぶんほう）", RubyInitial: "B", Year: 1317, Month: 2, Day: 3},
	{Name: "元応", Ruby: "げんおう（げんのう）", RubyInitial: "G", Year: 1319, Month: 4, Day: 28},
	{Name: "元亨", Ruby: "げんこう", RubyInitial: "G", Year: 1321, Month: 2, Day: 23},
	{Name: "正中", Ruby: "しょうちゅう", RubyInitial: "S", Year: 1324, Month: 12, Day: 9},
	{Name: "嘉暦", Ruby: "かりゃく", RubyInitial: "K", Year: 1326, Month: 4, Day: 26},
	{Name: "元徳", Ruby: "げんとく", RubyInitial: "G", Year: 1329, Month: 8, Day: 29},
	{Name: "元弘", Ruby: "げんこう", RubyInitial: "G", Year: 1331, Month: 8, Day: 9},
	{Name: "正慶", Ruby: "しょうきょう（しょうけい）", RubyInitial: "S", Year: 1332, Month: 4, Day: 28},
	{Name: "建武", Ruby: "けんむ（けんぶ）", RubyInitial: "K", Year: 1334, Month: 1, Day: 29},
	{Name: "延元", Ruby: "えんげん", RubyInitial: "E", Year: 1336, Month: 2, Day: 29},
	{Name: "暦応", Ruby: "りゃくおう（れきおう）", RubyInitial: "R", Year: 1338, Month: 8, Day: 28},
	{Name: "興国", Ruby: "こうこく", RubyInitial: "K", Year: 1340, Month: 4, Day: 28},
	{Name: "康永", Ruby: "こうえい", RubyInitial: "K", Year: 1342, Month: 4, Day: 27},
	{Name: "貞和", Ruby: "じょうわ（ていわ）", RubyInitial: "J", Year: 1345, Month: 10, Day: 21},
	{Name: "正平", Ruby: "しょうへい", RubyInitial: "S", Year: 1346, Month: 12, Day: 8},
	{Name: "観応", Ruby: "かんおう（かんのう）", RubyInitial: "K", Year: 1350, Month: 2, Day: 27},
	{Name: "文和", Ruby: "ぶんな（ぶんわ）", RubyInitial: "B", Year: 1352, Month: 9, Day: 27},
	{Name: "延文", Ruby: "えんぶん", RubyInitial: "E", Year: 1356, Month: 3, Day: 28},
	{Name: "康安", Ruby: "こうあん", RubyInitial: "K", Year: 1361, Month: 3, Day: 29},
	{Name: "貞治", Ruby: "じょうじ（ていじ）", RubyInitial: "J", Year: 1362, Month: 9, Day: 23},
	{Name: "応安", Ruby: "おうあん", RubyInitial: "O", Year: 1368, Month: 2, Day: 18},
	{Name: "建徳", Ruby: "けんとく", RubyInitial: "K", Year: 1370, Month: 7, Day: 24},
	{Name: "文中", Ruby: "ぶんちゅう", RubyInitial: "B", Year: 1372, Month: 4, Day: 1},
	{Name: "天授", Ruby: "てんじゅ", RubyInitial: "T", Year: 1375, Month: 5, Day: 27},
	{Name: "永和", Ruby: "えいわ", RubyInitial: "E", Year: 1375, Month: 2, Day: 27},
	{Name: "康暦", Ruby: "こうりゃく", RubyInitial: "K", Year: 1379, Month: 3, Day: 22},
	{Name: "弘和", Ruby: "こうわ", RubyInitial: "K", Year: 1381, Month: 2, Day: 10},
	{Name: "永徳", Ruby: "えいとく", RubyInitial: "E", Year: 1381, Month: 2, Day: 24},
	{Name: "元中", Ruby: "げんちゅう", RubyInitial: "G", Year: 1384, Month: 4, Day: 28},
	{Name: "至徳", Ruby: "しとく", RubyInitial: "S", Year: 1384, Month: 2, Day: 27},
	{Name: "嘉慶", Ruby: "かきょう（かけい）", RubyInitial: "K", Year: 1387, Month: 8, Day: 23},
	{Name: "康応", Ruby: "こうおう", RubyInitial: "K", Year: 1389, Month: 2, Day: 9},
	{Name: "明徳", Ruby: "めいとく", RubyInitial: "M", Year: 1390, Month: 3, Day: 26},
	{Name: "応永", Ruby: "おうえい", RubyInitial: "O", Year: 1394, Month: 7, Day: 5},
	{Name: "正長", Ruby: "しょうちょう", RubyInitial: "S", Year: 1428, Month: 4, Day: 27},
	{Name: "永享", Ruby: "えいきょう", RubyInitial: "E", Year: 1429, Month: 9, Day: 5},
	{Name: "嘉吉", Ruby: "かきつ（かきち）", RubyInitial: "K", Year: 1441, Month: 2, Day: 17},
	{Name: "文安", Ruby: "ぶんあん", RubyInitial: "B", Year: 1444, Month: 2, Day: 5},
	{Name: "宝徳", Ruby: "ほうとく", RubyInitial: "H", Year: 1449, Month: 7, Day: 28},
	{Name: "享徳", Ruby: "きょうとく", RubyInitial: "K", Year: 1452, Month: 7, Day: 25},
	{Name: "康正", Ruby: "こうしょう", RubyInitial: "K", Year: 1455, Month: 7, Day: 25},
	{Name: "長禄", Ruby: "ちょうろく", RubyInitial: "T", Year: 1457, Month: 9, Day: 28},
	{Name: "寛正", Ruby: "かんしょう", RubyInitial: "K", Year: 1460, Month: 12, Day: 21},
	{Name: "文正", Ruby: "ぶんしょう（もんしょう）", RubyInitial: "B", Year: 1466, Month: 2, Day: 28},
	{Name: "応仁", Ruby: "おうにん", RubyInitial: "O", Year: 1467, Month: 3, Day: 5},
	{Name: "文明", Ruby: "ぶんめい", RubyInitial: "B", Year: 1469, Month: 4, Day: 28},
	{Name: "長享", Ruby: "ちょうきょう", RubyInitial: "T", Year: 1487, Month: 7, Day: 20},
	{Name: "延徳", Ruby: "えんとく", RubyInitial: "E", Year: 1489, Month: 8, Day: 21},
	{Name: "明応", Ruby: "めいおう", RubyInitial: "M", Year: 1492, Month: 7, Day: 19},
	{Name: "文亀", Ruby: "ぶんき", RubyInitial: "B", Year: 1501, Month: 2, Day: 29},
	{Name: "永正", Ruby: "えいしょう", RubyInitial: "E", Year: 1504, Month: 2, Day: 30},
	{Name: "大永", Ruby: "だいえい", RubyInitial: "D", Year: 1521, Month: 8, Day: 23},
	{Name: "享禄", Ruby: "きょうろく", RubyInitial: "K", Year: 1528, Month: 8, Day: 20},
	{Name: "天文", Ruby: "てんぶん", RubyInitial: "T", Year: 1532, Month: 7, Day: 29},
	{Name: "弘治", Ruby: "こうじ", RubyInitial: "K", Year: 1555, Month: 10, Day: 23},
	{Name: "永禄", Ruby: "えいろく", RubyInitial: "E", Year: 1558, Month: 2, Day: 28},
	{Name: "元亀", Ruby: "げんき", RubyInitial: "G", Year: 1570, Month: 4, Day: 23},
	{Name: "文禄", Ruby: "ぶんろく", RubyInitial: "B", Year: 1592, Month: 12, Day: 8},
	{Name: "慶長", Ruby: "けいちょう（きょうちょう）", RubyInitial: "K", Year: 1596, Month: 10, Day: 27},
	{Name: "寛永", Ruby: "かんえい", RubyInitial: "K", Year: 1624, Month: 2, Day: 30},
	{Name: "正保", Ruby: "しょうほう", RubyInitial: "S", Year: 1644, Month: 12, Day: 16},
	{Name: "慶安", Ruby: "けいあん", RubyInitial: "K", Year: 1648, Month: 2, Day: 15},
	{Name: "承応", Ruby: "じょうおう（しょうおう）", RubyInitial: "J", Year: 1652, Month: 9, Day: 18},
	{Name: "明暦", Ruby: "めいれき（みょうりゃく）", RubyInitial: "M", Year: 1655, Month: 4, Day: 13},
	{Name: "万治", Ruby: "まんじ", RubyInitial: "M", Year: 1658, Month: 7, Day: 23},
	{Name: "寛文", Ruby: "かんぶん", RubyInitial: "K", Year: 1661, Month: 4, Day: 25},
	{Name: "延宝", Ruby: "えんぽう", RubyInitial: "E", Year: 1673, Month: 9, Day: 21},
	{Name: "天和", Ruby: "てんな", RubyInitial: "T", Year: 1681, Month: 9, Day: 29},
	{Name: "貞享", Ruby: "じょうきょう", RubyInitial: "J", Year: 1684, Month: 2, Day: 21},
	{Name: "元禄", Ruby: "げんろく", RubyInitial: "G", Year: 1688, Month: 9, Day: 30},
	{Name: "宝永", Ruby: "ほうえい", RubyInitial: "H", Year: 1704, Month: 3, Day: 13},
	{Name: "正徳", Ruby: "しょうとく", RubyInitial: "S", Year: 1711, Month: 4, Day: 25},
	{Name: "享保", Ruby: "きょうほう（きょうほ）", RubyInitial: "K", Year: 1716, Month: 6, Day: 22},
	{Name: "元文", Ruby: "げんぶん", RubyInitial: "G", Year: 1736, Month: 4, Day: 28},
	{Name: "寛保", Ruby: "かんぽう（かんほう）", RubyInitial: "K", Year: 1741, Month: 2, Day: 27},
	{Name: "延享", Ruby: "えんきょう", RubyInitial: "E", Year: 1744, Month: 2, Day: 21},
	{Name: "寛延", Ruby: "かんえん", RubyInitial: "K", Year: 1748, Month: 7, Day: 12},
	{Name: "宝暦", Ruby: "ほうれき（ほうりゃく）", RubyInitial: "H", Year: 1751, Month: 10, Day: 27},
	{Name: "明和", Ruby: "めいわ", RubyInitial: "M", Year: 1764, Month: 6, Day: 2},
	{Name: "安永", Ruby: "あんえい", RubyInitial: "A", Year: 1772, Month: 11, Day: 16},
	{Name: "天明", Ruby: "てんめい", RubyInitial: "T", Year: 1781, Month: 4, Day: 2},
	{Name: "寛政", Ruby: "かんせい", RubyInitial: "K", Year: 1789, Month: 1, Day: 25},
	{Name: "享和", Ruby: "きょうわ", RubyInitial: "K", Year: 1801, Month: 2, Day: 5},
	{Name: "文化", Ruby: "ぶんか", RubyInitial: "B", Year: 1804, Month: 2, Day: 11},
	{Name: "文政", Ruby: "ぶんせい", RubyInitial: "B", Year: 1818, Month: 4, Day: 22},
	{Name: "天保", Ruby: "てんぽう（てんほう）", RubyInitial: "T", Year: 1830, Month: 12, Day: 10},
	{Name: "弘化", Ruby: "こうか", RubyInitial: "K", Year: 1844, Month: 12, Day: 2},
	{Name: "嘉永", Ruby: "かえい", RubyInitial: "K", Year: 1848, Month: 2, Day: 28},
	{Name: "安政", Ruby: "あんせい", RubyInitial: "A", Year: 1854, Month: 11, Day: 27},
	{Name: "万延", Ruby: "まんえん", RubyInitial: "M", Year: 1860, Month: 3, Day: 18},
	{Name: "文久", Ruby: "ぶんきゅう", RubyInitial: "B", Year: 1861, Month: 2, Day: 19},
	{Name: "元治", Ruby: "げんじ", RubyInitial: "G", Year: 1864, Month: 2, Day: 20},
	{Name: "慶応", Ruby: "けいおう", RubyInitial: "K", Year: 1865, Month: 4, Day: 7},
	{Name: "明治", Ruby: "めいじ", RubyInitial: "M", Year: 1868, Month: 9, Day: 8},
	{Name: "大正", Ruby: "たいしょう", RubyInitial: "T", Year: 1912, Month: 7, Day: 30},
	{Name: "昭和", Ruby: "しょうわ", RubyInitial: "S", Year: 1926, Month: 12, Day: 25},
	{Name: "平成", Ruby: "へいせい", RubyInitial: "H", Year: 1989, Month: 1, Day: 8},
	{Name: "令和", Ruby: "れいわ", RubyInitial: "R", Year: 2019, Month: 5, Day: 1},
}
