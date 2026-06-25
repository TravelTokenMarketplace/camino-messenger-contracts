import { describe, expect, it } from "vitest";
import { keccak256, toBytes } from "viem";
import { BOOKINGTOKEN_ROLES, MANAGER_ROLES, ROLE_HASHES } from "./roles";

describe("ROLE_HASHES", () => {
  it("uses zero bytes32 for DEFAULT_ADMIN_ROLE", () => {
    expect(ROLE_HASHES.DEFAULT_ADMIN_ROLE).toBe(`0x${"0".repeat(64)}`);
  });

  it("hashes named roles with keccak256", () => {
    expect(ROLE_HASHES.BOT_ADMIN_ROLE).toBe(keccak256(toBytes("BOT_ADMIN_ROLE")));
  });
});

describe("manager & booking token roles", () => {
  it("manager roles include admin, pauser, upgrader, versioner, service registry admin", () => {
    expect([...MANAGER_ROLES]).toEqual(
      expect.arrayContaining([
        "DEFAULT_ADMIN_ROLE",
        "PAUSER_ROLE",
        "UPGRADER_ROLE",
        "VERSIONER_ROLE",
        "SERVICE_REGISTRY_ADMIN_ROLE",
      ]),
    );
  });

  it("booking token roles include admin, upgrader, min expiration admin", () => {
    expect([...BOOKINGTOKEN_ROLES]).toEqual(["DEFAULT_ADMIN_ROLE", "UPGRADER_ROLE", "MIN_EXPIRATION_ADMIN_ROLE"]);
  });

  it("hashes MIN_EXPIRATION_ADMIN_ROLE with keccak256", () => {
    expect(ROLE_HASHES.MIN_EXPIRATION_ADMIN_ROLE).toBe(keccak256(toBytes("MIN_EXPIRATION_ADMIN_ROLE")));
  });
});
