package main

import (
	"image/color"
	"time"

	"github.com/spearson78/tinyamifont"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/akashi"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/aldous"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/andover"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/avantg"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/basel"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/bigfont"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/blackboard32"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/blacksh"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/bookman"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/broadway"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/bubbles"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/c"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/camelot"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/celtic"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/celtich"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/cleanii"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/cobweb"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/colossal"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/countdown"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/courier"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/crown"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/cryfont4"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/cupertino"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/cutout"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/deborah"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/dempseyoutline"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/flow"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/future"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/geneva"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/ghost"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/gothic"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/ham"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/hercules"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/hollywood"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/ibm5"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/internat"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/italic"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/iva"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/lineal"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/lines"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/loderunner"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/losangel"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/microsoft"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/monaco"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/news"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/outline"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/paloalto"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/parkave"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/patricia"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/peignot"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/premiere"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/prime"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/pudgy"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/rangers"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/roma"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/romamed"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/sanfran"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/schalt18"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/serpentine"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/shadow"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/shadow2"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/sherbert"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/silicon"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/slant"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/slope"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/sparklers"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/stencil"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/stop"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/stripes"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/swansong"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/twothousandandone"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/typerite"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/vancouver"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/vari1"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/video1"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/video2"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/video3"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/whitesh"
	"github.com/spearson78/tinyamifont/fonts/fletcherfonts/zschool"
	"github.com/spearson78/tinyamifont/fonts/hkfonts8/digitalquill"
	"github.com/spearson78/tinyamifont/fonts/hkfonts8/foobar"
	"github.com/spearson78/tinyamifont/fonts/hkfonts8/marble"
	"github.com/spearson78/tinyamifont/fonts/hkfonts8/spaceage"
	"github.com/spearson78/tinyamifont/fonts/hkfonts8/ungabunga"
	"github.com/spearson78/tinyamifont/fonts/magicwb/xcourier"
	"github.com/spearson78/tinyamifont/fonts/magicwb/xen"
	"github.com/spearson78/tinyamifont/fonts/magicwb/xhelvetica"
	"github.com/spearson78/tinyamifont/fonts/micronano/micropaz"
	"github.com/spearson78/tinyamifont/fonts/micronano/nanopaz"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/alexismachine"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/antisocial"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/arrested"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/bar"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/basic"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/bdr"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/bluecons"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/bsquer"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/catchtwentythree"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/cerebral"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/clea"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/count"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/crossfire"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/crossfire2"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/cul8r"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/cutelittle"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/digital"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/edge"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/edie"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/fat"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/hysteria"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/indy"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/jetman8"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/karel"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/knightlore"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/knightmare"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/lifeforce"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/limaer"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/livingstone"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/logo"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/longine"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/maiden"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/malice"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/malicegothic2"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/marlboro"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/master"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/master6"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/medothic"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/micro"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/millius"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/murkybedlam"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/mykenos"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/nightmare"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/ocean1"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/ocean2"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/packey"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/parola"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/phenon"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/rafaello"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/rebelyell"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/ride"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/rideon"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/rockadisca"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/rodia"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/runaway"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/script"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/seventyfivec"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/shadowlord"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/sidewize"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/sorcery"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/spectrum"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/starquake"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/tales"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/theft"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/trovolco"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/ttt"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/turboiii"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/veryslim"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/victim"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/vulcan2"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/wednesday"
	"github.com/spearson78/tinyamifont/fonts/mlcfont/xecutor"
	"github.com/spearson78/tinyamifont/fonts/newwebfont/alfa3"
	"github.com/spearson78/tinyamifont/fonts/newwebfont/alfathin3"
	"github.com/spearson78/tinyamifont/fonts/newwebfont/voyager"
	"github.com/spearson78/tinyamifont/fonts/sicon/sicon"
	"github.com/spearson78/tinyamifont/fonts/supersans/supersansserif"
	"github.com/spearson78/tinyamifont/fonts/wbfont/clean"
	"github.com/spearson78/tinyamifont/fonts/wbfont/iconfont"
)

type demoFont struct {
	Name     string
	FontData tinyamifont.FontData
}

