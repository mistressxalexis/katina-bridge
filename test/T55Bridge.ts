// import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
// import { expect } from 'chai';
// import { ethers } from 'hardhat';

// describe("BridgeSEPO test", ()=>{
//     const deployBridgeSEPO = async () => {
//         const T55NFT = await ethers.getContractFactory("T55NFTs")
//         const t55NFT = await T55NFT.deploy();
//         await t55NFT.deployed();

//         const T54NFT = await ethers.getContractFactory("T54NFTs")
//         const t54NFT = await T54NFT.deploy();
//         await t54NFT.deployed();

//         const T55Bridge = await ethers.getContractFactory("T55Bridge")
//         const t55Bridge = await T55Bridge.deploy(t55NFT.address);
//         await t55Bridge.deployed();

//         const T54Bridge = await ethers.getContractFactory("T54Bridge")
//         const t54Bridge = await T54Bridge.deploy(t54NFT.address);
//         await t54Bridge.deployed();

//         const [sender, receiver] = await ethers.getSigners();
//         return {sender, receiver, t55NFT, t54NFT,t54Bridge, t55Bridge}
//     }

//     describe("deploy KAF token", () => {
//         it("should mint a new token", async ()=> {
//             const {sender, receiver, t55NFT, t54NFT,t54Bridge, t55Bridge} = await loadFixture(deployBridgeSEPO);
            
//             await t55NFT.mint(sender.address, "www.kafka.org")
//             const t55tid = await t55NFT.getTokenID()

//             await t54NFT.sync(t55tid);
            
//             await t55Bridge.deposit(t55tid)

//             await t54Bridge.syncNFTs(sender.address, t55tid, "www.kafka.org")
//             await t54Bridge.claimNFTs(sender.address, t55tid);

//             expect(await t54NFT.ownerOf(t55tid)).to.equal(sender.address)
//         })
//     })
// })