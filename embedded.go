/*
 * Copyright (c) 2014 Kurt Jung (Gmail: kurt.w.jung)
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gofpdf

// Embedded standard fonts

import (
	"strings"
)

var embeddedFontList = map[string]string{
	"courierBI":    `{"Tp":"Core","Name":"Courier-BoldOblique","Up":-100,"Ut":50,"I":256,"Cw":[600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600]}`,
	"courierB":     `{"Tp":"Core","Name":"Courier-Bold","Up":-100,"Ut":50,"I":256,"Cw":[600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600]}`,
	"courierI":     `{"Tp":"Core","Name":"Courier-Oblique","Up":-100,"Ut":50,"I":256,"Cw":[600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600]}`,
	"courier":      `{"Tp":"Core","Name":"Courier","Up":-100,"Ut":50,"I":256,"Cw":[600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600,600]}`,
	"helveticaBI":  `{"Tp":"Core","Name":"Helvetica-BoldOblique","Up":-100,"Ut":50,"Cw":[278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,333,474,556,556,889,722,238,333,333,389,584,278,333,278,278,556,556,556,556,556,556,556,556,556,556,333,333,584,584,584,611,975,722,722,722,722,667,611,778,722,278,556,722,611,833,722,778,667,778,722,667,611,722,667,944,667,667,611,333,278,333,584,556,333,556,611,556,611,556,333,611,611,278,278,556,278,889,611,611,611,611,389,556,333,611,556,778,556,556,500,389,280,389,584,350,556,350,278,556,500,1000,556,556,333,1000,667,333,1000,350,611,350,350,278,278,500,500,350,556,1000,333,1000,556,333,944,350,500,667,278,333,556,556,556,556,280,556,333,737,370,556,584,333,737,333,400,584,333,333,333,611,556,278,333,333,365,556,834,834,834,611,722,722,722,722,722,722,1000,722,667,667,667,667,278,278,278,278,722,722,778,778,778,778,778,584,778,722,722,722,722,667,667,611,556,556,556,556,556,556,889,556,556,556,556,556,278,278,278,278,611,611,611,611,611,611,611,584,611,611,611,611,611,556,611,556]}`,
	"helveticaB":   `{"Tp":"Core","Name":"Helvetica-Bold","Up":-100,"Ut":50,"Cw":[278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,333,474,556,556,889,722,238,333,333,389,584,278,333,278,278,556,556,556,556,556,556,556,556,556,556,333,333,584,584,584,611,975,722,722,722,722,667,611,778,722,278,556,722,611,833,722,778,667,778,722,667,611,722,667,944,667,667,611,333,278,333,584,556,333,556,611,556,611,556,333,611,611,278,278,556,278,889,611,611,611,611,389,556,333,611,556,778,556,556,500,389,280,389,584,350,556,350,278,556,500,1000,556,556,333,1000,667,333,1000,350,611,350,350,278,278,500,500,350,556,1000,333,1000,556,333,944,350,500,667,278,333,556,556,556,556,280,556,333,737,370,556,584,333,737,333,400,584,333,333,333,611,556,278,333,333,365,556,834,834,834,611,722,722,722,722,722,722,1000,722,667,667,667,667,278,278,278,278,722,722,778,778,778,778,778,584,778,722,722,722,722,667,667,611,556,556,556,556,556,556,889,556,556,556,556,556,278,278,278,278,611,611,611,611,611,611,611,584,611,611,611,611,611,556,611,556]}`,
	"helveticaI":   `{"Tp":"Core","Name":"Helvetica-Oblique","Up":-100,"Ut":50,"Cw":[278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,355,556,556,889,667,191,333,333,389,584,278,333,278,278,556,556,556,556,556,556,556,556,556,556,278,278,584,584,584,556,1015,667,667,722,722,667,611,778,722,278,500,667,556,833,722,778,667,778,722,667,611,722,667,944,667,667,611,278,278,278,469,556,333,556,556,500,556,556,278,556,556,222,222,500,222,833,556,556,556,556,333,500,278,556,500,722,500,500,500,334,260,334,584,350,556,350,222,556,333,1000,556,556,333,1000,667,333,1000,350,611,350,350,222,222,333,333,350,556,1000,333,1000,500,333,944,350,500,667,278,333,556,556,556,556,260,556,333,737,370,556,584,333,737,333,400,584,333,333,333,556,537,278,333,333,365,556,834,834,834,611,667,667,667,667,667,667,1000,722,667,667,667,667,278,278,278,278,722,722,778,778,778,778,778,584,778,722,722,722,722,667,667,611,556,556,556,556,556,556,889,500,556,556,556,556,278,278,278,278,556,556,556,556,556,556,556,584,611,556,556,556,556,500,556,500]}`,
	"helvetica":    `{"Tp":"Core","Name":"Helvetica","Up":-100,"Ut":50,"Cw":[278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,278,355,556,556,889,667,191,333,333,389,584,278,333,278,278,556,556,556,556,556,556,556,556,556,556,278,278,584,584,584,556,1015,667,667,722,722,667,611,778,722,278,500,667,556,833,722,778,667,778,722,667,611,722,667,944,667,667,611,278,278,278,469,556,333,556,556,500,556,556,278,556,556,222,222,500,222,833,556,556,556,556,333,500,278,556,500,722,500,500,500,334,260,334,584,350,556,350,222,556,333,1000,556,556,333,1000,667,333,1000,350,611,350,350,222,222,333,333,350,556,1000,333,1000,500,333,944,350,500,667,278,333,556,556,556,556,260,556,333,737,370,556,584,333,737,333,400,584,333,333,333,556,537,278,333,333,365,556,834,834,834,611,667,667,667,667,667,667,1000,722,667,667,667,667,278,278,278,278,722,722,778,778,778,778,778,584,778,722,722,722,722,667,667,611,556,556,556,556,556,556,889,500,556,556,556,556,278,278,278,278,556,556,556,556,556,556,556,584,611,556,556,556,556,500,556,500]}`,
	"timesBI":      `{"Tp":"Core","Name":"Times-BoldItalic","Up":-100,"Ut":50,"Cw":[250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,389,555,500,500,833,778,278,333,333,500,570,250,333,250,278,500,500,500,500,500,500,500,500,500,500,333,333,570,570,570,500,832,667,667,667,722,667,667,722,778,389,500,667,611,889,722,722,611,722,667,556,611,722,667,889,667,611,611,333,278,333,570,500,333,500,500,444,500,444,333,500,556,278,278,500,278,778,556,500,500,500,389,389,278,556,444,667,500,444,389,348,220,348,570,350,500,350,333,500,500,1000,500,500,333,1000,556,333,944,350,611,350,350,333,333,500,500,350,500,1000,333,1000,389,333,722,350,389,611,250,389,500,500,500,500,220,500,333,747,266,500,606,333,747,333,400,570,300,300,333,576,500,250,333,300,300,500,750,750,750,500,667,667,667,667,667,667,944,667,667,667,667,667,389,389,389,389,722,722,722,722,722,722,722,570,722,722,722,722,722,611,611,500,500,500,500,500,500,500,722,444,444,444,444,444,278,278,278,278,500,556,500,500,500,500,500,570,500,556,556,556,556,444,500,444]}`,
	"timesB":       `{"Tp":"Core","Name":"Times-Bold","Up":-100,"Ut":50,"Cw":[250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,333,555,500,500,1000,833,278,333,333,500,570,250,333,250,278,500,500,500,500,500,500,500,500,500,500,333,333,570,570,570,500,930,722,667,722,722,667,611,778,778,389,500,778,667,944,722,778,611,778,722,556,667,722,722,1000,722,722,667,333,278,333,581,500,333,500,556,444,556,444,333,500,556,278,333,556,278,833,556,500,556,556,444,389,333,556,500,722,500,500,444,394,220,394,520,350,500,350,333,500,500,1000,500,500,333,1000,556,333,1000,350,667,350,350,333,333,500,500,350,500,1000,333,1000,389,333,722,350,444,722,250,333,500,500,500,500,220,500,333,747,300,500,570,333,747,333,400,570,300,300,333,556,540,250,333,300,330,500,750,750,750,500,722,722,722,722,722,722,1000,722,667,667,667,667,389,389,389,389,722,722,778,778,778,778,778,570,778,722,722,722,722,722,611,556,500,500,500,500,500,500,722,444,444,444,444,444,278,278,278,278,500,556,500,500,500,500,500,570,500,556,556,556,556,500,556,500]}`,
	"timesI":       `{"Tp":"Core","Name":"Times-Italic","Up":-100,"Ut":50,"Cw":[250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,333,420,500,500,833,778,214,333,333,500,675,250,333,250,278,500,500,500,500,500,500,500,500,500,500,333,333,675,675,675,500,920,611,611,667,722,611,611,722,722,333,444,667,556,833,667,722,611,722,611,500,556,722,611,833,611,556,556,389,278,389,422,500,333,500,500,444,500,444,278,500,500,278,278,444,278,722,500,500,500,500,389,389,278,500,444,667,444,444,389,400,275,400,541,350,500,350,333,500,556,889,500,500,333,1000,500,333,944,350,556,350,350,333,333,556,556,350,500,889,333,980,389,333,667,350,389,556,250,389,500,500,500,500,275,500,333,760,276,500,675,333,760,333,400,675,300,300,333,500,523,250,333,300,310,500,750,750,750,500,611,611,611,611,611,611,889,667,611,611,611,611,333,333,333,333,722,667,722,722,722,722,722,675,722,722,722,722,722,556,611,500,500,500,500,500,500,500,667,444,444,444,444,444,278,278,278,278,500,500,500,500,500,500,500,675,500,500,500,500,500,444,500,444]}`,
	"times":        `{"Tp":"Core","Name":"Times-Roman","Up":-100,"Ut":50,"Cw":[250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,250,333,408,500,500,833,778,180,333,333,500,564,250,333,250,278,500,500,500,500,500,500,500,500,500,500,278,278,564,564,564,444,921,722,667,667,722,611,556,722,722,333,389,722,611,889,722,722,556,722,667,556,611,722,722,944,722,722,611,333,278,333,469,500,333,444,500,444,500,444,333,500,500,278,278,500,278,778,500,500,500,500,333,389,278,500,500,722,500,500,444,480,200,480,541,350,500,350,333,500,444,1000,500,500,333,1000,556,333,889,350,611,350,350,333,333,444,444,350,500,1000,333,980,389,333,722,350,444,722,250,333,500,500,500,500,200,500,333,760,276,500,564,333,760,333,400,564,300,300,333,500,453,250,333,300,310,500,750,750,750,444,722,722,722,722,722,722,889,667,611,611,611,611,333,333,333,333,722,722,722,722,722,722,722,564,722,722,722,722,722,722,556,500,444,444,444,444,444,444,667,444,444,444,444,444,278,278,278,278,500,500,500,500,500,500,500,564,500,500,500,500,500,500,500,500]}`,
	"zapfdingbats": `{"Tp":"Core","Name":"ZapfDingbats","Up":-100,"Ut":50,"Cw":[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,278,974,961,974,980,719,789,790,791,690,960,939,549,855,911,933,911,945,974,755,846,762,761,571,677,763,760,759,754,494,552,537,577,692,786,788,788,790,793,794,816,823,789,841,823,833,816,831,923,744,723,749,790,792,695,776,768,792,759,707,708,682,701,826,815,789,789,707,687,696,689,786,787,713,791,785,791,873,761,762,762,759,759,892,892,788,784,438,138,277,415,392,392,668,668,0,390,390,317,317,276,276,509,509,410,410,234,234,334,334,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,732,544,544,910,667,760,760,776,595,694,626,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,788,894,838,1016,458,748,924,748,918,927,928,928,834,873,828,924,924,917,930,931,463,883,836,836,867,867,696,696,874,0,874,760,946,771,865,771,888,967,888,831,873,927,970,918,0]}`,
}

func (f *Fpdf) coreFontReader(familyStr, styleStr string) (r *strings.Reader) {
	key := familyStr + styleStr
	str, ok := embeddedFontList[key]
	if ok {
		r = strings.NewReader(str)
	} else {
		f.SetErrorf("could not locate \"%s\" among embedded core font definition files", key)
	}
	return
}

var embeddedMapList = map[string]string{
	"cp1252": `
!00 U+0000 .notdef
!01 U+0001 .notdef
!02 U+0002 .notdef
!03 U+0003 .notdef
!04 U+0004 .notdef
!05 U+0005 .notdef
!06 U+0006 .notdef
!07 U+0007 .notdef
!08 U+0008 .notdef
!09 U+0009 .notdef
!0A U+000A .notdef
!0B U+000B .notdef
!0C U+000C .notdef
!0D U+000D .notdef
!0E U+000E .notdef
!0F U+000F .notdef
!10 U+0010 .notdef
!11 U+0011 .notdef
!12 U+0012 .notdef
!13 U+0013 .notdef
!14 U+0014 .notdef
!15 U+0015 .notdef
!16 U+0016 .notdef
!17 U+0017 .notdef
!18 U+0018 .notdef
!19 U+0019 .notdef
!1A U+001A .notdef
!1B U+001B .notdef
!1C U+001C .notdef
!1D U+001D .notdef
!1E U+001E .notdef
!1F U+001F .notdef
!20 U+0020 space
!21 U+0021 exclam
!22 U+0022 quotedbl
!23 U+0023 numbersign
!24 U+0024 dollar
!25 U+0025 percent
!26 U+0026 ampersand
!27 U+0027 quotesingle
!28 U+0028 parenleft
!29 U+0029 parenright
!2A U+002A asterisk
!2B U+002B plus
!2C U+002C comma
!2D U+002D hyphen
!2E U+002E period
!2F U+002F slash
!30 U+0030 zero
!31 U+0031 one
!32 U+0032 two
!33 U+0033 three
!34 U+0034 four
!35 U+0035 five
!36 U+0036 six
!37 U+0037 seven
!38 U+0038 eight
!39 U+0039 nine
!3A U+003A colon
!3B U+003B semicolon
!3C U+003C less
!3D U+003D equal
!3E U+003E greater
!3F U+003F question
!40 U+0040 at
!41 U+0041 A
!42 U+0042 B
!43 U+0043 C
!44 U+0044 D
!45 U+0045 E
!46 U+0046 F
!47 U+0047 G
!48 U+0048 H
!49 U+0049 I
!4A U+004A J
!4B U+004B K
!4C U+004C L
!4D U+004D M
!4E U+004E N
!4F U+004F O
!50 U+0050 P
!51 U+0051 Q
!52 U+0052 R
!53 U+0053 S
!54 U+0054 T
!55 U+0055 U
!56 U+0056 V
!57 U+0057 W
!58 U+0058 X
!59 U+0059 Y
!5A U+005A Z
!5B U+005B bracketleft
!5C U+005C backslash
!5D U+005D bracketright
!5E U+005E asciicircum
!5F U+005F underscore
!60 U+0060 grave
!61 U+0061 a
!62 U+0062 b
!63 U+0063 c
!64 U+0064 d
!65 U+0065 e
!66 U+0066 f
!67 U+0067 g
!68 U+0068 h
!69 U+0069 i
!6A U+006A j
!6B U+006B k
!6C U+006C l
!6D U+006D m
!6E U+006E n
!6F U+006F o
!70 U+0070 p
!71 U+0071 q
!72 U+0072 r
!73 U+0073 s
!74 U+0074 t
!75 U+0075 u
!76 U+0076 v
!77 U+0077 w
!78 U+0078 x
!79 U+0079 y
!7A U+007A z
!7B U+007B braceleft
!7C U+007C bar
!7D U+007D braceright
!7E U+007E asciitilde
!7F U+007F .notdef
!80 U+20AC Euro
!82 U+201A quotesinglbase
!83 U+0192 florin
!84 U+201E quotedblbase
!85 U+2026 ellipsis
!86 U+2020 dagger
!87 U+2021 daggerdbl
!88 U+02C6 circumflex
!89 U+2030 perthousand
!8A U+0160 Scaron
!8B U+2039 guilsinglleft
!8C U+0152 OE
!8E U+017D Zcaron
!91 U+2018 quoteleft
!92 U+2019 quoteright
!93 U+201C quotedblleft
!94 U+201D quotedblright
!95 U+2022 bullet
!96 U+2013 endash
!97 U+2014 emdash
!98 U+02DC tilde
!99 U+2122 trademark
!9A U+0161 scaron
!9B U+203A guilsinglright
!9C U+0153 oe
!9E U+017E zcaron
!9F U+0178 Ydieresis
!A0 U+00A0 space
!A1 U+00A1 exclamdown
!A2 U+00A2 cent
!A3 U+00A3 sterling
!A4 U+00A4 currency
!A5 U+00A5 yen
!A6 U+00A6 brokenbar
!A7 U+00A7 section
!A8 U+00A8 dieresis
!A9 U+00A9 copyright
!AA U+00AA ordfeminine
!AB U+00AB guillemotleft
!AC U+00AC logicalnot
!AD U+00AD hyphen
!AE U+00AE registered
!AF U+00AF macron
!B0 U+00B0 degree
!B1 U+00B1 plusminus
!B2 U+00B2 twosuperior
!B3 U+00B3 threesuperior
!B4 U+00B4 acute
!B5 U+00B5 mu
!B6 U+00B6 paragraph
!B7 U+00B7 periodcentered
!B8 U+00B8 cedilla
!B9 U+00B9 onesuperior
!BA U+00BA ordmasculine
!BB U+00BB guillemotright
!BC U+00BC onequarter
!BD U+00BD onehalf
!BE U+00BE threequarters
!BF U+00BF questiondown
!C0 U+00C0 Agrave
!C1 U+00C1 Aacute
!C2 U+00C2 Acircumflex
!C3 U+00C3 Atilde
!C4 U+00C4 Adieresis
!C5 U+00C5 Aring
!C6 U+00C6 AE
!C7 U+00C7 Ccedilla
!C8 U+00C8 Egrave
!C9 U+00C9 Eacute
!CA U+00CA Ecircumflex
!CB U+00CB Edieresis
!CC U+00CC Igrave
!CD U+00CD Iacute
!CE U+00CE Icircumflex
!CF U+00CF Idieresis
!D0 U+00D0 Eth
!D1 U+00D1 Ntilde
!D2 U+00D2 Ograve
!D3 U+00D3 Oacute
!D4 U+00D4 Ocircumflex
!D5 U+00D5 Otilde
!D6 U+00D6 Odieresis
!D7 U+00D7 multiply
!D8 U+00D8 Oslash
!D9 U+00D9 Ugrave
!DA U+00DA Uacute
!DB U+00DB Ucircumflex
!DC U+00DC Udieresis
!DD U+00DD Yacute
!DE U+00DE Thorn
!DF U+00DF germandbls
!E0 U+00E0 agrave
!E1 U+00E1 aacute
!E2 U+00E2 acircumflex
!E3 U+00E3 atilde
!E4 U+00E4 adieresis
!E5 U+00E5 aring
!E6 U+00E6 ae
!E7 U+00E7 ccedilla
!E8 U+00E8 egrave
!E9 U+00E9 eacute
!EA U+00EA ecircumflex
!EB U+00EB edieresis
!EC U+00EC igrave
!ED U+00ED iacute
!EE U+00EE icircumflex
!EF U+00EF idieresis
!F0 U+00F0 eth
!F1 U+00F1 ntilde
!F2 U+00F2 ograve
!F3 U+00F3 oacute
!F4 U+00F4 ocircumflex
!F5 U+00F5 otilde
!F6 U+00F6 odieresis
!F7 U+00F7 divide
!F8 U+00F8 oslash
!F9 U+00F9 ugrave
!FA U+00FA uacute
!FB U+00FB ucircumflex
!FC U+00FC udieresis
!FD U+00FD yacute
!FE U+00FE thorn
!FF U+00FF ydieresis
	`,
}