func main() {

	fonts := []demoFont{
		{"aldous.Regular12pt", aldous.Regular12pt},
		{"aldous.Regular18pt", aldous.Regular18pt},
		{"aldous.Regular24pt", aldous.Regular24pt},
		{"aldous.Regular36pt", aldous.Regular36pt},
		{"andover.Regular12pt", andover.Regular12pt},
		{"avantg.Regular10pt", avantg.Regular10pt},
		{"avantg.Regular12pt", avantg.Regular12pt},
		{"avantg.Regular14pt", avantg.Regular14pt},
		{"avantg.Regular18pt", avantg.Regular18pt},
		{"avantg.Regular24pt", avantg.Regular24pt},
		{"basel.Regular48pt", basel.Regular48pt},
		{"bigfont.Regular16pt", bigfont.Regular16pt},
		{"blacksh.Regular48pt", blacksh.Regular48pt},
		{"bookman.Regular10pt", bookman.Regular10pt},
		{"bookman.Regular12pt", bookman.Regular12pt},
		{"bookman.Regular14pt", bookman.Regular14pt},
		{"bookman.Regular18pt", bookman.Regular18pt},
		{"bookman.Regular24pt", bookman.Regular24pt},
		{"broadway.Regular24pt", broadway.Regular24pt},
		{"bubbles.Regular14pt", bubbles.Regular14pt},
		{"bubbles.Regular24pt", bubbles.Regular24pt},
		{"camelot.Regular12pt", camelot.Regular12pt},
		{"camelot.Regular24pt", camelot.Regular24pt},
		{"celtic.Regular12pt", celtic.Regular12pt},
		{"celtic.Regular18pt", celtic.Regular18pt},
		{"celtic.Regular24pt", celtic.Regular24pt},
		{"celtic.Regular36pt", celtic.Regular36pt},
		{"celtic.Regular48pt", celtic.Regular48pt},
		{"celtic.Regular72pt", celtic.Regular72pt},
		{"celtich.Regular72pt", celtich.Regular72pt},
		{"cleanii.Regular12pt", cleanii.Regular12pt},
		{"cleanii.Regular16pt", cleanii.Regular16pt},
		{"cleanii.Regular24pt", cleanii.Regular24pt},
		{"cleanii.Regular8pt", cleanii.Regular8pt},
		{"courier.Regular10pt", courier.Regular10pt},
		{"courier.Regular20pt", courier.Regular20pt},
		{"cryfont4.Regular8pt", cryfont4.Regular8pt},
		{"deborah.Regular45pt", deborah.Regular45pt},
		{"future.Regular12pt", future.Regular12pt},
		{"future.Regular18pt", future.Regular18pt},
		{"future.Regular24pt", future.Regular24pt},
		{"future.Regular9pt", future.Regular9pt},
		{"geneva.Regular36pt", geneva.Regular36pt},
		{"ghost.Regular27pt", ghost.Regular27pt},
		{"ghost.Regular35pt", ghost.Regular35pt},
		{"ham.Regular14pt", ham.Regular14pt},
		{"hollywood.Regular24pt", hollywood.Regular24pt},
		{"internat.Regular12pt", internat.Regular12pt},
		{"internat.Regular24pt", internat.Regular24pt},
		{"iva.Regular50pt", iva.Regular50pt},
		{"lineal.Regular18pt", lineal.Regular18pt},
		{"loderunner.Regular12pt", loderunner.Regular12pt},
		{"losangel.Regular12pt", losangel.Regular12pt},
		{"losangel.Regular24pt", losangel.Regular24pt},
		{"microsoft.Regular10pt", microsoft.Regular10pt},
		{"monaco.Regular10pt", monaco.Regular10pt},
		{"monaco.Regular12pt", monaco.Regular12pt},
		{"monaco.Regular14pt", monaco.Regular14pt},
		{"monaco.Regular18pt", monaco.Regular18pt},
		{"monaco.Regular20pt", monaco.Regular20pt},
		{"monaco.Regular24pt", monaco.Regular24pt},
		{"monaco.Regular8pt", monaco.Regular8pt},
		{"monaco.Regular9pt", monaco.Regular9pt},
		{"news.Regular32pt", news.Regular32pt},
		{"paloalto.Regular12pt", paloalto.Regular12pt},
		{"paloalto.Regular24pt", paloalto.Regular24pt},
		{"parkave.Regular18pt", parkave.Regular18pt},
		{"patricia.Regular54pt", patricia.Regular54pt},
		{"peignot.Regular48pt", peignot.Regular48pt},
		{"premiere.Regular24pt", premiere.Regular24pt},
		{"prime.Regular8pt", prime.Regular8pt},
		{"prime.Regular9pt", prime.Regular9pt},
		{"rangers.Regular18pt", rangers.Regular18pt},
		{"sanfran.Regular18pt", sanfran.Regular18pt},
		{"silicon.Regular12pt", silicon.Regular12pt},
		{"stencil.Regular12pt", stencil.Regular12pt},
		{"stencil.Regular24pt", stencil.Regular24pt},
		{"swansong.Regular12pt", swansong.Regular12pt},
		{"twothousandandone.Regular8pt", twothousandandone.Regular8pt},
		{"vancouver.Regular10pt", vancouver.Regular10pt},
		{"vancouver.Regular12pt", vancouver.Regular12pt},
		{"vancouver.Regular14pt", vancouver.Regular14pt},
		{"vancouver.Regular18pt", vancouver.Regular18pt},
		{"vancouver.Regular20pt", vancouver.Regular20pt},
		{"vancouver.Regular24pt", vancouver.Regular24pt},
		{"vancouver.Regular9pt", vancouver.Regular9pt},
		{"video1.Regular50pt", video1.Regular50pt},
		{"video2.Regular49pt", video2.Regular49pt},
		{"video3.Regular54pt", video3.Regular54pt},
		{"whitesh.Regular48pt", whitesh.Regular48pt},
		{"akashi.Regular22pt", akashi.Regular22pt},
		//{"basic.Regular28pt", basic.Regular28pt},
		//{"basic.Regular32pt", basic.Regular32pt},
		//{"big.Regular16pt", big.Regular16pt},
		{"blackboard32.Regular32brpt", blackboard32.Regular32brpt},
		{"c.Regular6pt", c.Regular6pt},
		{"cobweb.Regular32pt", cobweb.Regular32pt},
		{"colossal.Regular8pt", colossal.Regular8pt},
		{"countdown.Regular16pt", countdown.Regular16pt},
		{"crown.Regular8pt", crown.Regular8pt},
		{"cupertino.Regular16pt", cupertino.Regular16pt},
		{"cupertino.Regular32pt", cupertino.Regular32pt},
		{"cutout.Regular19pt", cutout.Regular19pt},
		{"cutout.Regular21pt", cutout.Regular21pt},
		{"dempseyoutline.Regular38pt", dempseyoutline.Regular38pt},
		{"dempseyoutline.Regular54pt", dempseyoutline.Regular54pt},
		//{"dempseyoutline.Regular77pt", dempseyoutline.Regular77pt},
		//{"dempseyoutline.Regular83pt", dempseyoutline.Regular83pt},
		//{"digital.Regular41pt", digital.Regular41pt},
		//{"dotty.Regular16bpt", dotty.Regular16bpt},
		{"flow.Regular8pt", flow.Regular8pt},
		{"gothic.Regular8pt", gothic.Regular8pt},
		//{"helvetica.Regular21pt", helvetica.Regular21pt},
		//{"helvetica.Regular27pt", helvetica.Regular27pt},
		//{"helvetica.Regular62pt", helvetica.Regular62pt},
		//{"helvetica.Regular82pt", helvetica.Regular82pt},
		{"hercules.Regular6pt", hercules.Regular6pt},
		{"ibm5.Regular8pt", ibm5.Regular8pt},
		{"italic.Regular13pt", italic.Regular13pt},
		{"italic.Regular15pt", italic.Regular15pt},
		{"lines.Regular21pt", lines.Regular21pt},
		{"outline.Regular21pt", outline.Regular21pt},
		{"pudgy.Regular8pt", pudgy.Regular8pt},
		{"roma.Regular13pt", roma.Regular13pt},
		{"roma.Regular15pt", roma.Regular15pt},
		{"roma.Regular17pt", roma.Regular17pt},
		{"roma.Regular19pt", roma.Regular19pt},
		{"romamed.Regular19pt", romamed.Regular19pt},
		{"schalt18.Regular20pt", schalt18.Regular20pt},
		{"serpentine.Regular15pt", serpentine.Regular15pt},
		{"shadow.Regular21pt", shadow.Regular21pt},
		{"shadow2.Regular26pt", shadow2.Regular26pt},
		{"sherbert.Regular48pt", sherbert.Regular48pt},
		{"sherbert.Regular61pt", sherbert.Regular61pt},
		{"slant.Regular8pt", slant.Regular8pt},
		{"slope.Regular16pt", slope.Regular16pt},
		{"sparklers.Regular30pt", sparklers.Regular30pt},
		{"stop.Regular8pt", stop.Regular8pt},
		{"stripes.Regular21pt", stripes.Regular21pt},
		{"typerite.Regular15pt", typerite.Regular15pt},
		{"typerite.Regular17pt", typerite.Regular17pt},
		{"vari1.Regular32pt", vari1.Regular32pt},
		{"zschool.Regular19pt", zschool.Regular19pt},
		{"zschool.Regular30pt", zschool.Regular30pt},
		{"zschool.Regular38pt", zschool.Regular38pt},
		{"digitalquill.Regular11pt", digitalquill.Regular11pt},
		{"digitalquill.Regular16pt", digitalquill.Regular16pt},
		{"digitalquill.Regular9pt", digitalquill.Regular9pt},
		{"marble.Regular12pt", marble.Regular12pt},
		{"spaceage.Regular13pt", spaceage.Regular13pt},
		{"ungabunga.Regular13pt", ungabunga.Regular13pt},
		{"foobar.Regular11pt", foobar.Regular11pt},
		{"foobar.Regular13pt", foobar.Regular13pt},
		{"xcourier.Regular11pt", xcourier.Regular11pt},
		{"xcourier.Regular13pt", xcourier.Regular13pt},
		{"xcourier.Regular15pt", xcourier.Regular15pt},
		{"xen.Regular11pt", xen.Regular11pt},
		{"xen.Regular8pt", xen.Regular8pt},
		{"xen.Regular9pt", xen.Regular9pt},
		{"xhelvetica.Regular11pt", xhelvetica.Regular11pt},
		{"xhelvetica.Regular13pt", xhelvetica.Regular13pt},
		{"xhelvetica.Regular9pt", xhelvetica.Regular9pt},
		{"micropaz.Regular6pt", micropaz.Regular6pt},
		{"nanopaz.Regular6pt", nanopaz.Regular6pt},
		{"catchtwentythree.Regular8pt", catchtwentythree.Regular8pt},
		{"alexismachine.Regular8pt", alexismachine.Regular8pt},
		{"antisocial.Regular8pt", antisocial.Regular8pt},
		{"arrested.Regular8pt", arrested.Regular8pt},
		{"bar.Regular8pt", bar.Regular8pt},
		{"basic.Regular17pt", basic.Regular17pt},
		{"bdr.Regular8pt", bdr.Regular8pt},
		{"bluecons.Regular8pt", bluecons.Regular8pt},
		{"bsquer.Regular8pt", bsquer.Regular8pt},
		{"cerebral.Regular8pt", cerebral.Regular8pt},
		{"clea.Regular8pt", clea.Regular8pt},
		{"count.Regular8pt", count.Regular8pt},
		{"crossfire.Regular8pt", crossfire.Regular8pt},
		{"crossfire2.Regular8pt", crossfire2.Regular8pt},
		{"cul8r.Regular8pt", cul8r.Regular8pt},
		{"cutelittle.Regular7pt", cutelittle.Regular7pt},
		{"digital.Regular8pt", digital.Regular8pt},
		{"edge.Regular8pt", edge.Regular8pt},
		{"edie.Regular8pt", edie.Regular8pt},
		{"fat.Regular8pt", fat.Regular8pt},
		{"hysteria.Regular8pt", hysteria.Regular8pt},
		{"indy.Regular8pt", indy.Regular8pt},
		{"jetman8.Regular8pt", jetman8.Regular8pt},
		{"karel.Regular24pt", karel.Regular24pt},
		{"knightlore.Regular8pt", knightlore.Regular8pt},
		{"knightmare.Regular8pt", knightmare.Regular8pt},
		{"lifeforce.Regular8pt", lifeforce.Regular8pt},
		{"limaer.Regular8pt", limaer.Regular8pt},
		{"livingstone.Regular8pt", livingstone.Regular8pt},
		{"logo.Regular8pt", logo.Regular8pt},
		{"longine.Regular16pt", longine.Regular16pt},
		{"maiden.Regular8pt", maiden.Regular8pt},
		{"malice.Regular8pt", malice.Regular8pt},
		{"malicegothic2.Regular19pt", malicegothic2.Regular19pt},
		{"marlboro.Regular32pt", marlboro.Regular32pt},
		{"master.Regular8pt", master.Regular8pt},
		{"master6.Regular8pt", master6.Regular8pt},
		{"medothic.Regular8pt", medothic.Regular8pt},
		{"micro.Regular5pt", micro.Regular5pt},
		{"millius.Regular8pt", millius.Regular8pt},
		{"murkybedlam.Regular8pt", murkybedlam.Regular8pt},
		{"mykenos.Regular8pt", mykenos.Regular8pt},
		{"nightmare.Regular8pt", nightmare.Regular8pt},
		{"ocean1.Regular8pt", ocean1.Regular8pt},
		{"ocean2.Regular8pt", ocean2.Regular8pt},
		{"packey.Regular8bpt", packey.Regular8bpt},
		{"parola.Regular8pt", parola.Regular8pt},
		{"phenon.Regular8pt", phenon.Regular8pt},
		{"rafaello.Regular8pt", rafaello.Regular8pt},
		{"rebelyell.Regular8pt", rebelyell.Regular8pt},
		{"ride.Regular8pt", ride.Regular8pt},
		{"rideon.Regular8bpt", rideon.Regular8bpt},
		{"rockadisca.Regular8pt", rockadisca.Regular8pt},
		{"rodia.Regular8pt", rodia.Regular8pt},
		{"runaway.Regular8pt", runaway.Regular8pt},
		{"script.Regular8pt", script.Regular8pt},
		{"seventyfivec.Regular8pt", seventyfivec.Regular8pt},
		{"shadowlord.Regular8pt", shadowlord.Regular8pt},
		{"sidewize.Regular8pt", sidewize.Regular8pt},
		{"sorcery.Regular8pt", sorcery.Regular8pt},
		{"spectrum.Regular8pt", spectrum.Regular8pt},
		{"starquake.Regular8pt", starquake.Regular8pt},
		{"tales.Regular8pt", tales.Regular8pt},
		{"theft.Regular6pt", theft.Regular6pt},
		{"trovolco.Regular24pt", trovolco.Regular24pt},
		{"ttt.Regular8pt", ttt.Regular8pt},
		{"turboiii.Regular8pt", turboiii.Regular8pt},
		{"veryslim.Regular8pt", veryslim.Regular8pt},
		{"victim.Regular8pt", victim.Regular8pt},
		{"vulcan2.Regular8pt", vulcan2.Regular8pt},
		{"wednesday.Regular8pt", wednesday.Regular8pt},
		{"xecutor.Regular8pt", xecutor.Regular8pt},
		{"alfa3.Regular11fpt", alfa3.Regular11fpt},
		{"alfa3.Regular13fpt", alfa3.Regular13fpt},
		{"alfa3.Regular15fpt", alfa3.Regular15fpt},
		{"alfathin3.Regular13fpt", alfathin3.Regular13fpt},
		{"alfathin3.Regular15fpt", alfathin3.Regular15fpt},
		{"voyager.Regular11pt", voyager.Regular11pt},
		{"voyager.Regular13pt", voyager.Regular13pt},
		{"voyager.Regular16pt", voyager.Regular16pt},
		{"voyager.Regular18pt", voyager.Regular18pt},
		{"voyager.Regular20pt", voyager.Regular20pt},
		{"voyager.Regular23pt", voyager.Regular23pt},
		{"voyager.Regular9pt", voyager.Regular9pt},
		{"sicon.Regular8pt", sicon.Regular8pt},
		{"supersansserif.Regular11pt", supersansserif.Regular11pt},
		{"supersansserif.Regular13pt", supersansserif.Regular13pt},
		{"supersansserif.Regular14pt", supersansserif.Regular14pt},
		{"supersansserif.Regular9pt", supersansserif.Regular9pt},
		{"iconfont.Regular9pt", iconfont.Regular9pt},
		{"clean.Regular10pt", clean.Regular10pt},
		{"clean.Regular11pt", clean.Regular11pt},
		{"clean.Regular13pt", clean.Regular13pt},
		{"clean.Regular9pt", clean.Regular9pt},
	}

	display := initHardware()

	time.Sleep(1000 * time.Millisecond)

	pageSleep := 500 * time.Millisecond

	_, height := display.Size()
	var style tinyamifont.Style
	for {
		style++
		y := int16(0)
		display.ClearDisplay()

		for _, demoFont := range fonts {

			//println("FONT", demoFont.Name)
			font := tinyamifont.MustLoadFont(demoFont.FontData)

			if y+(int16(font.Ysize)*2)+4+5 >= height {
				time.Sleep(pageSleep)
				display.ClearDisplay()
				y = 0
			}

			y += int16(font.Ysize) + 2
			tinyamifont.PrintString(&font, display, demoFont.Name, 20, y, color.RGBA{250, 250, 250, 255}, style)
			y += int16(font.Ysize) + 2
			tinyamifont.PrintString(&font, display, "The quick brown fox", 20, y, color.RGBA{250, 250, 250, 255}, style)
			y += int16(font.Ysize) + 2
			tinyamifont.PrintString(&font, display, "jumped over the lazy dog", 20, y, color.RGBA{250, 250, 250, 255}, style)
			y += 5

		}

		time.Sleep(pageSleep)
	}
}
