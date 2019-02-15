package eth_test

import (
	"encoding/json"
	"github.com/INFURA/eth/pkg/eth"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewHeadsNotificationParams(t *testing.T) {

	{
		params := eth.NewHeadsNotificationParams{}

		geth := `{"subscription":"0x5e68301d4a26b0610081dca8b41eb888","result":{"parentHash":"0xf997a57caff40e89e9f8716568de63f1c1632629d9de750f06c8f7c32bf2c379","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","miner":"0x829bd824b016326a401d083b33d092293333a830","stateRoot":"0x015cfcb599289caeffeb41087745a128768c46d41313e1c31d9133ceb638fbcd","transactionsRoot":"0x10e942a373a9383a6d183dd12d346ed5d692dca75597d23d3536dd4035f4b6ba","receiptsRoot":"0x1889c02d922caf787adccc53e4a48c865615fd02e113eb0e2fb76dbb49f7403a","logsBloom":"0x80000110a40001802020200141130000820013a00200ac0100a12ca121ac40600200041020022808882048480a0812284240808da98800020901ac10d07e4692c500514038100d098207aa080d6004b220400066c286528840004ec18800a3040404c0b02283074918d00000010034402b628180401a0600400c2431001001b40041000252000200404405400020480200020180402808044000900042101ea0020404254010f000514818020081a0808000290100a200291c80209c20982008840011223050100270051640081a1220040000a00500436104208015804058f831323000a0112140121040860030801084000000007200980004400244080081","difficulty":"0x9dd74b59cb6be","number":"0x6e3044","gasLimit":"0x7a4f2b","gasUsed":"0x47c15f","timestamp":"0x5c65fe8d","extraData":"0x7070796520e4b883e5bda9e7a59ee4bb99e9b1bc","mixHash":"0x4ed4476377a8246b7930dedcbbe3c967496a9794ebd0fc2310d1c7906f94db87","nonce":"0xd28cb2800a6b6fff","hash":"0xcf026edb3d84e540aed9ca11b3c1dfa678bd4bda2d4cc31953e006de6292d53a"}}`

		err := json.Unmarshal([]byte(geth), &params)
		require.NoError(t, err)
		require.Equal(t, eth.MustHash("0xcf026edb3d84e540aed9ca11b3c1dfa678bd4bda2d4cc31953e006de6292d53a"), params.Result.Hash)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(geth), j)
	}

	{
		params := eth.NewHeadsNotificationParams{}

		parity := `{"result":{"author":"0xb2930b35844a230f00e51431acae96fe543a0347","difficulty":"0x9dd8742e86a7e","extraData":"0x76697231","gasLimit":"0x7a121d","gasUsed":"0x79e63d","hash":"0xdcecb7c36c48c25886f26e21993ddd12c9e17ae015709bce520650cf6a79e757","logsBloom":"0x2880180008000a8002a00030000024044038020044102000522200001043200b0a0148040a206b048412041000c100a600004402800901160020100811a40981940249801d0160410180400c010181065e544100944e29310085011ca110009303901801122a0080002800084051ac0140208620015500407600049090c004400054103300e5080402a10040002080090500028000000440040100240853203002041008400c0002050408801081801b65b001009000c306006000410002037018b0500a20ae0729442800211205480444c1800fb204291080110821006160c28410200c00010a0014005004c4b4b0002420026780400048985ac04000320296","miner":"0xb2930b35844a230f00e51431acae96fe543a0347","mixHash":"0x00367fe1776614d931a76eeca13bac9680cbf9ac662250b688831d01fb8ad8ba","nonce":"0xd9790000a29655d0","number":"0x6e310b","parentHash":"0xc757d60368bd8ece87adffb9910b8fe5770d85e8c1a1b34482141c4d729780a0","receiptsRoot":"0xb1a98f7ba64eb612acfee18dbe3d3dc6826a95f3b8003251c812af1869701066","sealFields":["0x00367fe1776614d931a76eeca13bac9680cbf9ac662250b688831d01fb8ad8ba","0xd9790000a29655d0"],"sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x208","stateRoot":"0x402cb409a1efbfd5fc85d10a04eb49cf00b80d5812f81f85b328416d51d6f2ae","timestamp":"0x5c660e2c","transactionsRoot":"0xdbeb6bd49efffac42f43b5d4c9e739285b088e23e85b1f5ff00be9e7b60f2293"},"subscription":"0x03580d776ce9f51c"}`
		err := json.Unmarshal([]byte(parity), &params)
		require.NoError(t, err)
		require.Equal(t, *eth.MustAddress("0xb2930b35844a230f00e51431acae96fe543a0347"), params.Result.Author)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(parity), j)
	}

	{
		params := eth.NewHeadsNotificationParams{}

		kovan := `{"result":{"author":"0x0010f94b296a852aaac52ea6c5ac72e03afd032d","difficulty":"0xfffffffffffffffffffffffffffffffa","extraData":"0xde8302020a8f5061726974792d457468657265756d86312e33322e30826c69","gasLimit":"0x7a1200","gasUsed":"0x6f91e9","hash":"0xe14fdd50f0de3874a6542a48597ff2ffedff9cc0b9e2dc67ec8c09fa16e5644d","logsBloom":"0x00000000000000000000000000000000000000000400000000000000200000000000000000000000000000000000000400020000000000400000000000000000000000000000000000000408000000000000000000000002000000008000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000400040000000000000000000000000000000000000000001080000000000000000000000000000000000000000100000000102000000000000000000000000000000000000400000000800000000000000000000000004000000000000002000000000000000000000000000000000","miner":"0x0010f94b296a852aaac52ea6c5ac72e03afd032d","number":"0x9e2269","parentHash":"0xdb210dc3a36647b48087daf877d82d2184b69df8897d7c92c69811a12b2c758d","receiptsRoot":"0xcffeccbe902d26ce499262095c717c20af7d1e46758941762110c23bd5de8662","sealFields":["0x17198407","0x5144875e245f6d76f1540df1ce87206c2523e3a9f71665b628ac09231b0678231fe36b82aba8162fadeb2af69a9862032a867d6dbef25a2b4a10044fc779cefc00"],"sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","signature":"5144875e245f6d76f1540df1ce87206c2523e3a9f71665b628ac09231b0678231fe36b82aba8162fadeb2af69a9862032a867d6dbef25a2b4a10044fc779cefc00","size":"0x24a","stateRoot":"0xc074e559c3100eaf216a2b3b111182ec1caaed3ede49bf4c02cce0d1398dc428","step":"387548167","timestamp":"0x5c66101c","transactionsRoot":"0xca6c0e435a7dcd8dbe1acaecbaa2f5af52ddb0809f0527007f44eda1bb34eddb"},"subscription":"0xf45ad86ee72015ed"}`
		err := json.Unmarshal([]byte(kovan), &params)
		require.NoError(t, err)
		require.Equal(t, *eth.MustQuantity("0x6f91e9"), params.Result.GasUsed)

		j, err := json.Marshal(&params)
		require.NoError(t, err)

		RequireEqualJSON(t, []byte(kovan), j)
	}

}