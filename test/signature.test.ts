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
        // console.log("wallet", wallet.address)
        
        const message = "www.example.com";
        const signature = await wallet.signMessage(message);
        const sig = ethers.utils.splitSignature(signature)
        
        const r = sig.r
        const s = sig.s;
        const v = sig.v;

        const [sender, receiver] = await ethers.getSigners();
        return {sender, receiver, sen, senbridge, leo, leobridge, wallet, r, s, v}
    }

    describe("deploy sentinels nft", () => {
        it("should mint a new signature", async ()=> {
            const {sender, receiver, sen, senbridge, leo, leobridge, wallet, r, s, v} = await loadFixture(deployToken);
            const signature = await sender.signMessage("www.example.com")
            const verify = await senbridge.connect(sender).verify("www.example.com", signature)
            const address = await senbridge.recover("www.example.com", signature)
            expect(verify).to.be.true
            expect(address).to.equal(sender.address)

        })
    })
})