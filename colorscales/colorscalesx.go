package colorscales

import (
	"image/color"
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

type ColorScaleX []color.RGBA

func (c ColorScaleX) GetColor(val float64) color.RGBA {
	maxIdx := len(c) - 1
	idx := int(math.Round(float64(maxIdx) * val))
	return c[idx]
}

func Hex(hex string) color.RGBA {
	c, err := colorful.Hex(hex)
	if err != nil {
		panic(err)
	}
	r, g, b := c.RGB255()
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func RGB(R float64, G float64, B float64) color.RGBA {
	return color.RGBA{
		R: uint8(255 * R),
		G: uint8(255 * G),
		B: uint8(255 * B),
		A: 255,
	}
}

var PuBu9 = ColorScaleX{
	color.RGBA{255, 247, 251, 255},
	color.RGBA{236, 231, 242, 255},
	color.RGBA{208, 209, 230, 255},
	color.RGBA{166, 189, 219, 255},
	color.RGBA{116, 169, 207, 255},
	color.RGBA{54, 144, 192, 255},
	color.RGBA{5, 112, 176, 255},
	color.RGBA{4, 90, 141, 255},
	color.RGBA{2, 56, 88, 255},
}

var YlGn9 = ColorScaleX{
	color.RGBA{255, 255, 229, 255},
	color.RGBA{247, 252, 185, 255},
	color.RGBA{217, 240, 163, 255},
	color.RGBA{173, 221, 142, 255},
	color.RGBA{120, 198, 121, 255},
	color.RGBA{65, 171, 93, 255},
	color.RGBA{35, 132, 67, 255},
	color.RGBA{0, 104, 55, 255},
	color.RGBA{0, 69, 41, 255},
}

var GnBu9 = ColorScaleX{
	color.RGBA{247, 252, 240, 255},
	color.RGBA{224, 243, 219, 255},
	color.RGBA{204, 235, 197, 255},
	color.RGBA{168, 221, 181, 255},
	color.RGBA{123, 204, 196, 255},
	color.RGBA{78, 179, 211, 255},
	color.RGBA{43, 140, 190, 255},
	color.RGBA{8, 104, 172, 255},
	color.RGBA{8, 64, 129, 255},
}

var Aquamarine = ColorScaleX{ // not good
	RGB(0.68069, 0.735561, 0.850004),
	RGB(0.762233, 0.803087, 0.873397),
	RGB(0.738044, 0.801736, 0.849202),
	RGB(0.653851, 0.759285, 0.801252),
	RGB(0.555381, 0.703512, 0.753383),
	RGB(0.488359, 0.662194, 0.729431),
	RGB(0.498514, 0.663109, 0.75323),
	RGB(0.631571, 0.734035, 0.848615),
}

var Thermal = ColorScaleX{
	RGB(0.015556013331540799, 0.13824424546464084, 0.2018108864558305),
	RGB(0.01620183633850513, 0.14105074288662173, 0.20897651254408078),
	RGB(0.01685648942708359, 0.14382701436218343, 0.21623868044760436),
	RGB(0.017526400647825284, 0.14657172506679966, 0.223599699683326),
	RGB(0.01821871873545745, 0.14928346382380617, 0.23106186935280404),
	RGB(0.018941378369021544, 0.15196073496435802, 0.2386274839825403),
	RGB(0.01969967580211434, 0.15460145133852177, 0.2463049741539925),
	RGB(0.020503315127140914, 0.157203779485642, 0.2540971092002057),
	RGB(0.02136720981691052, 0.15976645008644966, 0.26199914590368084),
	RGB(0.022303406774609756, 0.16228755369440137, 0.27001321141123913),
	RGB(0.02332520459934066, 0.16476505482550094, 0.2781413941166567),
	RGB(0.024447278236844824, 0.16719678243498234, 0.28638572702162585),
	RGB(0.025685816291650603, 0.16958042033133544, 0.2947481661145329),
	RGB(0.027058671851373926, 0.17191349765849873, 0.30323056309247975),
	RGB(0.028585527643947253, 0.1741933796245943, 0.3118346316168388),
	RGB(0.03028807626872338, 0.17641725871610026, 0.32056190614579816),
	RGB(0.032190216097775136, 0.17858214671238393, 0.32941369222834893),
	RGB(0.03431826321592867, 0.1806848679100377, 0.33839100697329016),
	RGB(0.03670117943457189, 0.18272205408291475, 0.34749450822554256),
	RGB(0.039370815947033844, 0.1846901418460877, 0.35672441079352574),
	RGB(0.042304741163601826, 0.18658537326432548, 0.3660803878810173),
	RGB(0.0454412843165637, 0.18840380075243815, 0.37556145569341703),
	RGB(0.048798894592565444, 0.19014129756004725, 0.3851658390246699),
	RGB(0.05238564999909288, 0.1917935754201579, 0.39489081550577426),
	RGB(0.05620896927989653, 0.19335621127078723, 0.4047325361348015),
	RGB(0.060275609713176435, 0.19482468533005756, 0.4146858197468099),
	RGB(0.0645951913792415, 0.1961877478810098, 0.4247714594384869),
	RGB(0.06917293885235362, 0.19744583346763395, 0.4349572770994128),
	RGB(0.07401397924875777, 0.19859437346292247, 0.44523225498656266),
	RGB(0.07912632580773252, 0.1996251445642635, 0.45559655554449024),
	RGB(0.08452074659970571, 0.20052842125954135, 0.46605087322539746),
	RGB(0.09019392390291772, 0.2013079448693052, 0.47654787791351255),
	RGB(0.09616430834359416, 0.20194725295148624, 0.4871044455010897),
	RGB(0.10242539874668813, 0.2024520213616346, 0.4976646226688279),
	RGB(0.10899442702495121, 0.20280888558325388, 0.5082270885561683),
	RGB(0.11585973516266682, 0.2030273547727351, 0.5187245256305992),
	RGB(0.12304242634434309, 0.20309379703068342, 0.5291483757425789),
	RGB(0.13052766946136096, 0.20302178164248064, 0.5394201195013325),
	RGB(0.13830991252152922, 0.202819561373477, 0.5494767764101642),
	RGB(0.14637970601506853, 0.20250017692775046, 0.5592461312000727),
	RGB(0.15471863262808744, 0.20208506808526183, 0.56864226209659),
	RGB(0.16329704990775173, 0.20160546411933639, 0.5775676861911989),
	RGB(0.17207282317497322, 0.20110272912781796, 0.5859183845130114),
	RGB(0.18099176221810628, 0.20062705242843215, 0.593591818607864),
	RGB(0.18999021995742105, 0.20023415835648706, 0.6004970669758479),
	RGB(0.19899974419406802, 0.19998029273013757, 0.6065651516258429),
	RGB(0.2079529779081165, 0.19991642567386514, 0.6117570630961893),
	RGB(0.21678951602164598, 0.20008303402379002, 0.6160673924886362),
	RGB(0.22546042571503655, 0.20050670816546434, 0.6195227918364496),
	RGB(0.23393063020234023, 0.20119919840576933, 0.6221760825969699),
	RGB(0.24217906839492115, 0.20215871627233217, 0.6240979346009031),
	RGB(0.2501971277034635, 0.20337273145067664, 0.625368243352074),
	RGB(0.2579861074778474, 0.20482134175919747, 0.6260687943604991),
	RGB(0.2655544225081862, 0.20648046943023657, 0.6262779744769026),
	RGB(0.2729150389729214, 0.20832445576036884, 0.626067576730132),
	RGB(0.2800833915559314, 0.21032791942986045, 0.6255013265069015),
	RGB(0.28706750740330417, 0.21246913864072653, 0.6246420009551492),
	RGB(0.29388513606279454, 0.2147254304183257, 0.6235380183242871),
	RGB(0.30055772134613623, 0.21707565146197155, 0.6222251903624),
	RGB(0.3070843773107185, 0.219505980201521, 0.6207554984633753),
	RGB(0.3134916285549605, 0.22199831003913523, 0.6191452159520037),
	RGB(0.3197798380866584, 0.22454203566029374, 0.6174343435010864),
	RGB(0.3259695527216342, 0.22712412378631391, 0.6156329324235469),
	RGB(0.3320579089479565, 0.2297373625924511, 0.6137771467682859),
	RGB(0.3380660052498126, 0.2323715331127139, 0.6118660583919268),
	RGB(0.34399174205251376, 0.23502150765272706, 0.6099280657295181),
	RGB(0.34984606856187234, 0.23768089458696007, 0.6079694872801727),
	RGB(0.355638458018118, 0.2403444043084655, 0.6059953497614897),
	RGB(0.36136862955584764, 0.24300891538364394, 0.6040241385974032),
	RGB(0.3670434468015163, 0.24567078321674146, 0.6020607616318823),
	RGB(0.37267087856159326, 0.24832669647222555, 0.6001059364008434),
	RGB(0.3782547954216435, 0.2509743622345631, 0.5981656274234779),
	RGB(0.38379608226889167, 0.25361228346436304, 0.5962498766549472),
	RGB(0.3892988345534371, 0.2562387736058999, 0.594361804997624),
	RGB(0.3947690991030737, 0.2588520700720948, 0.5924996351953922),
	RGB(0.4002098946086878, 0.26145106597703455, 0.5906660913059953),
	RGB(0.40562409093791035, 0.2640348396856425, 0.5888633104220149),
	RGB(0.4110144223076953, 0.266602624643428, 0.5870929096766684),
	RGB(0.41638321728018857, 0.26915383449755276, 0.585356595138856),
	RGB(0.42173132310073763, 0.2716882525666768, 0.5836583632684496),
	RGB(0.42706343453096, 0.27420503062640733, 0.5819941042354171),
	RGB(0.43238181048971136, 0.2767038027770707, 0.5803639256565889),
	RGB(0.4376886243188534, 0.27918426251430606, 0.5787676418773346),
	RGB(0.44298596966437376, 0.2816461519867727, 0.5772048027503156),
	RGB(0.44827586549027204, 0.28408925312591227, 0.5756747186760452),
	RGB(0.45356026028639157, 0.2865133804028728, 0.5741764823913097),
	RGB(0.45884103552590655, 0.288918375000015, 0.5727089879419176),
	RGB(0.46412000842361484, 0.2913041002126984, 0.5712709472315609),
	RGB(0.4693989340425177, 0.2936704379215549, 0.5698609044983501),
	RGB(0.47467950679335086, 0.29601728599654775, 0.5684772490345112),
	RGB(0.47996336136971124, 0.2983445565121623, 0.5671182264323321),
	RGB(0.4852520731601698, 0.30065217466834365, 0.565781948610285),
	RGB(0.4905471581781304, 0.30294007832472053, 0.5644664028469251),
	RGB(0.49585007255021457, 0.30520821806641846, 0.5631694600262432),
	RGB(0.5011622116043697, 0.30745655772877606, 0.5618888822762295),
	RGB(0.5064849085998189, 0.3096850753156841, 0.560622330162177),
	RGB(0.5118194331420761, 0.31189376425240567, 0.5593673695773669),
	RGB(0.5171669893275794, 0.31408263491877814, 0.5581214784559523),
	RGB(0.52252871366388, 0.31625171641284816, 0.5568820534159195),
	RGB(0.5279056728126165, 0.31840105849846967, 0.5556464164236832),
	RGB(0.5332988612036094, 0.3205307336933521, 0.5544118215561066),
	RGB(0.5387091985692399, 0.3226408394566178, 0.5531754619204082),
	RGB(0.5441375274486375, 0.3247315004373091, 0.551934476777456),
	RGB(0.5495846107110636, 0.32680287074753156, 0.5506859588994192),
	RGB(0.5550511291471262, 0.3288551362262009, 0.549426962178624),
	RGB(0.5605376791749926, 0.33088851666170865, 0.5481545094908906),
	RGB(0.5660447559653562, 0.3329032737201366, 0.546865629096665),
	RGB(0.5715723572408056, 0.3348998704862336, 0.5455581199822508),
	RGB(0.577121310428943, 0.336878448643229, 0.5442280104688512),
	RGB(0.5826918609018116, 0.33883938062594116, 0.5428722728323092),
	RGB(0.588284161005071, 0.3407830824811499, 0.5414878910241607),
	RGB(0.5938982692324898, 0.3427100153754325, 0.5400718694374621),
	RGB(0.5995341497366565, 0.344620686926, 0.5386212416598933),
	RGB(0.605191672186175, 0.34651565234936754, 0.537133079145422),
	RGB(0.6108706119725975, 0.34839551542678215, 0.5356044997333839),
	RGB(0.6165706507631435, 0.35026092928951547, 0.534032675943247),
	RGB(0.6222913773881608, 0.3521125970311436, 0.5324148429745681),
	RGB(0.6280322890453093, 0.35395127215788635, 0.5307483063447554),
	RGB(0.6337927927958686, 0.35577775889170465, 0.5290304491020394),
	RGB(0.6395721854014133, 0.3575929224264805, 0.5272587829694529),
	RGB(0.6453697498455219, 0.3593976455700172, 0.5254307632896803),
	RGB(0.6511846984697031, 0.36119285495278036, 0.5235439096149941),
	RGB(0.6570160963803247, 0.36297955796848924, 0.5215959730192854),
	RGB(0.662862928356282, 0.364758813268274, 0.51958481096754),
	RGB(0.6687241012633005, 0.366531730670847, 0.5175083912494202),
	RGB(0.6745984464659754, 0.3682994710922304, 0.5153647953022982),
	RGB(0.6804847221871189, 0.3700632465224874, 0.513152220938634),
	RGB(0.6863816157645579, 0.3718243200760849, 0.5108689845040223),
	RGB(0.6922877457570544, 0.37358400614116233, 0.5085135225034096),
	RGB(0.6982016638532076, 0.3753436706510841, 0.5060843927435033),
	RGB(0.7041218565401858, 0.3771047314992945, 0.5035802750493091),
	RGB(0.7100468845018955, 0.3788685911907691, 0.5009996451582281),
	RGB(0.7159752518731293, 0.38063670179786635, 0.4983410633235164),
	RGB(0.721905025476091, 0.3824107548202066, 0.4956041045679996),
	RGB(0.7278344332015346, 0.3841923879692217, 0.492787948314018),
	RGB(0.7337616407032358, 0.3859832960516178, 0.4898918977085311),
	RGB(0.7396847510296571, 0.38778523244612817, 0.4869153794944617),
	RGB(0.7456018038508585, 0.38960001075795564, 0.48385794404037125),
	RGB(0.7515107742785041, 0.3914295066368485, 0.4807192656291402),
	RGB(0.7574095712835626, 0.3932756597383702, 0.4774991431076147),
	RGB(0.7632960357235062, 0.3951404758009958, 0.4741975009994948),
	RGB(0.7691679379984674, 0.3970260288041012, 0.4708143911830061),
	RGB(0.775022975364167, 0.39893446316360887, 0.46734999523312),
	RGB(0.7808587689384376, 0.40086799591301037, 0.4638046275250513),
	RGB(0.7866728604480875, 0.4028289188075652, 0.46017873919138147),
	RGB(0.7924627087737056, 0.404819600278673, 0.45647292301911163),
	RGB(0.7982256863620216, 0.4068424871537119, 0.4526879193650572),
	RGB(0.8039590755885861, 0.4089001060440321, 0.4488246231577861),
	RGB(0.8096600651680348, 0.41099506429037364, 0.44488409204143814),
	RGB(0.8153258420561761, 0.413130007258232, 0.44086720057071005),
	RGB(0.8209533846301936, 0.4153077124896902, 0.4367753702901882),
	RGB(0.8265394587133883, 0.4175310847974927, 0.4326106584748535),
	RGB(0.8320808476618816, 0.4198030498289601, 0.42837489120561767),
	RGB(0.8375742298072512, 0.42212660638942423, 0.4240701127101389),
	RGB(0.8430161781421202, 0.4245048216497959, 0.41969859978974927),
	RGB(0.8484031611801809, 0.42694082479318446, 0.4152628771038664),
	RGB(0.8537315452228416, 0.42943779890321876, 0.410765733130543),
	RGB(0.8589975982777037, 0.4319989708985696, 0.4062102365684474),
	RGB(0.8641974958823695, 0.43462759932722356, 0.4015997528869853),
	RGB(0.8693273290890854, 0.4373269598520186, 0.39693796066844417),
	RGB(0.874383114859191, 0.440100328287653, 0.3922288673203609),
	RGB(0.8793608090992057, 0.44295096109057586, 0.3874768236699652),
	RGB(0.884256322540282, 0.4458820732584556, 0.38268653688832266),
	RGB(0.8890655396175203, 0.44889681366634854, 0.3778630811333164),
	RGB(0.8937843693689974, 0.45199824049733944, 0.37301160144257006),
	RGB(0.8984086715918496, 0.45518929004870207, 0.3681382748140558),
	RGB(0.9029343899688814, 0.4584727367754223, 0.3632494215164466),
	RGB(0.9073575586033109, 0.4618511663085561, 0.3583516614288674),
	RGB(0.9116743277656134, 0.4653269421603677, 0.3534519722899168),
	RGB(0.9158809992051784, 0.4689021723136149, 0.34855766996190674),
	RGB(0.9199740625655836, 0.47257867672703685, 0.3436763818638891),
	RGB(0.9239502320744639, 0.4763579567194723, 0.3388160133952918),
	RGB(0.9278064825504417, 0.48024116725702753, 0.3339847073943023),
	RGB(0.9315400836682339, 0.48422909318759705, 0.3291907969273191),
	RGB(0.9351486313601244, 0.48832213043652256, 0.3244427519772631),
	RGB(0.9386300752174523, 0.49252027309055885, 0.3197491208755644),
	RGB(0.9419827407976891, 0.49682310715260664, 0.3151184675888816),
	RGB(0.9452053458453827, 0.5012298115494861, 0.31055930620928),
	RGB(0.9482970095990234, 0.5057391667267316, 0.306080034187972),
	RGB(0.9512572545755481, 0.5103495708806106, 0.30168886598175115),
	RGB(0.9540860004899644, 0.5150590635745327, 0.2973937688352702),
	RGB(0.9567835502644689, 0.5198653561843506, 0.2932024023936338),
	RGB(0.9593505683914355, 0.5247658683349885, 0.289122063726538),
	RGB(0.9617880522174246, 0.5297577692490467, 0.2851596391521898),
	RGB(0.9640972969908582, 0.5348380227429672, 0.28132156398730335),
	RGB(0.9662798557460374, 0.5400034344900019, 0.27761379103491124),
	RGB(0.9683374952663609, 0.5452507001277228, 0.2740417682745931),
	RGB(0.9702721494708056, 0.5505764528210163, 0.27061042586199074),
	RGB(0.9720889605050128, 0.5559744240639273, 0.26733230572851496),
	RGB(0.9737882391976737, 0.5614431075234689, 0.2642051375197136),
	RGB(0.9753722517954447, 0.5669791681305654, 0.26123210436073513),
	RGB(0.9768433321430202, 0.5725792962307278, 0.2584161030939571),
	RGB(0.978203755076359, 0.5782403184217296, 0.2557595328984586),
	RGB(0.9794556966363962, 0.5839592236834147, 0.25326432412228395),
	RGB(0.9806011989875686, 0.58973318383579, 0.2509319710696592),
	RGB(0.9816445119805891, 0.595557644502105, 0.24876746140531764),
	RGB(0.9825957966513862, 0.6014235301335163, 0.24678388862201112),
	RGB(0.9834484089207011, 0.6073354334336248, 0.24496691674741777),
	RGB(0.9842037798764917, 0.6132913094304376, 0.24331643158613364),
	RGB(0.984863090176353, 0.6192893396808168, 0.2418320393357376),
	RGB(0.9854335855591604, 0.6253232095344845, 0.2405207443125048),
	RGB(0.9859266054527726, 0.6313839151833325, 0.23939249556833203),
	RGB(0.986329078808121, 0.6374805817839867, 0.23842850052710096),
	RGB(0.9866412575432023, 0.6436122274808632, 0.2376272705850792),
	RGB(0.9868743031654043, 0.6497702193582077, 0.2369977788720154),
	RGB(0.9870387195949218, 0.6559468364085242, 0.2365460008515471),
	RGB(0.9871160198111447, 0.6621544902577213, 0.2362516434857292),
	RGB(0.9871054832314209, 0.6683929123881801, 0.23611250931701094),
	RGB(0.9870452158568818, 0.6746361467497961, 0.23615516408559517),
	RGB(0.9869008071972435, 0.6809069196703932, 0.23634896083634738),
	RGB(0.9866692878063003, 0.6872064512263131, 0.2366900769944742),
	RGB(0.9863929517876754, 0.6935072010148546, 0.2372025181179538),
	RGB(0.9860358081545987, 0.6998320059248219, 0.23785913082733845),
	RGB(0.9855936903889445, 0.7061828277187244, 0.23865534671373595),
	RGB(0.985116676068181, 0.7125285611733752, 0.23961378464975586),
	RGB(0.9845515404596056, 0.718901634870693, 0.24070274652398377),
	RGB(0.9839219208542443, 0.7252873073536283, 0.2419304971970358),
	RGB(0.9832411577231265, 0.7316776102841333, 0.24329811021760384),
	RGB(0.9824686360853926, 0.7380960010628512, 0.2447853973017023),
	RGB(0.9816675490572134, 0.7445058075421049, 0.2464118597749721),
	RGB(0.9807811261674898, 0.7509394073137782, 0.24815276245045803),
	RGB(0.9798377198237013, 0.7573804700508822, 0.2500137146942527),
	RGB(0.978839265458651, 0.7638279008720278, 0.2519905378906656),
	RGB(0.9777574675384567, 0.770296904750519, 0.2540716600569991),
	RGB(0.9766479350869351, 0.7767573771217612, 0.256266407083425),
	RGB(0.9754356977304895, 0.7832491288133777, 0.2585534361426328),
	RGB(0.9742107025732686, 0.7897246981741067, 0.26094721036697005),
	RGB(0.9728835679779427, 0.7962305102434533, 0.26342619832724345),
	RGB(0.9715297353896645, 0.8027276438790406, 0.2659994820979044),
	RGB(0.970086134837361, 0.8092479763321948, 0.2686528986078252),
	RGB(0.9686057099481874, 0.8157648637457888, 0.27138992938800466),
	RGB(0.967043069160827, 0.8223006128282059, 0.27420073280666335),
	RGB(0.9654377156641302, 0.8288358297838602, 0.2770858440538391),
	RGB(0.9637524161478507, 0.8353882936883591, 0.28003762625495093),
	RGB(0.9620231537308009, 0.8419407901977009, 0.28305544485779377),
	RGB(0.960210488502386, 0.8485116373841136, 0.2861325388392048),
	RGB(0.9583576301078739, 0.8550807238511935, 0.28926815137068224),
	RGB(0.956411752870728, 0.8616719638797333, 0.2924557053792813),
	RGB(0.9544348419615948, 0.8682572954460299, 0.2956947937778981),
	RGB(0.9523487086523742, 0.8748712522095133, 0.29897881641458274),
	RGB(0.9502464564071694, 0.8814728114137494, 0.3023077648025994),
	RGB(0.948014130728471, 0.8881110895508604, 0.30567519634043294),
	RGB(0.9457819798963852, 0.8947301765899248, 0.30908112084085126),
	RGB(0.943419864615293, 0.9013849121338212, 0.3125200216900484),
	RGB(0.941028615969167, 0.9080328518084608, 0.3159906396320989),
	RGB(0.9385331275251968, 0.9147047487143539, 0.3194893313796646),
	RGB(0.935971108343295, 0.921384812614979, 0.3230138415459749),
	RGB(0.9333375985220992, 0.928074834184405, 0.3265615469169363),
	RGB(0.930591565440312, 0.9347905093445135, 0.33012998103815777),
	RGB(0.9278142496620382, 0.9414998557883345, 0.3337168327083648),
	RGB(0.9248899165263079, 0.948247019741531, 0.33731996777333784),
	RGB(0.9219411401959985, 0.9549849147932995, 0.3409370492608515),
	RGB(0.9188613878011584, 0.9617532494963948, 0.34456629972324265),
	RGB(0.9156931782520092, 0.9685354904254356, 0.34820569007261515),
	RGB(0.9124490701578419, 0.9753266872784462, 0.3518533597970245),
	RGB(0.9090418416674036, 0.9821574063216706, 0.3555078064299531),
}

