import { ethers } from "hardhat";
import fs from 'fs';

function writeTextToFile(text: string, filePath: string): void {
  fs.writeFileSync(filePath, text);
  console.log(`Text written to ${filePath}`);
}


async function main() {
  // const LEO = await ethers.getContractFactory("LEO");
  // const leo = await LEO.deploy();
  // await leo.deployed();

  // const T54 = await ethers.getContractFactory("T54");
  // const t54 = await T54.deploy();

  // const Bridge = await ethers.getContractFactory("Bridge");
  // const bridge = await Bridge.deploy(leo.address, t54.address);

  // console.log("Leopard token deployed: ", leo.address);
  // console.log("T54 token deployed: ", t54.address);
  // console.log("Bridge deployed: ", bridge.address);

  const LeopardNFT = await ethers.getContractFactory("LeopardNFT")
  const leo = await LeopardNFT.deploy();
  await leo.deployed();
  console.log("Leopard nft deployed: ", leo.address)

  const T54NFT = await ethers.getContractFactory("T54NFT")
  const t54 = await T54NFT.deploy();
  await t54.deployed();
  console.log("T54 nft deployed: ", t54.address)

  const BridgeSEPO = await ethers.getContractFactory("BridgeSEPO")
  const bridgeSepo = await BridgeSEPO.deploy(leo.address);
  await bridgeSepo.deployed();
  console.log("BridgeSEPO contract deployed: ", bridgeSepo.address)

  const BridgeFTM = await ethers.getContractFactory("BridgeFTM")
  const bridgeFTM = await BridgeFTM.deploy(t54.address);
  await bridgeFTM.deployed();
  console.log("BridgeSEPO contract deployed: ", bridgeFTM.address)


  writeTextToFile(`LEO_ADDRESS=${leo.address}\nT54_ADDRESS=${t54.address}\nBRIDGE_SEPO_ADDR=${bridgeSepo.address}\nBRIDGE_FTM_ADDR=${bridgeFTM.address}\n`,'.env.sepo');

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
