import { describe, expect, it } from "vitest";
import { uniqueAddresses } from "./useMyAccounts";

describe("uniqueAddresses", () => {
  it("dedupes case-insensitively, preserving checksum of first seen", () => {
    expect(uniqueAddresses(["0xABC", "0xabc", "0xDEF"])).toEqual(["0xABC", "0xDEF"]);
  });
});
