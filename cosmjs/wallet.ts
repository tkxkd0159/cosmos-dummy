import {readFile, writeFile} from "fs/promises";
import { DirectSecp256k1HdWallet, DirectSecp256k1Wallet, AccountData, OfflineDirectSigner } from "@cosmjs/proto-signing"
import { stringToPath, HdPath } from "@cosmjs/crypto";
import { fromHex } from "@cosmjs/encoding"

const isGen = true

const generateKey = async (keypath: string): Promise<AccountData> => {
    const wallet: DirectSecp256k1HdWallet = await DirectSecp256k1HdWallet.generate(24)
    console.log(`Mnemonic: ${wallet.mnemonic}`)

    let encrypedWallet = await wallet.serialize("random")
    await writeFile(keypath, encrypedWallet, { encoding: 'utf8' })

    let loadEncrypedWallet = await readFile(keypath, { encoding: 'utf8' })
    let decryptedWallet = await DirectSecp256k1HdWallet.deserialize(loadEncrypedWallet, "random")
    console.log(`Deserialized Wallet: ${decryptedWallet.mnemonic}`)

    const accounts = await wallet.getAccounts()
    return accounts[0]
}

if (isGen) {
    const keydir = "keys"
    const keyname = "dummy"
    const keypath = `${keydir}/${keyname}.key`
    generateKey(keypath).then(r => {
    console.error(`\n${r.algo}, Key: ${r.pubkey}`)
    console.log(`Address: ${r.address}`)
})
}


const WalletFromMnemonic = async (mnemonicPath: string, prefix: string): Promise<OfflineDirectSigner> => {
    const p: HdPath = stringToPath("m/44'/118'/0'/0/0");
    const mnemonic = await readFile(mnemonicPath, { encoding: 'utf8' })
    return await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {prefix: prefix, hdPaths: [p]})
}

const WalletFromPrivkey = async (keypath: string): Promise<OfflineDirectSigner>  => {
    return DirectSecp256k1Wallet.fromKey(
        fromHex((await readFile(keypath)).toString()),
        "cosmos",
    )
}

export {
    WalletFromMnemonic,
    WalletFromPrivkey,
}