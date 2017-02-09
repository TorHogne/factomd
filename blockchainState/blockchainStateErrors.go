package blockchainState

import (
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

func (bs *BlockchainState) HandlePreBlockErrors(dBlockHash interfaces.IHash) error {
	bs.Init()
	hash := dBlockHash.String()

	//ECBlock discontinuity
	switch hash {
	case "7bd9704ec9e8e5238fe7669cc235b56c109052e7a8be4281b2a985922c24e038": //22881
		h, _ := primitives.HexToHash("f3fa09154eeba632d88542a20637a8b3eb86758f0a46b5cca364483b815e6b45")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("0d3b181f15177156e8381ffc21203b31114676902947987170ac73008d20a352")
		bs.ECBlockHeadHash = h
		break
	case "3957f3f30fa71d1586e0fe2c575a11775fc52bce08326084ab459b2cea28d7fb": //22884
		h, _ := primitives.HexToHash("c3e8fda5dfea0d7ea0dad5d59a6f6bd10ca18cae5f57545913915b635cb1667c")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("19c73168fed48738bfdc53c63dc10beb18cdf12d63339f3eefb9b2bdc35af1ee")
		bs.ECBlockHeadHash = h
		break
	case "59b2a648c5a32ec3dcf2254e6e50e492e7443e390edb36f3908fecd870e20ec5": //22941
		h, _ := primitives.HexToHash("7962567ff7b9cd4e9283e8120d20af03c221205035702ea517e4a9cf7874b828")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("916a5e9a653ae50597f6a98515687621925d4403b0582aa054f3f5363b784a46")
		bs.ECBlockHeadHash = h
		break
	case "4d19fde10fdaf3873b4dd827e5360a6ee9b848d73d469b274e6bd24a5ffa2096": //22950
		h, _ := primitives.HexToHash("04faa69b36313712c6d166a7f74af56cfe846128af20ee422126b900b10f0ddb")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("d369b72492ee9c91424be0b0866959c4c02bd3f888abc3fa71359a6b260f4c10")
		bs.ECBlockHeadHash = h
		break
	case "0c93039ccd3b9d685c354323c356ed56e12447b8663d69d539ed17104a1bbc4a": //22977
		h, _ := primitives.HexToHash("750bafb62bd5783c3459e65ec15a80d56c9007399ab8c78e3205085a99c13b4b")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("dc77eed730b88c0eca3e3d8157eb0376de2a9fe1bcf36acf9fa78f5994a83517")
		bs.ECBlockHeadHash = h
		break
	case "4e184bf1326e9cd5cea23d5a409bbbcf45450754d6e3c16938cb709d2cfebf1c": //22979
		h, _ := primitives.HexToHash("45146a4b32883031e1ffbfadc25688375a7d78964224a77e018c19224a8f9445")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("8ab1001168b5427c13064c849fa3e33e0ada01579a45a468aa40c419b7235e0d")
		bs.ECBlockHeadHash = h
		break
	case "8441b4e44f5b1696e17696ddfae81448b77f10347a72931a7b2f896a2deef063": //23269
		h, _ := primitives.HexToHash("88a7ee2eb0bbc3d66b64531ee877620d8b13b327628f47d914e19f6ee9075233")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("f878c85e0c22b4980423971ece2dd43636d67d46eeabbf8e394443d0d30cd885")
		bs.ECBlockHeadHash = h
		break
	case "043a2153b6905ffc9560c5aa36569010dbcba97fb86edb00315e1eb1f063cfa8": //31461
		h, _ := primitives.HexToHash("8df30f50b9e00fdbff33971f1de2b0ccdae68433064e50398368f8a232c26b7e")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("f3ef19fdac26528336e8edefd46ad1b56fe6ba824e855c7fead879c8e6c2cf65")
		bs.ECBlockHeadHash = h
		break
	case "ceb025d7b3e1e75adab24696d8c9c6f352f8b672e80ecca16631cce886b56001": //49237
		h, _ := primitives.HexToHash("a39f0ea8c6db864615afcd112450dbe2be01ae8d8e13f8be15a50dcef820b10d")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("651d784531c3d46e62911de5f0298a2e256b6308e9d91e10697b7228fe67a979")
		bs.ECBlockHeadHash = h
		break
	case "8a2fd42e9fa3bcaff68ecf3396991954688c766bba18d24cf4d8dbaaba016ca0": //50159
		h, _ := primitives.HexToHash("d3363256fc494c78b76218b936979b31ea83bc6c3684946ed990d0052151ea48")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("bc7e49eb3cb7e97d8fae8fc3542590687d5b477118f269a7d557b2e99902dce8")
		bs.ECBlockHeadHash = h
		break
	case "cb32117379e6cae4097b29ef7e9a40995fa21f38b0723c0c3ec2d1f489b137cd": //54355
		h, _ := primitives.HexToHash("4550edb9b41c2e92a4dd76d0648f3d45f852918113d3281fb522fb890b79ddcb")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("fe66be1fee8230c9efe3893de052f61f0575b23780e0fccdce676b2a087fd578")
		bs.ECBlockHeadHash = h
		break
	case "446bdeeb4cbb811ced0d784c1384279da9dd452cc9fbdfde58eeb11b7869d10a": //57216
		h, _ := primitives.HexToHash("82f7551efc2be9c12d58715cde6bf4dca7cc2974ecc7b59991590ca529ca3a39")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("2d51bda94923de4a838808b576d8a04be0adcf0a7253ef9187cc8c77709e38b5")
		bs.ECBlockHeadHash = h
		break
	case "960d90279b03b14d00a9bac86a69d5e8408d97bdd4da0120b79b1f9e1ab9bf18": //62783
		h, _ := primitives.HexToHash("c84ef4fdc379a62c39ba8b00602a8582eb8ec3d1340e6e49c7f6230774dfb055")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("fe63d6644cbf2ca581d9a0ba0b4d81d0e4b1a5d4c9e9b89ba15a356059fbac35")
		bs.ECBlockHeadHash = h
		break
	case "af55e8ab385b64f1ca2bc8202d4571c4832b1a58a50467e708571bf176cc9aa3": //67813
		h, _ := primitives.HexToHash("abc7ba807e13b92ff9e0fbae1f1f702ce4d2821506fc35b39fc016770f31714e")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("3c91b12d0a7a5372bd852df6f5cde7d0de1d3dce03cd1fbe58af9d79a1144dc7")
		bs.ECBlockHeadHash = h
		break
	case "6cb8e09d851c94425aea470d7fa771ccd56bf9408d52d6f76f8bdc022c81bfaf": //69088
		h, _ := primitives.HexToHash("8abc6032ad76242a6c3f1ed2adcf7cf4256c81f869f915b08a93a4601fe0f3b7")
		bs.ECBlockHeadKeyMR = h
		h, _ = primitives.HexToHash("89c1dcc0f7e4822e3fe6c9482e20b15c2b003fd37e3106bb25c27a1479b585ff")
		bs.ECBlockHeadHash = h
		break
	}

	//Expired Commits
	switch hash {
	case "90424572f5663cbe6d2fc6d97c0b97384c20d8c8b0182c9b0948ea1051f0dcf5": //5050
		eHash, _ := primitives.NewShaHashFromStr("75acd41e736406404bf6705a0ac3e5435442458954a413c32c472d7f46dc69af")
		ecHash, _ := primitives.NewShaHashFromStr("0fb30d88fb38ddaea3fdf99ba03f310e7b2a471a4390bfd4e19cd06f3ac49b62")
		bs.PushCommit(eHash, ecHash)
		eHash, _ = primitives.NewShaHashFromStr("280829e1be009fd5ff4b04028a4794dc027f2ba9667d7b3695b42c1039ee61b7")
		ecHash, _ = primitives.NewShaHashFromStr("21473e3bfd058ba941d2e2d4095237787c2de0e1cf32c9977366819f9385deec")
		bs.PushCommit(eHash, ecHash)
		break
	case "4b68855a7185bc100b9d7f106737e1277a9756a565d7b86e2cc282a08e7f4251": //5051
		eHash, _ := primitives.NewShaHashFromStr("7601c6a77560402999ba91b5be7439ad81e132c0b5b221bfb8afce3abb6de917")
		ecHash, _ := primitives.NewShaHashFromStr("3396cbd422bcbf1f45dca8e41e4814233b61f484460ca3ee5a67177f21745888")
		bs.PushCommit(eHash, ecHash)
		break
	case "d3b3663b218183ce89dbe177621f4d4214de2413b094713b7511a04aad638380": //25283
		eHash, _ := primitives.NewShaHashFromStr("b3742461ed073ff0f185b858d3ebe8ab1638f0c055c1950e0709964da77f71b2")
		ecHash, _ := primitives.NewShaHashFromStr("d11473f4a0cb0595e7c1dce8708dc01de0972f55e314e68bd656b0a0b96be5a8")
		bs.PushCommit(eHash, ecHash)
		break
	}

	//Pre-revealed entries
	pres := pre.GetPreErrors(hash)
	if len(pres) > 0 {
		for _, v := range pres {
			eHash, _ := primitives.NewShaHashFromStr(v.EntryID)
			ecHash, _ := primitives.NewShaHashFromStr(v.ECID)
			bs.PushCommit(eHash, ecHash)
		}
	}

	return nil
}

func (bs *BlockchainState) HandlePostBlockErrors(dBlockHash interfaces.IHash) error {
	bs.Init()
	hash := dBlockHash.String()

	//Pre-revealed entries
	pres := pre.GetPostErrors(hash)
	if len(pres) > 0 {
		for _, v := range pres {
			h, _ := primitives.NewShaHashFromStr(v.EntryID)
			bs.PopCommit(h)
		}
	}

	return nil
}

var pre *PreRevealedErrors

func init() {
	pre = new(PreRevealedErrors)
	pre.Entries = map[string][]*PRError{}
	pre.ECs = map[string][]*PRError{}

	pre.AddErrorPrePost("04e4f61c02ba862d7ad537ec6b9102f8197462ec62e3d4e04636e5899e0858d9", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "325da7e1db4379da4567a643c2e72ec5cc39182c41a1c709d0921244ed443937", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("08e3031aac2334411c4bd26176ca14daadcf04158c9e222a21b676cd8c38c98d", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "d5ec6690d183e1188dfa0039927d6ec95e928eb3e351a5c8d70f52f5c02e33c1", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("0fc878f3f8011034a66bea1dd6bb6e8e88f15ab36e00c736e2de48a74436798a", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "f4d14bf3900be35e495c59d61ae5b7f48c42a452508ed1a8341bf45e475b589c", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("165d626e1f41fab623d88b2816db3b6ca3e5ac3a774402125e36a3fb4ce4d116", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "7128982e4267d32df3a5c95e0116f4c9454b3c25f22735bd6ad82e82a49349c0", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("20a8f165bb58a9b1e93e88049cce1d7508e9e5292e62638f6df8152b82fd698c", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "38ab3d341fce102abbc2884ab1eea6a98a614f4acd773d446d6d889e4507788d", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("27f390a40190d18c49ad400b6f1e98b5d88c346e897dcdfe80e3379467e52255", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "54a7d2003ba842e2c735a6d6553778520d38348a4e98f6507272a999ac9d7297", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("28b486c4d748ee699a847eb0e284d5f8227bcb7d4a948a39df51fd8725579952", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "afa6f23426d65a5f182c88be55d01c7ba61303d3857133068af8a58bd1db0909", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("29aee0c092a3ad7abdac275a2de0b644510c77889eb4a8ff1db3f60d8a652375", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "3a687e05375beb6dbac26aa55edab01a8965d82f5d9711a2431470b50b7e4221", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("32ffef710041bdbf91a74a41287d0bf735f06703cab06f3f82379c68b993037c", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "70951d73bceaed711f6ea554e0205aef2bc8b0d3d16c08eb0de430a0afa4f478", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("3381c6f558c0cd72bb047a08bc5379d952e3319de10d006ce0bd465761190b18", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "c5e6c1d7a70c24476d238ab177aac5ace671cdd20833047cfa386972636b3a1f", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("43f4fc2a107b27e62d84b988872413654bce904498be8152dfce81b0a5f612d2", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ae9481c4e87e83f067f506ef2ef35e8c9a0b4a09ee65312f7b3d1acce1f8ce72", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("51f666f22cdd506cffacd4e81451d649dcafd80ac50154ffcdd04efb24d68778", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ae1d5c8f52e4155f582df88ef9d673f7ed9aa6d9054b75eb39c1228996e80b15", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("55ced301094e0ef4c0545768b8a30ec4137a47883e51614f681687c842141d18", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "842dfedf9edffdfa2d48285e1ccea288779f04cf80a6d95e94d4b5c661363f30", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("57befcfc4e236ab450285b4bb33b0f069ae15c994a4409dd705b9a50146b2117", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "af819ba83db0efb6cfa61593b65816d2ecbd6030ab01e485584e0e3acf663b33", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("5837bb414cf003a2ac8911bc9f98490cc07d516335161faf505104518d5fdcde", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "25eb10d300a53e32307980870b5b6b8d126989bc42f6878ab4864a206a115869", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("5a3d58dc5366b15e85b76b83805c29d73bf1044d03eaf0a5f0515ae6eaab56d1", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "627f0b1dabb0ab1b4e68677a638e0bdbc4c49d62e7c2150bfda1c371314a77a2", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("5d42f0f8fc6857ed207286aa68b32734a639d769fb5349d1f2b1a3d16515411f", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "6a7eaea303c9c697a88d94e144f1ea551b72694f3d143fcf02e3db4866f76352", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("607951a8fc6dfe384a770b56fc92c4d6384e629d67907ef44ea97f64065970b9", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "993722d663dd4880e048ad6b057eca71fc1f8fabc1e698f4cb0f5b2c29c75b6b", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("693a135f6d747b0cd46d9a9c83653e8b3d5598e2592250a5d50074390dcbcedb", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "37e7fe196f6404ef5d27b51b3754a3684b9b8c98f5cca45005c600f1a05f62b3", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("7aea0ad5db315af7ef25127570c774b75add27108d4f7723cbd0f0e8039b9784", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "b6a320f66151ee957a97adf11a45297d049fb82811b6c7b0ffbd2973243cc369", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("7ee6c3021bb3e95cc1495c47b4e505fb1e90792879067062ff1286b3595202e1", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "77071593576525f2b064ed9113ba61247610f83c862593b928bf57d3b284cea5", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("8f044304745e3df7594420f05f2e6330473eeda4c0559a242f227ca5b25002a1", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "67349e81fcd2a036bfec02e80c0c85bb0118f7508b9b20280498f81ff0809b77", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("940e372536e96d9742b23665f976c2fb610bc52a1edbe4bcc9f2fa6c1efa8f8a", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "5406b10e81dc65aa7062a611fbe035dc97221cdc6b19cd9c46ddd13398bf9fa7", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("959eb6a885fb4df07f410856ffdc971e1021ce777c918c7705252bb47253f663", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "61815982a985a4b18bbcd12866d3c9c4b43773b3d5db89b04e812d9412a1d44e", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("9b293d17991a95253ecc92c0e8089be1cde0cd1a21344e0c03b84abb68f29a8e", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "f526aa777cafa42f755346c0e9259e643ade7e3eddc3de584562bd97f449dbe6", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("9c4cb42a3250788bcfcca0b35e5bbf42bfe9fdd436e9394dec6d3a15e54ecc0b", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "7f5c6417dcff44028855fbe910322a320d3e8a02d9f7577c712858e972a54146", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("9d3877130398cac5a460fa5989e5251dcc76ab0bd39bd51e5840e42da5143ca6", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "f62a1fbe5f750c20fb119d725846f71c097526667735e8ee056ea3a491d8b198", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("a43b83d5fd59f444edec5475a9e446bcc8c2565e8f7e3ffe86981d2a42b0694c", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ac3d08965722db172c64ddc15fe3f447d0769553d0ec066b4fd3ca07de063bec", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("b735d05adbcff7057673d38f2b59e1c6c747e7095c6f51199d4e8bc80f520985", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "31f1115c1ae4a6babdd139871a90c86ab67bfca6a2ddfa13c7f43f653afcdc1f", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("bd2068e84d703c31a25880ab62263e161936c7b16d03e8724a8ac31100e6132b", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ce612f37de654845f764f6b431581bf246df2f2aa7fb4a5b3db3994ca39f0c4c", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("c2cb669c7c7f3beabf0b0e757291c672b0f0395ad93c0f827e8c857272922b1a", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ebcb1029845a568b6b197dc8243c9e78707d657fc24c764ca554232025e4de47", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("c4ea62b5e28fb116018537592a5bc8ae1b21aba51232a9cb2d231e52c3f7fd2e", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "9cb5fce50cd9346a8b865041cd0dc1538df5f753176f09e10a64f41aaf0bbe3b", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("c6f8b0884e6494289ad9fe46f0623a174caf40dfa21471bb2e746db0eea0ee73", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "2493bc70f43b54f8811c27fb062c91d1dfd03480475e67d292c7fdf2bf27d968", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("cf0b41b057bb2b1f088bcccd552304b21aede671aaa631c86e586e77f81276e9", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "8d7a0e09fdf7c5a8d76486595892a7d811eb87521727a2b2a65d57802351c934", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("db4cbf8a9148d7cf91e5369dd0c8d1b7f74e4190cdc25775c95f8b3868080629", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "33ae4af0b2680bd8a1863e79571b185d2f4ef2b5d832ec0ac1fe80da81c2a66a", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("dcbc532ce32761781e209f7eb2bc6a56d82469d5997c03745127ce000b653f93", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "121ddf55b4567914c608b6c098c2977040d98b7f0f0e572b22cb7bffc72e94fd", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("dcd3e45560b421650c5b8314229a4771eccea15179b04df8318ca4c7dfcc92dc", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "1b3340eb5a72090e113d60adab43e84d92d1dcd4732362e3cdf52fc24f053d5f", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("e6ff494b2468897e98d333c7685f000c6be6d4e1023c0b2d2b4080a33e911245", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "ac144075bfde97a33d9cb0173c869d92542214c0b6782a02c6b3beff0c880788", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("f29aa952ee1de9a64f702a3b6615f8de41ee31e7581b0ab2e2283d98b62f7876", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "7a46da65b37a77e83fb388c6d09883b3af0ddc6f28b30432b6f411379c03002a", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("f3d7683205a190b5577b1863190c7eac2152e3d7071b147d4935171cc287995a", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "6cba39384b4e10e1f674f67af5cd41f7a1570cfe43fe0edb352e9792fd661584", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")
	pre.AddErrorPrePost("f8142476974e66de99422b4a987ae0250d0c1c7d372b3d5150e80cc2ec726b03", "990f7976638e46ac673773361d07a033863468debdfc33856d6a4f9d3c3fece1", "4663887ffc44752f347b5ae34b8b096aeeb651b159b5341314ee2ab713113e74", "72b76ccb33b8aa27d618861eb7246971c8399ddd11bf4627227897a05b719ae5")

	pre.AddPreError("04e4f61c02ba862d7ad537ec6b9102f8197462ec62e3d4e04636e5899e0858d9", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "325da7e1db4379da4567a643c2e72ec5cc39182c41a1c709d0921244ed443937")
	pre.AddPreError("08e3031aac2334411c4bd26176ca14daadcf04158c9e222a21b676cd8c38c98d", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "d5ec6690d183e1188dfa0039927d6ec95e928eb3e351a5c8d70f52f5c02e33c1")
	pre.AddPreError("0fc878f3f8011034a66bea1dd6bb6e8e88f15ab36e00c736e2de48a74436798a", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "f4d14bf3900be35e495c59d61ae5b7f48c42a452508ed1a8341bf45e475b589c")
	pre.AddPreError("165d626e1f41fab623d88b2816db3b6ca3e5ac3a774402125e36a3fb4ce4d116", "a2139054b6c7f4f48440c22fd5fe189300bc7c442ffe825ff4e071e9b47d56d9", "7128982e4267d32df3a5c95e0116f4c9454b3c25f22735bd6ad82e82a49349c0")
	pre.AddPreError("20a8f165bb58a9b1e93e88049cce1d7508e9e5292e62638f6df8152b82fd698c", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "38ab3d341fce102abbc2884ab1eea6a98a614f4acd773d446d6d889e4507788d")
	pre.AddPreError("27f390a40190d18c49ad400b6f1e98b5d88c346e897dcdfe80e3379467e52255", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "54a7d2003ba842e2c735a6d6553778520d38348a4e98f6507272a999ac9d7297")
	pre.AddPreError("28b486c4d748ee699a847eb0e284d5f8227bcb7d4a948a39df51fd8725579952", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "afa6f23426d65a5f182c88be55d01c7ba61303d3857133068af8a58bd1db0909")
	pre.AddPreError("29aee0c092a3ad7abdac275a2de0b644510c77889eb4a8ff1db3f60d8a652375", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "3a687e05375beb6dbac26aa55edab01a8965d82f5d9711a2431470b50b7e4221")
	pre.AddPreError("32ffef710041bdbf91a74a41287d0bf735f06703cab06f3f82379c68b993037c", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "70951d73bceaed711f6ea554e0205aef2bc8b0d3d16c08eb0de430a0afa4f478")
	pre.AddPreError("3381c6f558c0cd72bb047a08bc5379d952e3319de10d006ce0bd465761190b18", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "c5e6c1d7a70c24476d238ab177aac5ace671cdd20833047cfa386972636b3a1f")
	pre.AddPreError("43f4fc2a107b27e62d84b988872413654bce904498be8152dfce81b0a5f612d2", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ae9481c4e87e83f067f506ef2ef35e8c9a0b4a09ee65312f7b3d1acce1f8ce72")
	pre.AddPreError("51f666f22cdd506cffacd4e81451d649dcafd80ac50154ffcdd04efb24d68778", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ae1d5c8f52e4155f582df88ef9d673f7ed9aa6d9054b75eb39c1228996e80b15")
	pre.AddPreError("55ced301094e0ef4c0545768b8a30ec4137a47883e51614f681687c842141d18", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "842dfedf9edffdfa2d48285e1ccea288779f04cf80a6d95e94d4b5c661363f30")
	pre.AddPreError("57befcfc4e236ab450285b4bb33b0f069ae15c994a4409dd705b9a50146b2117", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "af819ba83db0efb6cfa61593b65816d2ecbd6030ab01e485584e0e3acf663b33")
	pre.AddPreError("5837bb414cf003a2ac8911bc9f98490cc07d516335161faf505104518d5fdcde", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "25eb10d300a53e32307980870b5b6b8d126989bc42f6878ab4864a206a115869")
	pre.AddPreError("5a3d58dc5366b15e85b76b83805c29d73bf1044d03eaf0a5f0515ae6eaab56d1", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "627f0b1dabb0ab1b4e68677a638e0bdbc4c49d62e7c2150bfda1c371314a77a2")
	pre.AddPreError("5d42f0f8fc6857ed207286aa68b32734a639d769fb5349d1f2b1a3d16515411f", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "6a7eaea303c9c697a88d94e144f1ea551b72694f3d143fcf02e3db4866f76352")
	pre.AddPreError("607951a8fc6dfe384a770b56fc92c4d6384e629d67907ef44ea97f64065970b9", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "993722d663dd4880e048ad6b057eca71fc1f8fabc1e698f4cb0f5b2c29c75b6b")
	pre.AddPreError("693a135f6d747b0cd46d9a9c83653e8b3d5598e2592250a5d50074390dcbcedb", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "37e7fe196f6404ef5d27b51b3754a3684b9b8c98f5cca45005c600f1a05f62b3")
	pre.AddPreError("7aea0ad5db315af7ef25127570c774b75add27108d4f7723cbd0f0e8039b9784", "a2139054b6c7f4f48440c22fd5fe189300bc7c442ffe825ff4e071e9b47d56d9", "b6a320f66151ee957a97adf11a45297d049fb82811b6c7b0ffbd2973243cc369")
	pre.AddPreError("7ee6c3021bb3e95cc1495c47b4e505fb1e90792879067062ff1286b3595202e1", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "77071593576525f2b064ed9113ba61247610f83c862593b928bf57d3b284cea5")
	pre.AddPreError("8f044304745e3df7594420f05f2e6330473eeda4c0559a242f227ca5b25002a1", "a2139054b6c7f4f48440c22fd5fe189300bc7c442ffe825ff4e071e9b47d56d9", "67349e81fcd2a036bfec02e80c0c85bb0118f7508b9b20280498f81ff0809b77")
	pre.AddPreError("940e372536e96d9742b23665f976c2fb610bc52a1edbe4bcc9f2fa6c1efa8f8a", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "5406b10e81dc65aa7062a611fbe035dc97221cdc6b19cd9c46ddd13398bf9fa7")
	pre.AddPreError("959eb6a885fb4df07f410856ffdc971e1021ce777c918c7705252bb47253f663", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "61815982a985a4b18bbcd12866d3c9c4b43773b3d5db89b04e812d9412a1d44e")
	pre.AddPreError("9b293d17991a95253ecc92c0e8089be1cde0cd1a21344e0c03b84abb68f29a8e", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "f526aa777cafa42f755346c0e9259e643ade7e3eddc3de584562bd97f449dbe6")
	pre.AddPreError("9c4cb42a3250788bcfcca0b35e5bbf42bfe9fdd436e9394dec6d3a15e54ecc0b", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "7f5c6417dcff44028855fbe910322a320d3e8a02d9f7577c712858e972a54146")
	pre.AddPreError("9d3877130398cac5a460fa5989e5251dcc76ab0bd39bd51e5840e42da5143ca6", "a2139054b6c7f4f48440c22fd5fe189300bc7c442ffe825ff4e071e9b47d56d9", "f62a1fbe5f750c20fb119d725846f71c097526667735e8ee056ea3a491d8b198")
	pre.AddPreError("a43b83d5fd59f444edec5475a9e446bcc8c2565e8f7e3ffe86981d2a42b0694c", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ac3d08965722db172c64ddc15fe3f447d0769553d0ec066b4fd3ca07de063bec")
	pre.AddPreError("b735d05adbcff7057673d38f2b59e1c6c747e7095c6f51199d4e8bc80f520985", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "31f1115c1ae4a6babdd139871a90c86ab67bfca6a2ddfa13c7f43f653afcdc1f")
	pre.AddPreError("bd2068e84d703c31a25880ab62263e161936c7b16d03e8724a8ac31100e6132b", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ce612f37de654845f764f6b431581bf246df2f2aa7fb4a5b3db3994ca39f0c4c")
	pre.AddPreError("c2cb669c7c7f3beabf0b0e757291c672b0f0395ad93c0f827e8c857272922b1a", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ebcb1029845a568b6b197dc8243c9e78707d657fc24c764ca554232025e4de47")
	pre.AddPreError("c4ea62b5e28fb116018537592a5bc8ae1b21aba51232a9cb2d231e52c3f7fd2e", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "9cb5fce50cd9346a8b865041cd0dc1538df5f753176f09e10a64f41aaf0bbe3b")
	pre.AddPreError("c6f8b0884e6494289ad9fe46f0623a174caf40dfa21471bb2e746db0eea0ee73", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "2493bc70f43b54f8811c27fb062c91d1dfd03480475e67d292c7fdf2bf27d968")
	pre.AddPreError("cf0b41b057bb2b1f088bcccd552304b21aede671aaa631c86e586e77f81276e9", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "8d7a0e09fdf7c5a8d76486595892a7d811eb87521727a2b2a65d57802351c934")
	pre.AddPreError("db4cbf8a9148d7cf91e5369dd0c8d1b7f74e4190cdc25775c95f8b3868080629", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "33ae4af0b2680bd8a1863e79571b185d2f4ef2b5d832ec0ac1fe80da81c2a66a")
	pre.AddPreError("dcbc532ce32761781e209f7eb2bc6a56d82469d5997c03745127ce000b653f93", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "121ddf55b4567914c608b6c098c2977040d98b7f0f0e572b22cb7bffc72e94fd")
	pre.AddPreError("dcd3e45560b421650c5b8314229a4771eccea15179b04df8318ca4c7dfcc92dc", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "1b3340eb5a72090e113d60adab43e84d92d1dcd4732362e3cdf52fc24f053d5f")
	pre.AddPreError("e6ff494b2468897e98d333c7685f000c6be6d4e1023c0b2d2b4080a33e911245", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "ac144075bfde97a33d9cb0173c869d92542214c0b6782a02c6b3beff0c880788")
	pre.AddPreError("f29aa952ee1de9a64f702a3b6615f8de41ee31e7581b0ab2e2283d98b62f7876", "a2139054b6c7f4f48440c22fd5fe189300bc7c442ffe825ff4e071e9b47d56d9", "7a46da65b37a77e83fb388c6d09883b3af0ddc6f28b30432b6f411379c03002a")
	pre.AddPreError("f3d7683205a190b5577b1863190c7eac2152e3d7071b147d4935171cc287995a", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "6cba39384b4e10e1f674f67af5cd41f7a1570cfe43fe0edb352e9792fd661584")
	pre.AddPreError("f8142476974e66de99422b4a987ae0250d0c1c7d372b3d5150e80cc2ec726b03", "53e0a7d4c8bc355f15716a845d846169190e4b6dd4b0f7c18ed2a208cc653b68", "4663887ffc44752f347b5ae34b8b096aeeb651b159b5341314ee2ab713113e74")
}

type PreRevealedErrors struct {
	Entries map[string][]*PRError
	ECs     map[string][]*PRError
}

type PRError struct {
	EntryID     string
	EntryDBlock string

	ECID     string
	ECDBlock string
}

func (pre *PreRevealedErrors) AddErrorPrePost(entryID, dBlock string, commitID, commitDBlock string) {
	prError := new(PRError)

	prError.EntryID = entryID
	prError.EntryDBlock = dBlock

	prError.ECID = commitID
	prError.ECDBlock = commitDBlock

	pre.Entries[dBlock] = append(pre.Entries[dBlock], prError)
	pre.ECs[commitDBlock] = append(pre.ECs[commitDBlock], prError)
}

func (pre *PreRevealedErrors) AddPreError(entryID, dBlock, ecID string) {
	prError := new(PRError)

	prError.EntryID = entryID
	prError.EntryDBlock = dBlock
	prError.ECID = ecID

	pre.Entries[dBlock] = append(pre.Entries[dBlock], prError)
}

func (pre *PreRevealedErrors) GetPreErrors(dBlockHash string) []*PRError {
	return pre.Entries[dBlockHash]
}

func (pre *PreRevealedErrors) GetPostErrors(dBlockHash string) []*PRError {
	return pre.ECs[dBlockHash]
}
