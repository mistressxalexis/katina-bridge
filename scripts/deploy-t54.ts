import { ethers } from "hardhat";
import fs from 'fs';

function writeTextToFile(text: string, filePath: string): void {
  fs.writeFileSync(filePath, text);
  console.log(`Text written to ${filePath}`);
}


async function main() {
  const T54NFT = await ethers.getContractFactory("T54NFTs")
  const t54 = await T54NFT.deploy();
  await t54.deployed();
  console.log("T54 nft deployed: ", t54.address)

  const T54Bridge = await ethers.getContractFactory("T54Bridge")
  const t54Bridge = await T54Bridge.deploy(t54.address);
  await t54Bridge.deployed();
  console.log("T54Bridge contract deployed: ", t54Bridge.address)


  writeTextToFile(`T54_ADDRESS=${t54.address}\nT54BRIDGE_ADDR=${t54Bridge.address}`,'.env.t54');

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
