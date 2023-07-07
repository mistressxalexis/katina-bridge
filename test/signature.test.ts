import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { expect } from 'chai';
import { ethers } from 'hardhat';
import { ec as EC } from 'elliptic';

describe("setinels nfts signature", ()=>{
    const deployToken = async () => {

        const Sen = await ethers.getContractFactory("SEN")
        const sen = await Sen.deploy();
        await sen.deployed();

        const Bridge = await ethers.getContractFactory("SentinelsBridge")
        const senbridge = await Bridge.deploy(sen.address);
        await senbridge.deployed();

        const Leo = await ethers.getContractFactory("LEO")
        const leo = await Leo.deploy();
        await leo.deployed();

        const LeoBridge = await ethers.getContractFactory("LeopardBridge")
        const leobridge = await LeoBridge.deploy(leo.address);
        await leobridge.deployed();


        const wallet = ethers.Wallet.createRandom();
        console.log("wallet", wallet.address)
        
        const message = "www.example.com";
        const signature = await wallet.signMessage(message);
        const sig = ethers.utils.splitSignature(signature)
        
        const r = sig.r
        const s = sig.s;
        const v = sig.v;

        const [sender, receiver] = await ethers.getSigners();
        return {sender, receiver, sen, senbridge, leo, leobridge, wallet, r, s, v, signature}
    }

    describe("deploy sentinels nft", () => {
        it("should mint a new signature", async ()=> {
            const {sender, receiver, sen, senbridge, leo, leobridge, wallet, r, s, v, signature} = await loadFixture(deployToken);
            const _signature = await sender.signMessage("www.example.com")
            // const verify = await senbridge.connect(sender).verify("www.example.com", signature)
            // const address = await senbridge.recover("www.example.com", signature)
            // expect(verify).to.be.true
            // expect(address).to.equal(sender.address)
            // console.log(signature)
            const verify = await senbridge.SignVerify("www.example.com", _signature)
            console.log(verify)
            // expect(signature).to.equal("0x6aacbe0fb92d51a5dfc1399aa7b992175faca163d739a0b1cde48c5f5796c4313bf46f8188b50d9b1f338cf68fe1f2f57f547353706b99b369b9db420513250c1b")

            
        })
    })
})