import { describe, expect, it } from "vitest";
import { keccak256, toBytes } from "viem";
import { ROLE_HASHES } from "./roles";

describe("ROLE_HASHES", () => {
  it("uses zero bytes32 for DEFAULT_ADMIN_ROLE", () => {
    expect(ROLE_HASHES.DEFAULT_ADMIN_ROLE).toBe(`0x${"0".repeat(64)}`);
  });

  it("hashes named roles with keccak256", () => {
    expect(ROLE_HASHES.BOT_ADMIN_ROLE).toBe(keccak256(toBytes("BOT_ADMIN_ROLE")));
  });
});
