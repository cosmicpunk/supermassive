import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";

describe("SigningCosmWasmClient", () => {
    describe("connectWithSigner", () => {
        it("can be constructed", async () => {
            // pendingWithoutWasmd();
            const wallet = await DirectSecp256k1HdWallet.fromMnemonic(alice.mnemonic, { prefix: wasmd.prefix });
            const options = { prefix: wasmd.prefix };
            const client = await SigningCosmWasmClient.connectWithSigner(wasmd.endpoint, wallet, options);
            expect(client).toBeTruthy();
        });
    });
});
