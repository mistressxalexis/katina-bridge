import { ethers } from "hardhat";
import fs from 'fs';


async function main() {

    const Sentinel = await ethers.getContractFactory("SEN")
    const sen = await Sentinel.deploy();
    await sen.deployed();

    const SentinelsBridge = await ethers.getContractFactory("SentinelsBridge")
    const senbr = await SentinelsBridge.deploy(sen.address);
    await senbr.deployed();
    console.log("senbr nft deployed: ", senbr.address)
    console.log("sen nft deployed: ", sen.address)

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
