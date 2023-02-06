import {readFile} from "fs/promises";
import { DirectSecp256k1HdWallet, DirectSecp256k1Wallet, AccountData, OfflineDirectSigner } from "@cosmjs/proto-signing"
import { stringToPath, HdPath } from "@cosmjs/crypto";
import { fromHex } from "@cosmjs/encoding"

const isGen = false

const generateKey = async (): Promise<AccountData> => {
    const wallet: DirectSecp256k1HdWallet = await DirectSecp256k1HdWallet.generate(24)
    process.stdout.write(wallet.mnemonic)
    const accounts = await wallet.getAccounts()
    return accounts[0]
}

if (isGen) {
    generateKey().then(r => {
    console.error(`\n${r.algo}, Key: ${r.pubkey}`)
    console.log(`Address: ${r.address}`)
})
}


const WalletFromMnemonic = async (mnemonicPath: string): Promise<OfflineDirectSigner> => {
    const p: HdPath = stringToPath("m/44'/118'/0'/0/0");
    const mnemonic = await readFile(mnemonicPath, { encoding: 'utf8' })
    return await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {prefix: "cosmos", hdPaths: [p]})
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