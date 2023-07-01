import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import dotenv from 'dotenv'
require("hardhat-abi-exporter");
dotenv.config()
const config = {
  solidity: "0.8.18",
  networks: {
    sepolia: {
      url: `${process.env.URL_SEPOLIA}`,
      accounts: process.env.PRIVATE_KEY_SEPOLIA?.split(',')
    },
    ftm: {
      url: `${process.env.URL_FTM}`,
      accounts: process.env.PRIVATE_KEY_FTM?.split(','),
    },
    arb: {
      url: `${process.env.URL_ARB}`,
      accounts: process.env.PRIVATE_KEY_ARB?.split(','),
    }
  },
  abiExporter: {
    path: './app/pkg/resource',
    clear: true,
    flat: true,
    spacing: 2,
  }
};

export default config;
