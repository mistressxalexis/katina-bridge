// import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
// import { expect } from 'chai';
// import { ethers } from 'hardhat';

// describe("Kafrex Token", ()=>{
//     const deployToken = async () => {
//         const LEO = await ethers.getContractFactory("LEO")
//         const leo = await LEO.deploy();
//         await leo.deployed();

//         const T54 = await ethers.getContractFactory("LEO")
//         const t54 = await T54.deploy();
//         await t54.deployed();

//         const [sender, receiver] = await ethers.getSigners();
//         return {sender, receiver, leo, t54}
//     }

//     describe("deploy KAF token", () => {
//         it("should mint a new token", async ()=> {
//             const {sender, receiver, leo, t54} = await loadFixture(deployToken);
//             await leo.mint(receiver.address, 1000)
//             expect(await leo.balanceOf(receiver.address)).to.equal(1000);
//         })
//     })
// })