import { ethers } from "hardhat";
import fs from 'fs';

function writeTextToFile(text: string, filePath: string): void {
  fs.writeFileSync(filePath, text);
  console.log(`Text written to ${filePath}`);
}


async function main() {

  const T55NFTs = await ethers.getContractFactory("T55NFTs")
  const t55NFTs = await T55NFTs.deploy();
  await t55NFTs.deployed();
  console.log("T55 nft deployed: ", t55NFTs.address)
  
  const T55Bridge = await ethers.getContractFactory("T55Bridge")
  const t55Bridge = await T55Bridge.deploy(t55NFTs.address);
  await t55Bridge.deployed();
  console.log("T55Bridge contract deployed: ", t55Bridge.address)

  writeTextToFile(`T55_ADDRESS=${t55NFTs.address}\nT55BRIDGE_ADDR=${t55Bridge.address}`,'.env.t55');

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
