const Migrations = artifacts.require("Migrations");
const NativeBridge = artifacts.require("NativeBridge");
const NativeProxy = artifacts.require("NativeProxy");

module.exports = function(deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(NativeBridge);
  deployer.deploy(NativeProxy);
};
