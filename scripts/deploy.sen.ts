import { ethers } from "hardhat";
import fs from 'fs';

function writeTextToFile(text: string, filePath: string): void {
  fs.writeFileSync(filePath, text);
  console.log(`Text written to ${filePath}`);
}

async function main() {

  const Sentinel = await ethers.getContractFactory("SEN")
  const sen = await Sentinel.deploy();
  await sen.deployed();

  const SentinelsBridge = await ethers.getContractFactory("SentinelsBridge")
  const senbr = await SentinelsBridge.deploy(sen.address);
  await senbr.deployed();
  console.log("senbr nft deployed: ", senbr.address)
  console.log("sen nft deployed: ", sen.address)

  writeTextToFile(`SEN_ADDRESS=${sen.address}\nSEN_BRIDGE_ADDR=${senbr.address}`,'.env.sentinels');

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
