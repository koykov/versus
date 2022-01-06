package nlp

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type stageLP struct {
	lang, text string
}

var (
	stagesLPHW = []stageLP{
		{lang: "Afrikaans", text: "Hello Wêreld!"},
		{lang: "Albanian", text: "Përshendetje Botë!"},
		{lang: "Amharic", text: "ሰላም ልዑል!"},
		{lang: "Arabic", text: "مرحبا بالعالم!"},
		{lang: "Armenian", text: "Բարեւ աշխարհ!"},
		{lang: "Basque", text: "Kaixo Mundua!"},
		{lang: "Belarussian", text: "Прывітанне Сусвет!"},
		{lang: "Bengali", text: "ওহে বিশ্ব!"},
		{lang: "Bulgarian", text: "Здравей свят!"},
		{lang: "Catalan", text: "Hola món!"},
		{lang: "Chichewa", text: "Moni Dziko Lapansi!"},
		{lang: "Chinese", text: "你好世界！"},
		{lang: "Croatian", text: "Pozdrav svijete!"},
		{lang: "Czech", text: "Ahoj světe!"},
		{lang: "Danish", text: "Hej Verden!"},
		{lang: "Dutch", text: "Hallo Wereld!"},
		{lang: "English", text: "Hello World!"},
		{lang: "Estonian", text: "Tere maailm!"},
		{lang: "Finnish", text: "Hei maailma!"},
		{lang: "French", text: "Bonjour monde!"},
		{lang: "Frisian", text: "Hallo wrâld!"},
		{lang: "Georgian", text: "გამარჯობა მსოფლიო!"},
		{lang: "German", text: "Hallo Welt!"},
		{lang: "Greek", text: "Γειά σου Κόσμε!"},
		{lang: "Hausa", text: "Sannu Duniya!"},
		{lang: "Hebrew", text: "שלום עולם!"},
		{lang: "Hindi", text: "नमस्ते दुनिया!"},
		{lang: "Hungarian", text: "Helló Világ!"},
		{lang: "Icelandic", text: "Halló heimur!"},
		{lang: "Igbo", text: "Ndewo Ụwa!"},
		{lang: "Indonesian", text: "Halo Dunia!"},
		{lang: "Italian", text: "Ciao mondo!"},
		{lang: "Japanese", text: "こんにちは世界！"},
		{lang: "Kazakh", text: "Сәлем Әлем!"},
		{lang: "Khmer", text: "សួស្តី​ពិភពលោក!"},
		{lang: "Kyrgyz", text: "Салам дүйнө!"},
		{lang: "Lao", text: "ສະ​ບາຍ​ດີ​ຊາວ​ໂລກ!"},
		{lang: "Latvian", text: "Sveika pasaule!"},
		{lang: "Lithuanian", text: "Labas pasauli!"},
		{lang: "Luxemburgish", text: "Moien Welt!"},
		{lang: "Macedonian", text: "Здраво свету!"},
		{lang: "Malay", text: "Hai dunia!"},
		{lang: "Malayalam", text: "ഹലോ വേൾഡ്!"},
		{lang: "Mongolian", text: "Сайн уу дэлхий!"},
		{lang: "Myanmar", text: "မင်္ဂလာပါကမ္ဘာလောက!"},
		{lang: "Nepali", text: "नमस्कार संसार!"},
		{lang: "Norwegian", text: "Hei Verden!"},
		{lang: "Pashto", text: "سلام نړی!"},
		{lang: "Persian", text: "سلام دنیا!"},
		{lang: "Polish", text: "Witaj świecie!"},
		{lang: "Portuguese", text: "Olá Mundo!"},
		{lang: "Punjabi", text: "ਸਤਿ ਸ੍ਰੀ ਅਕਾਲ ਦੁਨਿਆ!"},
		{lang: "Romanian", text: "Salut Lume!"},
		{lang: "Russian", text: "Привет мир!"},
		{lang: "Scots Gaelic", text: "Hàlo a Shaoghail!"},
		{lang: "Serbian", text: "Здраво Свете!"},
		{lang: "Sesotho", text: "Lefatše Lumela!"},
		{lang: "Sinhala", text: "හෙලෝ වර්ල්ඩ්!"},
		{lang: "Slovenian", text: "Pozdravljen svet!"},
		{lang: "Spanish", text: "¡Hola Mundo!"},
		{lang: "Sundanese", text: "Halo Dunya!"},
		{lang: "Swahili", text: "Salamu Dunia!"},
		{lang: "Swedish", text: "Hej världen!"},
		{lang: "Tajik", text: "Салом Ҷаҳон!"},
		{lang: "Thai", text: "สวัสดีชาวโลก!"},
		{lang: "Turkish", text: "Selam Dünya!"},
		{lang: "Ukrainian", text: "Привіт Світ!"},
		{lang: "Uzbek", text: "Salom Dunyo!"},
		{lang: "Vietnamese", text: "Chào thế giới!"},
		{lang: "Welsh", text: "Helo Byd!"},
		{lang: "Xhosa", text: "Molo Lizwe!"},
		{lang: "Yiddish", text: "העלא וועלט!"},
		{lang: "Yoruba", text: "Mo ki O Ile Aiye!"},
		{lang: "Zulu", text: "Sawubona Mhlaba!"},
	}
	stageLPDS = make([]stageLP, 0, 3e5)
)

func init() {
	f, err := os.Open("testdata/language.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = f.Close() }()

	reader := csv.NewReader(f)
	for i := 0; ; i++ {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if i == 0 {
			// CSV header
			continue
		}
		stageLPDS = append(stageLPDS, stageLP{
			lang: rec[1],
			text: rec[0],
		})
	}
}
