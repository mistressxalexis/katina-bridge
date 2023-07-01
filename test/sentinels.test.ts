// import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
// import { expect } from 'chai';
// import { ethers } from 'hardhat';

// describe("setinels nfts", ()=>{
//     const deployToken = async () => {
//         const Sen = await ethers.getContractFactory("SEN")
//         const sen = await Sen.deploy();
//         await sen.deployed();

//         const Bridge = await ethers.getContractFactory("SentinelsBridge")
//         const senbridge = await Bridge.deploy(sen.address);
//         await senbridge.deployed();

//         const Leo = await ethers.getContractFactory("LEO")
//         const leo = await Leo.deploy();
//         await leo.deployed();

//         const LeoBridge = await ethers.getContractFactory("LeopardBridge")
//         const leobridge = await LeoBridge.deploy(leo.address);
//         await leobridge.deployed();

//         const [sender, receiver] = await ethers.getSigners();
//         return {sender, receiver, sen, senbridge, leo, leobridge}
//     }

//     describe("deploy sentinels nft", () => {
//         it("should mint a new token", async ()=> {
//             const {sender, receiver, sen, senbridge, leo, leobridge} = await loadFixture(deployToken);
            
//             console.log("sender: ", sender.address)
//             console.log("receiver: ", receiver.address)

//             // await sen.mintItem("www.example.com")
//             // const tokenID = await sen.connect(receiver).getTokenID();
//             // // await leobridge.connect(sender).awaitSyncData(tokenID, "www.example.com")
            
//             // expect(tokenID.toString()).to.equal("0");
//         })
//     })
// })