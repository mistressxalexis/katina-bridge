import { ethers } from "hardhat";
import fs from 'fs';

function writeTextToFile(text: string, filePath: string): void {
  fs.writeFileSync(filePath, text);
  console.log(`Text written to ${filePath}`);
}

async function main() {

  const Leopard = await ethers.getContractFactory("LEO")
  const leo = await Leopard.deploy();
  await leo.deployed();

  const LeopardBridge = await ethers.getContractFactory("LeopardBridge")
  const leobr = await LeopardBridge.deploy(leo.address);
  await leobr.deployed();
  console.log("senbr nft deployed: ", leobr.address)
  console.log("sen nft deployed: ", leo.address)

  writeTextToFile(`LEO_ADDRESS=${leo.address}\nLEO_BRIDGE_ADDR=${leobr.address}`,'.env.leopard');

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
